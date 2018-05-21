//using a loop to print out binary
package main

import "fmt"
import "strconv"

func main() {

	for i := 0; i < 16; i++ {
	n := int64(i)
	fmt.Println(strconv.FormatInt(n, 2)) // 1111011
	}
}
