// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var decimalNumberHandlerClass _DecimalNumberHandlerClass

func init() {
	decimalNumberHandlerClass = _DecimalNumberHandlerClass{objc.GetClass("NSDecimalNumberHandler")}
}

type _DecimalNumberHandlerClass struct {
	class objc.Class
}

type DecimalNumberHandler struct {
	objc.ID
}

func DecimalNumberHandlerFrom(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{
		ID: objc.ID(ptr),
	}
}




