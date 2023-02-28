package main

import(
	"blog/config"
	"github.com/gin-gonic/gin"
	"blog/controller"
)

//-- set cors policy allow
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, username, token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		c.Next()
	}
}
func main(){
	//- create DB instance
	config.InitDB()
	//- set route hadler
	r := gin.Default()
	r.Use(Cors())
	r.GET("/articles", controller.GetArticles)
	r.GET("/article/:id",controller.GetArticleById)
	r.POST("article/create",controller.CreateArticle)
r.Run(":8080")
}
