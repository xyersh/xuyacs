package trie

type TrieI interface {
	// добавление слова в дерево
	Insert(word string)

	// поиск слова в дереве
	Search(word string) bool

	// проверка, начинается ли слово с префикса
	StartsWith(prefix string) bool

	// удаление слова из дерева
	Delete(word string) bool

	// получение всех слов с заданным префиксом
	GetAllWordsWithPrefix(prefix string) []string
}

// TrieNode — узел дерева
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode создаёт новый узел
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie — структура дерева
type Trie struct {
	root *TrieNode
}

// NewTrie создаёт новое дерево
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert добавляет слово в дерево
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search проверяет, существует ли слово целиком
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith проверяет наличие хотя бы одного слова с данным префиксом
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Delete удаляет слово из дерева (если оно существует)
// Возвращает true, если слово было удалено
func (t *Trie) Delete(word string) bool {
	if !t.Search(word) {
		return false // Слово не существует
	}
	t.deleteHelper(t.root, word, 0)
	return true
}

// deleteHelper — рекурсивная вспомогательная функция
func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		// Мы на последнем символе
		if !node.isEnd {
			return false // Не должно произойти при корректном вызове
		}
		node.isEnd = false
		// Можно ли удалить этот узел? Только если нет потомков.
		return len(node.children) == 0
	}

	ch := rune(word[index])
	child := node.children[ch]

	shouldDeleteChild := t.deleteHelper(child, word, index+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		// Удаляем текущий узел, только если он не конец другого слова и не имеет других детей
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// GetAllWordsWithPrefix возвращает все слова, начинающиеся с prefix
func (t *Trie) GetAllWordsWithPrefix(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return nil // Префикс не найден
		}
		node = node.children[ch]
	}

	var results []string
	t.collectWords(node, prefix, &results)
	return results
}

// collectWords — рекурсивный обход поддерева для сбора слов
func (t *Trie) collectWords(node *TrieNode, current string, results *[]string) {
	if node.isEnd {
		*results = append(*results, current)
	}
	for ch, child := range node.children {
		t.collectWords(child, current+string(ch), results)
	}
}
