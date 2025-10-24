// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterAccessRestrictionEntryStruct */


/* debug [class_header]: Header for MTRAccessControlClusterAccessRestrictionEntryStruct */
// The class instance for the [MTRAccessControlClusterAccessRestrictionEntryStruct] class.
var (
	MTRAccessControlClusterAccessRestrictionEntryStructClass     _MTRAccessControlClusterAccessRestrictionEntryStructClass
	MTRAccessControlClusterAccessRestrictionEntryStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessRestrictionEntryStructClass() _MTRAccessControlClusterAccessRestrictionEntryStructClass {
	MTRAccessControlClusterAccessRestrictionEntryStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessRestrictionEntryStructClass = _MTRAccessControlClusterAccessRestrictionEntryStructClass{objc.GetClass("MTRAccessControlClusterAccessRestrictionEntryStruct")}
	})
	return MTRAccessControlClusterAccessRestrictionEntryStructClass
}

type _MTRAccessControlClusterAccessRestrictionEntryStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterAccessRestrictionEntryStruct */
// An interface definition for the [MTRAccessControlClusterAccessRestrictionEntryStruct] class.
type IMTRAccessControlClusterAccessRestrictionEntryStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterAccessRestrictionEntryStruct */
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	SetCluster(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterAccessRestrictionEntryStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterAccessRestrictionEntryStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessRestrictionEntryStructClass) Alloc() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterAccessRestrictionEntryStructClass) New() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Init() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Autorelease() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessRestrictionEntryStruct creates a new MTRAccessControlClusterAccessRestrictionEntryStruct instance.
func NewMTRAccessControlClusterAccessRestrictionEntryStruct() MTRAccessControlClusterAccessRestrictionEntryStruct {
	return getMTRAccessControlClusterAccessRestrictionEntryStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterAccessRestrictionEntryStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct
type MTRAccessControlClusterAccessRestrictionEntryStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessRestrictionEntryStructFrom constructs a [MTRAccessControlClusterAccessRestrictionEntryStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessRestrictionEntryStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessRestrictionEntryStruct {
	return MTRAccessControlClusterAccessRestrictionEntryStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterAccessRestrictionEntryStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterAccessRestrictionEntryStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterAccessRestrictionEntryStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterAccessRestrictionEntryStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterAccessRestrictionEntryStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}/* debug [instance_properties/getter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetCluster(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}/* debug [instance_properties/setter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionentrystruct/endpoint
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionentrystruct/endpoint
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionentrystruct/fabricindex
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionentrystruct/fabricindex
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterAccessRestrictionEntryStruct */



