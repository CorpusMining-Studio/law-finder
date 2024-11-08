package router

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Map of Chinese numerals to integer values
var chineseNumeralMap = map[rune]int{
	'零': 0, '一': 1, '二': 2, '三': 3, '四': 4,
	'五': 5, '六': 6, '七': 7, '八': 8, '九': 9,

	'兩': 2, // '兩' is also used for '2' in some contexts
}

// Map of Chinese multipliers to integer values
var chineseMultiplierMap = map[rune]int{
	'十': 10, '百': 100, '千': 1000,
}

func normalizeArticle(input string) (int, error) {
	// Remove non-essential characters and whitespace
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "第", "")
	input = strings.ReplaceAll(input, "條", "")
	input = strings.ReplaceAll(input, " ", "")

	// Try to parse the input directly as an integer
	if num, err := strconv.Atoi(input); err == nil {
		return num, nil
	}

	result, err := chineseNumeralToInt(input)
	if err != nil {
		return 0, fmt.Errorf("unable to parse article number from input: %s", input)
	}
	log.Printf("result: %d", result)
	return result, nil
}

// chineseNumeralToInt converts a Chinese numerical representation up to the hundreds digit to an integer
func chineseNumeralToInt(chineseNumeral string) (int, error) {
	result := 0
	temp := 0

	for _, char := range chineseNumeral {
		if value, ok := chineseNumeralMap[char]; ok {
			temp = value
		} else if multiplier, ok := chineseMultiplierMap[char]; ok {
			// If the multiplier is '十' and temp is 0, treat it as '一十' (i.e., 10)
			if temp == 0 && char == '十' {
				temp = 1
			}
			result += temp * multiplier
			temp = 0
		} else {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}

	// Add any remaining value
	result += temp

	return result, nil
}
