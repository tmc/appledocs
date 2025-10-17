// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Form] class.
var formClass = _FormClass{objc.GetClass("NSForm")}

type _FormClass struct {
	class objc.Class
}

// An interface definition for the [Form] class.
type IForm interface {
	IMatrix
}

// An object is a vertical matrix of objects to implement the fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSForm

type Form struct {
	Matrix
}

// FormFrom constructs a [Form] from an unsafe.Pointer.
//
// An object is a vertical matrix of objects to implement the fields.
func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		Matrix: MatrixFrom(ptr),
	}
}



