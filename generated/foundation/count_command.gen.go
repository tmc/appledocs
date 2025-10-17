// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountCommand] class.
var CountCommandClass = _CountCommandClass{objc.GetClass("NSCountCommand")}

type _CountCommandClass struct {
	class objc.Class
}

type CountCommand struct {
	objc.ID
}

func CountCommandFrom(ptr unsafe.Pointer) CountCommand {
	return CountCommand{
		ID: objc.ID(ptr),
	}
}




