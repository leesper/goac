package goac

import (
	"fmt"
	"testing"
)

func TestAhoCorasick(t *testing.T) {
	content := "Aho-Corasick是一种基于Trie树的确定性有穷状态自动机算法"
	ac := NewAhoCorasick()
	ac.AddPattern("树")
	ac.AddPattern("自动机")
	ac.AddPattern("Trie")
	ac.AddPattern("基于")
	ac.Build()
	results := ac.Scan(content)
	fmt.Println("内容: " + content)
	fmt.Println("匹配词: ")
	for _, result := range results {
		fmt.Println(string([]rune(content)[result.Start : result.End+1]))
	}
}
