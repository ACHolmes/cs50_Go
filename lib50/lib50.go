package lib50

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// func Getint(format string, a ...any) int {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Printf(format, a...)
// 		ok := scanner.Scan()

// 		if !ok {
// 			err := scanner.Err()
// 			if err != nil {
// 				log.Fatal("GetInt err,", err)
// 			}
// 		}

// 		i64, err := strconv.ParseInt(scanner.Text(), 10, 0)
// 		i := int(i64)
// 		if err != nil {
// 			continue
// 		}
// 		return i
// 	}
// }

func GetInt(format string, a ...any) int {
	for {
		str := GetString(format, a...)
		i64, err := strconv.ParseInt(str, 10, 0)
		i := int(i64)
		if err == nil {
			return i
		}
	}
}

func GetInt64(format string, a ...any) int64 {
	for {
		str := GetString(format, a...)
		i64, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return i64
		}
	}
}

func GetString(format string, a ...any) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(format, a...)
		ok := scanner.Scan()

		if !ok {
			err := scanner.Err()
			if err != nil {
				log.Fatal("Getstring err,", err)
			}
		}

		text := scanner.Text()
		if len(text) > 0 {
			return text
		}
	}
}

// func GetString(format string, a ...any) string {
// 	var line string

// 	for {
// 		fmt.Printf(format, a...)
// 		_, err := fmt.Scanln(&line)
// 		if err != nil {
// 			continue
// 		}
// 		return line
// 	}
// }
