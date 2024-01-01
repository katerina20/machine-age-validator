package models

type RequestData struct {
	Machines []Machine `json:"machines"`
}

type ResponseData struct {
	Outliers []Machine `json:"outliers"`
}
