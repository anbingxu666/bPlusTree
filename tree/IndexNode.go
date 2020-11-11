package tree

// 索引节点（也是中间结点，仅起到索引作用）
type IndexNode struct {
	Keys       []string     // 关键字，非叶子节点才有
	Children   []*IndexNode // 叶子结点没有children
	ParentNode *IndexNode

	IsLeaf    bool        // 是否是叶子结点
	DataNodes []*DataNode // 叶子结点才有数据结点
}

// 初始化一个node
func MallocNewIndexNode(isLeaf bool) *IndexNode {
	return &IndexNode{
		IsLeaf: isLeaf,
	}
}

// 初始化一个数据结点
func MallocNewDataNode(keyAndValue KeyAndValue) *DataNode {
	return &DataNode{
		KeyAndValue: keyAndValue}
}


/**
查找key 在 keys[] 中的位置
returns:
	previousKeyIndex: 前一个key的index
	nextKeyIndex: 后一个key的index
	如果 previousKeyIndex == nextKeyIndex && previousKeyIndex > 0 则key存在并且index为previousKeyIndex
*/
func (indexNode *IndexNode) binarySearchIndexKey(key string) (previousKeyIndex int, nextKeyIndex int) {
	if indexNode.Keys == nil || len(indexNode.Keys) <= 0 {
		return -1, -1
	}
	//fmt.Println(" 二分查找。indexkey。。")
	var low int = 0
	var height int = len(indexNode.Keys)

	for low <= height {
		//fmt.Println(" 。。。。比较。。")
		var mid int = low + (height-low)/2
		if indexNode.Keys[mid] == key {
			return mid, mid
		} else if key > indexNode.Keys[mid] { // 如果新的key 大于中间值的key,则查找右边
			if mid == height-1 || (mid < len(indexNode.Keys)-1 && key < indexNode.Keys[mid+1]) {
				return mid, mid + 1
			}
			low = mid + 1
		} else if key < indexNode.Keys[mid] { // 如果新的key 小于中间值的key,则查找左边

			if mid == 0 || (mid > 0 && key > indexNode.Keys[mid-1]) {

				return mid - 1, mid
			}
			height = mid - 1
		}
	}
	return -1, -1
}

/**
 查找key对应的indexNode 在 children[]中的位置
 returns:
		currentIndexNode: 当前索引节点
		indexAtCurrentIndexNode： key在当前索引节点的children[]中的index
*/
func (indexNode *IndexNode) binarySearchChildNode(key string) (currentIndexNode *IndexNode, indexAtCurrentIndexNode int) {

	if indexNode == nil {
		return nil, -1
	}
	if indexNode.IsLeaf {
		return indexNode, -1
	}

	//fmt.Println(" 二分查找 index node index。。。")
	var low int = 0
	var height int = len(indexNode.Keys)

	for low <= height {
		//fmt.Println(" 。。。。。比较。。")
		var mid int = low + (height-low)/2

		if indexNode.Keys[mid] == key { // 如果存在这个key

			return indexNode.Children[mid+1], mid + 1
		} else if key > indexNode.Keys[mid] { // 如果新的key 大于中间值的key,则查找右边
			if mid == len(indexNode.Keys)-1 || (mid < len(indexNode.Keys)-1 && key < indexNode.Keys[mid+1]) {

				return indexNode.Children[mid+1], mid + 1
			}
			low = mid + 1
		} else if key < indexNode.Keys[mid] { // 如果新的key 小于中间值的key,则查找左边

			if mid == 0 || (mid > 0 && key > indexNode.Keys[mid-1]) {

				return indexNode.Children[mid], mid
			}
			height = mid - 1
		}
	}

	return indexNode, -1
}

func (indexNode *IndexNode) insertKey(key string, previousIndex int) {
	// 合并Keys
	if previousIndex < 0 {
		indexNode.Keys = append([]string{key}, indexNode.Keys[:]...)
	} else {
		tempSlice := append([]string{}, indexNode.Keys[:previousIndex+1]...)
		tempSlice = append(tempSlice, key)
		tempSlice = append(tempSlice, indexNode.Keys[previousIndex+1:]...)
		indexNode.Keys = tempSlice
	}
}

func (indexNode *IndexNode) insertChild(child *IndexNode, previousIndex int) {
	child.ParentNode = indexNode
	tempChildSlice := append([]*IndexNode{}, indexNode.Children[:previousIndex+2]...)
	tempChildSlice = append(tempChildSlice, child)
	tempChildSlice = append(tempChildSlice, indexNode.Children[previousIndex+2:]...)
	indexNode.Children = tempChildSlice
}





// 根据索引删除孩子的方法
func (indexNode *IndexNode) removeChild(childIndex int) {
	tempChildSlice := append([]*IndexNode{}, indexNode.Children[:childIndex]...)
	tempChildSlice = append(tempChildSlice, indexNode.Children[childIndex+1:]...)
	indexNode.Children = tempChildSlice
}

// 根据索引删除key的方法
func (indexNode *IndexNode) removeKey(keyIndex int) {
	tempKeySlice := append([]string{}, indexNode.Keys[:keyIndex-1]...)
	tempKeySlice = append(tempKeySlice, indexNode.Keys[keyIndex:]...)
	indexNode.Keys = tempKeySlice
}