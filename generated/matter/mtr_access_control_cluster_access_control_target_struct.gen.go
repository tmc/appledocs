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
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	SetCluster(value objc.IObject /* cross-framework: NSNumber */)
	DeviceType() objc.IObject /* cross-framework: NSNumber */
	SetDeviceType(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/cluster
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/cluster
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetCluster(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/devicetype
func (m_ MTRAccessControlClusterAccessControlTargetStruct) DeviceType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("deviceType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/devicetype
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetDeviceType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/endpoint
func (m_ MTRAccessControlClusterAccessControlTargetStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontroltargetstruct/endpoint
func (m_ MTRAccessControlClusterAccessControlTargetStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



