package Model

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
	Type     int    `json:"type"`
}

type Song struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Duration float32 `json:"duration"`
	Singer   string  `json:"singer"`
}

type Playlist struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	State  bool      `json:"state"`
	UserID int       `json:"user"`
}

type DetailPlaylistSong struct {
	PlaylistID int `json:"playlist"`
	SongID     int `json:"song"`
	Played     int `json:"timeplayed"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SongsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Song `json:"data"`
}
