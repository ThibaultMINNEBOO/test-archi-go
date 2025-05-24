package users

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
