package model
type Book struct {
	BOOK_ID int `gorm:"primaryKey" json:"bookId"`
	BOOK_ISBN *string `form:"BOOK_ISBN" json:"bookIsbn"`
	BOOK_NAME *string `form:"BOOK_NAME" json:"bookName"`
	BOOK_QUANILTY *int `form:"BOOK_QUANILTY" json:"bookQuanilty"`

	AUTHOR_ID *int `form:"AUTHOR_ID" json:"authorId"`
	AUTHOR_CODE *string `form:"AUTHOR_CODE" json:"authorCode"`

	CATEGORY_ID  *int   `form:"CATEGORY_ID" json:"categoryId"`
	CATEGORY_CODE  *int   `form:"CATEGORY_CODE" json:"categoryCode"`
}
type ForeignkeyBook struct{
	AUTHOR_CODE *string `form:"AUTHOR_CODE" json:"authorCode"`
	CATEGORY_CODE  *int   `form:"CATEGORY_CODE" json:"categoryCode"`
}

type GetBook struct {
	BOOK_ID int `gorm:"primaryKey" json:"bookId"`
	BOOK_ISBN *string `json:"bookIsbn"`
	BOOK_NAME *string `json:"bookName"`
	BOOK_QUANILTY *int `form:"BOOK_QUANILTY" json:"bookQuanilty"`
	AUTHOR_ID  int   `json:"-"`
	CATEGORY_ID   int   `json:"-"`
	Author Author `json:"author" gorm:"index" gorm:"foreignkey:AUTHOR_ID"`
	Category Category `json:"category" gorm:"index" gorm:"foreignkey:CATEGORY_ID"`
}

func (book *Book) TableName() string {
	return "book"
}


type BookResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []GetBook
}
