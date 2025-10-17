// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedSet] class.
var orderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}

type _OrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objectivec.IObject
}

// A static, ordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet

type OrderedSet struct {
	objectivec.Object
}

// OrderedSetFrom constructs a [OrderedSet] from an unsafe.Pointer.
//
// A static, ordered collection of unique objects.
func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{objectivec.Object{objc.ID(ptr)}}
}



