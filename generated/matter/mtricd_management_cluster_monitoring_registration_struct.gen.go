// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRICDManagementClusterMonitoringRegistrationStruct] class.
type IMTRICDManagementClusterMonitoringRegistrationStruct interface {
	objectivec.IObject
	// properties:
	CheckInNodeID() objc.IObject /* cross-framework: NSNumber */
	SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ClientType() objc.IObject /* cross-framework: NSNumber */
	SetClientType(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	MonitoredSubject() objc.IObject /* cross-framework: NSNumber */
	SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct
type MTRICDManagementClusterMonitoringRegistrationStruct struct {
	objectivec.Object
}

// MTRICDManagementClusterMonitoringRegistrationStructFrom constructs a [MTRICDManagementClusterMonitoringRegistrationStruct] from an unsafe.Pointer.
func MTRICDManagementClusterMonitoringRegistrationStructFrom(ptr unsafe.Pointer) MTRICDManagementClusterMonitoringRegistrationStruct {
	return MTRICDManagementClusterMonitoringRegistrationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterMonitoringRegistrationStructClass) Alloc() MTRICDManagementClusterMonitoringRegistrationStruct {
	rv := objc.Send[MTRICDManagementClusterMonitoringRegistrationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/checkInNodeID
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) CheckInNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/checkInNodeID
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/clientType
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) ClientType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clientType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/clientType
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetClientType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/monitoredSubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) MonitoredSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/monitoredSubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}



