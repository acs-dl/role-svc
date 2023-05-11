package requests

import (
	"encoding/json"
	"net/http"

	"github.com/acs-dl/role-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddUsersRequest struct {
	Data resources.Request `json:"data"`
}

func NewCreateRequestRequest(r *http.Request) (AddUsersRequest, error) {
	var request AddUsersRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, err
	}

	return request, request.validate()
}

func (r *AddUsersRequest) validate() error {
	err := validation.Errors{
		"module":       validation.Validate(&r.Data.Attributes.Module, validation.Required),
		"link":         validation.Validate(&r.Data.Attributes.Link, validation.Required),
		"access_level": validation.Validate(&r.Data.Attributes.AccessLevel, validation.Required),
		"users":        validation.Validate(&r.Data.Attributes.Users, validation.Required),
		"from_user":    validation.Validate(&r.Data.Attributes.FromUser, validation.Required),
		"to_user":      validation.Validate(&r.Data.Attributes.ToUser, validation.Required),
	}.Filter()

	if err != nil {
		return err
	}

	return validateUsers(r.Data.Attributes.Users)
}

func validateUsers(users []resources.User) error {
	for _, user := range users {
		phoneValidationCase := validation.When(user.Attributes.Username == nil, validation.Required.Error("phone is required if username is not set"))
		usernameValidationCase := validation.When(user.Attributes.Phone == nil, validation.Required.Error("username is required if phone is not set"))

		err := validation.Errors{
			"username": validation.Validate(&user.Attributes.Username, usernameValidationCase),
			"phone":    validation.Validate(&user.Attributes.Phone, phoneValidationCase),
			"user_id":  validation.Validate(&user.Attributes.UserId, validation.Required),
		}.Filter()

		if err != nil {
			return err
		}
	}

	return nil
}
