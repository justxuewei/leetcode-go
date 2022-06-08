package jzoffer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	btree "github.com/xavier-niu/leetcode/btree"
)

type bTreeCodecTestSuite struct {
	t     *testing.T
	codec *BTreeCodec
}

func (s *bTreeCodecTestSuite) testSerialize(root *TreeNode, expected string) string {
	str := s.codec.serialize(root)
	assert.Equal(s.t, expected, str)
	return str
}

func (s *bTreeCodecTestSuite) testDeserialize(str string) *TreeNode {
	return s.codec.deserialize(str)
}

func TestBTreeCodec(t *testing.T) {
	suite := bTreeCodecTestSuite{t: t, codec: &BTreeCodec{}}

	// case1
	case1, _ := btree.GenerateBinaryTree([]int{1, 2, 3, btree.Null, btree.Null, 4, 5})
	case1str := suite.testSerialize(case1, "1,2,3,n,n,4,5")
	case1root := suite.testDeserialize(case1str)
	_ = suite.testSerialize(case1root, "1,2,3,n,n,4,5")

	// case2
	case2, _ := btree.GenerateBinaryTree([]int{1, 2, 3, btree.Null, btree.Null, 4, 5, 6, 7})
	case2str := suite.testSerialize(case2, "1,2,3,n,n,4,5,6,7")
	case2root := suite.testDeserialize(case2str)
	_ = suite.testSerialize(case2root, "1,2,3,n,n,4,5,6,7")
}
