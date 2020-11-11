package tree

import "fmt"

// 叶子的数据节点
type DataNode struct {
	KeyAndValue      KeyAndValue
	ParentNode       *IndexNode
	PreviousDataNode *DataNode
	NextDataNode     *DataNode
}
// 真正存放数据的结构体
type KeyAndValue struct {
	Key   string // key 关键字
	Value string // value是对应数据在levelDB中的id、地址或者nil
}

// 插入叶子结点的方法
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

// 二分查找
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


func (t *BPTree) Traversal() {
	fmt.Println()
	p := t.Head
	// 遍历
	for p != nil {
		if p.ParentNode.ParentNode == nil {
			fmt.Printf("key %s: value %s  \n", p.KeyAndValue.Key, p.KeyAndValue.Value)
		} else {

			fmt.Printf("key %s: value %s , parent keys:%s \n", p.KeyAndValue.Key, p.KeyAndValue.Value, p.ParentNode.ParentNode.Keys)
		}
		p = p.NextDataNode
	}
	fmt.Println()
}

func (t *BPTree) UpToDownPrint() {
	p := t.Root

	if p != nil {
		if p.IsLeaf {
			// 打印 DataNode
			for _, dataNode := range p.DataNodes {
				fmt.Printf("%s ", dataNode.KeyAndValue.Key)
			}
			fmt.Println()
		} else {
			fmt.Println(p.Keys)
			var tempArray []*IndexNode
			// 打印child的
			for _, child := range p.Children {
				if child.IsLeaf {
					for _, dataNode := range child.DataNodes {
						fmt.Printf("%s ", dataNode.KeyAndValue.Key)
					}
					fmt.Print("|")
				} else {
					tempArray = append(tempArray, child)
				}
			}

			for len(tempArray) > 0 {
				var newTempArray []*IndexNode
				for _, node := range tempArray {
					if node.IsLeaf {
						for _, dataNode := range node.DataNodes {
							fmt.Printf("%s ", dataNode.KeyAndValue.Key)
						}
						fmt.Print("|")
					} else {
						fmt.Print(node.Keys)
						for _, newChild := range node.Children {
							newTempArray = append(newTempArray, newChild)
						}

					}
				}
				tempArray = newTempArray
				fmt.Println()
			}

		}
	}
}
func (t *BPTree) FindLeft() *IndexNode {
	if len(t.Root.Children) > 0 {
		p := t.Root.Children[0]
		for len(p.Children) > 0 {
			p = p.Children[0]
		}
		return p
	}
	return t.Root
}




func (indexNode *IndexNode) removeDataNode(removeDataNode *DataNode, dataNodeIndex int) {
	// 修改被删除dataNode的前后dataNode的左右连接关系
	if removeDataNode.PreviousDataNode != nil { // 删除的是叶子结点中的处于中间位置的结点 维系双向链表的关系1
		removeDataNode.PreviousDataNode.NextDataNode = removeDataNode.NextDataNode
	}
	if removeDataNode.NextDataNode != nil { // 维系双向链表的关系2
		removeDataNode.NextDataNode.PreviousDataNode = removeDataNode.PreviousDataNode
	}
	// 删除节点 ？通过且切片的方式删除
	tempSlice := append([]*DataNode{}, indexNode.DataNodes[:dataNodeIndex]...)
	tempSlice = append(tempSlice, indexNode.DataNodes[dataNodeIndex+1:]...)
	indexNode.DataNodes = tempSlice
}




