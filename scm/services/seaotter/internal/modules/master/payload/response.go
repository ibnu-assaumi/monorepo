package payload

import "monorepo/services/seaotter/pkg/shared/model"

// ResponseSOPrefix type
type ResponseSOPrefix struct{}

// NewResponseSOPrefix type
func NewResponseSOPrefix(data model.MasterSOPrefix) ResponseSOPrefix {
	return ResponseSOPrefix{}
}
