package entity

type Post struct {
	ID	   int  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"Body"`
}
