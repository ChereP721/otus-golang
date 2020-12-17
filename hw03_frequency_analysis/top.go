package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

const maxResultCount = 10

type wordRepeat struct {
	word  string
	count int
}

func Top10(str string) []string {
	str = strings.ToLower(str)
	re := regexp.MustCompile(`[a-zA-Zа-яА-Я-ёЁ]+`)

	strWordList := re.FindAllString(str, -1)

	countMap := make(map[string]int)
	for _, word := range strWordList {
		countMap[word]++
	}
	delete(countMap, "-")

	topCountList := []wordRepeat{}
	for word, count := range countMap {
		if count == 1 { // слова встречающиеся по 1 разу нельзя считать чатсотными т.к. меньше 1 раза слово встречаться не может
			continue
		}
		topCountList = append(topCountList, wordRepeat{word: word, count: count})
	}
	sort.Slice(topCountList, func(i, j int) bool {
		return topCountList[i].count > topCountList[j].count
	})

	resultCount := len(topCountList)
	if resultCount > maxResultCount {
		resultCount = maxResultCount
	}
	topStringList := make([]string, resultCount)
	for i := 0; i < resultCount; i++ {
		topStringList[i] = topCountList[i].word
	}

	return topStringList
}
