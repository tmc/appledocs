// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Movie] class.
var MovieClass objc.Class

func init() {
	MovieClass = objc.GetClass("NSMovie")
}

type Movie struct {
	objc.ID
}

func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{
		ID: objc.ID(ptr),
	}
}




