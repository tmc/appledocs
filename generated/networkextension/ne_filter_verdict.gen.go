// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterVerdict] class.
var (
	NEFilterVerdictClass     _NEFilterVerdictClass
	NEFilterVerdictClassOnce sync.Once
)

func getNEFilterVerdictClass() _NEFilterVerdictClass {
	NEFilterVerdictClassOnce.Do(func() {
		NEFilterVerdictClass = _NEFilterVerdictClass{objc.GetClass("NEFilterVerdict")}
	})
	return NEFilterVerdictClass
}

type _NEFilterVerdictClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterVerdict] class.
type INEFilterVerdict interface {
	objectivec.IObject
}

// The abstract base class for filter verdict classes.
//
// Filter providers use instances this class to inform the system about how to handle flows of network data.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict
type NEFilterVerdict struct {
	objectivec.Object
}

// NEFilterVerdictFrom constructs a [NEFilterVerdict] from an unsafe.Pointer.
//
// The abstract base class for filter verdict classes.
func NEFilterVerdictFrom(ptr unsafe.Pointer) NEFilterVerdict {
	return NEFilterVerdict{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterVerdictClass) Alloc() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterVerdictClass) New() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterVerdict) Init() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterVerdict) Autorelease() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterVerdict creates a new NEFilterVerdict instance.
func NewNEFilterVerdict() NEFilterVerdict {
	return getNEFilterVerdictClass().New()
}




