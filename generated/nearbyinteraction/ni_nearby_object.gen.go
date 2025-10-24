// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NINearbyObject */


/* debug [class_header]: Header for NINearbyObject */
// The class instance for the [NINearbyObject] class.
var (
	NINearbyObjectClass     _NINearbyObjectClass
	NINearbyObjectClassOnce sync.Once
)

func getNINearbyObjectClass() _NINearbyObjectClass {
	NINearbyObjectClassOnce.Do(func() {
		NINearbyObjectClass = _NINearbyObjectClass{objc.GetClass("NINearbyObject")}
	})
	return NINearbyObjectClass
}

type _NINearbyObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NINearbyObject */
// An interface definition for the [NINearbyObject] class.
type ININearbyObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NINearbyObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NINearbyObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NINearbyObject */
// Alloc allocates a new instance without initialization.
func (nc _NINearbyObjectClass) Alloc() NINearbyObject {
	rv := objc.Send[NINearbyObject](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NINearbyObjectClass) New() NINearbyObject {
	rv := objc.Send[NINearbyObject](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NINearbyObject) Init() NINearbyObject {
	rv := objc.Send[NINearbyObject](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NINearbyObject) Autorelease() NINearbyObject {
	rv := objc.Send[NINearbyObject](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNINearbyObject creates a new NINearbyObject instance.
func NewNINearbyObject() NINearbyObject {
	return getNINearbyObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NINearbyObject */
// Location information for a peer device in an interaction session.
//
// A nearby object refers to a peer Apple device or third-party accessory. When the framework is ready to provide your app with information about a nearby object’s relative position, it calls your delegate’s implementation. If a session can’t provide peer direction or distance, it sets the values to . In Objective-C, the session uses the and values to indicate missing direction or distance. For more information, see .


// Location information for a peer device in an interaction session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject
type NINearbyObject struct {
	objectivec.Object
}

// NINearbyObjectFrom constructs a [NINearbyObject] from an unsafe.Pointer.
//
// Location information for a peer device in an interaction session.
func NINearbyObjectFrom(ptr unsafe.Pointer) NINearbyObject {
	return NINearbyObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NINearbyObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NINearbyObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NINearbyObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NINearbyObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NINearbyObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NINearbyObject */


