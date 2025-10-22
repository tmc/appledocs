// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [wantsInvalid] class.
var (
	WantsInvalidClass     _wantsInvalidClass
	WantsInvalidClassOnce sync.Once
)

func getwantsInvalidClass() _wantsInvalidClass {
	WantsInvalidClassOnce.Do(func() {
		WantsInvalidClass = _wantsInvalidClass{objc.GetClass("wantsInvalid")}
	})
	return WantsInvalidClass
}

type _wantsInvalidClass struct {
	class objc.Class
}

// An interface definition for the [wantsInvalid] class.
type IwantsInvalid interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/wantsInvalid
type wantsInvalid struct {
	objectivec.Object
}

// wantsInvalidFrom constructs a [wantsInvalid] from an unsafe.Pointer.
func wantsInvalidFrom(ptr unsafe.Pointer) wantsInvalid {
	return wantsInvalid{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _wantsInvalidClass) Alloc() wantsInvalid {
	rv := objc.Send[wantsInvalid](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _wantsInvalidClass) New() wantsInvalid {
	rv := objc.Send[wantsInvalid](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ wantsInvalid) Init() wantsInvalid {
	rv := objc.Send[wantsInvalid](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ wantsInvalid) Autorelease() wantsInvalid {
	rv := objc.Send[wantsInvalid](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwantsInvalid creates a new wantsInvalid instance.
func NewwantsInvalid() wantsInvalid {
	return getwantsInvalidClass().New()
}




