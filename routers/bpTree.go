package routers

import (
	"BPTree_Web/bptree"
	"BPTree_Web/pkg/app"
	"BPTree_Web/pkg/e"
	"BPTree_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Tree struct {
	Id       int      `json:"id" binding:"required"`
	Order    int      `json:"order" binding:"required"`
	KeyArray []string `json:"key_array" binding:"required"`
	ValArray []string `json:"val_array" binding:"required"`
}

var TreeMap map[int]*bptree.BPTree

func Init() {
	TreeMap = make(map[int]*bptree.BPTree)
}

// @Summary Show the B Plus Tree
// @Param id formData int true "B+树id"
// @Success 200 {string} string "B+树的结构"
// @Failure 400 {string} string "B+树不存在"
// @Router /show [post]
func ShowTree(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; ok {
		appG.ResponseString(http.StatusOK, TreeMap[id].UpToDownPrint())
	} else {
		appG.ResponseString(http.StatusBadRequest, e.GetMsg(e.TREE_NOT_EXISTS))
	}
}

// @Summary Create a B Plus Tree
// @Accept  json
// @Produce  json
// @Param t body Tree true "B+树结构"
// @Success 200 {object} app.Response "创建成功"
// @Failure 404 {object} app.Response "创建失败"
// @Router /create [post]
func CreateTree(c *gin.Context) {
	json := Tree{}
	appG := app.Gin{C: c}

	err := c.BindJSON(&json)

	if err != nil {
		appG.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if _, ok := TreeMap[json.Id]; ok {
		appG.ResponseJson(http.StatusBadRequest, e.ALREADY_EXISTS, nil)
	} else {
		TreeMap[json.Id] = service.BuildBPTree(json.Order, json.KeyArray, json.ValArray)
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

// @Summary Get a B Plus Tree Key and Value
// @Param id formData int true "B+树id"
// @Param key formData string true "B+树的key"
// @Success 200 {string} string "key对应的value"
// @Failure 400 {string} string "错误信息"
// @Router /get [post]
func Get(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseString(http.StatusBadRequest, e.GetMsg(e.TREE_NOT_EXISTS))
		return
	}
	appG.ResponseString(http.StatusOK, service.GetNode(TreeMap[id], key))
}

// @Summary Update a B Plus Tree Key and Value
// @Produce  json
// @Param id formData int true "B+树id"
// @Param key formData string true "B+树的key"
// @Param val formData string true "B+树的value"
// @Success 200 {object} app.Response "修改成功"
// @Failure 404 {object} app.Response "修改失败"
// @Router /update [post]
func Update(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")
	val := c.PostForm("val")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusBadRequest, e.TREE_NOT_EXISTS, nil)
		return
	}

	ok, err := service.UpdateNode(TreeMap[id], key, val)
	if !ok {
		appG.ResponseJson(http.StatusBadRequest, e.UPDATE_FAILED, err.Error())
	} else {
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

// @Summary Remove a B Plus Tree Key and Value
// @Produce  json
// @Param id formData int true "B+树id"
// @Param key formData string true "B+树的key"
// @Success 200 {object} app.Response "删除成功"
// @Failure 404 {object} app.Response "删除失败"
// @Router /Remove [post]
func Remove(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusBadRequest, e.TREE_NOT_EXISTS, nil)
		return
	}
	ok, err := service.RemoveNode(TreeMap[id], key)
	if !ok {
		appG.ResponseJson(http.StatusBadRequest, e.REMOVE_FAILED, err.Error())
	} else {
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

// @Summary Serialized B Plus Tree Struct
// @Param id formData int true "B+树id"
// @Success 200 {string} string "B+树的序列化信息"
// @Failure 400 {string} string "错误信息"
// @Router /marshal [post]
func Marshal(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseString(http.StatusBadRequest, e.GetMsg(e.TREE_NOT_EXISTS))
		return
	}

	data, err := service.MarshalTree(TreeMap[id])

	if err != nil {
		appG.ResponseString(http.StatusBadRequest, err.Error())
		return
	}

	appG.ResponseString(http.StatusOK, data)
}

// @Summary Traversing leaf nodes of B Plus Tree
// @Param id formData int true "B+树id"
// @Success 200 {string} string "B+树的叶子节点遍历信息"
// @Failure 400 {string} string "错误信息"
// @Router /travel [post]
func Travel(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseString(http.StatusBadRequest, e.GetMsg(e.TREE_NOT_EXISTS))
		return
	}
	appG.ResponseString(http.StatusOK, service.TraversingLeafNodes(TreeMap[id]))
}
