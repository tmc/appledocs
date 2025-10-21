// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCVirtualControllerElementConfiguration] class.
var (
	GCVirtualControllerElementConfigurationClass     _GCVirtualControllerElementConfigurationClass
	GCVirtualControllerElementConfigurationClassOnce sync.Once
)

func getGCVirtualControllerElementConfigurationClass() _GCVirtualControllerElementConfigurationClass {
	GCVirtualControllerElementConfigurationClassOnce.Do(func() {
		GCVirtualControllerElementConfigurationClass = _GCVirtualControllerElementConfigurationClass{objc.GetClass("GCVirtualControllerElementConfiguration")}
	})
	return GCVirtualControllerElementConfigurationClass
}

type _GCVirtualControllerElementConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [GCVirtualControllerElementConfiguration] class.
type IGCVirtualControllerElementConfiguration interface {
	objectivec.IObject
}

// The properties of a virtual controller’s element that you can customize.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/ElementConfiguration
type GCVirtualControllerElementConfiguration struct {
	objectivec.Object
}

// GCVirtualControllerElementConfigurationFrom constructs a [GCVirtualControllerElementConfiguration] from an unsafe.Pointer.
//
// The properties of a virtual controller’s element that you can customize.
func GCVirtualControllerElementConfigurationFrom(ptr unsafe.Pointer) GCVirtualControllerElementConfiguration {
	return GCVirtualControllerElementConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerElementConfigurationClass) Alloc() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCVirtualControllerElementConfigurationClass) New() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCVirtualControllerElementConfiguration) Init() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCVirtualControllerElementConfiguration) Autorelease() GCVirtualControllerElementConfiguration {
	rv := objc.Send[GCVirtualControllerElementConfiguration](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCVirtualControllerElementConfiguration creates a new GCVirtualControllerElementConfiguration instance.
func NewGCVirtualControllerElementConfiguration() GCVirtualControllerElementConfiguration {
	return getGCVirtualControllerElementConfigurationClass().New()
}


// A Boolean value that determines whether the thumbstick element behaves as a touchpad.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/actsastouchpad
func (g_ GCVirtualControllerElementConfiguration) ActsAsTouchpad() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("actsAsTouchpad"))
	return rv
}


// SetActsAsTouchpad sets the value of the actsAsTouchpad property.
// A Boolean value that determines whether the thumbstick element behaves as a touchpad.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/actsastouchpad
func (g_ GCVirtualControllerElementConfiguration) SetActsAsTouchpad(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActsAsTouchpad:"), value)
}

// The Bezier path for the shape of an element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/path
func (g_ GCVirtualControllerElementConfiguration) Path() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The Bezier path for the shape of an element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/path
func (g_ GCVirtualControllerElementConfiguration) SetPath(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPath:"), value)
}

// A Boolean value that determines whether the virtual controller hides the element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/ishidden
func (g_ GCVirtualControllerElementConfiguration) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that determines whether the virtual controller hides the element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/elementconfiguration/ishidden
func (g_ GCVirtualControllerElementConfiguration) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}




