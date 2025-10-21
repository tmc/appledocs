// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTREndpointInfo] class.
type IMTREndpointInfo interface {
	objectivec.IObject
}

// Meta-data about an endpoint of a Matter node.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTREndpointInfoClass) Alloc() MTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The direct children of this endpoint. This excludes indirect descendants even if they are listed in the PartsList attribute of this endpoint due to the Full-Family Pattern being used. Refer to Endpoint Composition Patterns in the Matter specification for details.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo/children
func (m_ MTREndpointInfo) Children() []MTREndpointInfo {
	rv := objc.Send[[]MTREndpointInfo](m_.ID, objc.Sel("children"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo/deviceTypes
func (m_ MTREndpointInfo) DeviceTypes() []MTRDeviceTypeRevision {
	rv := objc.Send[[]MTRDeviceTypeRevision](m_.ID, objc.Sel("deviceTypes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo/endpointID
func (m_ MTREndpointInfo) EndpointID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endpointID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREndpointInfo/partsList
func (m_ MTREndpointInfo) PartsList() []foundation.NSNumber {
	rv := objc.Send[[]foundation.NSNumber](m_.ID, objc.Sel("partsList"))
	return rv
}



