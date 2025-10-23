// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EAAccessoryManager] class.
var (
	EAAccessoryManagerClass     _EAAccessoryManagerClass
	EAAccessoryManagerClassOnce sync.Once
)

func getEAAccessoryManagerClass() _EAAccessoryManagerClass {
	EAAccessoryManagerClassOnce.Do(func() {
		EAAccessoryManagerClass = _EAAccessoryManagerClass{objc.GetClass("EAAccessoryManager")}
	})
	return EAAccessoryManagerClass
}

type _EAAccessoryManagerClass struct {
	class objc.Class
}

// An interface definition for the [EAAccessoryManager] class.
type IEAAccessoryManager interface {
	objectivec.IObject
	ShowBluetoothAccessoryPickerWithNameFilterCompletion(predicate foundation.IPredicate, completion unsafe.Pointer)
	EAAccessoryKey() string
	ConnectedAccessories() EAAccessory
	SetConnectedAccessories(value IEAAccessory)
	EAAccessorySelectedKey() string
	EABluetoothAccessoryPickerErrorDomain() string
}

// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
//
// An object coordinates the attached accessories for an iOS-based device. Use the shared accessory manager to retrieve a list of connected accessories, and start and stop the delivery of connection and disconnection notifications.


// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager
type EAAccessoryManager struct {
	objectivec.Object
}

// EAAccessoryManagerFrom constructs a [EAAccessoryManager] from an unsafe.Pointer.
//
// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
func EAAccessoryManagerFrom(ptr unsafe.Pointer) EAAccessoryManager {
	return EAAccessoryManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EAAccessoryManagerClass) Alloc() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAAccessoryManagerClass) New() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAAccessoryManager) Init() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAAccessoryManager) Autorelease() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAAccessoryManager creates a new EAAccessoryManager instance.
func NewEAAccessoryManager() EAAccessoryManager {
	return getEAAccessoryManagerClass().New()
}



// Displays an alert that allows the user to pair the device with a Bluetooth accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/showBluetoothAccessoryPicker(withNameFilter:completion:)
func (e_ EAAccessoryManager) ShowBluetoothAccessoryPickerWithNameFilterCompletion(predicate foundation.IPredicate, completion unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("showBluetoothAccessoryPickerWithNameFilter:completion:"), predicate, completion)
}


// A key that indicates the accessory object whose status changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessorykey
func (e_ EAAccessoryManager) EAAccessoryKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("EAAccessoryKey"))
	return rv
}


// The accessory objects corresponding to the list of currently connected accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessorymanager/connectedaccessories
func (e_ EAAccessoryManager) ConnectedAccessories() EAAccessory {
	rv := objc.Send[EAAccessory](e_.ID, objc.Sel("connectedAccessories"))
	return rv
}


// The accessory objects corresponding to the list of currently connected accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessorymanager/connectedaccessories
func (e_ EAAccessoryManager) SetConnectedAccessories(value IEAAccessory) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConnectedAccessories:"), value)
}


// A key that indicates the accessory object that the user selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessoryselectedkey
func (e_ EAAccessoryManager) EAAccessorySelectedKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("EAAccessorySelectedKey"))
	return rv
}


// The domain for errors passed to a Bluetooth picker completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eabluetoothaccessorypickererrordomain
func (e_ EAAccessoryManager) EABluetoothAccessoryPickerErrorDomain() string {
	rv := objc.Send[string](e_.ID, objc.Sel("EABluetoothAccessoryPickerErrorDomain"))
	return rv
}



