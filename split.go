package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)
type Login struct {
	Message  string `json:"message"`

}
type Word struct{
	Word string `json:"word"`
	Length int `json:"length"`
}
func main() {
	r := gin.Default()
	r.POST("/lab4", func(c *gin.Context) {
		var message Login
		c.ShouldBindJSON(&message)

		m := strings.Split(message.Message, " ")
		var words []word
		for i := 0; i < len(m); i++ {
			words = append(words, word{Word: m[i], Length: len(m[i])})
			fmt.Print(m[i])
		}
		c.JSON(200, words)
	})
	r.Run(":8000")

}