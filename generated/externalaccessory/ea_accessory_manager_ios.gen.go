//go:build darwin && ios

// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for EAAccessoryManager


// Displays an alert that allows the user to pair the device with a Bluetooth accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/showBluetoothAccessoryPicker(withNameFilter:completion:)
func (e_ EAAccessoryManager) ShowBluetoothAccessoryPickerWithNameFilterCompletion(predicate foundation.Predicate, completion unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("showBluetoothAccessoryPickerWithNameFilter:completion:"), predicate, completion)
}

// iOS-only properties





