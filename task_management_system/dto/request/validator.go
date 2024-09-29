package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"task_management_system/errors"

	"gopkg.in/validator.v2"
)

type IValidator interface {
	Validate(ctx context.Context) errors.IError
}

// DecodeAndValidate decodes request payload and validates if required fields are filled
func DecodeAndValidate(ctx context.Context, r *http.Request, v IValidator) errors.IError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[DecodeAndValidate|failed] err:%s", err.Error())
		err2 := errors.New(errors.FailedToDecodeRequestBodyID, errors.FailedToDecodeRequestBodyCode, err.Error())
		return err2
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	d := json.NewDecoder(r.Body)

	defer r.Body.Close()
	d.UseNumber()
	err = d.Decode(v)
	if err != nil {
		log.Printf("[Decode|failed] err:%s", err.Error())
		err2 := errors.New(errors.FailedToDecodeRequestBodyID, errors.FailedToDecodeRequestBodyCode, err.Error())
		return err2
	}
	return v.Validate(r.Context())
}

// validateFields checks if the required fields in a model is filled.
func ValidateFields(model interface{}) errors.IError {
	err := validator.Validate(model)
	if err != nil {
		log.Printf("[Validate|failed] err:%s", err.Error())
		errs, ok := err.(validator.ErrorMap)
		if ok {
			var errOuts []string
			for f, e := range errs {
				msg := e.Error()
				if e.Error() == validator.ErrZeroValue.Error() {
					msg = errors.ValidatorZeroValueMsg
				}
				errOuts = append(errOuts, fmt.Sprintf("%s (%v)", f, msg))
			}
			return errors.New(errors.MissingArgumentGenericID, errors.MissingArgumentGenericCode, fmt.Sprintf(errors.ValidatorMissingArgsMsg, strings.Join(errOuts, ", ")))
		} else {
			err2 := errors.New(errors.MissingArgumentGenericID, errors.MissingArgumentGenericCode, err.Error())
			return err2
		}
	}
	return nil
}
