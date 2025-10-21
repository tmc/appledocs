// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [arrayLength] class.
var (
	ArrayLengthClass     _arrayLengthClass
	ArrayLengthClassOnce sync.Once
)

func getarrayLengthClass() _arrayLengthClass {
	ArrayLengthClassOnce.Do(func() {
		ArrayLengthClass = _arrayLengthClass{objc.GetClass("arrayLength")}
	})
	return ArrayLengthClass
}

type _arrayLengthClass struct {
	class objc.Class
}

// An interface definition for the [arrayLength] class.
type IarrayLength interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/arrayLength-c.ivar
type arrayLength struct {
	objectivec.Object
}

// arrayLengthFrom constructs a [arrayLength] from an unsafe.Pointer.
func arrayLengthFrom(ptr unsafe.Pointer) arrayLength {
	return arrayLength{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _arrayLengthClass) Alloc() arrayLength {
	rv := objc.Send[arrayLength](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _arrayLengthClass) New() arrayLength {
	rv := objc.Send[arrayLength](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ arrayLength) Init() arrayLength {
	rv := objc.Send[arrayLength](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ arrayLength) Autorelease() arrayLength {
	rv := objc.Send[arrayLength](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewarrayLength creates a new arrayLength instance.
func NewarrayLength() arrayLength {
	return getarrayLengthClass().New()
}




