package utils

import "testing"

// --------------------- 测试代码 ---------------------

// 运行 `go test` 时会自动执行该测试
func TestSet(t *testing.T) {
	strSet := NewSet([]string{"apple", "banana", "cherry"})

	if !strSet.Contains("banana") {
		t.Errorf("Expected 'banana' to be in the set")
	}
	if strSet.Contains("grape") {
		t.Errorf("Expected 'grape' to NOT be in the set")
	}

	intSet := NewSet([]int{1, 2, 3, 4, 5})
	if !intSet.Contains(3) {
		t.Errorf("Expected 3 to be in the set")
	}
	if intSet.Contains(10) {
		t.Errorf("Expected 10 to NOT be in the set")
	}
}
