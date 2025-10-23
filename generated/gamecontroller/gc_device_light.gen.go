// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCDeviceLight] class.
var (
	GCDeviceLightClass     _GCDeviceLightClass
	GCDeviceLightClassOnce sync.Once
)

func getGCDeviceLightClass() _GCDeviceLightClass {
	GCDeviceLightClassOnce.Do(func() {
		GCDeviceLightClass = _GCDeviceLightClass{objc.GetClass("GCDeviceLight")}
	})
	return GCDeviceLightClass
}

type _GCDeviceLightClass struct {
	class objc.Class
}

// An interface definition for the [GCDeviceLight] class.
type IGCDeviceLight interface {
	objectivec.IObject
	Color() IGCColor
	SetColor(value IGCColor)
}

// The colored light on a device.


// The colored light on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceLight
type GCDeviceLight struct {
	objectivec.Object
}

// GCDeviceLightFrom constructs a [GCDeviceLight] from an unsafe.Pointer.
//
// The colored light on a device.
func GCDeviceLightFrom(ptr unsafe.Pointer) GCDeviceLight {
	return GCDeviceLight{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCDeviceLightClass) Alloc() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCDeviceLightClass) New() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceLight) Init() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceLight) Autorelease() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceLight creates a new GCDeviceLight instance.
func NewGCDeviceLight() GCDeviceLight {
	return getGCDeviceLightClass().New()
}



// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicelight/color
func (g_ GCDeviceLight) Color() IGCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("color"))
	return rv
}


// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicelight/color
func (g_ GCDeviceLight) SetColor(value IGCColor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColor:"), value)
}



