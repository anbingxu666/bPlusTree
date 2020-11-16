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

func ShowTree(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; ok {
		appG.ResponseString(http.StatusOK, TreeMap[id].UpToDownPrint())
	} else {
		appG.ResponseString(http.StatusOK, "BPlusTree not exist!")
	}
}

func CreateTree(c *gin.Context) {
	json := Tree{}
	appG := app.Gin{C: c}

	err := c.BindJSON(&json)

	if err != nil {
		appG.ResponseJson(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if _, ok := TreeMap[json.Id]; ok {
		appG.ResponseJson(http.StatusOK, e.ALREADY_EXISTS, nil)
	} else {
		TreeMap[json.Id] = service.BuildBPTree(json.Order, json.KeyArray, json.ValArray)
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

func Get(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusOK, e.TREE_NOT_EXISTS, nil)
		return
	}
	appG.ResponseString(http.StatusOK, service.GetNode(TreeMap[id], key))
}

func Update(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")
	val := c.PostForm("val")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusOK, e.TREE_NOT_EXISTS, nil)
		return
	}

	ok, err := service.UpdateNode(TreeMap[id], key, val)
	if !ok {
		appG.ResponseJson(http.StatusOK, e.UPDATE_FAILED, err.Error())
	} else {
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

func Remove(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusOK, e.TREE_NOT_EXISTS, nil)
		return
	}
	ok, err := service.RemoveNode(TreeMap[id], key)
	if !ok {
		appG.ResponseJson(http.StatusOK, e.REMOVE_FAILED, err.Error())
	} else {
		appG.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	}
}

func Marshal(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusOK, e.TREE_NOT_EXISTS, nil)
		return
	}

	data, err := service.MarshalTree(TreeMap[id])

	if err != nil {
		appG.ResponseString(http.StatusOK, err.Error())
		return
	}

	appG.ResponseString(http.StatusOK, data)
}

func Travel(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.PostForm("id"))

	if _, ok := TreeMap[id]; !ok {
		appG.ResponseJson(http.StatusOK, e.TREE_NOT_EXISTS, nil)
		return
	}
	appG.ResponseString(http.StatusOK, service.TraversingLeafNodes(TreeMap[id]))
}
