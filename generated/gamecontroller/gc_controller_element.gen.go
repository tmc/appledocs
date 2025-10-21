// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCControllerElement] class.
var (
	GCControllerElementClass     _GCControllerElementClass
	GCControllerElementClassOnce sync.Once
)

func getGCControllerElementClass() _GCControllerElementClass {
	GCControllerElementClassOnce.Do(func() {
		GCControllerElementClass = _GCControllerElementClass{objc.GetClass("GCControllerElement")}
	})
	return GCControllerElementClass
}

type _GCControllerElementClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerElement] class.
type IGCControllerElement interface {
	objectivec.IObject
}

// An input for a physical control, such as a button or thumbstick.
//
// is an abstract superclass for specific types of elements that represent controls on a game controller. Use the respective subclasses to either get the input of an element directly or set a handler that the element calls when the user changes a value. This class provides support for common features. For complex elements that have subelements, you can get the containing element using the property. For example, a direction pad ( ) has two axis control and four button subelements. If the user binds a controller element to a system gesture, the system sends the input to the system gesture recognizer first. If it doesn’t recognize a gesture, the system sends the input to your app but with a delay. If it does recognize a gesture, it doesn’t send any input to your app. To change this default behavior, you can set the property to to receive the input simultaneously without delay. Alternatively, set it to to disable the system gesture and receive the input exclusively. Use the property to check whether the user included an element in a system gesture. Use the property to determine whether an element’s input value is a range of values or a discrete digital value.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement
type GCControllerElement struct {
	objectivec.Object
}

// GCControllerElementFrom constructs a [GCControllerElement] from an unsafe.Pointer.
//
// An input for a physical control, such as a button or thumbstick.
func GCControllerElementFrom(ptr unsafe.Pointer) GCControllerElement {
	return GCControllerElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerElementClass) Alloc() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerElementClass) New() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerElement) Init() GCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerElement) Autorelease() GCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerElement creates a new GCControllerElement instance.
func NewGCControllerElement() GCControllerElement {
	return getGCControllerElementClass().New()
}


// The element’s aliases you use when accessing it with the subscript notation.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/aliases
func (g_ GCControllerElement) Aliases() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("aliases"))
	return rv
}

// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/isAnalog
func (g_ GCControllerElement) Analog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("analog"))
	return rv
}



