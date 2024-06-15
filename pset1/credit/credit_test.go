package main

import (
	"regexp"
	"testing"
)

func TestCredit(t *testing.T) {
	tests := map[string]struct {
		input   int64
		result  string
		message string
	}{
		"test1": {
			input:   378282246310005,
			result:  "AMEX",
			message: "identifies 378282246310005 as AMEX",
		},
		"test2": {
			input:   371449635398431,
			result:  "AMEX",
			message: "identifies 371449635398431 as AMEX",
		},
		"test3": {
			input:   5555555555554444,
			result:  "MASTERCARD",
			message: "identifies 5555555555554444 as MASTERCARD",
		},
		"test4": {
			input:   5105105105105100,
			result:  "MASTERCARD",
			message: "identifies 5105105105105100 as MASTERCARD",
		},
		"test5": {
			input:   4111111111111111,
			result:  "VISA",
			message: "identifies 4111111111111111 as VISA",
		},
		"test6": {
			input:   4012888888881881,
			result:  "VISA",
			message: "identifies 4012888888881881 as VISA",
		},
		"test7": {
			input:   4222222222222,
			result:  "VISA",
			message: "identifies 4222222222222 as VISA",
		},
		"test8": {
			input:   1234567890,
			result:  "INVALID",
			message: "identifies 1234567890 as INVALID (invalid length, checksum, identifying digits)",
		},
		"test9": {
			input:   369421438430814,
			result:  "INVALID",
			message: "identifies 369421438430814 as INVALID (invalid identifying digits)",
		},
		"test10": {
			input:   4062901840,
			result:  "INVALID",
			message: "identifies 4062901840 as INVALID (invalid length)",
		},
		"test11": {
			input:   5673598276138003,
			result:  "INVALID",
			message: "identifies 5673598276138003 as INVALID (invalid identifying digits)",
		},
		"test12": {
			input:   4111111111111113,
			result:  "INVALID",
			message: "identifies 4111111111111113 as INVALID (invalid checksum)",
		},
		"test13": {
			input:   4222222222223,
			result:  "INVALID",
			message: "identifies 4222222222223 as INVALID (invalid checksum)",
		},
		"test14": {
			input:   3400000000000620,
			result:  "INVALID",
			message: "identifies 3400000000000620 as INVALID (AMEX identifying digits, VISA/Mastercard length)",
		},
		"test15": {
			input:   430000000000000,
			result:  "INVALID",
			message: "identifies 430000000000000 as INVALID (VISA identifying digits, AMEX length)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			match, err := regexp.MatchString(credit(test.input), test.result)
			if err != nil {
				t.Errorf(test.message)
			}
			if match {
				t.Logf(test.message)
			} else {
				t.Errorf(test.message)
			}
		})
	}
}
