package euler001

func MultiplesOf3And5(n int32) int64 {
	n = n - 1
	n3 := int64(n / 3)
	n5 := int64(n / 5)
	n15 := int64(n / 15)
	set3 := int64(0)
	set5 := int64(0)
	set15 := int64(0)
	if n3 > 0 {
		set3 = (3 + 3*n3) * n3 / 2
	}
	if n5 > 0 {
		set5 = (5 + 5*n5) * n5 / 2
	}

	if n15 > 0 {
		set15 = (15 + 15*n15) * n15 / 2
	}

	return (set3 + set5 - set15)
}
