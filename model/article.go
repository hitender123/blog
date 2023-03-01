package model

import(
	"blog/config"
	"fmt"
)

type Article struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func GetArticleList(limit,offset int)([]*Article, error){
	query:="select id,title,content,author from article limit ? ?"
	rows,err:=config.DB.Query(query,offset,limit)
	if err!=nil{
		return nil, fmt.Errorf("No record found")
	}
	article:=make([]*Article,0)
	defer rows.Close()
	for  rows.Next(){
		 temp :=new(Article)
			err:= rows.Scan(&temp.Id,&temp.Title,&temp.Content,&temp.Author)
			if err !=nil{
				return nil,fmt.Errorf("Error in Scan data: %v",err)	
			}
		article=append(article,temp)
	}
 return article,nil
}

func GetArticleById(id int)(*Article,error){
	article:=new(Article)
	query:="select id,title,content,author from article where id=?"
	err:=config.DB.QueryRow(query,id).Scan(&article.Id,&article.Title,&article.Content,&article.Author)
	if err!=nil{
		return nil,fmt.Errorf("No record found")
	}
return article,nil
}

func CreateArticle(data Article)(int64,error){
	// var stmt *sql.Stmt
	stmt,err:=config.DB.Prepare("Insert Into(title,content,author)Values(?,?,?,)")
	if err !=nil{
		return 0,fmt.Errorf("Error inerting Article: %v",err)
	}
	result,errR:=stmt.Exec(data.Title,data.Content,data.Author)
	if errR!=nil{
		return 0,fmt.Errorf("Error inserting Article %v",errR)
	}
	Id,_:=result.LastInsertId()
	return Id,nil
}