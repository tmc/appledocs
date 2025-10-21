// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTableStruct
type MTRThreadNetworkDiagnosticsClusterNeighborTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Init() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct creates a new MTRThreadNetworkDiagnosticsClusterNeighborTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) MleFrameCounter() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mleFrameCounter"))
	return rv
}


// SetMleFrameCounter sets the value of the mleFrameCounter property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetMleFrameCounter(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMleFrameCounter:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Lqi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqi"))
	return rv
}


// SetLqi sets the value of the lqi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetLqi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) FullThreadDevice() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fullThreadDevice"))
	return rv
}


// SetFullThreadDevice sets the value of the fullThreadDevice property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetFullThreadDevice(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullThreadDevice:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) LinkFrameCounter() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("linkFrameCounter"))
	return rv
}


// SetLinkFrameCounter sets the value of the linkFrameCounter property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetLinkFrameCounter(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkFrameCounter:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) LastRssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lastRssi"))
	return rv
}


// SetLastRssi sets the value of the lastRssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetLastRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) RxOnWhenIdle() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rxOnWhenIdle"))
	return rv
}


// SetRxOnWhenIdle sets the value of the rxOnWhenIdle property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetRxOnWhenIdle(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRxOnWhenIdle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) IsChild() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("isChild"))
	return rv
}


// SetIsChild sets the value of the isChild property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetIsChild(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsChild:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) FullNetworkData() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fullNetworkData"))
	return rv
}


// SetFullNetworkData sets the value of the fullNetworkData property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetFullNetworkData(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullNetworkData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Rloc16() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rloc16"))
	return rv
}


// SetRloc16 sets the value of the rloc16 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetRloc16(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) FrameErrorRate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("frameErrorRate"))
	return rv
}


// SetFrameErrorRate sets the value of the frameErrorRate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetFrameErrorRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrameErrorRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Age() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("age"))
	return rv
}


// SetAge sets the value of the age property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetAge(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) MessageErrorRate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("messageErrorRate"))
	return rv
}


// SetMessageErrorRate sets the value of the messageErrorRate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetMessageErrorRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageErrorRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) ExtAddress() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extAddress"))
	return rv
}


// SetExtAddress sets the value of the extAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetExtAddress(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) AverageRssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("averageRssi"))
	return rv
}


// SetAverageRssi sets the value of the averageRssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortablestruct/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) SetAverageRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageRssi:"), value)
}



