package models

type ExternalAnalyte struct {
	Unit                   string         `json:"unit"`
	Plan_name              string         `json:"plan_name"`
	Sample_type            string         `json:"sample_type"`
	Sample_collection_type string         `json:"sample_collection_type"`
	Subtitle               string         `json:"subtitle"`
	Name                   string         `json:"name"`
	Description            string         `json:"description"`
	Abv                    string         `json:"abv"`
	Metric                 string         `json:"metric"`
	Method                 string         `json:"method"`
	Comercial_name         string         `json:"comercial"`
	DescTISS               string         `json:"descTISS"`
	Duration               string         `json:"duration"`
	Notes                  string         `json:"notes"`
	Info                   string         `json:"info"`
	Limitations            string         `json:"limitations"`
	ReferenceValues        string         `json:"referenceValues"`
	Device                 string         `json:"device"`
	Supplier               string         `json:"supplier"`
	Limits                 ExternalLimits `json:"Limits"`
}
