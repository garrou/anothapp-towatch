package models

type ShowInfos struct {
	Show struct {
		Seasons string `json:"seasons"`
		Status  string `json:"status"`
	} `json:"show"`
}

type ShowDto struct {
	Id int

	NbToWatch int

	Ended bool
}
