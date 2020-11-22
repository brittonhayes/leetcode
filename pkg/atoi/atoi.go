package atoi

import (
	"math"
	"regexp"
	"strconv"
)

func StringToInt(s string) int {

	start := regexp.MustCompile(`^[\d ()+-]`)
	validStart := start.MatchString(s)
	minPlus, _ := regexp.MatchString(`(\+-|-\+|\+\s|-\s|\+\+|--)`, s)
	hasNum, _ := regexp.MatchString(`(\d)`, s)
	end, _ := regexp.MatchString(`(\S$)`, s)
	if !validStart || minPlus || !hasNum || !end {
		return 0
	}

	pattern2 := regexp.MustCompile("-?[0-9]+")
	numbers := pattern2.FindString(s)

	if v, err := strconv.Atoi(numbers); err == nil {
		if v <= math.MinInt32 {
			return math.MinInt32
		}

		if v >= math.MaxInt32 {
			return math.MaxInt32
		}

		return v
	}

	return math.MaxInt32
}
