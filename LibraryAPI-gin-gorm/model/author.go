package model

type Author struct {
	AUTHOR_ID      int    `gorm:"primaryKey" form:"AUTHOR_ID" json:"authorId"`
	AUTHOR_CODE    *string `form:"AUTHOR_CODE" json:"authorCode"`
	AUTHOR_NAME    *string `form:"AUTHOR_NAME" json:"authorName"`
	AUTHOR_ADDRESS *string `form:"AUTHOR_ADDRESS" json:"authorAddress"`
}

func (author *Author) TableName() string {
	return "author"
}

type AuthorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Author
}
