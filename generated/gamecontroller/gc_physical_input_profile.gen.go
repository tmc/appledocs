// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCPhysicalInputProfile] class.
var (
	GCPhysicalInputProfileClass     _GCPhysicalInputProfileClass
	GCPhysicalInputProfileClassOnce sync.Once
)

func getGCPhysicalInputProfileClass() _GCPhysicalInputProfileClass {
	GCPhysicalInputProfileClassOnce.Do(func() {
		GCPhysicalInputProfileClass = _GCPhysicalInputProfileClass{objc.GetClass("GCPhysicalInputProfile")}
	})
	return GCPhysicalInputProfileClass
}

type _GCPhysicalInputProfileClass struct {
	class objc.Class
}

// An interface definition for the [GCPhysicalInputProfile] class.
type IGCPhysicalInputProfile interface {
	objectivec.IObject
	Capture() unsafe.Pointer
	MappedElementAliasForPhysicalInputName(inputName string) unsafe.Pointer
	MappedPhysicalInputNamesForElementAlias(elementAlias string) unsafe.Pointer
}

// The base class for controller profiles that support physical buttons, thumbsticks, and directional pads.
//
// This class provides properties and methods for accessing common elements of controllers, and for creating snapshots of profiles.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile
type GCPhysicalInputProfile struct {
	objectivec.Object
}

// GCPhysicalInputProfileFrom constructs a [GCPhysicalInputProfile] from an unsafe.Pointer.
//
// The base class for controller profiles that support physical buttons, thumbsticks, and directional pads.
func GCPhysicalInputProfileFrom(ptr unsafe.Pointer) GCPhysicalInputProfile {
	return GCPhysicalInputProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCPhysicalInputProfileClass) Alloc() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCPhysicalInputProfileClass) New() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCPhysicalInputProfile) Init() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCPhysicalInputProfile) Autorelease() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCPhysicalInputProfile creates a new GCPhysicalInputProfile instance.
func NewGCPhysicalInputProfile() GCPhysicalInputProfile {
	return getGCPhysicalInputProfileClass().New()
}


// Returns a snapshot of the profile with its current element values.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/capture()
func (g_ GCPhysicalInputProfile) Capture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("capture"))
	return rv
}

// Returns the name of the input element to which the user remaps the given physical element.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedElementAlias(forPhysicalInputName:)
func (g_ GCPhysicalInputProfile) MappedElementAliasForPhysicalInputName(inputName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mappedElementAliasForPhysicalInputName:"), objc.String(inputName))
	return rv
}

// Returns the physical input elements to which the user remaps the given input element.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedPhysicalInputNames(forElementAlias:)
func (g_ GCPhysicalInputProfile) MappedPhysicalInputNamesForElementAlias(elementAlias string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mappedPhysicalInputNamesForElementAlias:"), objc.String(elementAlias))
	return rv
}

// The axes in the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allAxes
func (g_ GCPhysicalInputProfile) AllAxes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allAxes"))
	return rv
}

// The buttons in the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allButtons
func (g_ GCPhysicalInputProfile) AllButtons() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allButtons"))
	return rv
}

// The directional pads in the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allDpads
func (g_ GCPhysicalInputProfile) AllDpads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allDpads"))
	return rv
}

// The elements in the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allElements
func (g_ GCPhysicalInputProfile) AllElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allElements"))
	return rv
}

// The touchpads in the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allTouchpads
func (g_ GCPhysicalInputProfile) AllTouchpads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allTouchpads"))
	return rv
}

// A Boolean value that indicates whether the user remaps elements in this profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/hasRemappedElements
func (g_ GCPhysicalInputProfile) HasRemappedElements() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasRemappedElements"))
	return rv
}



