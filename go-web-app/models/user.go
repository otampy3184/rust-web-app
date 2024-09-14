package models

// Json出力時のキーも指定したデータモデル
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
