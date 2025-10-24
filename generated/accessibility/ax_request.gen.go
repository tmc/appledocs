// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXRequest */


/* debug [class_header]: Header for AXRequest */
// The class instance for the [AXRequest] class.
var (
	AXRequestClass     _AXRequestClass
	AXRequestClassOnce sync.Once
)

func getAXRequestClass() _AXRequestClass {
	AXRequestClassOnce.Do(func() {
		AXRequestClass = _AXRequestClass{objc.GetClass("AXRequest")}
	})
	return AXRequestClass
}

type _AXRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXRequest */
// An interface definition for the [AXRequest] class.
type IAXRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXRequest */
	// properties:
	Technology() AXTechnology /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXRequest */
// Alloc allocates a new instance without initialization.
func (ac _AXRequestClass) Alloc() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXRequestClass) New() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXRequest) Init() AXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXRequest) Autorelease() AXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXRequest creates a new AXRequest instance.
func NewAXRequest() AXRequest {
	return getAXRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest
type AXRequest struct {
	objectivec.Object
}

// AXRequestFrom constructs a [AXRequest] from an unsafe.Pointer.
func AXRequestFrom(ptr unsafe.Pointer) AXRequest {
	return AXRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/current
func (ac _AXRequestClass) CurrentRequest() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("currentRequest"))
	return rv
}/* debug [class_properties_class/property]: currentRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/current
func (a_ AXRequest) CurrentRequest() IAXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("currentRequest"))
	return rv
}/* debug [instance_properties/getter]: currentRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/technology
func (a_ AXRequest) Technology() AXTechnology /* typedef */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("technology"))
	return rv
}/* debug [instance_properties/getter]: technology */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXRequest */



