package controller


import(
//"fmt"
//"blog/config"
"github.com/gin-gonic/gin"
"blog/model"
"strconv"
)

func GetArticles(c *gin.Context){
	limit,_:=strconv.Atoi(c.Param("l"))
	offset,_ :=strconv.Atoi(c.Param("o"))
	var nextoffset int
	if limit >10{
		limit=10
	}
	nextoffset=limit + offset
	res,err:=model.GetArticleList(limit,offset)
	if err!=nil{
		c.JSON(200,gin.H{"status":101,"error":"No record found"})
		return
	}
	c.JSON(200,gin.H{"status":200,"message":"success","data":res,"offset":nextoffset,"limit":limit})
	return
}

func GetArticleById(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	if id<=0{
		c.JSON(200, gin.H{"status":102,"error":"Param missing"})
		return
	}
	res,err:=model.GetArticleById(id)
	if err!=nil{
		c.JSON(200,gin.H{"status":200,"message":"success","data":res})
		return
	}
}

func CreateArticle(c *gin.Context){
	var article model.Article
	var id int64
	err:=c.BindJSON(&article)
	if err !=nil{
		c.JSON(200,gin.H{"status":103,"error":"Invalid json input"})
		return
	}
	id,err=model.CreateArticle(article)
	data:=make(map[string]string,0)
	data["id"]=strconv.FormatInt(id, 10)
	if err!=nil{
		c.JSON(200,gin.H{"status":200,"message":"Success","data":data})
		return
	}
}