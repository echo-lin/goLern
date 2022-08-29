package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golern/common"
	"golern/dto"
	"golern/model"
	"golern/response"
	"golern/util"
)

type getChapterListByDocsId struct {
	Id int `json:"id"` // 文档id
}

// AddChapter
// @Description	添加章节
// @Date 2022-08-18 10:41:12
// @param c
func AddChapter(c *gin.Context) {
	chapter := model.Chapter{}
	c.Bind(&chapter)
	fmt.Println(chapter.Name, chapter.Did)
	if chapter.Name == "" || chapter.Did == 0 {
		response.Fail(c, "请正确填写数据", gin.H{})
		return
	}

	chapterData := model.Chapter{
		Name: chapter.Name,
		Pid:  chapter.Pid,
		Did:  chapter.Did,
	}
	dbRes := common.DB.Create(&chapterData)
	if dbRes.Error != nil {
		response.Fail(c, "操作失败", gin.H{})
		return
	}
	response.Success(c, "操作成功", gin.H{})
}

// ChapterList chapterList
// @Description	章节列表(无限级分类)
// @Date 2022-08-18 10:40:38
// @param c
func ChapterList(c *gin.Context) {
	var chapter []model.Chapter
	document := model.Document{}
	reqData := getChapterListByDocsId{}
	c.Bind(&reqData)
	fmt.Println(reqData)
	if reqData.Id == 0 {
		common.DB.First(&document)
	} else {
		common.DB.Where("id = ?", reqData.Id).First(&document)
	}
	if document.ID == 0 {
		response.Success(c, "操作成功", gin.H{"data": "", "msg": "没有相关文档"})
		return
	}
	common.DB.Where("did = ?", document.ID).Find(&chapter)
	response.Success(c, "操作成功", gin.H{"data": util.UnlimitedClass(dto.ToChapterDto(chapter)), "msg": document.DocName})

}
