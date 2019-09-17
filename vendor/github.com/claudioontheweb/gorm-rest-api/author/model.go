package author

type Author struct {
	ID   int `gorm:"primary_key" json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}