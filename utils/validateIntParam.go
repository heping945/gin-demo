package utils

import (
	"reflect"
	"strconv"
)

// 接收一个不为nil的interface{}，是否可以转化为int，输出为int和bool
func transToint(arg interface{}) (int, bool) {
	if reflect.TypeOf(arg).Kind() == reflect.Int {
		result := arg.(int)
		return result, true
	}
	return -1, false
}

func ValidateParam(param string, rules ...interface{}) (int, bool) {
	length := len(rules)
	if length > 2 {
		panic("参数限制少于两个，可选nil和整型数字类型")
	}

	if res, err := strconv.Atoi(param); err != nil { //参数不可转化为整数
		return 0, false
	} else {
		// 参数可以转化为整数
		if length == 2 { //规则长度为2
			min := rules[0]
			max := rules[1]
			switch {
			case min == nil && max == nil: // 两个nil
				return res, true
			case min == nil: // min为nil另一个不是nil的max,如果可以转化就要小于，不然一律通过
				if result, ok := transToint(max); ok {
					return res, res < result
				} else {
					return res, true
				}
			case max == nil: // max为nil，min不是
				if result, ok := transToint(min); ok {
					return res, res > result
				} else {
					return res, true
				}
			case min != nil || max != nil:
				m1, o1 := transToint(min)
				m2, o2 := transToint(max)
				if o1 && o2 {
					if m1 > m2 {
						return res, false
					} else {
						return res, m1 < res && res < m2
					}
				} else if o1 {
					return res, m1 < res
				} else if o2 {
					return res, res < m2
				} else {
					return res, false
				}
			default:
				return res, false

			}
		} else if length == 1 {
			min := rules[0]
			result, ok := transToint(min)
			if ok {
				return res, res > result
			} else {
				return res, false
			}
		} else {
			return res, true
		}
	}
}
