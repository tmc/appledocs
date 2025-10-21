// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WaterSubmersionManager] class.
var (
	WaterSubmersionManagerClass     _WaterSubmersionManagerClass
	WaterSubmersionManagerClassOnce sync.Once
)

func getWaterSubmersionManagerClass() _WaterSubmersionManagerClass {
	WaterSubmersionManagerClassOnce.Do(func() {
		WaterSubmersionManagerClass = _WaterSubmersionManagerClass{objc.GetClass("CMWaterSubmersionManager")}
	})
	return WaterSubmersionManagerClass
}

type _WaterSubmersionManagerClass struct {
	class objc.Class
}

// An interface definition for the [WaterSubmersionManager] class.
type IWaterSubmersionManager interface {
	objectivec.IObject
}

// An object for managing the collection of pressure and temperature data during submersion.
//
// Use this class to receive live depth, water pressure, and water temperature data on Apple Watch Ultra. Start by assigning a usage description using the key in your app target’s information property list. You also need to include an entitlement to access the live submersion data. To access data for dives with a maximum depth of 6 m, add the Shallow Depth and Pressure capability to your app. For more information, see . To enable a maximum depth of 40 m, you must apply for the full Submerged Depth and Pressure entitlement. For more information, see . Next, check whether submersion data is available. If the property is , instantiate a object and assign a delegate. Your delegate then begins receiving updates from the system. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager
type WaterSubmersionManager struct {
	objectivec.Object
}

// WaterSubmersionManagerFrom constructs a [WaterSubmersionManager] from an unsafe.Pointer.
//
// An object for managing the collection of pressure and temperature data during submersion.
func WaterSubmersionManagerFrom(ptr unsafe.Pointer) WaterSubmersionManager {
	return WaterSubmersionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WaterSubmersionManagerClass) Alloc() WaterSubmersionManager {
	rv := objc.Send[WaterSubmersionManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WaterSubmersionManagerClass) New() WaterSubmersionManager {
	rv := objc.Send[WaterSubmersionManager](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterSubmersionManager) Init() WaterSubmersionManager {
	rv := objc.Send[WaterSubmersionManager](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterSubmersionManager) Autorelease() WaterSubmersionManager {
	rv := objc.Send[WaterSubmersionManager](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterSubmersionManager creates a new WaterSubmersionManager instance.
func NewWaterSubmersionManager() WaterSubmersionManager {
	return getWaterSubmersionManagerClass().New()
}


// A value indicating whether the app has user authorization to receive submersion data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/authorizationStatus
func (wc _WaterSubmersionManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(wc.class), objc.Sel("authorizationStatus"))
	return rv
}
// A Boolean value indicating whether the current device supports the submersion manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/waterSubmersionAvailable
func (wc _WaterSubmersionManagerClass) WaterSubmersionAvailable() bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("waterSubmersionAvailable"))
	return rv
}
// A value indicating whether the app has user authorization to receive submersion data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/authorizationStatus
func (w_ WaterSubmersionManager) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](w_.ID, objc.Sel("authorizationStatus"))
	return rv
}

// The object that receives updates about submersion data and events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/delegate
func (w_ WaterSubmersionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that receives updates about submersion data and events.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/delegate
func (w_ WaterSubmersionManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}

// The maximum depth supported by the water submersion manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/maximumDepth
func (w_ WaterSubmersionManager) MaximumDepth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("maximumDepth"))
	return rv
}

// A Boolean value indicating whether the current device supports the submersion manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/waterSubmersionAvailable
func (w_ WaterSubmersionManager) WaterSubmersionAvailable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("waterSubmersionAvailable"))
	return rv
}



