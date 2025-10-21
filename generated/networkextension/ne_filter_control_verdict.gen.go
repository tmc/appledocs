// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterControlVerdict] class.
var (
	NEFilterControlVerdictClass     _NEFilterControlVerdictClass
	NEFilterControlVerdictClassOnce sync.Once
)

func getNEFilterControlVerdictClass() _NEFilterControlVerdictClass {
	NEFilterControlVerdictClassOnce.Do(func() {
		NEFilterControlVerdictClass = _NEFilterControlVerdictClass{objc.GetClass("NEFilterControlVerdict")}
	})
	return NEFilterControlVerdictClass
}

type _NEFilterControlVerdictClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterControlVerdict] class.
type INEFilterControlVerdict interface {
	INEFilterNewFlowVerdict
}

// The result from a filter control provider.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict
type NEFilterControlVerdict struct {
	NEFilterNewFlowVerdict
}

// NEFilterControlVerdictFrom constructs a [NEFilterControlVerdict] from an unsafe.Pointer.
//
// The result from a filter control provider.
func NEFilterControlVerdictFrom(ptr unsafe.Pointer) NEFilterControlVerdict {
	return NEFilterControlVerdict{
		NEFilterNewFlowVerdict: NEFilterNewFlowVerdictFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterControlVerdictClass) Alloc() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterControlVerdictClass) New() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterControlVerdict) Init() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterControlVerdict) Autorelease() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterControlVerdict creates a new NEFilterControlVerdict instance.
func NewNEFilterControlVerdict() NEFilterControlVerdict {
	return getNEFilterControlVerdictClass().New()
}




