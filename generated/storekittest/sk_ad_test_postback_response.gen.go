// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKAdTestPostbackResponse */


/* debug [class_header]: Header for SKAdTestPostbackResponse */
// The class instance for the [AdTestPostbackResponse] class.
var (
	AdTestPostbackResponseClass     _AdTestPostbackResponseClass
	AdTestPostbackResponseClassOnce sync.Once
)

func getAdTestPostbackResponseClass() _AdTestPostbackResponseClass {
	AdTestPostbackResponseClassOnce.Do(func() {
		AdTestPostbackResponseClass = _AdTestPostbackResponseClass{objc.GetClass("SKAdTestPostbackResponse")}
	})
	return AdTestPostbackResponseClass
}

type _AdTestPostbackResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AdTestPostbackResponse */
// An interface definition for the [AdTestPostbackResponse] class.
type IAdTestPostbackResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AdTestPostbackResponse */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AdTestPostbackResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AdTestPostbackResponse */
// Alloc allocates a new instance without initialization.
func (ac _AdTestPostbackResponseClass) Alloc() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdTestPostbackResponseClass) New() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestPostbackResponse) Init() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestPostbackResponse) Autorelease() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestPostbackResponse creates a new AdTestPostbackResponse instance.
func NewAdTestPostbackResponse() AdTestPostbackResponse {
	return getAdTestPostbackResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AdTestPostbackResponse */
// The status and error information for a postback that the system sends in the testing environment.


// The status and error information for a postback that the system sends in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse
type AdTestPostbackResponse struct {
	objectivec.Object
}

// AdTestPostbackResponseFrom constructs a [AdTestPostbackResponse] from an unsafe.Pointer.
//
// The status and error information for a postback that the system sends in the testing environment.
func AdTestPostbackResponseFrom(ptr unsafe.Pointer) AdTestPostbackResponse {
	return AdTestPostbackResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AdTestPostbackResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AdTestPostbackResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AdTestPostbackResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AdTestPostbackResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AdTestPostbackResponse */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKAdTestPostbackResponse */


