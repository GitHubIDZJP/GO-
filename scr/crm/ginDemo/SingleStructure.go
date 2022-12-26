package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type sProductModel struct {
	Name     string
	CommentS string
	Age      int
	//OrderId  string    //`json:"order_id,omitempty"`
	//CommentS []CommentModel //`json:"comment_s,omitempty"`
	//OneComment *Comment  `json:"one_comment,omitempty"`

}

func SingleStructure() {

	r := gin.Default()

	r.GET("/JS", func(c *gin.Context) {
		P_M := [2]sProductModel{}
		nameArr := []string{"小刘", "小黄"}
		ageArr := []int{20, 90}
		contentArr := []string{"评论1:你今天真美", "评论2:你明天要去爬山么?"}
		for i := 0; i < len(P_M); i++ {
			P_M[i] = sProductModel{
				Name:     nameArr[i],
				CommentS: contentArr[i],
				Age:      ageArr[i],
			}
		}

		data := map[string]interface{}{

			"data": P_M,
		}

		c.JSONP(http.StatusOK, gin.H{"code": http.StatusOK,
			"message": "响应成功",
			"status":  "normal",
			"result":  data,
		})
	})
	r.Run(":8080")
}
