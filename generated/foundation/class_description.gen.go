// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ClassDescription] class.
var classDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}

type _ClassDescriptionClass struct {
	class objc.Class
}

// An abstract class that provides the interface for querying the relationships and properties of a class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassDescription

type ClassDescription struct {
	objectivec.Object
}

// ClassDescriptionFrom constructs a [ClassDescription] from an unsafe.Pointer.
//
// An abstract class that provides the interface for querying the relationships and properties of a class.
func ClassDescriptionFrom(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{objectivec.Object{objc.ID(ptr)}}
}



