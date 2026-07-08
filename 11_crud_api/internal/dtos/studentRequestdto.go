package dtos

// it is useful when we start reading JSON from the request
type CreateStudentRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
