// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexPath] class.
var indexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}

type _IndexPathClass struct {
	class objc.Class
}

// A list of indexes that together represent the path to a specific location in a tree of nested arrays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath

type IndexPath struct {
	objectivec.Object
}

// IndexPathFrom constructs a [IndexPath] from an unsafe.Pointer.
//
// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{objectivec.Object{objc.ID(ptr)}}
}



