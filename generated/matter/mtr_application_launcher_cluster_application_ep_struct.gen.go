// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterApplicationEPStruct */


/* debug [class_header]: Header for MTRApplicationLauncherClusterApplicationEPStruct */
// The class instance for the [MTRApplicationLauncherClusterApplicationEPStruct] class.
var (
	MTRApplicationLauncherClusterApplicationEPStructClass     _MTRApplicationLauncherClusterApplicationEPStructClass
	MTRApplicationLauncherClusterApplicationEPStructClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationEPStructClass() _MTRApplicationLauncherClusterApplicationEPStructClass {
	MTRApplicationLauncherClusterApplicationEPStructClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationEPStructClass = _MTRApplicationLauncherClusterApplicationEPStructClass{objc.GetClass("MTRApplicationLauncherClusterApplicationEPStruct")}
	})
	return MTRApplicationLauncherClusterApplicationEPStructClass
}

type _MTRApplicationLauncherClusterApplicationEPStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterApplicationEPStruct */
// An interface definition for the [MTRApplicationLauncherClusterApplicationEPStruct] class.
type IMTRApplicationLauncherClusterApplicationEPStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterApplicationEPStruct */
	// properties:
	Application() IMTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterApplicationEPStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterApplicationEPStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationEPStructClass) Alloc() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterApplicationEPStructClass) New() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Init() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Autorelease() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationEPStruct creates a new MTRApplicationLauncherClusterApplicationEPStruct instance.
func NewMTRApplicationLauncherClusterApplicationEPStruct() MTRApplicationLauncherClusterApplicationEPStruct {
	return getMTRApplicationLauncherClusterApplicationEPStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterApplicationEPStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct
type MTRApplicationLauncherClusterApplicationEPStruct struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterApplicationEPStructFrom constructs a [MTRApplicationLauncherClusterApplicationEPStruct] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationEPStructFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationEPStruct {
	return MTRApplicationLauncherClusterApplicationEPStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterApplicationEPStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterApplicationEPStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterApplicationEPStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterApplicationEPStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterApplicationEPStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct/application
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Application() IMTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}/* debug [instance_properties/getter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct/application
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}/* debug [instance_properties/setter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterApplicationEPStruct */



