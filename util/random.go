package util

import (
	"math/rand"
	"strings"
	"time"
)
const alphabets = "abcdefghijklmnopqrstuvwxyz"
// automatically called when package is first used
func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInteger(min,max int64) int64{
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string{
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInteger(0,1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}