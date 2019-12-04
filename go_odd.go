package is_odd

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"math"
	"strconv"
)

// IsOdd 判断奇偶数
func IsOdd(value interface{}) (bool, error) {
	b, m := isInteger(value)
	if !b {
		return false, errors.New("请输入数字整数或者字符串整数")
	}
	r := int(math.Abs(float64(m.(int))))
	return r%2 == 1, nil
}

// isInteger 判断是不是数字整数或者字符串整数
func isInteger(val interface{}) (bool, interface{}) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true, val
	case string:
		if valid.IsInt(val.(string)) {
			if v, err := strconv.Atoi(val.(string)); err == nil {
				return true, v
			}
		}
	}
	return false, nil
}
