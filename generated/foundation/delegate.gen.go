// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [delegate] class.
var (
	DelegateClass     _delegateClass
	DelegateClassOnce sync.Once
)

func getdelegateClass() _delegateClass {
	DelegateClassOnce.Do(func() {
		DelegateClass = _delegateClass{objc.GetClass("delegate")}
	})
	return DelegateClass
}

type _delegateClass struct {
	class objc.Class
}

// An interface definition for the [delegate] class.
type Idelegate interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/delegate-c.ivar

type delegate struct {
	objectivec.Object
}

// delegateFrom constructs a [delegate] from an unsafe.Pointer.
func delegateFrom(ptr unsafe.Pointer) delegate {
	return delegate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _delegateClass) Alloc() delegate {
	rv := objc.Send[delegate](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _delegateClass) New() delegate {
	rv := objc.Send[delegate](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ delegate) Init() delegate {
	rv := objc.Send[delegate](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ delegate) Autorelease() delegate {
	rv := objc.Send[delegate](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdelegate creates a new delegate instance.
func Newdelegate() delegate {
	return getdelegateClass().New()
}




