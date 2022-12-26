package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductModel struct {
	Name string
	//OrderId  string    //`json:"order_id,omitempty"`
	CommentS []CommentModel //`json:"comment_s,omitempty"`
	//OneComment *Comment  `json:"one_comment,omitempty"`

}
type CommentModel struct {
	Id      string //`json:"id,omitempty"`
	Content string //`json:"content,omitempty"`
}

func DoubleStructure() {

	r := gin.Default()

	//r.GET("/JSONP?callback=x", func(c *gin.Context) {
	r.GET("/JS", func(c *gin.Context) {
		P_M := [2]ProductModel{}
		nameArr := []string{"小刘", "小黄"}
		idArr := []string{"10", "20"}
		contentArr := []string{"评论1:你今天真美", "评论2:你明天要去爬山么?"}
		for i_c := 0; i_c < 2; i_c++ {
			contentArrObj := CommentModel{Id: idArr[i_c], Content: contentArr[i_c]}
			fetchContent := []CommentModel{contentArrObj}
			P_M[i_c] = ProductModel{
				Name:     nameArr[i_c],
				CommentS: fetchContent,
			}
		}

		data := map[string]interface{}{
			//"foo": "bar",
			"data": P_M,
		}

		c.JSONP(http.StatusOK, gin.H{"code": http.StatusOK,
			"message": "响应成功",
			"status":  "normal",
			"result":  data,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
