package bytedance

import "testing"

func TestLRUCache(t *testing.T) {
	case1 := Constructor(2)
	var case1Output []int
	case1.Put(1, 1)
	case1.Put(2, 2)
	case1Output = append(case1Output, case1.Get(1))
	case1.Put(3, 3)
	case1Output = append(case1Output, case1.Get(2))
	case1.Put(4, 4)
	case1Output = append(case1Output, case1.Get(1))
	case1Output = append(case1Output, case1.Get(3))
	case1Output = append(case1Output, case1.Get(4))
	t.Log(case1Output)
}
