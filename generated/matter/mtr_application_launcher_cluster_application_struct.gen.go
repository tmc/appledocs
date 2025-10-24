// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterApplicationStruct */


/* debug [class_header]: Header for MTRApplicationLauncherClusterApplicationStruct */
// The class instance for the [MTRApplicationLauncherClusterApplicationStruct] class.
var (
	MTRApplicationLauncherClusterApplicationStructClass     _MTRApplicationLauncherClusterApplicationStructClass
	MTRApplicationLauncherClusterApplicationStructClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationStructClass() _MTRApplicationLauncherClusterApplicationStructClass {
	MTRApplicationLauncherClusterApplicationStructClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationStructClass = _MTRApplicationLauncherClusterApplicationStructClass{objc.GetClass("MTRApplicationLauncherClusterApplicationStruct")}
	})
	return MTRApplicationLauncherClusterApplicationStructClass
}

type _MTRApplicationLauncherClusterApplicationStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterApplicationStruct */
// An interface definition for the [MTRApplicationLauncherClusterApplicationStruct] class.
type IMTRApplicationLauncherClusterApplicationStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterApplicationStruct */
	// properties:
	ApplicationId() objc.IObject /* cross-framework: NSString */
	SetApplicationId(value objc.IObject /* cross-framework: NSString */)
	ApplicationID() objc.IObject /* cross-framework: NSString */
	SetApplicationID(value objc.IObject /* cross-framework: NSString */)
	CatalogVendorID() objc.IObject /* cross-framework: NSNumber */
	SetCatalogVendorID(value objc.IObject /* cross-framework: NSNumber */)
	CatalogVendorId() objc.IObject /* cross-framework: NSNumber */
	SetCatalogVendorId(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterApplicationStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterApplicationStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationStructClass) Alloc() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterApplicationStructClass) New() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationStruct) Init() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationStruct) Autorelease() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationStruct creates a new MTRApplicationLauncherClusterApplicationStruct instance.
func NewMTRApplicationLauncherClusterApplicationStruct() MTRApplicationLauncherClusterApplicationStruct {
	return getMTRApplicationLauncherClusterApplicationStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterApplicationStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct
type MTRApplicationLauncherClusterApplicationStruct struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterApplicationStructFrom constructs a [MTRApplicationLauncherClusterApplicationStruct] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationStructFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationStruct {
	return MTRApplicationLauncherClusterApplicationStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterApplicationStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterApplicationStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterApplicationStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterApplicationStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterApplicationStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/applicationId-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationId"))
	return rv
}/* debug [instance_properties/getter]: applicationId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/applicationId-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationId(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationId:"), value)
}/* debug [instance_properties/setter]: applicationId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/applicationID-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationID"))
	return rv
}/* debug [instance_properties/getter]: applicationID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/applicationID-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationID:"), value)
}/* debug [instance_properties/setter]: applicationID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/catalogVendorID-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("catalogVendorID"))
	return rv
}/* debug [instance_properties/getter]: catalogVendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/catalogVendorID-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorID:"), value)
}/* debug [instance_properties/setter]: catalogVendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/catalogVendorId-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("catalogVendorId"))
	return rv
}/* debug [instance_properties/getter]: catalogVendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct/catalogVendorId-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorId:"), value)
}/* debug [instance_properties/setter]: catalogVendorId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterApplicationStruct */



