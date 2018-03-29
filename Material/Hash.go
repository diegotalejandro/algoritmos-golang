package chained

type node struct {
	key   int
	value string
	next  *node
}

type table struct {
	array []*node
}

func New(n int) *table {
	return &table{array: make([]*node, n)}
}

func hash(v, n int) int {
	return (v * 1240559) % n
}

func (t *table) insert(key int, value string) {
	h := hash(key, len(t.array))
	t.array[h] = &node{
		key:   key,
		value: value,
		next:  t.array[h],
	}
}

func (t *table) search(key int) (string, bool) {
	h := hash(key, len(t.array))
	aux := t.array[h]
	for aux != nil {
		if aux.key == key {
			return aux.value, true
		}
		aux = aux.next
	}
	return "", false
}
