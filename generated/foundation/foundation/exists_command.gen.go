// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ExistsCommand] class.
var ExistsCommandClass objc.Class

func init() {
	ExistsCommandClass = objc.GetClass("NSExistsCommand")
}

type ExistsCommand struct {
	objc.ID
}

func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ID: objc.ID(ptr),
	}
}




