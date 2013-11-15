package TagTest

import "fmt"

func Tag() string {
	return "Second commit."
}

func init() {
	fmt.Println(Tag())
}
