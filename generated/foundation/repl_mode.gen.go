// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [replMode] class.
var (
	ReplModeClass     _replModeClass
	ReplModeClassOnce sync.Once
)

func getreplModeClass() _replModeClass {
	ReplModeClassOnce.Do(func() {
		ReplModeClass = _replModeClass{objc.GetClass("replMode")}
	})
	return ReplModeClass
}

type _replModeClass struct {
	class objc.Class
}

// An interface definition for the [replMode] class.
type IreplMode interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/replMode

type replMode struct {
	objectivec.Object
}

// replModeFrom constructs a [replMode] from an unsafe.Pointer.
func replModeFrom(ptr unsafe.Pointer) replMode {
	return replMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _replModeClass) Alloc() replMode {
	rv := objc.Send[replMode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _replModeClass) New() replMode {
	rv := objc.Send[replMode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ replMode) Init() replMode {
	rv := objc.Send[replMode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ replMode) Autorelease() replMode {
	rv := objc.Send[replMode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreplMode creates a new replMode instance.
func NewreplMode() replMode {
	return getreplModeClass().New()
}




