package process

import (
	"golang.org/x/exp/slices"
	"sort"
	"strings"
)

func GetTopWord(s string, k int) ([]string, error) {
	stopWords := []string{
		"a", "an", "and", "are", "as", "at", "be", "by",
		"is", "in", "of", "on", "or", "that", "the", "this",
		"to", "was", "what", "when", "where", "who", "will",
		"i", "you", "he", "she", "it", "we", "they", "my",
		"your", "his", "her", "its", "our", "their", "me",
		"", "him", "us", "them", "mine", "yours", "his",
		"hers", "ours", "theirs", "myself", "yourself",
		"himself", "herself", "itself", "ourselves",
		"themselves", "what", "which", "who", "whom",
		"whose", "this", "that", "these", "those", "am",
		"\n", "\t", " ", "!", "@", "#", "$", "%", "^", "&",
	}

	ss := strings.TrimSpace(s)
	ss = strings.ToLower(ss)
	ss = strings.Replace(ss, ".", " ", -1)
	ss = strings.Replace(ss, ",", " ", -1)
	ss = strings.Replace(ss, "!", " ", -1)
	ss = strings.Replace(ss, "?", " ", -1)
	ss = strings.Replace(ss, ":", " ", -1)
	ss = strings.Replace(ss, ";", " ", -1)
	ss = strings.Replace(ss, "(", " ", -1)
	ss = strings.Replace(ss, ")", " ", -1)
	ss = strings.Replace(ss, "[", " ", -1)
	ss = strings.Replace(ss, "]", " ", -1)
	ss = strings.Replace(ss, "{", " ", -1)
	ss = strings.Replace(ss, "}", " ", -1)
	ss = strings.Replace(ss, "'", " ", -1)
	ss = strings.Replace(ss, "\"", " ", -1)
	ss = strings.Replace(ss, "\\", " ", -1)
	ss = strings.Replace(ss, "/", " ", -1)
	ss = strings.Replace(ss, "-", " ", -1)
	ss = strings.Replace(ss, "_", " ", -1)
	ss = strings.Replace(ss, "=", " ", -1)
	ss = strings.Replace(ss, "+", " ", -1)
	ss = strings.Replace(ss, "\n", " ", -1)
	ss = strings.Replace(ss, "\t", " ", -1)

	type wordCount struct {
		word  string
		count int
	}
	result := make(map[string]int)
	for _, word := range strings.Split(ss, " ") {
		if word != " " && !slices.Contains(stopWords, word) {
			result[word]++
		}
	}

	var wc []wordCount
	for k, v := range result {
		wc = append(wc, wordCount{k, v})
	}
	sort.Slice(wc, func(i, j int) bool {
		return wc[i].count > wc[j].count
	})

	var top []string
	if k > len(wc) {
		k = len(wc)
	}
	for i := 0; i < k; i++ {
		top = append(top, wc[i].word)
	}
	return top, nil
}
