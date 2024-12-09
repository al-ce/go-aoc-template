package day05

import (
	"fmt"
	"slices"

	h "go-aoc-template/internal/helpers"
)

// Separate input into the page ordering rules and the updates
func separateInput(lines []string) ([]string, []string) {
	idx := slices.Index(lines, "")
	return lines[:idx], lines[idx+1:]
}

// Make a hashmap with page ordering rules.
// Keys are page, values are int slices of pages that must follow the key
func makeRules(rulesInput []string) map[int][]int {
	rules := map[int][]int{}
	for _, rule := range rulesInput {
		split := h.ParseIntString(rule, "|")
		key, value := split[0], split[1]
		if rule, exists := rules[key]; exists {
			rules[key] = append(rule, value)
		} else {
			rules[key] = []int{value}
		}
	}
	return rules
}

func checkOrder(rules map[int][]int, update []int) bool {
	for i := len(update) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !slices.Contains(rules[update[j]], update[i]) {
				return false
			}
		}
	}
	return true
}

func orderUpdate(rules map[int][]int, update []int) {
	for i := 0; i < len(update); i++ {
		for p, q := len(update)-2, len(update)-1; p >= i; p, q = p-1, q-1 {
			pval := update[p]
			if rule, exists := rules[pval]; !exists || !slices.Contains(rule, update[q]) {
				h.Swap(update, p, q)
			}
		}
	}
}

func PartOne(lines []string) string {
	rulesInput, updates := separateInput(lines)
	rules := makeRules(rulesInput)
	sum := 0
	for _, u := range updates {
		update := h.ParseIntString(u, ",")
		if checkOrder(rules, update) {
			sum += update[len(update)/2]
		}

	}
	return fmt.Sprintf("%d", sum)
}

func PartTwo(lines []string) string {
	rulesInput, updates := separateInput(lines)
	rules := makeRules(rulesInput)
	sum := 0
	for _, u := range updates {
		update := h.ParseIntString(u, ",")
		original := make([]int, len(update))
		copy(original, update)
		orderUpdate(rules, update)
		if !slices.Equal(original, update) {
			sum += update[len(update)/2]
		}
	}
	return fmt.Sprintf("%d", sum)
}
