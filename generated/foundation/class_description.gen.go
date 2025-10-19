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

// An interface definition for the [ClassDescription] class.
type IClassDescription interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (cc _ClassDescriptionClass) Alloc() ClassDescription {
	rv := objc.Send[ClassDescription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _ClassDescriptionClass) New() ClassDescription {
	rv := objc.Send[ClassDescription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClassDescription) Init() ClassDescription {
	rv := objc.Send[ClassDescription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClassDescription) Autorelease() ClassDescription {
	rv := objc.Send[ClassDescription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClassDescription creates a new ClassDescription instance.
func NewClassDescription() ClassDescription {
	return classDescriptionClass.New()
}




