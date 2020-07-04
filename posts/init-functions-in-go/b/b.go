package b

import (
	"fmt"
	"init-functions/a"
)

var B = "B"

func init() {
	fmt.Println("package B")
	fmt.Println("Print var B from pkg a: ", a.A)
}
