// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var countCommandClass _CountCommandClass

func init() {
	countCommandClass = _CountCommandClass{objc.GetClass("NSCountCommand")}
}

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




