
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Matrix] class.
var MatrixClass _MatrixClass

func init() {
	MatrixClass = _MatrixClass{objc.GetClass("NSMatrix")}
}

type _MatrixClass struct {
	objc.Class
}

// An interface definition for the [Matrix] class.
type IMatrix interface {
	ID() objc.ID
}

type Matrix struct {
	id objc.ID
}

func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ Matrix) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MatrixClass) New() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMatrix creates and returns a new initialized instance.
func NewMatrix() Matrix {
	return MatrixClass.New()
}

// Init initializes the instance.
func (m_ Matrix) Init() Matrix {
	rv := objc.Send[Matrix](m_.ID(), selInit)
	return rv
}
