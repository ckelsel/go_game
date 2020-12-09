package core

import (
	"fmt"
	"testing"
)

func TestNewGridManager(t *testing.T) {
	mgr := NewGridManager(0, 250, 5, 0, 250, 5)
	fmt.Println(mgr.String())
}
