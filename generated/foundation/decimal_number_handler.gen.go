// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DecimalNumberHandler] class.
var DecimalNumberHandlerClass objc.Class

func init() {
	DecimalNumberHandlerClass = objc.GetClass("NSDecimalNumberHandler")
}

type DecimalNumberHandler struct {
	objc.ID
}

func DecimalNumberHandlerFrom(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{
		ID: objc.ID(ptr),
	}
}



