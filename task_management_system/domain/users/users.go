package users

type Users struct {
	Id        *string `db:"id"`
	Name      *string `db:"name"`
	EmailId   *string `db:"email"`
	IsDeleted *bool   `db:"is_deleted"`
	// CreatedAt *time.Time `db:"created_at"`
	// UpdatedAt *time.Time `db:"updated_at"`
}

type UserColumn string

const (
	ColID        UserColumn = "id"
	ColName      UserColumn = "name"
	ColEmail     UserColumn = "email"
	ColIsDeleted UserColumn = "is_deleted"
	ColCreatedAt UserColumn = "created_at"
	ColUpdatedAt UserColumn = "updated_at"
)
