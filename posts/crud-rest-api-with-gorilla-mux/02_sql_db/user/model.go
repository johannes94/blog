package user

type User struct {
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
