// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREndpointInfo */


/* debug [class_header]: Header for MTREndpointInfo */
// The class instance for the [MTREndpointInfo] class.
var (
	MTREndpointInfoClass     _MTREndpointInfoClass
	MTREndpointInfoClassOnce sync.Once
)

func getMTREndpointInfoClass() _MTREndpointInfoClass {
	MTREndpointInfoClassOnce.Do(func() {
		MTREndpointInfoClass = _MTREndpointInfoClass{objc.GetClass("MTREndpointInfo")}
	})
	return MTREndpointInfoClass
}

type _MTREndpointInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREndpointInfo */
// An interface definition for the [MTREndpointInfo] class.
type IMTREndpointInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREndpointInfo */
	// properties:
	Children() []MTREndpointInfo
	DeviceTypes() objc.IObject /* cross-framework: MTRDeviceTypeRevision */
	SetDeviceTypes(value objc.IObject /* cross-framework: MTRDeviceTypeRevision */)
	EndpointID() objc.IObject /* cross-framework: NSNumber */
	SetEndpointID(value objc.IObject /* cross-framework: NSNumber */)
	PartsList() objc.IObject /* cross-framework: NSNumber */
	SetPartsList(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREndpointInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREndpointInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTREndpointInfoClass) Alloc() MTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREndpointInfoClass) New() MTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREndpointInfo) Init() MTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREndpointInfo) Autorelease() MTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREndpointInfo creates a new MTREndpointInfo instance.
func NewMTREndpointInfo() MTREndpointInfo {
	return getMTREndpointInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREndpointInfo */
// Meta-data about an endpoint of a Matter node.


// Meta-data about an endpoint of a Matter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo
type MTREndpointInfo struct {
	objectivec.Object
}

// MTREndpointInfoFrom constructs a [MTREndpointInfo] from an unsafe.Pointer.
//
// Meta-data about an endpoint of a Matter node.
func MTREndpointInfoFrom(ptr unsafe.Pointer) MTREndpointInfo {
	return MTREndpointInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREndpointInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREndpointInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREndpointInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREndpointInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREndpointInfo */

// The direct children of this endpoint. This excludes indirect descendants even if they are listed in the PartsList attribute of this endpoint due to the Full-Family Pattern being used. Refer to Endpoint Composition Patterns in the Matter specification for details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo/children
func (m_ MTREndpointInfo) Children() []MTREndpointInfo {
	rv := objc.Send[[]MTREndpointInfo](m_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/devicetypes
func (m_ MTREndpointInfo) DeviceTypes() objc.IObject /* cross-framework: MTRDeviceTypeRevision */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("deviceTypes"))
	return rv
}/* debug [instance_properties/getter]: deviceTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/devicetypes
func (m_ MTREndpointInfo) SetDeviceTypes(value objc.IObject /* cross-framework: MTRDeviceTypeRevision */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceTypes:"), value)
}/* debug [instance_properties/setter]: deviceTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/endpointid
func (m_ MTREndpointInfo) EndpointID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpointID"))
	return rv
}/* debug [instance_properties/getter]: endpointID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/endpointid
func (m_ MTREndpointInfo) SetEndpointID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpointID:"), value)
}/* debug [instance_properties/setter]: endpointID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/partslist
func (m_ MTREndpointInfo) PartsList() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("partsList"))
	return rv
}/* debug [instance_properties/getter]: partsList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrendpointinfo/partslist
func (m_ MTREndpointInfo) SetPartsList(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPartsList:"), value)
}/* debug [instance_properties/setter]: partsList */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREndpointInfo */



