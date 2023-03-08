package data

const (
	ModuleName    = "role"
	AddUserAction = "add_user"
)

type AddUserPayload struct {
	AccessLevel interface{} `json:"access_level,omitempty"`
	Action      string      `json:"action"`
	Link        string      `json:"link"`
	UserId      string      `json:"user_id"`
	Username    *string     `json:"username"`
	Phone       *string     `json:"phone"`
}
