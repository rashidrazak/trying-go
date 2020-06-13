package main

// Multiple packages must be separated by a space not comma
import (
	"fmt"
	"github.com/rashidrazak/trying-go/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(3.6))
	fmt.Println(math.Ceil(3.2))
	fmt.Println(math.Sqrt(81))

	fmt.Println(strutil.Reverse("hello"))
}
