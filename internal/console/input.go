package console

import "fmt"

func GetUserInput() int64 {

	var inputNumber int64

	_, _ = fmt.Scan(&inputNumber)

	return inputNumber
}
