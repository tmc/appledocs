// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Matrix] class.
var (
	matrixClass     _MatrixClass
	matrixClassOnce sync.Once
)

func getMatrixClass() _MatrixClass {
	matrixClassOnce.Do(func() {
		matrixClass = _MatrixClass{objc.GetClass("NSMatrix")}
	})
	return matrixClass
}

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
// Alloc allocates a new instance without initialization.
func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MatrixClass) New() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Matrix) Init() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Matrix) Autorelease() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrix creates a new Matrix instance.
func NewMatrix() Matrix {
	return getMatrixClass().New()
}




