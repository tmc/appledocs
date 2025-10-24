// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterAccessRestrictionStruct */


/* debug [class_header]: Header for MTRAccessControlClusterAccessRestrictionStruct */
// The class instance for the [MTRAccessControlClusterAccessRestrictionStruct] class.
var (
	MTRAccessControlClusterAccessRestrictionStructClass     _MTRAccessControlClusterAccessRestrictionStructClass
	MTRAccessControlClusterAccessRestrictionStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessRestrictionStructClass() _MTRAccessControlClusterAccessRestrictionStructClass {
	MTRAccessControlClusterAccessRestrictionStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessRestrictionStructClass = _MTRAccessControlClusterAccessRestrictionStructClass{objc.GetClass("MTRAccessControlClusterAccessRestrictionStruct")}
	})
	return MTRAccessControlClusterAccessRestrictionStructClass
}

type _MTRAccessControlClusterAccessRestrictionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterAccessRestrictionStruct */
// An interface definition for the [MTRAccessControlClusterAccessRestrictionStruct] class.
type IMTRAccessControlClusterAccessRestrictionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterAccessRestrictionStruct */
	// properties:
	Id() objc.IObject /* cross-framework: NSNumber */
	SetId(value objc.IObject /* cross-framework: NSNumber */)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterAccessRestrictionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterAccessRestrictionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessRestrictionStructClass) Alloc() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterAccessRestrictionStructClass) New() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Init() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Autorelease() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessRestrictionStruct creates a new MTRAccessControlClusterAccessRestrictionStruct instance.
func NewMTRAccessControlClusterAccessRestrictionStruct() MTRAccessControlClusterAccessRestrictionStruct {
	return getMTRAccessControlClusterAccessRestrictionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterAccessRestrictionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct
type MTRAccessControlClusterAccessRestrictionStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessRestrictionStructFrom constructs a [MTRAccessControlClusterAccessRestrictionStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessRestrictionStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessRestrictionStruct {
	return MTRAccessControlClusterAccessRestrictionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterAccessRestrictionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterAccessRestrictionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterAccessRestrictionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterAccessRestrictionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterAccessRestrictionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/id
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}/* debug [instance_properties/getter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/id
func (m_ MTRAccessControlClusterAccessRestrictionStruct) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}/* debug [instance_properties/setter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionstruct/type
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccessrestrictionstruct/type
func (m_ MTRAccessControlClusterAccessRestrictionStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterAccessRestrictionStruct */



