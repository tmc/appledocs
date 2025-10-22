// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CellularData] class.
var (
	CellularDataClass     _CellularDataClass
	CellularDataClassOnce sync.Once
)

func getCellularDataClass() _CellularDataClass {
	CellularDataClassOnce.Do(func() {
		CellularDataClass = _CellularDataClass{objc.GetClass("CTCellularData")}
	})
	return CellularDataClass
}

type _CellularDataClass struct {
	class objc.Class
}

// An interface definition for the [CellularData] class.
type ICellularData interface {
	objectivec.IObject
	CellularDataRestrictionDidUpdateNotifier() unsafe.Pointer
	SetCellularDataRestrictionDidUpdateNotifier(value unsafe.Pointer)
	RestrictedState() CellularDataRestrictedState
	SetRestrictedState(value CellularDataRestrictedState)
}

// An object indicating whether the app can access cellular data.
//
// This property represents all access to cellular data. If the is , the app cannot use the cellular network.


// An object indicating whether the app can access cellular data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularData

type CellularData struct {
	objectivec.Object
}

// CellularDataFrom constructs a [CellularData] from an unsafe.Pointer.
//
// An object indicating whether the app can access cellular data.
func CellularDataFrom(ptr unsafe.Pointer) CellularData {
	return CellularData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CellularDataClass) Alloc() CellularData {
	rv := objc.Send[CellularData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CellularDataClass) New() CellularData {
	rv := objc.Send[CellularData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularData) Init() CellularData {
	rv := objc.Send[CellularData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularData) Autorelease() CellularData {
	rv := objc.Send[CellularData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularData creates a new CellularData instance.
func NewCellularData() CellularData {
	return getCellularDataClass().New()
}



// A block that handles cellular data restriction state changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellulardata/cellulardatarestrictiondidupdatenotifier

func (c_ CellularData) CellularDataRestrictionDidUpdateNotifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cellularDataRestrictionDidUpdateNotifier"))
	return rv
}


// A block that handles cellular data restriction state changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellulardata/cellulardatarestrictiondidupdatenotifier

func (c_ CellularData) SetCellularDataRestrictionDidUpdateNotifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellularDataRestrictionDidUpdateNotifier:"), value)
}


// The current state of cellular data restrictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellulardata/restrictedstate

func (c_ CellularData) RestrictedState() CellularDataRestrictedState {
	rv := objc.Send[CellularDataRestrictedState](c_.ID, objc.Sel("restrictedState"))
	return rv
}


// The current state of cellular data restrictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellulardata/restrictedstate

func (c_ CellularData) SetRestrictedState(value CellularDataRestrictedState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRestrictedState:"), value)
}



