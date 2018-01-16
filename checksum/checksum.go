package checksum

func checksum(spreadsheet [][]int) int {
	var sum int

	for _, row := range spreadsheet {
		sum += highestDiff(row)

	}
	return sum

}

func highestDiff(row []int) int {
	var min, max = row[0], row[0]

	for _, v := range row[1:] {

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	return max - min
}
