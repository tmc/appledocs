// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterDataVerdict] class.
var (
	NEFilterDataVerdictClass     _NEFilterDataVerdictClass
	NEFilterDataVerdictClassOnce sync.Once
)

func getNEFilterDataVerdictClass() _NEFilterDataVerdictClass {
	NEFilterDataVerdictClassOnce.Do(func() {
		NEFilterDataVerdictClass = _NEFilterDataVerdictClass{objc.GetClass("NEFilterDataVerdict")}
	})
	return NEFilterDataVerdictClass
}

type _NEFilterDataVerdictClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterDataVerdict] class.
type INEFilterDataVerdict interface {
	INEFilterVerdict
}

// The result from a filter data provder for subsequent chunks of data on a flow.
//
// Return this verdict type from the various methods of .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict
type NEFilterDataVerdict struct {
	NEFilterVerdict
}

// NEFilterDataVerdictFrom constructs a [NEFilterDataVerdict] from an unsafe.Pointer.
//
// The result from a filter data provder for subsequent chunks of data on a flow.
func NEFilterDataVerdictFrom(ptr unsafe.Pointer) NEFilterDataVerdict {
	return NEFilterDataVerdict{
		NEFilterVerdict: NEFilterVerdictFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterDataVerdictClass) Alloc() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterDataVerdictClass) New() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterDataVerdict) Init() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterDataVerdict) Autorelease() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterDataVerdict creates a new NEFilterDataVerdict instance.
func NewNEFilterDataVerdict() NEFilterDataVerdict {
	return getNEFilterDataVerdictClass().New()
}




