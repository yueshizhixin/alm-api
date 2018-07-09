package note

type Note struct {
	Id         uint32    `json:"-" gorm:"primary_key;"`
}

