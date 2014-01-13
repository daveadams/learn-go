package main
import (
	"fmt"
	"github.com/daveadams/learn-go/newmath"
)

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("Sqrt(%d) = %v\n", i, newmath.Sqrt(float64(i)))
	}
}
