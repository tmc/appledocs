// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSNibControlConnector */


/* debug [class_header]: Header for NSNibControlConnector */
// The class instance for the [NibControlConnector] class.
var (
	NibControlConnectorClass     _NibControlConnectorClass
	NibControlConnectorClassOnce sync.Once
)

func getNibControlConnectorClass() _NibControlConnectorClass {
	NibControlConnectorClassOnce.Do(func() {
		NibControlConnectorClass = _NibControlConnectorClass{objc.GetClass("NSNibControlConnector")}
	})
	return NibControlConnectorClass
}

type _NibControlConnectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NibControlConnector */
// An interface definition for the [NibControlConnector] class.
type INibControlConnector interface {
	INibConnector
	
/* debug [class_interface_properties]: Properties for NibControlConnector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NibControlConnector */
	// methods:
	EstablishConnection()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NibControlConnector */
// Alloc allocates a new instance without initialization.
func (nc _NibControlConnectorClass) Alloc() NibControlConnector {
	rv := objc.Send[NibControlConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NibControlConnector */
// A control connection between two Interface Builder objects.


// A control connection between two Interface Builder objects.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NibControlConnector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NibControlConnector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NibControlConnector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NibControlConnector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibControlConnector/establishConnection
func (n_ NibControlConnector) EstablishConnection() {
	objc.Send[objc.ID](n_.ID, objc.Sel("establishConnection"))
}/* debug [instance_methods/method]: EstablishConnection */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NibControlConnector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNibControlConnector */



