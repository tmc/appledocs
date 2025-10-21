// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTable")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTable interface {
	IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable
type MTRThreadNetworkDiagnosticsClusterNeighborTable struct {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStruct
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTable] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNeighborTableFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return MTRThreadNetworkDiagnosticsClusterNeighborTable{
		MTRThreadNetworkDiagnosticsClusterNeighborTableStruct: MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Init() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTable creates a new MTRThreadNetworkDiagnosticsClusterNeighborTable instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTable() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Age() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("age"))
	return rv
}


// SetAge sets the value of the age property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAge(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) AverageRssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("averageRssi"))
	return rv
}


// SetAverageRssi sets the value of the averageRssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAverageRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) ExtAddress() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extAddress"))
	return rv
}


// SetExtAddress sets the value of the extAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetExtAddress(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FrameErrorRate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("frameErrorRate"))
	return rv
}


// SetFrameErrorRate sets the value of the frameErrorRate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFrameErrorRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrameErrorRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullNetworkData() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fullNetworkData"))
	return rv
}


// SetFullNetworkData sets the value of the fullNetworkData property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullNetworkData(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullNetworkData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullThreadDevice() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fullThreadDevice"))
	return rv
}


// SetFullThreadDevice sets the value of the fullThreadDevice property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullThreadDevice(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullThreadDevice:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) IsChild() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("isChild"))
	return rv
}


// SetIsChild sets the value of the isChild property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetIsChild(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsChild:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LastRssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lastRssi"))
	return rv
}


// SetLastRssi sets the value of the lastRssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLastRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LinkFrameCounter() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("linkFrameCounter"))
	return rv
}


// SetLinkFrameCounter sets the value of the linkFrameCounter property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLinkFrameCounter(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkFrameCounter:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Lqi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqi"))
	return rv
}


// SetLqi sets the value of the lqi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLqi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MessageErrorRate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("messageErrorRate"))
	return rv
}


// SetMessageErrorRate sets the value of the messageErrorRate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMessageErrorRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageErrorRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MleFrameCounter() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mleFrameCounter"))
	return rv
}


// SetMleFrameCounter sets the value of the mleFrameCounter property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMleFrameCounter(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMleFrameCounter:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Rloc16() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rloc16"))
	return rv
}


// SetRloc16 sets the value of the rloc16 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRloc16(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) RxOnWhenIdle() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rxOnWhenIdle"))
	return rv
}


// SetRxOnWhenIdle sets the value of the rxOnWhenIdle property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRxOnWhenIdle(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRxOnWhenIdle:"), value)
}



