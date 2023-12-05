package bot_test_dep_a

import (
	"fmt"

	dummy "github.com/phoenix2x/bot-test-dep-b"
)

func F() {
	fmt.Println(dummy.Dummy)
	fmt.Println("something")
}
