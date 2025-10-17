// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BlockOperation] class.
var BlockOperationClass objc.Class

func init() {
	BlockOperationClass = objc.GetClass("NSBlockOperation")
}

type BlockOperation struct {
	objc.ID
}

func BlockOperationFrom(ptr unsafe.Pointer) BlockOperation {
	return BlockOperation{
		ID: objc.ID(ptr),
	}
}



