// convet.go
package goutils

import (
	"errors"
	"fmt"
	"strconv"
)

func ArrIntToString(arr []int) (result string, err error) {

	result = ""

	for indx, a := range arr {
		if indx == 0 {
			result = fmt.Sprintf("%v", a)
		} else {
			result = fmt.Sprintf("%v,%v", result, a)
		}
	}

	if result == "" {
		err = errors.New(fmt.Sprintf("array is empty, err: %v", arr))
		return
	}

	return

}

func ArrIntToStringMust(arr []int) (result string) {

	result, _ = ArrIntToString(arr)

	return

}

func ArrInt64ToString(arr []int64) (result string, err error) {

	result = ""

	for indx, a := range arr {
		if indx == 0 {
			result = fmt.Sprintf("%v", a)
		} else {
			result = fmt.Sprintf("%v,%v", result, a)
		}
	}

	if result == "" {
		err = errors.New(fmt.Sprintf("array is empty, err: %v", arr))
	}

	return

}

func ArrInt64ToStringMust(arr []int64) (result string) {

	result, _ = ArrInt64ToString(arr)

	return

}

func ArrStringsToString(arr []string) (result string, err error) {

	result = ""

	for indx, a := range arr {
		if indx == 0 {
			result = fmt.Sprintf("\"%v\"", a)
		} else {
			result = fmt.Sprintf("%v,\"%v\"", result, a)
		}
	}

	if result == "" {
		err = errors.New(fmt.Sprintf("array is empty, err: %v", arr))
	}

	return

}

func ArrStringsToStringMust(arr []string) (result string) {

	result, _ = ArrStringsToString(arr)

	return

}

func ArrStringsToArrInt(arrStr []string) (arrInt []int) {

	arrInt = []int{}

	for _, str := range arrStr {
		i, _ := strconv.Atoi(str)
		arrInt = append(arrInt, i)
	}

	return arrInt

}

func ArrIntToArrString(arrInt []int) (arrStr []string) {

	arrStr = []string{}

	for _, i := range arrInt {
		str := IntToString(i)
		arrStr = append(arrStr, str)
	}

	return arrStr

}

func StringToInt(str string) (err error, result int) {

	result, err = strconv.Atoi(str)

	return

}
func StringToIntMust(str string) (result int) {

	result, _ = strconv.Atoi(str)

	return

}

func IntToString(intNum int) (result string) {

	result = strconv.Itoa(intNum)

	return

}

func Float64ToString(f float64) (result string) {

	result = strconv.FormatFloat(f, 'f', -1, 64)

	return

}
