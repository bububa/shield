package shield

import (
	"regexp"
	"strings"
)

type enTokenizer struct {
}

func NewEnglishTokenizer() Tokenizer {
	return &enTokenizer{}
}

func split(re *regexp.Regexp, s string, n int) []string {

    if n == 0 {
        return nil
    }

    if len(s) == 0 {
        return []string{""}
    }

    matches := re.FindAllStringIndex(s, n)
    strings := make([]string, 0, len(matches))

    beg := 0
    end := 0
    for _, match := range matches {
        if n > 0 && len(strings) >= n-1 {
            break
        }

        end = match[0]
        if match[1] != 0 {
            strings = append(strings, s[beg:end])
        }
        beg = match[1]
    }

    if end != len(s) {
        strings = append(strings, s[beg:])
    }

    return strings
}

func (t *enTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	for _, w := range split(splitTokenRx, text, -1) {
		if len(w) > 2 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

// Spamassassin stoplist
//
// http://wiki.apache.org/spamassassin/BayesStopList
//
var splitTokenRx = regexp.MustCompile(`[^\w]+|able|all|already|and|any|are|because|both|can|come|each|email|even|few|first|for|from|give|has|have|http|information|into|it's|just|know|like|long|look|made|mail|mailing|mailto|make|many|more|most|much|need|not|now|number|off|one|only|out|own|people|place|right|same|see|such|that|the|this|through|time|using|web|where|why|with|without|work|world|year|years|you|you're|your`)
