package bptree

type leafNode interface {
	binarySearchDataNode(indexNode *IndexNode, key string) (currentIndexNode *IndexNode, previousLeafIndexAtCurrentIndexNode int, nextLeafIndexAtCurrentIndexNode int)
	insertDataNode(dataNode *DataNode, previousIndex int)
	removeDataNode(dataNodeIndex int)
}

// 叶子的数据节点
type DataNode struct {
	KeyAndValue      KeyAndValue
	ParentNode       *IndexNode
	PreviousDataNode *DataNode
	NextDataNode     *DataNode
}

type KeyAndValue struct {
	Key   string // key 关键字
	Value string // value是对应数据在levelDB中的id、地址或者nil
}

func (indexNode *IndexNode) insertDataNode(newDataNode *DataNode, previousDataNodeIndex int) {
	newDataNode.ParentNode = indexNode
	// 新增节点的Key小于当前组的所有值
	if previousDataNodeIndex < 0 {

		// 将数据结点按顺序左右链接起来
		newDataNode.PreviousDataNode = indexNode.DataNodes[0].PreviousDataNode
		newDataNode.NextDataNode = indexNode.DataNodes[0]
		indexNode.DataNodes[0].PreviousDataNode = newDataNode

		// 合并到DataNodes[]中
		indexNode.DataNodes = append([]*DataNode{newDataNode}, indexNode.DataNodes[:]...)

	} else {

		// 将叶子结点按顺序左右链接起来
		newDataNode.PreviousDataNode = indexNode.DataNodes[previousDataNodeIndex]
		newDataNode.NextDataNode = indexNode.DataNodes[previousDataNodeIndex].NextDataNode

		// 合并到DataNodes[]中
		tempSlice := append([]*DataNode{}, indexNode.DataNodes[:previousDataNodeIndex+1]...)
		tempSlice = append(tempSlice, newDataNode)
		tempSlice = append(tempSlice, indexNode.DataNodes[previousDataNodeIndex+1:]...)
		indexNode.DataNodes = tempSlice
	}

	if newDataNode.PreviousDataNode != nil {
		newDataNode.PreviousDataNode.NextDataNode = newDataNode
	}

	if newDataNode.NextDataNode != nil {
		newDataNode.NextDataNode.PreviousDataNode = newDataNode
	}

}

func (indexNode *IndexNode) removeDataNode(removeDataNode *DataNode, dataNodeIndex int) {
	// 修改被删除dataNode的前后dataNode的左右连接关系
	if removeDataNode.PreviousDataNode != nil {
		removeDataNode.PreviousDataNode.NextDataNode = removeDataNode.NextDataNode
	}
	if removeDataNode.NextDataNode != nil {
		removeDataNode.NextDataNode.PreviousDataNode = removeDataNode.PreviousDataNode
	}
	// 删除节点
	tempSlice := append([]*DataNode{}, indexNode.DataNodes[:dataNodeIndex]...)
	tempSlice = append(tempSlice, indexNode.DataNodes[dataNodeIndex+1:]...)
	indexNode.DataNodes = tempSlice
}

/** 二分查找 定位key对应的leafNode在indexNode节点的DataNodes[]中的位置
  returns:
       currentIndexNode: 当前叶子结点
       previousLeafIndexAtCurrentIndexNode: 前一个leafNode的index
       nextLeafIndexAtCurrentIndexNode: 后一个leafNode的index
	如果 previousLeafIndexAtCurrentIndexNode == nextLeafIndexAtCurrentIndexNode
			&& previousLeafIndexAtCurrentIndexNode > 0 则key对应的leafNode存在，并且index为previousLeafIndexAtCurrentIndexNode
**/
func (indexNode *IndexNode) binarySearchDataNode(key string) (currentIndexNode *IndexNode, previousLeafIndexAtCurrentIndexNode int, nextLeafIndexAtCurrentIndexNode int) {
	if indexNode == nil || len(indexNode.DataNodes) <= 0 {
		return indexNode, -1, -1
	}
	//fmt.Println(" 二分查找 leafkey。。。")

	DataNodes := indexNode.DataNodes
	var low int = 0
	var height int = len(indexNode.DataNodes)

	for low <= height {
		//fmt.Println(" 。。。。比较。。")
		var mid int = low + (height-low)/2

		if DataNodes[mid].KeyAndValue.Key == key { // 如果存在这个key
			// 如果是叶子结点
			return indexNode, mid, mid

		} else if key > DataNodes[mid].KeyAndValue.Key { // 如果新的key 大于中间值的key,则查找右边
			if mid == len(DataNodes)-1 || (mid < len(DataNodes)-1 && key < DataNodes[mid+1].KeyAndValue.Key) {
				return indexNode, mid, mid + 1
			}

			low = mid + 1
		} else if key < DataNodes[mid].KeyAndValue.Key { // 如果新的key 小于中间值的key,则查找左边

			if mid == 0 || (mid > 0 && key > DataNodes[mid-1].KeyAndValue.Key) {

				return indexNode, mid - 1, mid

			}
			height = mid - 1
		}
	}
	return nil, -1, -1
}
