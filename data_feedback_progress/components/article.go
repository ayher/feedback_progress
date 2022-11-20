package components

import (
	"data_feedback_progress/model"
	"data_feedback_progress/public/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context){
	fmt.Println(model.FpArticleMgr(database.GormFp()).GetFromID(1))
	Sucess(c,"成功")
}
