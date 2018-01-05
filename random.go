//random
package goutils

import (
	srand "crypto/rand"
	"math/rand"
	"time"
)

const dictionaryString = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const dictionaryInt = "0123456789"

func CreateRandom(randomType string, randomLength int) string {

	var dictionary string
	if randomType == "string" {
		dictionary = dictionaryString
	} else {
		dictionary = dictionaryInt
	}

	b := make([]byte, randomLength)
	l := len(dictionary)

	_, err := srand.Read(b)

	if err != nil {
		// fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)

}
