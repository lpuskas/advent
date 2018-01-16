package captcha

/**
http://adventofcode.com/2017/day/1
*/
var (
	input string
)

func getSum(series string) int {
	var sum int
	chars := []rune(series)

	if chars[0] == chars[len(chars)-1] {
		sum = int(chars[0] - '0')
	}

	for i, k := range chars {

		if i > 0 && chars[i] == chars[i-1] {
			sum += int(k - '0')
		}
	}

	return sum
}
