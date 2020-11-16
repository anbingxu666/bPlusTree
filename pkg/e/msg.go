package e

var Messages = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "fail",
	INVALID_PARAMS:  "请求参数错误",
	ALREADY_EXISTS:  "B+树已经存在",
	TREE_NOT_EXISTS: "B+树不存在",
	UPDATE_FAILED:   "修改结点失败",
	REMOVE_FAILED:   "删除节点失败",
}

func GetMsg(code int) string {
	msg, ok := Messages[code]
	if ok {
		return msg
	}
	return Messages[ERROR]
}
