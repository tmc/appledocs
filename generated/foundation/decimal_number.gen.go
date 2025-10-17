// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DecimalNumber] class.
var DecimalNumberClass objc.Class

func init() {
	DecimalNumberClass = objc.GetClass("NSDecimalNumber")
}

type DecimalNumber struct {
	objc.ID
}

func DecimalNumberFrom(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		ID: objc.ID(ptr),
	}
}



