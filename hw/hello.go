package hello 

import "fmt"
import "rsc.io/quote"
import "math/rand"

func main() {
  fmt.Println(quote.Go())
}

func helloArrays(insertCount int) [3]int {
	var a [3]int
	for i := 0; i < insertCount; i++ {
		a[i] = rand.Int()
	}

	fmt.Print(a)
	return a
}
