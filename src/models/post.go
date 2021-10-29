package models

type Post struct {
	Model
	Title 		string	`json:"title"`
	SubTitle 	string	`json:"subtitle"`
	Content 	string	`json:"content"`
	AuthorId 	uint	`json:"author_id"`
	ImageUrl 	string	`json:"image_url"`
	ReleaseDate	string	`json:"release_date"`
	CategoryId  uint	`json:"category_id"`
}