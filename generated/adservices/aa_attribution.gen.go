// Code generated from Apple documentation for AdServices. DO NOT EDIT.

package adservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AAAttribution */


/* debug [class_header]: Header for AAAttribution */
// The class instance for the [AAAttribution] class.
var (
	AAAttributionClass     _AAAttributionClass
	AAAttributionClassOnce sync.Once
)

func getAAAttributionClass() _AAAttributionClass {
	AAAttributionClassOnce.Do(func() {
		AAAttributionClass = _AAAttributionClass{objc.GetClass("AAAttribution")}
	})
	return AAAttributionClass
}

type _AAAttributionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AAAttribution */
// An interface definition for the [AAAttribution] class.
type IAAAttribution interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AAAttribution */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AAAttribution */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AAAttribution */
// Alloc allocates a new instance without initialization.
func (ac _AAAttributionClass) Alloc() AAAttribution {
	rv := objc.Send[AAAttribution](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AAAttributionClass) New() AAAttribution {
	rv := objc.Send[AAAttribution](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AAAttribution) Init() AAAttribution {
	rv := objc.Send[AAAttribution](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AAAttribution) Autorelease() AAAttribution {
	rv := objc.Send[AAAttribution](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAAAttribution creates a new AAAttribution instance.
func NewAAAttribution() AAAttribution {
	return getAAAttributionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AAAttribution */
// The parent class that the framework uses to request a token.


// The parent class that the framework uses to request a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttribution
type AAAttribution struct {
	objectivec.Object
}

// AAAttributionFrom constructs a [AAAttribution] from an unsafe.Pointer.
//
// The parent class that the framework uses to request a token.
func AAAttributionFrom(ptr unsafe.Pointer) AAAttribution {
	return AAAttribution{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AAAttribution *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AAAttribution */

// Generates a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttribution/attributionToken()
func (ac _AAAttributionClass) AttributionTokenWithError(error_ unsafe.Pointer) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(ac.class), objc.Sel("attributionTokenWithError:"), error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributionTokenWithError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AAAttribution */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AAAttribution */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AAAttribution */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AAAttribution */






