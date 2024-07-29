package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// func main() {
// 	//签名密钥
// 	mySignedKeys := []byte("woshilinzy")
// 	c := MyClaims{
// 		UserName: "linzy",
// 		StandardClaims: jwt.StandardClaims{
// 			//什么时间生效
// 			NotBefore: time.Now().Unix() - 60,
// 			//什么时间失效
// 			ExpiresAt: time.Now().Unix() + 500,
// 			//签发者
// 			Issuer: "linzy",
// 		},
// 	}
// 	//NewWithClaims使用指定的签名方法和声明创建一个新的令牌。
// 	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
// 	fmt.Println(t)
// 	// SignedString创建并返回一个完整的有符号的JWT。
// 	//使用令牌中指定的SigningMethod对令牌进行签名。
// 	s, err := t.SignedString(mySignedKeys)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	mySignedKeys := []byte("woshilinzy")

// 	//刚刚加密的JWT
// 	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJsaW56eSIsImV4cCI6MTcyMTU0NDUxOSwiaXNzIjoibGluenkiLCJuYmYiOjE3MjE1NDM5NTl9.QE8Q3lH3sqqsQGUbXe8URlgP9O7IJEiv0TM3O-DFenA"
// 	//解析JWT
// 	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return mySignedKeys, nil
// 	})
// 	//使用断言获取MyClaims结构体
// 	fmt.Println(token.Claims.(*MyClaims))
// 	fmt.Println(err)
// }

func a() {
	fmt.Println("ok")
}

func main() {
	a()

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", featchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		// v1.GET("/11", featchAllTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}

	router.Run()
}

// createTodo add new todo
func createTodo(c *gin.Context) {
	c.Request.ParseForm()
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	fmt.Println(c.Request.PostForm)
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created!", "resourceId": todo.ID})
}

func featchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no todo found!"})
		return
	}

	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}
		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func fetchSingleTodo(c *gin.Context) {
	var todos todoModel

	todoID := c.Param("id")

	db.First(&todos, todoID)

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no todo found!"})
		return
	}

	completed := false
	if todos.Completed == 1 {
		completed = true
	}
	_todos := transformedTodo{ID: todos.ID, Title: todos.Title, Completed: completed}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func updateTodo(c *gin.Context) {
	var todos todoModel

	todoID := c.Param("id")

	db.First(&todos, todoID)

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no todo found!"})
		return
	}

	db.Model(&todos).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todos).Update("completed", completed)

	fmt.Println(todos)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update ok."})
}

func deleteTodo(c *gin.Context) {
	var todos todoModel

	todoID := c.Param("id")

	db.First(&todos, todoID)

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no todo found!"})
		return
	}

	db.Delete(todos)

	fmt.Println(todos)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "delete ok."})
}
