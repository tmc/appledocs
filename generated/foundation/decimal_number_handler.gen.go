// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DecimalNumberHandler] class.
var decimalNumberHandlerClass = _DecimalNumberHandlerClass{objc.GetClass("NSDecimalNumberHandler")}

type _DecimalNumberHandlerClass struct {
	class objc.Class
}

// An interface definition for the [DecimalNumberHandler] class.
type IDecimalNumberHandler interface {
	objectivec.IObject
}

// A class that adopts the decimal number behaviors protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler

type DecimalNumberHandler struct {
	objectivec.Object
}

// DecimalNumberHandlerFrom constructs a [DecimalNumberHandler] from an unsafe.Pointer.
//
// A class that adopts the decimal number behaviors protocol.
func DecimalNumberHandlerFrom(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{objectivec.Object{objc.ID(ptr)}}
}



