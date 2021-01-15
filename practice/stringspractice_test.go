package practice

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintLastIndex(t *testing.T) {
	PrintLastIndex("abc:dae:hi", ":")
	PrintLastIndex("abc:dae:h:i", ":")
	PrintLastIndex("abc:daehi", ":")
}

func PrintLastIndex(s, substr string) {
	// substr最后一次出现在s中的位置的下标
	n := strings.LastIndex(s, substr)
	fmt.Printf("s=%s substr=%s n=%d\n", s, substr, n)
}
