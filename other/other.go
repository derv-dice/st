package other

import (
	"github.com/google/uuid"
	"strings"
)

// UniqueSliceString - delete non unique elements from slice
func UniqueSliceString(arr []string) (result []string) {
	u := map[string]bool{}
	for i := range arr {
		if !u[arr[i]] {
			result = append(result, arr[i])
			u[arr[i]] = true
		}
	}
	return
}

// UniqueSliceInt - delete non unique elements from slice
func UniqueSliceInt(arr []int) (result []int) {
	u := map[int]bool{}
	for i := range arr {
		if !u[arr[i]] {
			result = append(result, arr[i])
			u[arr[i]] = true
		}
	}
	return
}

// UUIDv4 - new uuid v4
func UUIDv4(hyphens bool) string {
	if hyphens {
		return uuid.New().String()
	}
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
