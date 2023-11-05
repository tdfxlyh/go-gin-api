package utils

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T) {
	fmt.Println(GetFirstPinYin(""))
	fmt.Println(GetFirstPinYin("wxy"))
	fmt.Println(GetFirstPinYin("明珠"))
	fmt.Println(GetFirstPinYin("往文龙sfa"))
	fmt.Println(GetFirstPinYin("#sd"))
	fmt.Println(GetFirstPinYin("#明珠"))
}
