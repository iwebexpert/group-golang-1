package models

// PostItem ...
type PostItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// PostItemSlice ...
type PostItemSlice []PostItem
