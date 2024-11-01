package dao

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	Desc = "DESC" // Desc order
	Asc  = "ASC"  // Asc order
)

func DescOrder(column string) string { return fmt.Sprintf("`%s` %s", column, Desc) }

func AscOrder(column string) string { return fmt.Sprintf("`%s` %s", column, Asc) }

func StructField(obj interface{}) string {
	if obj == nil {
		return ""
	}
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return ""
	}
	jvs := make([]string, 0)
	for i := 0; i < value.NumField(); i++ {
		jv := value.Type().Field(i).Tag.Get("json")
		if jv == "" || jv == "-" || value.Type().Field(i).Tag.Get("gorm") == "-" {
			continue
		}
		if i := strings.Index(jv, ","); i != -1 {
			jv = jv[0:i]
		}
		jvs = append(jvs, "`"+jv+"`")
	}
	return strings.Join(jvs, ",")
}
