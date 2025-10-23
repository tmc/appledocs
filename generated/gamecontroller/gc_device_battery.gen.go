// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCDeviceBattery] class.
var (
	GCDeviceBatteryClass     _GCDeviceBatteryClass
	GCDeviceBatteryClassOnce sync.Once
)

func getGCDeviceBatteryClass() _GCDeviceBatteryClass {
	GCDeviceBatteryClassOnce.Do(func() {
		GCDeviceBatteryClass = _GCDeviceBatteryClass{objc.GetClass("GCDeviceBattery")}
	})
	return GCDeviceBatteryClass
}

type _GCDeviceBatteryClass struct {
	class objc.Class
}

// An interface definition for the [GCDeviceBattery] class.
type IGCDeviceBattery interface {
	objectivec.IObject
	BatteryLevel() float32
	BatteryState() GCDeviceBatteryState
}

// The charge level and state of a device’s battery.
//
// Use this class to display the state of a device’s battery to a player.


// The charge level and state of a device’s battery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery
type GCDeviceBattery struct {
	objectivec.Object
}

// GCDeviceBatteryFrom constructs a [GCDeviceBattery] from an unsafe.Pointer.
//
// The charge level and state of a device’s battery.
func GCDeviceBatteryFrom(ptr unsafe.Pointer) GCDeviceBattery {
	return GCDeviceBattery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCDeviceBatteryClass) Alloc() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCDeviceBatteryClass) New() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceBattery) Init() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceBattery) Autorelease() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceBattery creates a new GCDeviceBattery instance.
func NewGCDeviceBattery() GCDeviceBattery {
	return getGCDeviceBatteryClass().New()
}



// The charge level of a device’s battery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/batteryLevel
func (g_ GCDeviceBattery) BatteryLevel() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("batteryLevel"))
	return rv
}


// The state of a device’s battery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/batteryState
func (g_ GCDeviceBattery) BatteryState() GCDeviceBatteryState {
	rv := objc.Send[GCDeviceBatteryState](g_.ID, objc.Sel("batteryState"))
	return rv
}



