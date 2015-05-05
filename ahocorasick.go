// Package rng implements an Aho-Corasick automata for
// multi-pattern string matching use. You can use it to scan
// an object content to find whether there are forbidden
// words in it.
// Author: Leesper
// Inspired by: pyAhocorasick
package goac

// a Node represents a character of pattern word
type Node struct {
	fail      *Node
	isPattern bool
	next      map[rune]*Node
}

func newNode() *Node {
	return &Node{
		fail:      nil,
		isPattern: false,
		next:      map[rune]*Node{},
	}
}

// AhoCorasick is the ac-automata
type AhoCorasick struct {
	root *Node
}

// NewAhoCorasick returns an empty ac-automata
// one can call addPattern after creating it
func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: newNode(),
	}
}

// A SearchResult represents a position of found pattern in content
type SearchResult struct {
	Start, End int
}

// AddPattern adds a pattern, making a trie tree of patterns
func (ac *AhoCorasick) AddPattern(pattern string) {
	/* build trie tree */
	chars := []rune(pattern)
	iter := ac.root
	for _, c := range chars {
		if _, ok := iter.next[c]; !ok {
			iter.next[c] = newNode()
		}
		iter = iter.next[c]
	}
	iter.isPattern = true
}

// Build the trie tree into ac-automata by adding fail states
func (ac *AhoCorasick) Build() {
	/* traverse trie tree level by level, adding fail state, make it into ac-automata */
	queue := []*Node{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		// deque
		queue = append(queue[:0], queue[1:]...)

		for char, child := range parent.next {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				if _, ok := parent.fail.next[char]; ok {
					child.fail = parent.fail.next[char]
				} else {
					child.fail = parent.fail
				}
			}
			queue = append(queue, child)
		}
	}
}

// Scan the content, return a slice of searched results
func (ac *AhoCorasick) Scan(content string) (results []SearchResult) {
	chars := []rune(content)
	iter := ac.root
	var start, end int
	for i, c := range chars {
		_, ok := iter.next[c]
		// traverse fail state
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.next[c]; ok {
			if iter == ac.root { // this is the first match, record the start position
				start = i
			}
			iter = iter.next[c]
			if iter.isPattern {
				end = i // this is the end match, record one result
				results = append(results, SearchResult{start, end})
			}
		}
	}
	return
}
