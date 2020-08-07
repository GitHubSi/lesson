package skiplist

// skip list header node
type HeaderNode struct {
	Down *HeaderNode
	Next *DataNode
}

func (header *HeaderNode) GetDown() *HeaderNode {
	return header.Down
}

func (header *HeaderNode) SetDown(node *HeaderNode) {
	header.Down = node.Down
}

func (header *HeaderNode) GetNext() *DataNode {
	return header.Next
}

func (header *HeaderNode) SetNext(node *DataNode) {
	header.Next = node
}

// skip list data node
type DataNode struct {
	Down  *DataNode
	Next  *DataNode
	Key   int64
	Value string
}

func (data *DataNode) GetDown() *DataNode {
	return data.Down
}

func (data *DataNode) SetDown(node *DataNode) {
	data.Down = node
}

func (data *DataNode) GetNext() *DataNode {
	return data.Next
}

func (data *DataNode) SetNext(node *DataNode) {
	data.Next = node
}
