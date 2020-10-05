package assistant

import (
	"fmt"
	"testing"
)

func TestGenerateID(t *testing.T) {
	var s int64 = 10
	var i int64
	for i = 0;i < s;i++{
		fmt.Println(GenerateID(i))
	}
}
