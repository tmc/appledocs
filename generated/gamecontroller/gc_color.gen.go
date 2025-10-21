// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCColor] class.
var (
	GCColorClass     _GCColorClass
	GCColorClassOnce sync.Once
)

func getGCColorClass() _GCColorClass {
	GCColorClassOnce.Do(func() {
		GCColorClass = _GCColorClass{objc.GetClass("GCColor")}
	})
	return GCColorClass
}

type _GCColorClass struct {
	class objc.Class
}

// An interface definition for the [GCColor] class.
type IGCColor interface {
	objectivec.IObject
}

// The color of a device light.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor
type GCColor struct {
	objectivec.Object
}

// GCColorFrom constructs a [GCColor] from an unsafe.Pointer.
//
// The color of a device light.
func GCColorFrom(ptr unsafe.Pointer) GCColor {
	return GCColor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCColorClass) Alloc() GCColor {
	rv := objc.Send[GCColor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCColorClass) New() GCColor {
	rv := objc.Send[GCColor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCColor) Init() GCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCColor) Autorelease() GCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCColor creates a new GCColor instance.
func NewGCColor() GCColor {
	return getGCColorClass().New()
}




// Creates a color with the specified red, green, and blue values.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/init(red:green:blue:)
func NewGCColorWithRedGreenBlue(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer) GCColor {
	instance := getGCColorClass().Alloc()
	rv := objc.Send[GCColor](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	rv.Autorelease()
	return rv
}


// The normalized value of the blue component ranging from 0 to 1.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/blue
func (g_ GCColor) Blue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("blue"))
	return rv
}

// The normalized value of the green component ranging from 0 to 1.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/green
func (g_ GCColor) Green() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("green"))
	return rv
}


