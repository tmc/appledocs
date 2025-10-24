// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCStylus] class.
var (
	GCStylusClass     _GCStylusClass
	GCStylusClassOnce sync.Once
)

func getGCStylusClass() _GCStylusClass {
	GCStylusClassOnce.Do(func() {
		GCStylusClass = _GCStylusClass{objc.GetClass("GCStylus")}
	})
	return GCStylusClass
}

type _GCStylusClass struct {
	class objc.Class
}

// An interface definition for the [GCStylus] class.
type IGCStylus interface {
	objectivec.IObject
	// properties:
	GCInputStylusPrimaryButton() objc.IObject /* cross-framework: NSString */
	SetGCInputStylusPrimaryButton(value objc.IObject /* cross-framework: NSString */)
	GCInputStylusSecondaryButton() objc.IObject /* cross-framework: NSString */
	SetGCInputStylusSecondaryButton(value objc.IObject /* cross-framework: NSString */)
	GCInputStylusTip() objc.IObject /* cross-framework: NSString */
	SetGCInputStylusTip(value objc.IObject /* cross-framework: NSString */)
	Input() unsafe.Pointer
	SetInput(value unsafe.Pointer)
	// methods:
}

// An object that represents a physical stylus connected to the device.
//
// Use the property to get the currently connect stylus accessories when your application starts. Register for and to get notified when a stylus connects of disconnects while your application is running. Check the to determine the type of stylus. A spatial stylus - capable of 6DoF tracking by Apple Vision Pro - has a category. Use the property to get the input profile of the stylus. A spatial stylus includes a pressure sensitive tip and an input cluster composed of two buttons. The primary button ( ) is the front button (closest to the stylus tip) in the input cluster of the stylus. This button is frequently used grab virtual objects. The secondary button ( ) is the middle button in the input cluster. It can measures pressure/force levels. It’s intended to be used for controlling in-air drawing, selection, and generic interactions. The tip is also represented as a button ( ). Use the property to get the haptics profile of the stylus. A spatial stylus may optionally support haptic feedback to a single locality - .


// An object that represents a physical stylus connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCStylus
type GCStylus struct {
	objectivec.Object
}

// GCStylusFrom constructs a [GCStylus] from an unsafe.Pointer.
//
// An object that represents a physical stylus connected to the device.
func GCStylusFrom(ptr unsafe.Pointer) GCStylus {
	return GCStylus{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCStylusClass) Alloc() GCStylus {
	rv := objc.Send[GCStylus](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCStylusClass) New() GCStylus {
	rv := objc.Send[GCStylus](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCStylus) Init() GCStylus {
	rv := objc.Send[GCStylus](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCStylus) Autorelease() GCStylus {
	rv := objc.Send[GCStylus](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCStylus creates a new GCStylus instance.
func NewGCStylus() GCStylus {
	return getGCStylusClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylusprimarybutton-18g2p
func (g_ GCStylus) GCInputStylusPrimaryButton() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputStylusPrimaryButton"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylusprimarybutton-18g2p
func (g_ GCStylus) SetGCInputStylusPrimaryButton(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputStylusPrimaryButton:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylussecondarybutton-6r3q
func (g_ GCStylus) GCInputStylusSecondaryButton() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputStylusSecondaryButton"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylussecondarybutton-6r3q
func (g_ GCStylus) SetGCInputStylusSecondaryButton(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputStylusSecondaryButton:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylustip-1rhuw
func (g_ GCStylus) GCInputStylusTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputStylusTip"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputstylustip-1rhuw
func (g_ GCStylus) SetGCInputStylusTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputStylusTip:"), value)
}


// Gets the input profile for the stylus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcstylus/input
func (g_ GCStylus) Input() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("input"))
	return rv
}


// Gets the input profile for the stylus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcstylus/input
func (g_ GCStylus) SetInput(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInput:"), value)
}


