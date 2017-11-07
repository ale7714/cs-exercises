package trie

import (
  "fmt"
)

type Trie struct {
  children map[string]*Trie
  isWord bool
}

func NewTrie() *Trie {
  return &Trie{isWord: false, children: make(map[string]*Trie)}

}

func (t *Trie) Insert(word string) {
  children := t.children
  for index, char := range word {
    if(children[string(char)] == nil) {
      children[string(char)] = &Trie{isWord: false, children: make(map[string]*Trie)}
    }

    if(len(word) == index + 1) {
      children[string(char)].isWord = true
    }
    children = children[string(char)].children
  }
}


func (t *Trie) Search(word string) (found bool) {
  return t.searchHelper(word, true)
}

func (t *Trie) StartsWith(word string) (found bool) {
  return t.searchHelper(word, false)
}

func (t *Trie) searchHelper(word string, wordRequired bool) (found bool) {
  children := t.children
  found = false
  for index, char := range word {
    if(children[string(char)] == nil) {
      break
    } else {
      if(len(word) == index + 1 && requiresWord(wordRequired,children[string(char)].isWord)) {
        found = true
      }
    }
    children = children[string(char)].children
  }
  return found
}

func requiresWord(condition bool, isWord bool) bool {
  if(condition == false) {
    return true
  }
  return condition && isWord
}
