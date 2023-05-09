package model

type Category struct {
	CATEGORY_ID  int `gorm:"primaryKey" json:"categoryId"`
	CATEGORY_CODE *string `json:"categoryCode"`
	CATEGORY_NAME *string `json:"categoryName"`
}
func (category *Category) TableName() string {
	return "category"
}
