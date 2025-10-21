// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRCommissioneeInfo] class.
var (
	MTRCommissioneeInfoClass     _MTRCommissioneeInfoClass
	MTRCommissioneeInfoClassOnce sync.Once
)

func getMTRCommissioneeInfoClass() _MTRCommissioneeInfoClass {
	MTRCommissioneeInfoClassOnce.Do(func() {
		MTRCommissioneeInfoClass = _MTRCommissioneeInfoClass{objc.GetClass("MTRCommissioneeInfo")}
	})
	return MTRCommissioneeInfoClass
}

type _MTRCommissioneeInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissioneeInfo] class.
type IMTRCommissioneeInfo interface {
	objectivec.IObject
}

// Information read from the commissionee device during commissioning.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo
type MTRCommissioneeInfo struct {
	objectivec.Object
}

// MTRCommissioneeInfoFrom constructs a [MTRCommissioneeInfo] from an unsafe.Pointer.
//
// Information read from the commissionee device during commissioning.
func MTRCommissioneeInfoFrom(ptr unsafe.Pointer) MTRCommissioneeInfo {
	return MTRCommissioneeInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissioneeInfoClass) Alloc() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissioneeInfoClass) New() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissioneeInfo) Init() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissioneeInfo) Autorelease() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissioneeInfo creates a new MTRCommissioneeInfo instance.
func NewMTRCommissioneeInfo() MTRCommissioneeInfo {
	return getMTRCommissioneeInfoClass().New()
}


// Endpoint information for all endpoints of the commissionee. Will be present only if readEndpointInformation is set to YES on MTRCommissioningParameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo/endpointsById
func (m_ MTRCommissioneeInfo) EndpointsById() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endpointsById"))
	return rv
}

// The product identity (VID / PID) of the commissionee.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo/productIdentity
func (m_ MTRCommissioneeInfo) ProductIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("productIdentity"))
	return rv
}

// Endpoint information for the root endpoint of the commissionee. Will be present only if readEndpointInformation is set to YES on MTRCommissioningParameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo/rootEndpoint
func (m_ MTRCommissioneeInfo) RootEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootEndpoint"))
	return rv
}



