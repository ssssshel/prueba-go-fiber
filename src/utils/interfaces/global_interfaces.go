package global_interfaces

type Response struct {
	Message string
	Data    any
}

type User struct {
	Id        string
	Firstname string
	Lastname  string
}
