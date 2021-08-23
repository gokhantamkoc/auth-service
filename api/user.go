package api

type User struct {
	ID			string 					`json:"id"`
	FirstName 	string 					`json:"firstName"`
	LastName 	string 					`json:"lastName"`
	Email		string 					`json:"email"`
	Username 	string 					`json:"username"`
	Attributes 	*map[string][]string	`json:"attributes"`
}