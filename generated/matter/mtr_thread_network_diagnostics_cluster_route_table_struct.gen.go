// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClass     _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass() _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableStructClass = _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTableStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTableStruct
type MTRThreadNetworkDiagnosticsClusterRouteTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return MTRThreadNetworkDiagnosticsClusterRouteTableStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) New() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Init() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct creates a new MTRThreadNetworkDiagnosticsClusterRouteTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LqiOut() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqiOut"))
	return rv
}


// SetLqiOut sets the value of the lqiOut property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLqiOut(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiOut:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Allocated() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("allocated"))
	return rv
}


// SetAllocated sets the value of the allocated property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetAllocated(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllocated:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) RouterId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("routerId"))
	return rv
}


// SetRouterId sets the value of the routerId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetRouterId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRouterId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Age() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("age"))
	return rv
}


// SetAge sets the value of the age property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetAge(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Rloc16() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rloc16"))
	return rv
}


// SetRloc16 sets the value of the rloc16 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetRloc16(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) NextHop() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nextHop"))
	return rv
}


// SetNextHop sets the value of the nextHop property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetNextHop(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextHop:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LqiIn() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqiIn"))
	return rv
}


// SetLqiIn sets the value of the lqiIn property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLqiIn(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiIn:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) ExtAddress() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extAddress"))
	return rv
}


// SetExtAddress sets the value of the extAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetExtAddress(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LinkEstablished() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("linkEstablished"))
	return rv
}


// SetLinkEstablished sets the value of the linkEstablished property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLinkEstablished(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkEstablished:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) PathCost() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("pathCost"))
	return rv
}


// SetPathCost sets the value of the pathCost property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetPathCost(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPathCost:"), value)
}



