package model

/*
 * Responses
 */
// user
type UserResponse struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

/*
 * Requests
 */
// user
type UserLoginRequest struct {
	Id       string `json:"id" validate:"required,max=256"`
	Password string `json:"password" validate:"required,max=256"`
}
