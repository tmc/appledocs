// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MutableSet] class.
var (
	MutableSetClass     _MutableSetClass
	MutableSetClassOnce sync.Once
)

func getMutableSetClass() _MutableSetClass {
	MutableSetClassOnce.Do(func() {
		MutableSetClass = _MutableSetClass{objc.GetClass("NSMutableSet")}
	})
	return MutableSetClass
}

type _MutableSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableSet] class.
type IMutableSet interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type MutableSet struct {
	objectivec.Object
}

// MutableSetFrom constructs a [MutableSet] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableSetClass) Alloc() MutableSet {
	rv := objc.Send[MutableSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableSetClass) New() MutableSet {
	rv := objc.Send[MutableSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableSet) Init() MutableSet {
	rv := objc.Send[MutableSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableSet) Autorelease() MutableSet {
	rv := objc.Send[MutableSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableSet creates a new MutableSet instance.
func NewMutableSet() MutableSet {
	return getMutableSetClass().New()
}




