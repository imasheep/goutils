// if.go
package goutils

func If(condition bool, trueVal, falseVal interface{}) interface{} {

	if condition {
		return trueVal
	}

	return falseVal

}
