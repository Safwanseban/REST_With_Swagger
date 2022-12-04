package models



type User struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Books      []Books
}
type Books struct {
	Book_ID     uint   `json:"book_id"`
	Author_name string `json:"author_name"`
	Book_name   string `json:"book_name"`
}
// func (b *Books) create() []any {
// v

// // sl=append(sl, Books{Book_ID: 1, Book_name: "Wings of fire", Author_name: "APJ Abdul-Kalam"})
// // sl=append(sl, Books{Book_ID: 2, Book_name: "The Story of My Experiments with Truth",
// // 		Author_name: "Mahatma Gandhi"})
	

// return sl
// }

