package main
import "fmt"
func main(){
	fmt.Println("Hello, world. 你好，世界！")
	count()
	printPoint()
}

func count(){
	fmt.Println("counting")
	
			for i := 0; i < 10; i++ {
					defer fmt.Println(i)
			}
	
			fmt.Println("done")
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func printPoint() {
	fmt.Println(v1, p, v2, v3)
}