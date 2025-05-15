package api

import (
	"fmt"
)

type MeasureResponse struct {
	Coverage               Coverage               `json:"coverage"`
	Policies              []Policy               `json:"policies"`
	RequestID             string                 `json:"requestId"`
	TotalEmissions        float64               `json:"totalEmissions"`
	TotalEmissionsBreakdown EmissionsBreakdown    `json:"totalEmissionsBreakdown"`
	Rows                  []Row                  `json:"rows"`
}

type Coverage struct {
	AdFormats struct {
		Generic        int    `json:"generic"`
		Metric         string `json:"metric"`
		Unknown        int    `json:"unknown"`
		VendorSpecific int    `json:"vendorSpecific"`
	} `json:"adFormats"`
	Channels struct {
		Deprecated int    `json:"deprecated"`
		Metric     string `json:"metric"`
		Modeled    int    `json:"modeled"`
		Unknown    int    `json:"unknown"`
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
		Metric  string `json:"metric"`
		Modeled int    `json:"modeled"`
		Skipped int    `json:"skipped"`
	} `json:"totalImpressions"`
	TotalRows struct {
		Metric  string `json:"metric"`
		Modeled int    `json:"modeled"`
		Skipped int    `json:"skipped"`
	} `json:"totalRows"`
}

// Policy represents a top-level policy with integer compliant/noncompliant counts
type Policy struct {
	Compliant    int    `json:"compliant"`
	Noncompliant int    `json:"noncompliant"`
	Policy       string `json:"policy"`
	PolicyOwner  string `json:"policyOwner"`
}

// RowPolicy represents a row-level policy with boolean compliant field
type RowPolicy struct {
	Compliant   bool   `json:"compliant"`
	Policy      string `json:"policy"`
	PolicyOwner string `json:"policyOwner"`
}

type EmissionsBreakdown struct {
	Framework string `json:"framework"`
	Totals    struct {
		AdSelection       float64 `json:"adSelection"`
		CreativeDelivery  float64 `json:"creativeDelivery"`
		MediaDistribution float64 `json:"mediaDistribution"`
	} `json:"totals"`
}

type Row struct {
	Coverage          RowCoverage          `json:"coverage"`
	EmissionsBreakdown RowEmissionsBreakdown `json:"emissionsBreakdown"`
	InventoryCoverage  string               `json:"inventoryCoverage"`
	Policies           []RowPolicy          `json:"policies"`  // Changed to RowPolicy
	RowIdentifier      string               `json:"rowIdentifier"`
	TotalEmissions     float64              `json:"totalEmissions"`
	Internal           Internal             `json:"internal"`
}

type RowCoverage struct {
	AdFormat struct {
		Name     string `json:"name"`
		Value    string `json:"value"`
		Verified bool   `json:"verified"`
	} `json:"adFormat"`
	Channel struct {
		Value string `json:"value"`
	} `json:"channel"`
	CompensationProvider struct {
		Value string `json:"value"`
	} `json:"compensationProvider"`
	Impressions struct {
		Modeled   int `json:"modeled"`
		Processed int `json:"processed"`
		Skipped   int `json:"skipped"`
	} `json:"impressions"`
	Property struct {
		Value string `json:"value"`
	} `json:"property"`
	SupplyGraph struct {
		Logical struct {
			AverageDepth float64 `json:"averageDepth"`
			MaxDepth     int     `json:"maxDepth"`
			MinDepth     int     `json:"minDepth"`
		} `json:"logical"`
		Technical struct {
			AverageDepth float64 `json:"averageDepth"`
			MaxDepth     int     `json:"maxDepth"`
			MinDepth     int     `json:"minDepth"`
		} `json:"technical"`
		TotalCount int `json:"totalCount"`
	} `json:"supplyGraph"`
}

type RowEmissionsBreakdown struct {
	Breakdown struct {
		AdSelection struct {
			Breakdown struct {
				AdPlatform struct {
					Emissions float64 `json:"emissions"`
				} `json:"adPlatform"`
				DataTransfer struct {
					Emissions float64 `json:"emissions"`
				} `json:"dataTransfer"`
			} `json:"breakdown"`
			Total float64 `json:"total"`
		} `json:"adSelection"`
		Compensated struct {
			Breakdown struct {
				Compensation struct {
					Emissions float64 `json:"emissions"`
					Provider  string  `json:"provider"`
				} `json:"compensation"`
			} `json:"breakdown"`
			Total float64 `json:"total"`
		} `json:"compensated"`
		CreativeDelivery struct {
			Breakdown struct {
				AdPlatform struct {
					Emissions float64 `json:"emissions"`
				} `json:"adPlatform"`
				DataTransfer struct {
					Emissions float64 `json:"emissions"`
				} `json:"dataTransfer"`
			} `json:"breakdown"`
			Total float64 `json:"total"`
		} `json:"creativeDelivery"`
		MediaDistribution struct {
			Breakdown struct {
				Corporate struct {
					Emissions float64 `json:"emissions"`
				} `json:"corporate"`
				DataTransfer struct {
					Emissions float64 `json:"emissions"`
				} `json:"dataTransfer"`
			} `json:"breakdown"`
			Total float64 `json:"total"`
		} `json:"mediaDistribution"`
	} `json:"breakdown"`
	Framework string `json:"framework"`
}

type Internal struct {
	CountryRegionGCO2PerKwh int     `json:"countryRegionGCO2PerKwh"`
	CountryRegionCountry    string  `json:"countryRegionCountry"`
	Channel                 string  `json:"channel"`
	DeviceType             string  `json:"deviceType"`
	PropertyId             int     `json:"propertyId"`
	PropertyInventoryType  string  `json:"propertyInventoryType"`
	PropertyName           string  `json:"propertyName"`
	BenchmarkPercentile    int     `json:"benchmarkPercentile"`
	IsMFA                  bool    `json:"isMFA"`
	PolicyEvaluationData   PolicyEvaluationData `json:"policyEvaluationData"`
}

type PolicyEvaluationData struct {
	PropertyId          int    `json:"propertyId"`
	IsMFA              bool   `json:"isMFA"`
	IsInventory        bool   `json:"isInventory"`
	Channel            string `json:"channel"`
	ChannelStatus      string `json:"channelStatus"`
	BenchmarksPercentile int  `json:"benchmarksPercentile"`
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
