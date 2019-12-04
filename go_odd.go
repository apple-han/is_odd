package is_odd

import (
	"errors"
	"math"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

// IsOdd 判断奇偶数
func IsOdd(value interface{}) (bool, error) {
	b, m := isInteger(value)
	if b == nil {
		return false, m
	}
	r := int(math.Abs(float64(b.(int))))
	return r%2 == 1, nil
}

// isInteger 判断是不是数字整数或者字符串整数
func isInteger(val interface{}) (interface{}, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return val, nil
	case string:
		if valid.IsInt(val.(string)) {
			v, err := strconv.Atoi(val.(string))
			if err != nil {
				return nil, err
			}
			return v, nil
		}
	}
	return nil, errors.New("请输入数字整数或者字符串整数")
}
