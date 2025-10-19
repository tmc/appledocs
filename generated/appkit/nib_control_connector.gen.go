// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NibControlConnector] class.
var (
	nibControlConnectorClass     _NibControlConnectorClass
	nibControlConnectorClassOnce sync.Once
)

func getNibControlConnectorClass() _NibControlConnectorClass {
	nibControlConnectorClassOnce.Do(func() {
		nibControlConnectorClass = _NibControlConnectorClass{objc.GetClass("NSNibControlConnector")}
	})
	return nibControlConnectorClass
}

type _NibControlConnectorClass struct {
	class objc.Class
}

// An interface definition for the [NibControlConnector] class.
type INibControlConnector interface {
	INibConnector
}

// A control connection between two Interface Builder objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibControlConnector
type NibControlConnector struct {
	NibConnector
}

// NibControlConnectorFrom constructs a [NibControlConnector] from an unsafe.Pointer.
//
// A control connection between two Interface Builder objects.
func NibControlConnectorFrom(ptr unsafe.Pointer) NibControlConnector {
	return NibControlConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NibControlConnectorClass) Alloc() NibControlConnector {
	rv := objc.Send[NibControlConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NibControlConnectorClass) New() NibControlConnector {
	rv := objc.Send[NibControlConnector](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NibControlConnector) Init() NibControlConnector {
	rv := objc.Send[NibControlConnector](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NibControlConnector) Autorelease() NibControlConnector {
	rv := objc.Send[NibControlConnector](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNibControlConnector creates a new NibControlConnector instance.
func NewNibControlConnector() NibControlConnector {
	return getNibControlConnectorClass().New()
}




