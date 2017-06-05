package main

const (
	OpenDataUrl string = "http://ptx.transportdata.tw/MOTC/v2/Rail/THSR/Station?$top=30&$format=JSON"
)

type THSRAvailSeat struct {
	Result struct {
		StationID       string         `json:"StationID"`
		StationName     []StationInfo1 `jsoin:"StationName"`
		StationPosition []StationInfo2 `json:"StationPosition"`
		StationAddress  string         `json:"StationAddress"`
		OperatorID      string         `jsoin:"OperatorID"`
	} //`json:"Result"`
}

type StationInfo1 struct {
	Zhtw string `jsoin:"Zh_tw"`
	En   string `json:"En"`
}

type StationInfo2 struct {
	PositionLat string `jsoin:"PositionLat"`
	PositionLon string `json:"PositionLon"`
}
