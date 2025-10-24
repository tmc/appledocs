// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterApplicationEP */


/* debug [class_header]: Header for MTRApplicationLauncherClusterApplicationEP */
// The class instance for the [MTRApplicationLauncherClusterApplicationEP] class.
var (
	MTRApplicationLauncherClusterApplicationEPClass     _MTRApplicationLauncherClusterApplicationEPClass
	MTRApplicationLauncherClusterApplicationEPClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationEPClass() _MTRApplicationLauncherClusterApplicationEPClass {
	MTRApplicationLauncherClusterApplicationEPClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationEPClass = _MTRApplicationLauncherClusterApplicationEPClass{objc.GetClass("MTRApplicationLauncherClusterApplicationEP")}
	})
	return MTRApplicationLauncherClusterApplicationEPClass
}

type _MTRApplicationLauncherClusterApplicationEPClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterApplicationEP */
// An interface definition for the [MTRApplicationLauncherClusterApplicationEP] class.
type IMTRApplicationLauncherClusterApplicationEP interface {
	IMTRApplicationLauncherClusterApplicationEPStruct
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterApplicationEP */
	// properties:
	Application() IMTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterApplicationEP */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterApplicationEP */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationEPClass) Alloc() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterApplicationEPClass) New() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationEP) Init() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationEP) Autorelease() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationEP creates a new MTRApplicationLauncherClusterApplicationEP instance.
func NewMTRApplicationLauncherClusterApplicationEP() MTRApplicationLauncherClusterApplicationEP {
	return getMTRApplicationLauncherClusterApplicationEPClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterApplicationEP */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP
type MTRApplicationLauncherClusterApplicationEP struct {
	MTRApplicationLauncherClusterApplicationEPStruct
}

// MTRApplicationLauncherClusterApplicationEPFrom constructs a [MTRApplicationLauncherClusterApplicationEP] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationEPFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationEP {
	return MTRApplicationLauncherClusterApplicationEP{
		MTRApplicationLauncherClusterApplicationEPStruct: MTRApplicationLauncherClusterApplicationEPStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterApplicationEP *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterApplicationEP */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterApplicationEP */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterApplicationEP */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterApplicationEP */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP/application
func (m_ MTRApplicationLauncherClusterApplicationEP) Application() IMTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}/* debug [instance_properties/getter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP/application
func (m_ MTRApplicationLauncherClusterApplicationEP) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}/* debug [instance_properties/setter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEP) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEP) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterApplicationEP */



