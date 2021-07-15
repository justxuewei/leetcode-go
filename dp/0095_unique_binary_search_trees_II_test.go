package dp

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	case1 := generateTrees(2)
	fmt.Println(case1)

	case2 := generateTrees(3)
	fmt.Println(case2)
}
