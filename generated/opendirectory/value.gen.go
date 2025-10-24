// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [value] class.
var (
	ValueClass     _valueClass
	ValueClassOnce sync.Once
)

func getvalueClass() _valueClass {
	ValueClassOnce.Do(func() {
		ValueClass = _valueClass{objc.GetClass("value")}
	})
	return ValueClass
}

type _valueClass struct {
	class objc.Class
}

// An interface definition for the [value] class.
type Ivalue interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-c.ivar
type value struct {
	objectivec.Object
}

// valueFrom constructs a [value] from an unsafe.Pointer.
func valueFrom(ptr unsafe.Pointer) value {
	return value{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _valueClass) Alloc() value {
	rv := objc.Send[value](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _valueClass) New() value {
	rv := objc.Send[value](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ value) Init() value {
	rv := objc.Send[value](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ value) Autorelease() value {
	rv := objc.Send[value](v_.ID, objc.Sel("autorelease"))
	return rv
}

// Newvalue creates a new value instance.
func Newvalue() value {
	return getvalueClass().New()
}




