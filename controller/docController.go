package controller

import (
	"github.com/gin-gonic/gin"
	"golern/common"
	"golern/model"
	"golern/response"
	"net/http"
)

// AddDoc
// @Description	添加文档名称
// @Date 2022-08-16 08:58:30
// @param c
func AddDoc(c *gin.Context) {
	var document model.Document
	DB := common.GetDB()
	docData := model.Document{}
	c.Bind(&docData)
	DB.Where("doc_name = ?", docData.DocName).First(&document)
	if document.ID != 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "文档名称重复", gin.H{})
		return
	}
	documentCreate := model.Document{
		DocName: docData.DocName,
	}

	if documentCreate.DocName == "" {
		response.Fail(c, "请正确填写数据", gin.H{})
		return
	}

	dbRes := DB.Create(&documentCreate)
	if dbRes.Error != nil {
		response.Fail(c, "操作失败", gin.H{})
	}

	response.Success(c, "请求成功", gin.H{})
}

// DocList
// @Description	获取文档列表
// @Date 2022-08-16 10:59:42
// @param c
func DocList(c *gin.Context) {
	var document []model.Document
	DB := common.GetDB()
	DB.Order("id desc").Find(&document)
	response.Success(c, "请求成功", gin.H{"data": document})
}
