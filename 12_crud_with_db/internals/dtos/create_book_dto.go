package create_book_dto

type CreateBookRequest struct {
    Title       string `json:"title"`
    Author      string `json:"author"`
    Description string `json:"description"`
    Price       int    `json:"price"`
    Quantity    int    `json:"quantity"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}