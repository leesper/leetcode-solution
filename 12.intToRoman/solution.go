package leetcode12

import "bytes"

// Symbol    Value
// I          1
// V          5
// X          10
// L          50
// C          100
// D          500
// M          1,000
// 输入：一个数字num
// 输出：一个字符串roman
// 条件：
// （1）roman是num的罗马数字串
// （2）1 <= num <= 3999
// （3）不会出现多于3个相同罗马字母进行表示的情况，进一步展开：
// （3-1）4不表示为IIII而表示为IV；9不表示为VIIII而表示为IX
// （3-2）40不表示为XXXX而表示为XL；90不表示为LXXXX而表示为XC
// （3-3）400不表示为CCCC而表示为CD；900不表示为DCCCC而表示为CM
// 因为题目给出的取值范围最大到3999，所以4000和9000不予考虑
func Solution(number int) string {
	// convert the number to a roman numeral
	specials := map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}

	numerals := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	var appender bytes.Buffer
	kilos, hundreds, decades, units := decompose(number)

	appender.WriteString(parse(kilos, 4000, 1000, numerals, specials))
	appender.WriteString(parse(hundreds, 500, 100, numerals, specials))
	appender.WriteString(parse(decades, 50, 10, numerals, specials))
	appender.WriteString(parse(units, 5, 1, numerals, specials))

	return appender.String()
}

func parse(num, upperCutOff, lowerCutOff int, numerals, specials map[int]string) string {
	if _, ok := specials[num]; ok {
		return specials[num]
	}

	if _, ok := numerals[num]; ok {
		return numerals[num]
	}

	var romanAppender bytes.Buffer
	for num > upperCutOff {
		num -= upperCutOff
		romanAppender.WriteString(numerals[upperCutOff])
	}

	for num >= lowerCutOff {
		num -= lowerCutOff
		romanAppender.WriteString(numerals[lowerCutOff])
	}

	return romanAppender.String()
}

func decompose(number int) (kilos, hundreds, decades, units int) {
	kilos = number / 1000 * 1000
	number -= kilos
	hundreds = number / 100 * 100
	number -= hundreds
	decades = number / 10 * 10
	number -= decades
	units = number
	return
}
