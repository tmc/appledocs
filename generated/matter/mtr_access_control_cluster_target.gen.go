// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterTarget] class.
var (
	MTRAccessControlClusterTargetClass     _MTRAccessControlClusterTargetClass
	MTRAccessControlClusterTargetClassOnce sync.Once
)

func getMTRAccessControlClusterTargetClass() _MTRAccessControlClusterTargetClass {
	MTRAccessControlClusterTargetClassOnce.Do(func() {
		MTRAccessControlClusterTargetClass = _MTRAccessControlClusterTargetClass{objc.GetClass("MTRAccessControlClusterTarget")}
	})
	return MTRAccessControlClusterTargetClass
}

type _MTRAccessControlClusterTargetClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterTarget] class.
type IMTRAccessControlClusterTarget interface {
	IMTRAccessControlClusterAccessControlTargetStruct
	Cluster() foundation.Number
	SetCluster(value foundation.INumber)
	DeviceType() foundation.Number
	SetDeviceType(value foundation.INumber)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterTarget
type MTRAccessControlClusterTarget struct {
	MTRAccessControlClusterAccessControlTargetStruct
}

// MTRAccessControlClusterTargetFrom constructs a [MTRAccessControlClusterTarget] from an unsafe.Pointer.
func MTRAccessControlClusterTargetFrom(ptr unsafe.Pointer) MTRAccessControlClusterTarget {
	return MTRAccessControlClusterTarget{
		MTRAccessControlClusterAccessControlTargetStruct: MTRAccessControlClusterAccessControlTargetStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterTargetClass) Alloc() MTRAccessControlClusterTarget {
	rv := objc.Send[MTRAccessControlClusterTarget](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterTargetClass) New() MTRAccessControlClusterTarget {
	rv := objc.Send[MTRAccessControlClusterTarget](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterTarget) Init() MTRAccessControlClusterTarget {
	rv := objc.Send[MTRAccessControlClusterTarget](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterTarget) Autorelease() MTRAccessControlClusterTarget {
	rv := objc.Send[MTRAccessControlClusterTarget](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterTarget creates a new MTRAccessControlClusterTarget instance.
func NewMTRAccessControlClusterTarget() MTRAccessControlClusterTarget {
	return getMTRAccessControlClusterTargetClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/cluster
func (m_ MTRAccessControlClusterTarget) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/cluster
func (m_ MTRAccessControlClusterTarget) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/devicetype
func (m_ MTRAccessControlClusterTarget) DeviceType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("deviceType"))
	return rv
}


// SetDeviceType sets the value of the deviceType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/devicetype
func (m_ MTRAccessControlClusterTarget) SetDeviceType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/endpoint
func (m_ MTRAccessControlClusterTarget) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclustertarget/endpoint
func (m_ MTRAccessControlClusterTarget) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



