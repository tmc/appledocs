// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlTargetStruct] class.
var (
	MTRAccessControlClusterAccessControlTargetStructClass     _MTRAccessControlClusterAccessControlTargetStructClass
	MTRAccessControlClusterAccessControlTargetStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlTargetStructClass() _MTRAccessControlClusterAccessControlTargetStructClass {
	MTRAccessControlClusterAccessControlTargetStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlTargetStructClass = _MTRAccessControlClusterAccessControlTargetStructClass{objc.GetClass("MTRAccessControlClusterAccessControlTargetStruct")}
	})
	return MTRAccessControlClusterAccessControlTargetStructClass
}

type _MTRAccessControlClusterAccessControlTargetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlTargetStruct] class.
type IMTRAccessControlClusterAccessControlTargetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlTargetStruct
type MTRAccessControlClusterAccessControlTargetStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlTargetStructFrom constructs a [MTRAccessControlClusterAccessControlTargetStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlTargetStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlTargetStruct {
	return MTRAccessControlClusterAccessControlTargetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlTargetStructClass) Alloc() MTRAccessControlClusterAccessControlTargetStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlTargetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlTargetStructClass) New() MTRAccessControlClusterAccessControlTargetStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlTargetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Init() MTRAccessControlClusterAccessControlTargetStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlTargetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Autorelease() MTRAccessControlClusterAccessControlTargetStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlTargetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlTargetStruct creates a new MTRAccessControlClusterAccessControlTargetStruct instance.
func NewMTRAccessControlClusterAccessControlTargetStruct() MTRAccessControlClusterAccessControlTargetStruct {
	return getMTRAccessControlClusterAccessControlTargetStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/cluster
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/cluster
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/devicetype
func (m_ MTRAccessControlClusterAccessControlTargetStruct) DeviceType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("deviceType"))
	return rv
}


// SetDeviceType sets the value of the deviceType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/devicetype
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetDeviceType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/endpoint
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/endpoint
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



