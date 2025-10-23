// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCVirtualControllerConfiguration] class.
var (
	GCVirtualControllerConfigurationClass     _GCVirtualControllerConfigurationClass
	GCVirtualControllerConfigurationClassOnce sync.Once
)

func getGCVirtualControllerConfigurationClass() _GCVirtualControllerConfigurationClass {
	GCVirtualControllerConfigurationClassOnce.Do(func() {
		GCVirtualControllerConfigurationClass = _GCVirtualControllerConfigurationClass{objc.GetClass("GCVirtualControllerConfiguration")}
	})
	return GCVirtualControllerConfigurationClass
}

type _GCVirtualControllerConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [GCVirtualControllerConfiguration] class.
type IGCVirtualControllerConfiguration interface {
	objectivec.IObject
	Hidden() bool
	SetHidden(value bool)
	Elements() string
	SetElements(value string)
	IsHidden() bool
	SetIsHidden(value bool)
}

// The configuration of a virtual controller.
//
// You configure a virtual controller by specifying the input elements it contains. Then using the method, you can customize individual elements.


// The configuration of a virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration
type GCVirtualControllerConfiguration struct {
	objectivec.Object
}

// GCVirtualControllerConfigurationFrom constructs a [GCVirtualControllerConfiguration] from an unsafe.Pointer.
//
// The configuration of a virtual controller.
func GCVirtualControllerConfigurationFrom(ptr unsafe.Pointer) GCVirtualControllerConfiguration {
	return GCVirtualControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerConfigurationClass) Alloc() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCVirtualControllerConfigurationClass) New() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCVirtualControllerConfiguration) Init() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCVirtualControllerConfiguration) Autorelease() GCVirtualControllerConfiguration {
	rv := objc.Send[GCVirtualControllerConfiguration](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCVirtualControllerConfiguration creates a new GCVirtualControllerConfiguration instance.
func NewGCVirtualControllerConfiguration() GCVirtualControllerConfiguration {
	return getGCVirtualControllerConfigurationClass().New()
}



// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration/isHidden
func (g_ GCVirtualControllerConfiguration) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}


// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration/isHidden
func (g_ GCVirtualControllerConfiguration) SetHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHidden:"), value)
}


// The input elements of a virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/elements
func (g_ GCVirtualControllerConfiguration) Elements() string {
	rv := objc.Send[string](g_.ID, objc.Sel("elements"))
	return rv
}


// The input elements of a virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/elements
func (g_ GCVirtualControllerConfiguration) SetElements(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setElements:"), objc.String(value))
}


// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/ishidden
func (g_ GCVirtualControllerConfiguration) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcvirtualcontroller/configuration/ishidden
func (g_ GCVirtualControllerConfiguration) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}



