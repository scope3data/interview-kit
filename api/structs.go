package api

import (
	"fmt"
	"net/url"
	"strconv"
)

type MeasureResponse struct {
    Coverage struct {
        AdFormats struct {
            Generic        int     `json:"generic"`
            Metric        string   `json:"metric"`
            Unknown       int      `json:"unknown"`
            VendorSpecific int     `json:"vendorSpecific"`
        } `json:"adFormats"`
        Channels struct {
            Deprecated int    `json:"deprecated"`
            Metric    string  `json:"metric"`
            Modeled   int     `json:"modeled"`
            Unknown   int     `json:"unknown"`
        } `json:"channels"`
        MediaOwners struct {
            Metric  string `json:"metric"`
            Modeled int    `json:"modeled"`
            Unknown int    `json:"unknown"`
        } `json:"mediaOwners"`
        Properties struct {
            Metric  string `json:"metric"`
            Modeled int    `json:"modeled"`
            Unknown int    `json:"unknown"`
        } `json:"properties"`
        Sellers struct {
            Metric  string `json:"metric"`
            Modeled int    `json:"modeled"`
            Unknown int    `json:"unknown"`
        } `json:"sellers"`
        TotalImpressions struct {
            Metric   string `json:"metric"`
            Modeled  int    `json:"modeled"`
            Skipped  int    `json:"skipped"`
        } `json:"totalImpressions"`
        TotalRows struct {
            Metric  string `json:"metric"`
            Modeled int    `json:"modeled"`
            Skipped int    `json:"skipped"`
        } `json:"totalRows"`
    } `json:"coverage"`
    Policies []struct {
        Compliant    int    `json:"compliant"`
        Noncompliant int    `json:"noncompliant"`
        Policy       string `json:"policy"`
        PolicyOwner  string `json:"policyOwner"`
    } `json:"policies"`
    RequestID string  `json:"requestId"`
    TotalEmissions float64 `json:"totalEmissions"`
    TotalEmissionsBreakdown struct {
        Framework string `json:"framework"`
        Totals struct {
            AdSelection      float64 `json:"adSelection"`
            CreativeDelivery float64 `json:"creativeDelivery"`
            MediaDistribution float64 `json:"mediaDistribution"`
        } `json:"totals"`
    } `json:"totalEmissionsBreakdown"`
}

func (r *MeasureResponse) String() string {
    return fmt.Sprintf(
        "RequestID: %s\nTotal Emissions: %.4f\nBreakdown:\n"+
        "  - Ad Selection: %.4f\n"+
        "  - Creative Delivery: %.4f\n"+
        "  - Media Distribution: %.4f\n",
        r.RequestID,
        r.TotalEmissions,
        r.TotalEmissionsBreakdown.Totals.AdSelection,
        r.TotalEmissionsBreakdown.Totals.CreativeDelivery,
        r.TotalEmissionsBreakdown.Totals.MediaDistribution,
    )
}

type MeasureRow struct {
	RowIdentifier string `json:"rowIdentifier"`
	Impressions   int    `json:"impressions"`
	UtcDatetime   string `json:"utcDatetime"`
	InventoryId   string `json:"inventoryId"`
	AppStore    string `json:"appStore"`
	Country     string `json:"country"`
	Region      string `json:"region"`
	DeviceType  string `json:"deviceType"`
	Channel     string `json:"channel"`
	Network     string `json:"network"`
}

type RequestRow struct {
	InventoryId string `json:"inventoryId"`
	Impressions int    `json:"impressions"`
	DeviceType  string `json:"deviceType"`
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
    IncludeRows bool   `url:"include_rows"` // default is true
    Latest      bool   `url:"latest"` // default is false
    Fields      string `url:"fields"` // default is "all"
    Framework   string `url:"framework"` // default is "scope3"
}

func (p *MeasureQueryParams) ToQueryString() string {
    values := url.Values{}
    values.Add("include_rows", strconv.FormatBool(p.IncludeRows))
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
