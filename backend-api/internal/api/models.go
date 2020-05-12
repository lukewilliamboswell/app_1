package api

// Thing Model
type Thing struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Features []string `json:"features"`
}
