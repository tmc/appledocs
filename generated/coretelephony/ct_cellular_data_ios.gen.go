//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CellularData


// iOS-only properties

// A block that handles cellular data restriction state changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularData/cellularDataRestrictionDidUpdateNotifier
func (c_ CellularData) CellularDataRestrictionDidUpdateNotifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("cellularDataRestrictionDidUpdateNotifier"))
	return rv
}
func (c_ CellularData) SetCellularDataRestrictionDidUpdateNotifier(value objectivec.IObject) {
	c_.ID.Send(objc.RegisterName("setCellularDataRestrictionDidUpdateNotifier:"), value)
}

// The current state of cellular data restrictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularData/restrictedState
func (c_ CellularData) RestrictedState() CellularDataRestrictedState {
	rv := objc.Send[CellularDataRestrictedState](c_.ID, objc.Sel("restrictedState"))
	return rv
}





