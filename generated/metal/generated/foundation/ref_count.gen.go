// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [refCount] class.
var (
	RefCountClass     _refCountClass
	RefCountClassOnce sync.Once
)

func getrefCountClass() _refCountClass {
	RefCountClassOnce.Do(func() {
		RefCountClass = _refCountClass{objc.GetClass("refCount")}
	})
	return RefCountClass
}

type _refCountClass struct {
	class objc.Class
}

// An interface definition for the [refCount] class.
type IrefCount interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/refCount
type refCount struct {
	objectivec.Object
}

// refCountFrom constructs a [refCount] from an unsafe.Pointer.
func refCountFrom(ptr unsafe.Pointer) refCount {
	return refCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _refCountClass) Alloc() refCount {
	rv := objc.Send[refCount](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _refCountClass) New() refCount {
	rv := objc.Send[refCount](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ refCount) Init() refCount {
	rv := objc.Send[refCount](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ refCount) Autorelease() refCount {
	rv := objc.Send[refCount](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrefCount creates a new refCount instance.
func NewrefCount() refCount {
	return getrefCountClass().New()
}




