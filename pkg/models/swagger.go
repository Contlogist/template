package models

// swagger:parameters myRoute
type MultiFile struct {
	//in: formData
	//swagger:file
	Files interface{} `json:"files"`
}
