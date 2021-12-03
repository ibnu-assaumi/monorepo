package constant

import (
	"context"
	"strings"

	"github.com/Bhinneka/candi/candishared"
)

// Error constant
const (
	ErrContextCancelled = "context cancelled"
)

// KeyContextLang constant
const KeyContextLang = "lang"

var errormap = map[string]string{
	"record_not_found_en": "record not found",
	"record_not_found_id": "data tidak ditemukan",
	"invalid_data_en":     "invalid data",
	"invalid_data_id":     "data tidak valid",
}

// GetErrorMessage : get error mesage langunage profiling
func GetErrorMessage(ctx context.Context, key string) string {
	lang, _ := candishared.GetValueFromContext(ctx, KeyContextLang).(string)
	if strings.TrimSpace(lang) == "" {
		lang = "id"
	}
	if val, ok := errormap[key+"_"+strings.ToLower(lang)]; ok {
		return val
	}
	return "error"
}
