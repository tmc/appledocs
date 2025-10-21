// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAttributeValueWaiter] class.
var (
	MTRAttributeValueWaiterClass     _MTRAttributeValueWaiterClass
	MTRAttributeValueWaiterClassOnce sync.Once
)

func getMTRAttributeValueWaiterClass() _MTRAttributeValueWaiterClass {
	MTRAttributeValueWaiterClassOnce.Do(func() {
		MTRAttributeValueWaiterClass = _MTRAttributeValueWaiterClass{objc.GetClass("MTRAttributeValueWaiter")}
	})
	return MTRAttributeValueWaiterClass
}

type _MTRAttributeValueWaiterClass struct {
	class objc.Class
}

// An interface definition for the [MTRAttributeValueWaiter] class.
type IMTRAttributeValueWaiter interface {
	objectivec.IObject
	Cancel()
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeValueWaiter
type MTRAttributeValueWaiter struct {
	objectivec.Object
}

// MTRAttributeValueWaiterFrom constructs a [MTRAttributeValueWaiter] from an unsafe.Pointer.
func MTRAttributeValueWaiterFrom(ptr unsafe.Pointer) MTRAttributeValueWaiter {
	return MTRAttributeValueWaiter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeValueWaiterClass) Alloc() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAttributeValueWaiterClass) New() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeValueWaiter) Init() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeValueWaiter) Autorelease() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeValueWaiter creates a new MTRAttributeValueWaiter instance.
func NewMTRAttributeValueWaiter() MTRAttributeValueWaiter {
	return getMTRAttributeValueWaiterClass().New()
}


// Cancel the wait for the set of attribute path/value pairs represented by this MTRAttributeValueWaiter. If the completion has not been called yet, it will becalled with MTRErrorCodeCancelled.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeValueWaiter/cancel()
func (m_ MTRAttributeValueWaiter) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeValueWaiter/uuid
func (m_ MTRAttributeValueWaiter) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("UUID"))
	return rv
}



