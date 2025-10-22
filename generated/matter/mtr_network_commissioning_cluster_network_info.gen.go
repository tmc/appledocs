// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterNetworkInfo] class.
var (
	MTRNetworkCommissioningClusterNetworkInfoClass     _MTRNetworkCommissioningClusterNetworkInfoClass
	MTRNetworkCommissioningClusterNetworkInfoClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterNetworkInfoClass() _MTRNetworkCommissioningClusterNetworkInfoClass {
	MTRNetworkCommissioningClusterNetworkInfoClassOnce.Do(func() {
		MTRNetworkCommissioningClusterNetworkInfoClass = _MTRNetworkCommissioningClusterNetworkInfoClass{objc.GetClass("MTRNetworkCommissioningClusterNetworkInfo")}
	})
	return MTRNetworkCommissioningClusterNetworkInfoClass
}

type _MTRNetworkCommissioningClusterNetworkInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterNetworkInfo] class.
type IMTRNetworkCommissioningClusterNetworkInfo interface {
	IMTRNetworkCommissioningClusterNetworkInfoStruct
	Connected() foundation.Number
	SetConnected(value foundation.INumber)
	NetworkID() foundation.Data
	SetNetworkID(value foundation.IData)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterNetworkInfo
type MTRNetworkCommissioningClusterNetworkInfo struct {
	MTRNetworkCommissioningClusterNetworkInfoStruct
}

// MTRNetworkCommissioningClusterNetworkInfoFrom constructs a [MTRNetworkCommissioningClusterNetworkInfo] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterNetworkInfoFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterNetworkInfo {
	return MTRNetworkCommissioningClusterNetworkInfo{
		MTRNetworkCommissioningClusterNetworkInfoStruct: MTRNetworkCommissioningClusterNetworkInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterNetworkInfoClass) Alloc() MTRNetworkCommissioningClusterNetworkInfo {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterNetworkInfoClass) New() MTRNetworkCommissioningClusterNetworkInfo {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterNetworkInfo) Init() MTRNetworkCommissioningClusterNetworkInfo {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterNetworkInfo) Autorelease() MTRNetworkCommissioningClusterNetworkInfo {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterNetworkInfo creates a new MTRNetworkCommissioningClusterNetworkInfo instance.
func NewMTRNetworkCommissioningClusterNetworkInfo() MTRNetworkCommissioningClusterNetworkInfo {
	return getMTRNetworkCommissioningClusterNetworkInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkinfo/connected
func (m_ MTRNetworkCommissioningClusterNetworkInfo) Connected() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("connected"))
	return rv
}


// SetConnected sets the value of the connected property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkinfo/connected
func (m_ MTRNetworkCommissioningClusterNetworkInfo) SetConnected(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConnected:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkinfo/networkid
func (m_ MTRNetworkCommissioningClusterNetworkInfo) NetworkID() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("networkID"))
	return rv
}


// SetNetworkID sets the value of the networkID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkinfo/networkid
func (m_ MTRNetworkCommissioningClusterNetworkInfo) SetNetworkID(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkID:"), value)
}



