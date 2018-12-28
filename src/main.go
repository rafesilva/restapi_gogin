package main
	
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "GET Home",
			})
		}

func PostHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"Post Home",
		})
}

func main() {
	fmt.Println("I'm working")

	r := gin.Default()
	r.GET("/", Home) 
	r.POST("/", PostHome)
	r.Run()
	}
