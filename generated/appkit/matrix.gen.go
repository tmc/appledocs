// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Matrix] class.
var matrixClass = _MatrixClass{objc.GetClass("NSMatrix")}

type _MatrixClass struct {
	class objc.Class
}

// An interface definition for the [Matrix] class.
type IMatrix interface {
	IControl
}

// A legacy interface for grouping radio buttons or other types of cells together. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix

type Matrix struct {
	Control
}

// MatrixFrom constructs a [Matrix] from an unsafe.Pointer.
//
// A legacy interface for grouping radio buttons or other types of cells together.
func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		Control: ControlFrom(ptr),
	}
}



