// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Matrix] class.
var MatrixClass objc.Class

func init() {
	MatrixClass = objc.GetClass("NSMatrix")
}

type Matrix struct {
	objc.ID
}

func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		ID: objc.ID(ptr),
	}
}




