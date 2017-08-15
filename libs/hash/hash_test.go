package hash

import "testing"

func TestToHashNum(t *testing.T) {
	result := ToHashNum("asdasdasd")
	t.Log("result 1:",result)

	result = ToHashNum("asdasdasd")
	t.Log("result 2:",result)
}
