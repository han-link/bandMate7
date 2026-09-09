package utils

import "github.com/google/uuid"

func ContainsDuplicates(arr []uuid.UUID) bool {
	record := make(map[any]int)
	for _, el := range arr {
		if _, ok := record[el]; !ok {
			record[el] = 1
		} else {
			return true
		}
	}
	return false
}
