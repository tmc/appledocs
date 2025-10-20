// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ExistsCommandClass _ExistsCommandClass

func init() {
	ExistsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}
}

type _ExistsCommandClass struct {
	class objc.Class
}

type ExistsCommand struct {
	objc.ID
}

func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ID: objc.ID(ptr),
	}
}




