// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Exception] class.
var exceptionClass = _ExceptionClass{objc.GetClass("NSException")}

type _ExceptionClass struct {
	class objc.Class
}

// An interface definition for the [Exception] class.
type IException interface {
	objectivec.IObject
}

// An object that represents a special condition that interrupts the normal flow of program execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException

type Exception struct {
	objectivec.Object
}

// ExceptionFrom constructs a [Exception] from an unsafe.Pointer.
//
// An object that represents a special condition that interrupts the normal flow of program execution.
func ExceptionFrom(ptr unsafe.Pointer) Exception {
	return Exception{objectivec.Object{objc.ID(ptr)}}
}



