package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}
	for i := 1; i <= nb/2; i++ {
		if i*i == nb {
			return i
		}
		if i*i > nb {
			break
		}
	}
	return 0
}
