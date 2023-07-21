package main

import (
	"net/http"
	"github.com/gin-gonic/gin"//must be downlowded into app
	"errors"
)

//first is defining how our data structure is going to look by setting the todo struct
type todo struct {
	ID						string `json:"id"`
	Item 					string `json:"item"`
	Completed			bool	 `json:"completed"`
	//when server sends data back to the client and client to server `json:__` helps to convert to json format
}

//creating todo array
var todos = []todo {
		{ID: "1", Item: "Clean Room", Completed: false},
		{ID: "2", Item: "Read book", Completed: false},
		{ID: "3", Item: "Record video", Completed: false},
}

//get todos takes one parameter(will contain the inf about the incoming http req)  of type *gin.Context
//gin.Context is a standard package in Golang that is used to access or share data
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
	//gonna convert the var todos variable into json format
	//in brackets we are supplying it with the status of upcoming req (http is imported package)
}



///nil is a predefined identifier in Go that represents zero values of many types.

//context.BindJSON -- going to take json from the request body and bind it into our todo var with todo type. it probably will return an error if it doesn't have the json format like in our todo struct. In this case our code will not be executed an we will RETURN back.
func addTodo(context *gin.Context) {
	var newTodo todo

	if err:= context.BindJSON(&newTodo); err != nil {
			return
	}

  todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}


func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
				return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}


func main() {
		router:= gin.Default()//creating a server
		router.GET("/todos", getTodos)//in colons we have a path that we are appending to url(lolaclhost) 2nd parameter is the function
		router.GET("/todos/:id", getTodo)
		router.PATCH("/todos/:id", toggleTodoStatus)
		router.POST("/todos", addTodo)
		router.Run("localhost:9090")//running the server
}
