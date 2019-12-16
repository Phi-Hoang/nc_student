package db

// Student contains information for a student
type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

// Error with code and message
type Error struct {
	Code int
	Msg  string
}

// StudentUpdateRequest contains information for a request to update a student
type StudentUpdateRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
