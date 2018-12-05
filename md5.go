// md5
package goutils

import (
	"crypto/md5"
	"fmt"
)

func Md5SumStr(cont string) (md5Str string) {
	data := []byte(cont)
	has := md5.Sum(data)
	md5Str = fmt.Sprintf("%x", has)
	return
}
