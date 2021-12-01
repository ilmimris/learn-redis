package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  `json:"address"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  `json:"company"`
}
