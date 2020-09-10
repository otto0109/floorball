package dto

type Article struct {
	ID        int64  `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Article   string `json:"article"`
	Teams     []Team `json:"teams"`
	Category  string `json:"category"`
}
