package shield

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	tokenizer := shield.NewChineseTokenizer("/var/code/go/src/github.com/huichen/sego/data/dictionary.txt", true)
	text := "高级版如果选择超过5款宝贝的提示信息要更改"
	m := tokenizer.Tokenize(text)
	x := fmt.Sprintf("%v", m)
	if x != `map[款:1 宝贝:1 更改:1 高级:1 提示:1 选择:1 信息:1 提示信息:1 版:1]` {
		t.Fatal(x)
	}
}
