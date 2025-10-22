// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Coder] class.
var (
	CoderClass     _CoderClass
	CoderClassOnce sync.Once
)

func getCoderClass() _CoderClass {
	CoderClassOnce.Do(func() {
		CoderClass = _CoderClass{objc.GetClass("NSCoder")}
	})
	return CoderClass
}

type _CoderClass struct {
	class objc.Class
}

// An interface definition for the [Coder] class.
type ICoder interface {
	objectivec.IObject
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Coder struct {
	objectivec.Object
}

// CoderFrom constructs a [Coder] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoderClass) Alloc() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoderClass) New() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Coder) Init() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Coder) Autorelease() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoder creates a new Coder instance.
func NewCoder() Coder {
	return getCoderClass().New()
}




