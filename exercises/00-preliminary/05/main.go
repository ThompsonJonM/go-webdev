package main

import (
	"fmt"
	"log"
)

func main() {
	s := setTime(0, 31)
	fmt.Println(s)
}

func setTime(h, m int) string {
	var s string
	t := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21: "twenty-one",
		22: "twenty-two",
		23: "twenty-three",
		24: "twenty-four",
		25: "twenty-five",
		26: "twenty-six",
		27: "twenty-seven",
		28: "twenty-eight",
		29: "twenty-nine",
		30: "thirty",
	}

	if h <= 12 {
		if h == 0 {
			log.Fatalln("Time cannot be 0")
		}

		if m == 0 {
			s = fmt.Sprintf("%s o'clock", t[h])
		} else if m == 1 {
			s = fmt.Sprintf("%s minute past %s", t[m], t[h])
		} else if m < 30 {
			if m == 15 {
				s = fmt.Sprintf("quarter past %s", t[h])
			} else {
				s = fmt.Sprintf("%s minutes past %s", t[m], t[h])
			}
		} else if m == 30 {
			s = fmt.Sprintf("half past %s", t[h])
		} else if m > 30 {
			if m == 59 {
				s = fmt.Sprintf("one minute to %s", t[h])
			} else if m == 45 {
				s = fmt.Sprintf("quarter to %s", t[h])
			} else {
				if h == 12 {
					s = fmt.Sprintf("%s minutes to %s", t[60-m], t[1])
				} else {
					s = fmt.Sprintf("%s minutes to %s", t[60-m], t[h+1])
				}
			}
		}
	} else {
		log.Fatalln("Time cannot be more than 12")
	}

	return s
}
