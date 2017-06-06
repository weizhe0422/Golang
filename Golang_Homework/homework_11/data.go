package main

const (
	OpenDataUrl string = "http://ptx.transportdata.tw/MOTC/v2/Rail/THSR/Station?$top=30&$format=JSON"
	//OpenDataUrl string = "http://data.taipei/opendata/datalist/apiAccess?scope=resourceAquire&rid=f4a75ba9-7721-4363-884d-c3820b0b917c"
)

type THSRAvailSeat []struct {
	StationID   string `json:"StationID"`
	StationName struct {
		ZhTw string `json:"Zh_tw"`
		En   string `json:"En"`
	} `json:"StationName"`
	StationPosition struct {
		PositionLat float64 `json:"PositionLat"`
		PositionLon float64 `json:"PositionLon"`
	} `json:"StationPosition"`
	StationAddress string `json:"StationAddress"`
	OperatorID     string `json:"OperatorID"`
}

/*
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
}*/
