// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexSet] class.
var indexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}

type _IndexSetClass struct {
	class objc.Class
}

// An interface definition for the [IndexSet] class.
type IIndexSet interface {
	objectivec.IObject
}

// An immutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet

type IndexSet struct {
	objectivec.Object
}

// IndexSetFrom constructs a [IndexSet] from an unsafe.Pointer.
//
// An immutable collection of unique integer values that represent indexes in another collection.
func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{objectivec.Object{objc.ID(ptr)}}
}



