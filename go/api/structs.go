package api

import (
	"net/url"
	"strconv"
)

type MeasureRow struct {
	RowIdentifier string `json:"rowIdentifier"`
	Impressions   int    `json:"impressions"`
	UtcDatetime   string `json:"utcDatetime"`
	InventoryId   string `json:"inventoryId"`
	AppStore      string `json:"appStore"`
	Country       string `json:"country"`
	Region        string `json:"region"`
	DeviceType    string `json:"deviceType"`
	Channel       string `json:"channel"`
	Network       string `json:"network"`
}

type RequestRow struct {
	InventoryId   string `json:"inventoryId"`
	Impressions   int    `json:"impressions"`
	DeviceType    string `json:"deviceType"`
	RowIdentifier string `json:"rowIdentifier"`
	UtcDatetime   string `json:"utcDatetime"`
}

func NewRequestRow(inventoryId string, utcDatetime string) *RequestRow {
	return &RequestRow{
		InventoryId:   inventoryId,
		Impressions:   1,
		DeviceType:    "pc",
		RowIdentifier: inventoryId,
		UtcDatetime:   utcDatetime,
	}
}

type MeasureQueryParams struct {
	IncludeRows bool   `url:"includeRows"` // default is true
	Latest      bool   `url:"latest"`      // default is false
	Fields      string `url:"fields"`      // default is "all"
	Framework   string `url:"framework"`   // default is "scope3"
}

func (p *MeasureQueryParams) ToQueryString() string {
	values := url.Values{}
	values.Add("includeRows", strconv.FormatBool(p.IncludeRows))
	values.Add("latest", strconv.FormatBool(p.Latest))
	values.Add("fields", p.Fields)
	values.Add("framework", p.Framework)
	return values.Encode()
}

func NewMeasureQueryParams() *MeasureQueryParams {
	return &MeasureQueryParams{
		IncludeRows: true,
		Latest:      false,
		Fields:      "all",
		Framework:   "scope3",
	}
}
