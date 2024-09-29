package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"

	"task_management_system/api/controller"
	"task_management_system/api/respwriter"
	"task_management_system/api/router"
	"task_management_system/application"
	"task_management_system/config"
	domainServices "task_management_system/domain/services"
	"task_management_system/domain/task_details"
	"task_management_system/domain/users"
	infrasql "task_management_system/infrastructure/mysql"

	"github.com/spf13/viper"
)

// IApp defines the app interface
type IApp interface {
	Init(configPath string) error
	Start() error
	GetRouter() router.IRouter
}

type ServiceApp struct {
	config *config.GeneralConfig
	router router.IRouter

	//repositories
	sqlStore        infrasql.IMysqlStore
	usersRepo       users.IUsersRepo
	taskDetailsRepo task_details.ITaskDetailsRepo

	//Services
	taskDetailsServices domainServices.ITaskDetailsService
	usersServices       domainServices.IUsersService

	//application
	taskDetailsApplication application.ITaskApplication
	userAppication         application.IUsersApplication
}

// NewApplication call
func NewApplication() IApp {
	return &ServiceApp{}
}

// Init function Initialize the notification service app
func (n *ServiceApp) Init(configPath string) error {
	var err error
	n.config, err = loadConfig(configPath)
	if err != nil {
		log.Printf("failed to loadConfig: %s", err)
		return err
	}
	n.initializeDependencies(*n.config)
	return nil
}

// Start is to start the application
func (n *ServiceApp) Start() error {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", n.config.ApplicationPort),
		Handler: n.router.GetMux(),
	}
	startServer(httpServer)
	//gracefully shut down
	shutDownServer(httpServer)
	return nil
}

func (n *ServiceApp) GetRouter() router.IRouter {
	return n.router
}

func startServer(server *http.Server) {
	log.Printf("########## Server starting at %+v ##########", time.Now())
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("server failed to start:%s", err.Error())
		}
	}()
}
func shutDownServer(server *http.Server) {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("########### Server gracefully shutdown fail", err)
		}
	}
}
func loadConfig(filePath string) (*config.GeneralConfig, error) {

	config := &config.GeneralConfig{}

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("failed to read config file: %s", err)
		return nil, err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Printf("failed to deserialize config file: %s", err)
		return nil, err
	}
	return config, nil
}

func (n *ServiceApp) initializeDependencies(config config.GeneralConfig) {
	n.initSqlDatastores(*n.config.MySqlDBConfig)
	n.localDependencyInjection(&config)
}

func (n *ServiceApp) localDependencyInjection(config *config.GeneralConfig) {
	n.taskDetailsServices = domainServices.NewTaskDetailsService(config, n.taskDetailsRepo)
	n.usersServices = domainServices.NewUsersService(config, n.taskDetailsRepo, n.usersRepo)

	n.taskDetailsApplication = application.NewTaskApp(n.taskDetailsServices)
	n.userAppication = application.NewUserApp(n.usersServices)

	baseController := controller.NewBaseController(respwriter.NewHTTPResponseService())

	n.router = router.NewRouter(n.config)
	n.router.InitRoutes(router.InitRouterDto{
		HealthController:      controller.NewHealthController(baseController),
		UsersController:       controller.NewUsersController(baseController, n.userAppication),
		TaskDetailsController: controller.NewTaskDetailsController(baseController, n.taskDetailsApplication),
	})
}

// initSqlDatastores initialize sql datastore
func (s *ServiceApp) initSqlDatastores(config config.SQLConfig) {
	if !config.Enabled {
		return
	}
	s.sqlStore = infrasql.NewDatastore(&config)

	s.taskDetailsRepo = infrasql.NewTaskDetailsDatastoreApi(s.sqlStore)
	s.usersRepo = infrasql.NewUsersDatastoreApi(s.sqlStore)

}
