package models

// RoleUsers Structure of the role users
type RoleUsers struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	RoleID    string `json:"role_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
}
