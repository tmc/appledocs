// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [functionAttributes] class.
var (
	FunctionAttributesClass     _functionAttributesClass
	FunctionAttributesClassOnce sync.Once
)

func getfunctionAttributesClass() _functionAttributesClass {
	FunctionAttributesClassOnce.Do(func() {
		FunctionAttributesClass = _functionAttributesClass{objc.GetClass("functionAttributes")}
	})
	return FunctionAttributesClass
}

type _functionAttributesClass struct {
	class objc.Class
}

// An interface definition for the [functionAttributes] class.
type IfunctionAttributes interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-c.ivar
type functionAttributes struct {
	objectivec.Object
}

// functionAttributesFrom constructs a [functionAttributes] from an unsafe.Pointer.
func functionAttributesFrom(ptr unsafe.Pointer) functionAttributes {
	return functionAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _functionAttributesClass) Alloc() functionAttributes {
	rv := objc.Send[functionAttributes](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _functionAttributesClass) New() functionAttributes {
	rv := objc.Send[functionAttributes](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ functionAttributes) Init() functionAttributes {
	rv := objc.Send[functionAttributes](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ functionAttributes) Autorelease() functionAttributes {
	rv := objc.Send[functionAttributes](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewfunctionAttributes creates a new functionAttributes instance.
func NewfunctionAttributes() functionAttributes {
	return getfunctionAttributesClass().New()
}




