package functions

import (
	"sort"
	"strconv"
	"strings"
)

var categories = []string{
	"General",
	"Philosophy, Psychology",
	"Religion",
	"Social Sciences",
	"Language",
	"Science, Math",
	"Applied Sciences",
	"Arts",
	"Literature",
	"Geography, History",
}

func compareCategory(a, b string) bool {
	if strings.Compare(a, b) <= 0 {
		return true
	} else {
		return false
	}
}

func SortBooks(books []string) []string {
	res := make([]string, len(books))
	copy(res, books)
	sort.SliceStable(res, func(a, b int) bool {
		ca, _ := strconv.Atoi(string(res[a][0]))
		cb, _ := strconv.Atoi(string(res[b][0]))
		if categories[ca] < categories[cb] {
			return true
		} else {
			if ca == cb {
				ha, _ := strconv.Atoi(string(res[a][2:4]))
				hb, _ := strconv.Atoi(string(res[b][2:4]))
				if ha > hb {
					return true
				}
			}
		}
		return false
	})
	return res
}
