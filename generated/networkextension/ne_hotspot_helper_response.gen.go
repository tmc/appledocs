// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotHelperResponse */


/* debug [class_header]: Header for NEHotspotHelperResponse */
// The class instance for the [NEHotspotHelperResponse] class.
var (
	NEHotspotHelperResponseClass     _NEHotspotHelperResponseClass
	NEHotspotHelperResponseClassOnce sync.Once
)

func getNEHotspotHelperResponseClass() _NEHotspotHelperResponseClass {
	NEHotspotHelperResponseClassOnce.Do(func() {
		NEHotspotHelperResponseClass = _NEHotspotHelperResponseClass{objc.GetClass("NEHotspotHelperResponse")}
	})
	return NEHotspotHelperResponseClass
}

type _NEHotspotHelperResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotHelperResponse */
// An interface definition for the [NEHotspotHelperResponse] class.
type INEHotspotHelperResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotHelperResponse */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotHelperResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotHelperResponse */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperResponseClass) Alloc() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEHotspotHelperResponseClass) New() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelperResponse) Init() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelperResponse) Autorelease() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelperResponse creates a new NEHotspotHelperResponse instance.
func NewNEHotspotHelperResponse() NEHotspotHelperResponse {
	return getNEHotspotHelperResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotHelperResponse */
// The hotspot helper’s response to a command.


// The hotspot helper’s response to a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResponse
type NEHotspotHelperResponse struct {
	objectivec.Object
}

// NEHotspotHelperResponseFrom constructs a [NEHotspotHelperResponse] from an unsafe.Pointer.
//
// The hotspot helper’s response to a command.
func NEHotspotHelperResponseFrom(ptr unsafe.Pointer) NEHotspotHelperResponse {
	return NEHotspotHelperResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotHelperResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotHelperResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotHelperResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotHelperResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotHelperResponse */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotHelperResponse */


