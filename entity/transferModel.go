package entity

type TransferInfo_XieLinXuan struct {
	ID int `sql:"id" json:"id" form:"id"`
	ToAddress string `sql:"to_address" json:"to_address" form:"ToAddress"`
	Award float64 `sql:"award" json:"award" form:"Award"`
	GetTime uint `sql:"get_time" json:"get_time" form:"GetTime"`
}

type TransferLog_XieLinXuan struct {
	Time int `sql:"time" json:"time" form:"time"`
	OutNum float64 `sql:"out_num" json:"out_num" form:"out_num"`
	OutCount int `sql:"out_count" json:"out_count" form:"out_count"`
}
