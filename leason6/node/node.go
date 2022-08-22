package node

type Node struct {
	data int
	prev *Node
	next *Node
}

type ListNode struct {
	first   *Node
	last    *Node
	current *Node
	list    []*Node
}

func (nodeList *ListNode) Add(value int) {
	node := Node{value, nodeList.last, nil}
	// Инициализация первый элемент
	if len(nodeList.list) == 0 {
		nodeList.last = &node
		nodeList.first = &node
		nodeList.current = &node
	} else {
		// Добавляем следующий элемент и связываем
		nodeList.last.next = &node
		node.prev = nodeList.last
		nodeList.last = &node
	}
	data := append(nodeList.list, &node)
	nodeList.list = data
}

func (nodeList *ListNode) GetAll() []int {
	var dataList []int
	for _, data := range nodeList.list {
		dataList = append(dataList, data.data)
	}

	return dataList
}

func (nodeList *ListNode) GetAllv2() []int {
	var dataList []int
	node := nodeList.current
	for true {
		dataList = append(dataList, node.data)
		node = node.next

		if node == nil {
			return dataList
		}
	}

	return dataList
}
