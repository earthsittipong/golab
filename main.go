package main

import ("github.com/gin-gonic/gin"
"strconv")

func main() {
	r := gin.Default()
	r.GET("/lab3", func(c *gin.Context) {
		point := c.DefaultQuery("point", "0")
		p, _ := strconv.Atoi(point)
		if p > 90 {
			c.JSON(200, gin.H{
				"grade": "A",
				"point": point,
			})
		} else if p > 80 {
			c.JSON(200, gin.H{
				"grade": "B",
				"point": point,
			})
		} else if p > 70 {
			c.JSON(200, gin.H{
				"grade": "C",
				"point": point,
			})
		} else if p > 60 {
			c.JSON(200, gin.H{
				"grade": "D",
				"point": point,
			})
		} else if p <= 60 {
			c.JSON(200, gin.H{
				"grade": "F",
				"point": point,
			})
		}
	})
	r.Run(":8000")
}