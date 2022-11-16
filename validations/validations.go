package validations

import "strconv"

func IsNewExpenseValid(name, price, category string) bool {
	if len(name) == 0 || len(price) == 0 || len(category) == 0 {
		return false
	}

	if _, err := strconv.ParseFloat(price, 64); err != nil {
		return false
	}

	return true
}
