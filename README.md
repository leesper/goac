goac --- Aho-Corasick多模式字符串匹配算法Go语言实现
===========================================

### An Aho-Corasick multi-pattern string matching lib written in Golang

### Author:
[Leesper](pascal7718@gmail.com)

Inspired by:
[pyAhocorasick](https://github.com/metadata1984/pyAhocorasick)

Usage Example
----------------------

	package goac
	import "goac"

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

	[Output]
	内容: Aho-Corasick是一种基于Trie树的确定性有穷状态自动机算法
	匹配词: 
	基于
	Trie
	树
	自动机
	PASS
	ok  	goac	0.002s
