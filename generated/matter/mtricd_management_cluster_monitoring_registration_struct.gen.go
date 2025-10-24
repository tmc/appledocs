// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterMonitoringRegistrationStruct */


/* debug [class_header]: Header for MTRICDManagementClusterMonitoringRegistrationStruct */
// The class instance for the [MTRICDManagementClusterMonitoringRegistrationStruct] class.
var (
	MTRICDManagementClusterMonitoringRegistrationStructClass     _MTRICDManagementClusterMonitoringRegistrationStructClass
	MTRICDManagementClusterMonitoringRegistrationStructClassOnce sync.Once
)

func getMTRICDManagementClusterMonitoringRegistrationStructClass() _MTRICDManagementClusterMonitoringRegistrationStructClass {
	MTRICDManagementClusterMonitoringRegistrationStructClassOnce.Do(func() {
		MTRICDManagementClusterMonitoringRegistrationStructClass = _MTRICDManagementClusterMonitoringRegistrationStructClass{objc.GetClass("MTRICDManagementClusterMonitoringRegistrationStruct")}
	})
	return MTRICDManagementClusterMonitoringRegistrationStructClass
}

type _MTRICDManagementClusterMonitoringRegistrationStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterMonitoringRegistrationStruct */
// An interface definition for the [MTRICDManagementClusterMonitoringRegistrationStruct] class.
type IMTRICDManagementClusterMonitoringRegistrationStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterMonitoringRegistrationStruct */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	CheckInNodeID() objc.IObject /* cross-framework: NSNumber */
	SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ClientType() objc.IObject /* cross-framework: NSNumber */
	SetClientType(value objc.IObject /* cross-framework: NSNumber */)
	MonitoredSubject() objc.IObject /* cross-framework: NSNumber */
	SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterMonitoringRegistrationStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterMonitoringRegistrationStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterMonitoringRegistrationStructClass) Alloc() MTRICDManagementClusterMonitoringRegistrationStruct {
	rv := objc.Send[MTRICDManagementClusterMonitoringRegistrationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterMonitoringRegistrationStructClass) New() MTRICDManagementClusterMonitoringRegistrationStruct {
	rv := objc.Send[MTRICDManagementClusterMonitoringRegistrationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) Init() MTRICDManagementClusterMonitoringRegistrationStruct {
	rv := objc.Send[MTRICDManagementClusterMonitoringRegistrationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) Autorelease() MTRICDManagementClusterMonitoringRegistrationStruct {
	rv := objc.Send[MTRICDManagementClusterMonitoringRegistrationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterMonitoringRegistrationStruct creates a new MTRICDManagementClusterMonitoringRegistrationStruct instance.
func NewMTRICDManagementClusterMonitoringRegistrationStruct() MTRICDManagementClusterMonitoringRegistrationStruct {
	return getMTRICDManagementClusterMonitoringRegistrationStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterMonitoringRegistrationStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct
type MTRICDManagementClusterMonitoringRegistrationStruct struct {
	objectivec.Object
}

// MTRICDManagementClusterMonitoringRegistrationStructFrom constructs a [MTRICDManagementClusterMonitoringRegistrationStruct] from an unsafe.Pointer.
func MTRICDManagementClusterMonitoringRegistrationStructFrom(ptr unsafe.Pointer) MTRICDManagementClusterMonitoringRegistrationStruct {
	return MTRICDManagementClusterMonitoringRegistrationStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterMonitoringRegistrationStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterMonitoringRegistrationStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterMonitoringRegistrationStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterMonitoringRegistrationStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterMonitoringRegistrationStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/checkinnodeid
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) CheckInNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}/* debug [instance_properties/getter]: checkInNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/checkinnodeid
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}/* debug [instance_properties/setter]: checkInNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/clienttype
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) ClientType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clientType"))
	return rv
}/* debug [instance_properties/getter]: clientType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/clienttype
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetClientType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}/* debug [instance_properties/setter]: clientType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/monitoredsubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) MonitoredSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}/* debug [instance_properties/getter]: monitoredSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclustermonitoringregistrationstruct/monitoredsubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}/* debug [instance_properties/setter]: monitoredSubject */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterMonitoringRegistrationStruct */



