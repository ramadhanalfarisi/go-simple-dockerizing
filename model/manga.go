package model

type Manga struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Volumes uint8 `json:"volumes"`
	Chapters uint16 `json:"chapters"`
	Author string `json:"author"`
}

type PostManga struct {
	Title string `json:"title" binding:"required"`
	Genre string `json:"genre"`
	Volumes uint8 `json:"volumes"`
	Chapters uint16 `json:"chapters"`
	Author string `json:"author"`
}