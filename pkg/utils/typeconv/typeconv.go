package typeconv

import (
	"fmt"
	"strconv"
)

func ToInt(key interface{}) int {
	v, err := strconv.Atoi((key).(string))
	if err != nil {
		return 0
	}
	return v
}

func ToFloat(key interface{}) float64 {
	v, err := strconv.ParseFloat(key.(string), 64)
	if err != nil {
		return 0
	}
	return v
}

func ToInt64(key interface{}) int64 {
	v := int64(key.(float64))
	return v
}

func IChkStr(obj interface{}) string {
	if str, ok := obj.(string); ok {
		return str
	} else {
		return ""
	}
}

func IChkF64(obj interface{}) float64 {
	if str, ok := obj.(float64); ok {
		return str
	} else {
		return 0.0
	}
}

func IChkF64s(obj interface{}) string {
	if str, ok := obj.(float64); ok {
		//return strconv.FormatFloat(str, '', -1, 64)
		return fmt.Sprintf("%.8f", str)
	} else {
		return ""
	}
}

func IChkI64(obj interface{}) int64 {
	if str, ok := obj.(int64); ok {
		return str
	} else {
		return 0
	}
}
