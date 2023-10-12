package roman

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func ToNumber(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		c := string(n[i])
		cD := num[c]
		if i < len(n)-1 {
			cNext := string(n[i+1])
			cDNext := num[cNext]
			if cD < cDNext {
				res += cDNext - cD
				i++
			} else {
				res += cD
			}
		} else {
			res += cD
		}
	}
	return res
}

func NumberToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}
