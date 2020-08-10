package skiplist

type SkipList struct {
	*DataNode
}

func (skipList *SkipList) Insert(key int64, value string) {
	// if skip list is nil
	if skipList.DataNode == nil {
		skipList.DataNode = &DataNode{

		}
		dataNode := &DataNode{
			Down:  nil,
			Next:  nil,
			Key:   key,
			Value: value,
		}
		skipList.SetNext(dataNode)

		if IsCreateLevel() {

		}
	}

	// waiting insert point
	saveWaitingPoint := make([]*DataNode, 0)
	current := skipList.DataNode
	for {
		if current == nil {
			continue
		}

		// only exist empty header
		if current.GetNext() == nil {
			saveWaitingPoint = append(saveWaitingPoint, current)
			current = current.Down
			continue
		}

		if current.GetNext().Key > key {
			saveWaitingPoint = append(saveWaitingPoint, current)
			current = current.Down
			continue
		}
		current = current.GetNext()
	}

	// insert
	top := len(saveWaitingPoint) - 1
	lowestLevelPoint := saveWaitingPoint[top]
	aimNode := &DataNode{
		Down:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	}
	aimNode.SetNext(lowestLevelPoint.GetNext())
	lowestLevelPoint.SetNext(aimNode)
	top--

	if IsCreateLevel() {
		for i := top; i >= 0; i-- {

		}
	}
}

// print structure of skip list
func (skipList *SkipList) Display() {

}
