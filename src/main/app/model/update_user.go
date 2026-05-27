package model

type UpdateUserModel struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Gmail    string `json:"gmail"`
}
