package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestPage(t *testing.T) {
	paginator := acoolTools.PageUtils.Paginator(5, 20, 500)
	fmt.Println(paginator)
}
