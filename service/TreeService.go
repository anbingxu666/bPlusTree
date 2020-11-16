package service

import (
	"BPTree_Web/bptree"
	"encoding/json"
)

func BuildBPTree(order int, keys, vals []string) *bptree.BPTree {
	t := bptree.MallocNewBPTree(order)

	for i := 0; i < len(keys); i++ {
		keyAndValue := bptree.KeyAndValue{keys[i], vals[i]}
		t.Insert(keyAndValue)
	}

	return t
}

func GetNode(t *bptree.BPTree, key string) string {
	dataNode, _ := t.Get(key)
	if dataNode == nil {
		return "Not exist"
	}
	return dataNode.KeyAndValue.Value
}

func UpdateNode(t *bptree.BPTree, key, val string) (bool, error) {
	KV := bptree.KeyAndValue{
		Key:   key,
		Value: val,
	}

	return t.Update(KV)
}

func RemoveNode(t *bptree.BPTree, key string) (bool, error) {
	return t.Remove(key)
}

func MarshalTree(t *bptree.BPTree) (string, error) {
	data, err := json.Marshal(t)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func TraversingLeafNodes(t *bptree.BPTree) string {
	return t.Traversal()
}
