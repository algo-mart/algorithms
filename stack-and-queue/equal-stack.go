package main

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	sum1, sum2, sum3 := int32(0), int32(0), int32(0)
	i := 0
	for i < len(h1) {
		sum1 += h1[i]
		i++
	}
	i = 0
	for i < len(h2) {
		sum2 += h2[i]
		i++
	}
	i = 0
	for i < len(h3) {
		sum3 += h3[i]
		i++
	}
	for len(h1) > 0 && len(h2) > 0 && len(h3) > 0 {
		if sum1 == sum2 && sum2 == sum3 {
			return sum2
		}
		last1 := h1[0]
		last2 := h2[0]
		last3 := h3[0]
		if sum1 < sum2 && sum1 < sum3 {
			h2 = append(h2[1:])

			sum2 -= last2
			h3 = append(h3[1:])
			sum3 -= last3
		} else if sum2 < sum1 && sum2 < sum3 {
			h1 = append(h1[1:])
			sum1 -= last1
			h3 = append(h3[1:])
			sum3 -= last3

		} else if sum3 < sum1 && sum3 < sum2 {
			h1 = append(h1[1:])
			sum1 -= last1
			h2 = append(h2[1:])
			sum2 -= last2
		} else if sum1 < sum2 && sum1 == sum3 {
			h2 = append(h2[1:])
			sum2 -= last2
		} else if sum2 < sum1 && sum2 == sum3 {
			h1 = append(h1[1:])
			sum1 -= last1
		} else if sum2 < sum3 && sum2 == sum1 {
			h3 = append(h3[1:])
			sum3 -= last3
		}

	}
	return 0
}
func main() {

}
