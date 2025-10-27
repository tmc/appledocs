// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NibOutletConnector] class.
var (
	NibOutletConnectorClass     _NibOutletConnectorClass
	NibOutletConnectorClassOnce sync.Once
)

func getNibOutletConnectorClass() _NibOutletConnectorClass {
	NibOutletConnectorClassOnce.Do(func() {
		NibOutletConnectorClass = _NibOutletConnectorClass{objc.GetClass("NSNibOutletConnector")}
	})
	return NibOutletConnectorClass
}

type _NibOutletConnectorClass struct {
	class objc.Class
}





// An interface definition for the [NibOutletConnector] class.
type INibOutletConnector interface {
	INibConnector
	

	// properties:


	

	// methods:
	EstablishConnection()


}





// Alloc allocates a new instance without initialization.
func (nc _NibOutletConnectorClass) Alloc() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NibOutletConnectorClass) New() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NibOutletConnector) Init() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NibOutletConnector) Autorelease() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNibOutletConnector creates a new NibOutletConnector instance.
func NewNibOutletConnector() NibOutletConnector {
	return getNibOutletConnectorClass().New()
}





// An outlet connection between Interface Builder objects.


// An outlet connection between Interface Builder objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector
type NibOutletConnector struct {
	NibConnector
}

// NibOutletConnectorFrom constructs a [NibOutletConnector] from an unsafe.Pointer.
//
// An outlet connection between Interface Builder objects.
func NibOutletConnectorFrom(ptr unsafe.Pointer) NibOutletConnector {
	return NibOutletConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector/establishConnection
func (n_ NibOutletConnector) EstablishConnection() {
	objc.Send[objc.ID](n_.ID, objc.Sel("establishConnection"))
}













