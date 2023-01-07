package global_interfaces

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
