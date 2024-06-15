package main

import (
	"errors"
	"fmt"
	"math"

	"example.com/lib50"
)

func luhn(card_num int64) bool {
	var sum int
	var digit int

	even_index := false
	for card_num > 0 {
		digit = int(card_num) % 10
		if even_index {
			digit *= 2
			if digit >= 10 {
				digit = (digit / 10) + (digit % 10)
			}
		}
		sum += digit

		// Prepare for next iter
		card_num = card_num / 10
		even_index = !even_index
	}
	return sum%10 == 0
}

func lenCardNum(num int64) (int, error) {
	if num < 0 {
		return -1, errors.New("lenCardNum requires a positive integer")
	}

	var len int
	for num > 0 {
		num /= 10
		len++
	}
	return len, nil
}

func categorize(card_num int64) string {
	len, err := lenCardNum(card_num)
	if err != nil {
		return "INVALID"
	}
	switch len {
	case 13:
		if card_num/int64(math.Pow(10, 12)) == 4 {
			return "VISA"
		}
	case 15:
		first_digits := card_num / int64(math.Pow(10, 13))
		if first_digits == 34 || first_digits == 37 {
			return "AMEX"
		}
	case 16:

		first_digits := card_num / int64(math.Pow(10, 14))
		if first_digits >= 51 && first_digits <= 55 {
			return "MASTERCARD"
		}
		if first_digits/10 == 4 {
			return "VISA"
		}
	default:
		return "INVALID"
	}
	return "INVALID"
}

func credit(card_num int64) string {

	if !luhn(card_num) {
		return "INVALID"
	}
	return categorize(card_num)
}

func main() {
	var card_num int64
	for {
		card_num = lib50.GetInt64("Number: ")
		if card_num >= 0 {
			fmt.Println(credit(card_num))
			return
		}
	}
}
