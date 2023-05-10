package model

import (
	"strings"
	"time"
)

type Borrow struct {
	BORROW_ID int `gorm:"primaryKey" form:"BORROW_ID" json:"borrowId"`

	BOOK_ID int `form:"BOOK_NAME" json:"bookName"`
	BOOK_ISBN *string `form:"BOOK_ISBN" json:"bookIsbn"`
	
	STUDENT_ID int `form:"STUDENT_ID" json:"studentId"`
	STUDENT_CODE *string `form:"STUDENT_CODE" json:"studentCode"`

	DATE_BORROW *CustomTime `form:"DATE_BORROW"  gorm:"type:DATETIME" json:"dateBorrow"`
	DATE_RETURN *CustomTime `form:"DATE_RETURN"  gorm:"type:DATETIME" json:"dateReturn"`
	BORROW_QUANTITY int `form:"BORROW_QUANTITY" json:"borrowQuantity"`

	STATUS_ID int `form:"STATUS_ID" json:"statusId"`
	STATUS_CODE *string `form:"STATUS_CODE" json:"statusCode"`
}

// เป็นการแปลงค่า 2023-05-01 ไปเป็น type time
type CustomTime struct {
    time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
    str := strings.Trim(string(b), "\"")
    t, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    c.Time = t
    return nil
}
//-------------------------------//

type GetBorrow struct {
	BORROW_ID int `gorm:"primaryKey" form:"BORROW_ID" json:"borrowId"`
	BOOK_NAME *string `form:"BOOK_NAME" json:"bookName"`
	STUDENT_NAME *string `form:"STUDENT_NAME" json:"studentName"`
	DATE_BORROW *string `form:"DATE_BORROW" json:"dateBorrow"`
	DATE_RETURN *string `form:"DATE_RETURN" json:"dateReturn"`
	BORROW_QUANTITY *int `form:"BORROW_QUANTITY" json:"borrowQuantity"`
	STATUS_NAME *string `form:"STATUS_NAME" json:"statusBorrow"`
}
func (borrow *Borrow) TableName() string {
	return "borrow"
}
type BorrowResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []GetBorrow
}
