package models

type ExternalLimits struct {
	Index                  int     `json:"index"`
	Name                   string  `json:"name"`
	Superior               float64 `json:"superior"`
	Iinferior              float64 `json:"inferior"`
	Description            string  `json:"description"`
	Label                  string  `json:"label"`
	Result_meaning         string  `json:"result_meaning"`
	Sex                    string  `json:"sex"`
	Max_age                int     `json:"max_age"`
	Min_age                int     `json:"min_age"`
	Category               string  `json:"category"`
	ReferenceValue         string  `json:"referenceValue"`
	ReferenceValueRelative string  `json:"referenceValueRelative"`
	Color                  string  `json:"color"`
	WrongRange             bool    `json:"wrongRange"`
}
