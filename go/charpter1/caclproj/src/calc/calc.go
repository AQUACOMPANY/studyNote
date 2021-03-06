package main
/* import	"os" */
import	"fmt"
/* import	"simplemath"
import	"strconv" */
import	"github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

var Usage = func(){
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\n The commands are:\n\tadd\tAddtion of two values.\n\tsqr\tSquare root of a non-negative values.")
}

/* func main(){
	args := os.Args[1:]
	if args == nil || len(args) < 2{
		Usage()
		return
	}
	switch args[0]{
		case "add":
			if len(args) !=3 {
				fmt.Println("USAGE: calc add <integer1><integer2>")
				return
			}
			v1, err1 := strconv.Atoi(args[1])
			v2, err2 := strconv.Atoi(args[2])
			if err1 != nil || err2 != nil{
				fmt.Println("USAGE: calc add <integer1><integer2>")
				return
			}
			ret :=simplemath.Add(v1,v2)
			fmt.Println("Result: ", ret)
			case "sqrt":
				if len(args) !=2 {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			v, err := strconv.Atoi(args[1])
				if err != nil{
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			ret :=simplemath.Sqrt(v)
			fmt.Println("Result: ", ret)
			default:
				Usage()
	}
} */