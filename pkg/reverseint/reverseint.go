package reverseint

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Reverse(x int) int {

	strInt := strconv.Itoa(x)
	negative := false
	rev := ""

	for i := len(strInt); i > 0; i-- {
		val := string(strInt[i-1])
		if val == "-" {
			negative = true
		} else {
			rev += val
		}
	}

	if negative == true {
		rev = fmt.Sprintf("-%s", rev)
	}

	revNum, err := strconv.Atoi(rev)
	if err != nil {
		log.Fatal(err)
	}

	if revNum <= math.MinInt32 || revNum >= math.MaxInt32 {
		return 0
	}

	return revNum
}
