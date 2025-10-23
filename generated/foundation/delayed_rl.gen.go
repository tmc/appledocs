// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [delayedRL] class.
var (
	DelayedRLClass     _delayedRLClass
	DelayedRLClassOnce sync.Once
)

func getdelayedRLClass() _delayedRLClass {
	DelayedRLClassOnce.Do(func() {
		DelayedRLClass = _delayedRLClass{objc.GetClass("delayedRL")}
	})
	return DelayedRLClass
}

type _delayedRLClass struct {
	class objc.Class
}

// An interface definition for the [delayedRL] class.
type IdelayedRL interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/delayedRL
type delayedRL struct {
	objectivec.Object
}

// delayedRLFrom constructs a [delayedRL] from an unsafe.Pointer.
func delayedRLFrom(ptr unsafe.Pointer) delayedRL {
	return delayedRL{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _delayedRLClass) Alloc() delayedRL {
	rv := objc.Send[delayedRL](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _delayedRLClass) New() delayedRL {
	rv := objc.Send[delayedRL](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ delayedRL) Init() delayedRL {
	rv := objc.Send[delayedRL](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ delayedRL) Autorelease() delayedRL {
	rv := objc.Send[delayedRL](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdelayedRL creates a new delayedRL instance.
func NewdelayedRL() delayedRL {
	return getdelayedRLClass().New()
}




