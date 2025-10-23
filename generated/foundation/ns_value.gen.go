// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Value] class.
var (
	ValueClass     _ValueClass
	ValueClassOnce sync.Once
)

func getValueClass() _ValueClass {
	ValueClassOnce.Do(func() {
		ValueClass = _ValueClass{objc.GetClass("NSValue")}
	})
	return ValueClass
}

type _ValueClass struct {
	class objc.Class
}

// An interface definition for the [Value] class.
type IValue interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Value struct {
	objectivec.Object
}

// ValueFrom constructs a [Value] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ValueClass) Alloc() Value {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ValueClass) New() Value {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Value) Init() Value {
	rv := objc.Send[Value](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Value) Autorelease() Value {
	rv := objc.Send[Value](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewValue creates a new Value instance.
func NewValue() Value {
	return getValueClass().New()
}




