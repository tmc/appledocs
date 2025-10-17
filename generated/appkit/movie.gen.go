// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Movie] class.
var movieClass = _MovieClass{objc.GetClass("NSMovie")}

type _MovieClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie

type Movie struct {
	objectivec.Object
}

// MovieFrom constructs a [Movie] from an unsafe.Pointer.
func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{objectivec.Object{objc.ID(ptr)}}
}



