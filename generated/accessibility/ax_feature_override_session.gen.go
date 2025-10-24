// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXFeatureOverrideSession */


/* debug [class_header]: Header for AXFeatureOverrideSession */
// The class instance for the [AXFeatureOverrideSession] class.
var (
	AXFeatureOverrideSessionClass     _AXFeatureOverrideSessionClass
	AXFeatureOverrideSessionClassOnce sync.Once
)

func getAXFeatureOverrideSessionClass() _AXFeatureOverrideSessionClass {
	AXFeatureOverrideSessionClassOnce.Do(func() {
		AXFeatureOverrideSessionClass = _AXFeatureOverrideSessionClass{objc.GetClass("AXFeatureOverrideSession")}
	})
	return AXFeatureOverrideSessionClass
}

type _AXFeatureOverrideSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXFeatureOverrideSession */
// An interface definition for the [AXFeatureOverrideSession] class.
type IAXFeatureOverrideSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXFeatureOverrideSession */
	// properties:
	AXFeatureOverrideSessionErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXFeatureOverrideSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXFeatureOverrideSession */
// Alloc allocates a new instance without initialization.
func (ac _AXFeatureOverrideSessionClass) Alloc() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXFeatureOverrideSessionClass) New() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXFeatureOverrideSession) Init() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXFeatureOverrideSession) Autorelease() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXFeatureOverrideSession creates a new AXFeatureOverrideSession instance.
func NewAXFeatureOverrideSession() AXFeatureOverrideSession {
	return getAXFeatureOverrideSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXFeatureOverrideSession */
// A token object that represents an override session held by your app.


// A token object that represents an override session held by your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession
type AXFeatureOverrideSession struct {
	objectivec.Object
}

// AXFeatureOverrideSessionFrom constructs a [AXFeatureOverrideSession] from an unsafe.Pointer.
//
// A token object that represents an override session held by your app.
func AXFeatureOverrideSessionFrom(ptr unsafe.Pointer) AXFeatureOverrideSession {
	return AXFeatureOverrideSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXFeatureOverrideSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXFeatureOverrideSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXFeatureOverrideSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXFeatureOverrideSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXFeatureOverrideSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axfeatureoverridesessionerrordomain
func (a_ AXFeatureOverrideSession) AXFeatureOverrideSessionErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AXFeatureOverrideSessionErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: AXFeatureOverrideSessionErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXFeatureOverrideSession */



