package main

import (
	"BPlusProject/tree"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("B Plus Tree")

	// 初始化一颗5阶B树
	bpTree := tree.MallocNewBPTree(4)

	keyArray := []int{1,2,3,4,5,6,7}
	//keyArray := []int{55, 34, 15, 95, 99, 98, 81, 16, 99, 14, 36, 13, 77, 57, 37, 2, 39, 3, 89, 76}

	for _, key := range keyArray {

		//for n := 0; n < 50; n++ {
		//	rand.Seed(time.Now().UnixNano())
		//	key := rand.Intn(100)
		keystr := strconv.Itoa(key)
		keyAndValue := tree.KeyAndValue{
			"k" + keystr,
			"v" + keystr,
		}
		//fmt.Printf("开始插入： key:%s  \n\n", keyAndValue.Key)
		//
		bpTree.Insert(keyAndValue)
		//bpTree.UpToDownPrint()
		//t1.Traversal()
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("-----------------")
	fmt.Println()
	bpTree.Remove("k7")
	bpTree.UpToDownPrint()

}
