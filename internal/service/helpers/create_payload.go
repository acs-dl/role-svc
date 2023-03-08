package helpers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/acs/role-svc/internal/data"
	"gitlab.com/distributed_lab/acs/role-svc/resources"
	"strconv"
)

func CreateAddUserPayload(link string, user resources.User, accessLevel string) (json.RawMessage, error) {
	payload := data.AddUserPayload{
		Action:      data.AddUserAction,
		Link:        link,
		UserId:      user.Attributes.UserId,
		Username:    user.Attributes.Username,
		Phone:       user.Attributes.Phone,
		AccessLevel: accessLevel,
	}

	if accessLevelInt, err := strconv.ParseInt(accessLevel, 10, 64); err == nil {
		payload.AccessLevel = accessLevelInt
	}

	return json.Marshal(payload)
}
