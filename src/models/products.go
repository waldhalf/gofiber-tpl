package models

type Product struct {
	Model
	Title 		string 	`json:"title"`
	Description string 	`json:"Description"`
	Image 		string 	`json:"image"`
	Price 		float64	`json:"price"`
}