package model

type Status struct {
	STATUS_ID   int     `gorm:"primaryKey" form:"statusId" json:"statusId"`
	STATUS_CODE *string `form:"statusCode" json:"statusCode"`
	STATUS_NAME *string ` form:"statusName" json:"statusName"`
}
func (status *Status) TableName() string {
	return "status"
}

type StatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Status
}