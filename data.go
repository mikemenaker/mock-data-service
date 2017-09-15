package main

type StatsData struct {
	Data []int `json:"data"`
}

type TextData struct {
	Data string `json:"data"`
}

type TableDataItem struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Value2 string `json:"value2"`
}

type TableData struct {
	Data []TableDataItem `json:"data"`
}
