package TagTest

import "fmt"

func Tag() string {
	return "Tagged v0.0.0"
}

func init() {
	fmt.Println(Tag())
}
