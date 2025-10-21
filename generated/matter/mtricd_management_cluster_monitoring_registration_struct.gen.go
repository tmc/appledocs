// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/checkInNodeID
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) CheckInNodeID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}


// SetCheckInNodeID sets the value of the checkInNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/checkInNodeID
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetCheckInNodeID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/clientType
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) ClientType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("clientType"))
	return rv
}


// SetClientType sets the value of the clientType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/clientType
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetClientType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) FabricIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/fabricIndex
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetFabricIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/monitoredSubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) MonitoredSubject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}


// SetMonitoredSubject sets the value of the monitoredSubject property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterMonitoringRegistrationStruct/monitoredSubject
func (m_ MTRICDManagementClusterMonitoringRegistrationStruct) SetMonitoredSubject(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}


