// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */


/* debug [class_header]: Header for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
// The class instance for the [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] class.
var (
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass     _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClassOnce sync.Once
)

func getMTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass() _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass {
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClassOnce.Do(func() {
		MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass = _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass{objc.GetClass("MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct")}
	})
	return MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass
}

type _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
// An interface definition for the [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] class.
type IMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	SetCluster(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass) Alloc() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass) New() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Init() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Autorelease() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct creates a new MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct instance.
func NewMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	return getMTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct
type MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterCommissioningAccessRestrictionEntryStructFrom constructs a [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] from an unsafe.Pointer.
func MTRAccessControlClusterCommissioningAccessRestrictionEntryStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	return MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}/* debug [instance_properties/getter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) SetCluster(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}/* debug [instance_properties/setter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustercommissioningaccessrestrictionentrystruct/endpoint
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustercommissioningaccessrestrictionentrystruct/endpoint
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct */



