package utils

import (
	"fmt"
	"reflect"

	"CodeX/backend/common"
)

func GetConfigFieldValue(config *common.IConfig, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(config).Elem() // 解引用，拿到 struct 值
	fieldVal := v.FieldByName(fieldName)
	if !fieldVal.IsValid() {
		return nil, fmt.Errorf("No such field: %s", fieldName)
	}
	return fieldVal.Interface(), nil
}
