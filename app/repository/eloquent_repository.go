package repository

import (
	"gorm.io/gorm"
	"strings"
)

func getQueryLike(result *gorm.DB, key string, value string) *gorm.DB {
	valueTrim := strings.TrimSpace(value)
	valueTrimLen := len(valueTrim)
	if valueTrim[:1] == `"` && valueTrim[(valueTrimLen-1):] == `"` {
		result = result.Where(key+" = ?", valueTrim[1:(valueTrimLen-1)])
	} else {
		result = result.Where(key+" LIKE ?", "%"+convertStringQueryLike(valueTrim)+"%")
	}
	return result
}

func convertStringQueryLike(value string) string {
	return strings.ReplaceAll(value, "%", `\%`)
}
