package backend 

import (
	"testing"
	"fmt"

	"github.com/MichaelDao/version-bump/pkg/dummy"
)

type UltimateTest struct {
	BackendImpl Backend
}


func TestInterfaces(t *testing.T) {
	test := UltimateTest{}
	test.BackendImpl = dummy.Dummy {}

	fmt.Println("Passing dummy interface: " + test.BackendImpl.Lock())
	fmt.Println("Passing dummy interface: " + test.BackendImpl.Unlock())
}
