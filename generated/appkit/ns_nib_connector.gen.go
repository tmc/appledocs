// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNibConnector */


/* debug [class_header]: Header for NSNibConnector */
// The class instance for the [NibConnector] class.
var (
	NibConnectorClass     _NibConnectorClass
	NibConnectorClassOnce sync.Once
)

func getNibConnectorClass() _NibConnectorClass {
	NibConnectorClassOnce.Do(func() {
		NibConnectorClass = _NibConnectorClass{objc.GetClass("NSNibConnector")}
	})
	return NibConnectorClass
}

type _NibConnectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NibConnector */
// An interface definition for the [NibConnector] class.
type INibConnector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NibConnector */
	// properties:
	Destination() objc.ID
	SetDestination(value objc.ID)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Source() objc.ID
	SetSource(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NibConnector */
	// methods:
	EstablishConnection()
	ReplaceObjectWithObject(oldObject objc.IObject, newObject objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NibConnector */
// Alloc allocates a new instance without initialization.
func (nc _NibConnectorClass) Alloc() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NibConnectorClass) New() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NibConnector) Init() NibConnector {
	rv := objc.Send[NibConnector](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NibConnector) Autorelease() NibConnector {
	rv := objc.Send[NibConnector](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNibConnector creates a new NibConnector instance.
func NewNibConnector() NibConnector {
	return getNibConnectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NibConnector */
// A connection between two nibs.


// A connection between two nibs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector
type NibConnector struct {
	objectivec.Object
}

// NibConnectorFrom constructs a [NibConnector] from an unsafe.Pointer.
//
// A connection between two nibs.
func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NibConnector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NibConnector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NibConnector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NibConnector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/establishConnection
func (n_ NibConnector) EstablishConnection() {
	objc.Send[objc.ID](n_.ID, objc.Sel("establishConnection"))
}/* debug [instance_methods/method]: EstablishConnection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/replaceObject:withObject:
func (n_ NibConnector) ReplaceObjectWithObject(oldObject objc.IObject, newObject objc.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("replaceObject:withObject:"), oldObject, newObject)
}/* debug [instance_methods/method]: ReplaceObjectWithObject */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NibConnector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/destination
func (n_ NibConnector) Destination() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/destination
func (n_ NibConnector) SetDestination(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/label
func (n_ NibConnector) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/label
func (n_ NibConnector) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/source
func (n_ NibConnector) Source() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/source
func (n_ NibConnector) SetSource(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSource:"), value)
}/* debug [instance_properties/setter]: source */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNibConnector */



