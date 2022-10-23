package rules

import (
	"strconv"
	"strings"
)

func parseRule(rawRule string) (rule *Rule, err error) {
	split := strings.Split(rawRule, "/")
	rule = &Rule{
		survival: make(map[int]bool),
		birth:    make(map[int]bool),
	}

	rule.survival = parseCommaSeparatedValuesWithRanges(split[0])
	rule.birth = parseCommaSeparatedValuesWithRanges(split[1])

	state, _ := strconv.Atoi(split[2])
	rule.state = state

	switch split[3] {
	case "M":
		rule.neighbourhood = Moore
	case "N":
		rule.neighbourhood = Neumann
	}

	return rule, nil
}

func parseCommaSeparatedValuesWithRanges(rawString string) map[int]bool {
	split := strings.Split(rawString, ",")
	result := make(map[int]bool)

	for _, value := range split {
		isRange := strings.Contains(value, "-")
		if isRange {
			split := strings.Split(value, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			for i := start; i <= end; i++ {
				result[i] = true
			}
		} else {
			v, _ := strconv.Atoi(value)
			result[v] = true
		}
	}

	return result
}
