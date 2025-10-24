// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [isQueueing] class.
var (
	IsQueueingClass     _isQueueingClass
	IsQueueingClassOnce sync.Once
)

func getisQueueingClass() _isQueueingClass {
	IsQueueingClassOnce.Do(func() {
		IsQueueingClass = _isQueueingClass{objc.GetClass("isQueueing")}
	})
	return IsQueueingClass
}

type _isQueueingClass struct {
	class objc.Class
}





// An interface definition for the [isQueueing] class.
type IisQueueing interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _isQueueingClass) Alloc() isQueueing {
	rv := objc.Send[isQueueing](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isQueueingClass) New() isQueueing {
	rv := objc.Send[isQueueing](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isQueueing) Init() isQueueing {
	rv := objc.Send[isQueueing](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isQueueing) Autorelease() isQueueing {
	rv := objc.Send[isQueueing](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisQueueing creates a new isQueueing instance.
func NewisQueueing() isQueueing {
	return getisQueueingClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isQueueing
type isQueueing struct {
	objectivec.Object
}

// isQueueingFrom constructs a [isQueueing] from an unsafe.Pointer.
func isQueueingFrom(ptr unsafe.Pointer) isQueueing {
	return isQueueing{objectivec.Object{objc.ID(ptr)}}
}































