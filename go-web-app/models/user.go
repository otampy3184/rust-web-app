package models

// Json出力時のキーも指定したデータモデル
type User struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
}
