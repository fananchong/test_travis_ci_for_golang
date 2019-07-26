package sumup2

import (
	"testing"

	// 随便加的一个库，为了引入 golang.org/x/crypto 这种墙外库，看是否正常
	_ "github.com/micro/go-plugins/registry/etcdv3"
)

func TestSum2(t *testing.T) {
	reference := 40 + 2
	result := Sum2(40, 2)

	if result != reference {
		t.Errorf("Sum should return %d instead of %d", reference, result)
	}
}
