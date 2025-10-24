// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSNibOutletConnector */


/* debug [class_header]: Header for NSNibOutletConnector */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NibOutletConnector */
// An interface definition for the [NibOutletConnector] class.
type INibOutletConnector interface {
	INibConnector
	
/* debug [class_interface_properties]: Properties for NibOutletConnector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NibOutletConnector */
	// methods:
	EstablishConnection()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NibOutletConnector */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NibOutletConnector */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NibOutletConnector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NibOutletConnector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NibOutletConnector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NibOutletConnector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector/establishConnection
func (n_ NibOutletConnector) EstablishConnection() {
	objc.Send[objc.ID](n_.ID, objc.Sel("establishConnection"))
}/* debug [instance_methods/method]: EstablishConnection */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NibOutletConnector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNibOutletConnector */



