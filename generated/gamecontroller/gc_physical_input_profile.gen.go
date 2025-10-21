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
	MappedElementAliasForPhysicalInputName(inputName string) string
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
func (g_ GCPhysicalInputProfile) MappedElementAliasForPhysicalInputName(inputName string) string {
	rv := objc.Send[string](g_.ID, objc.Sel("mappedElementAliasForPhysicalInputName:"), objc.String(inputName))
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

// The extended gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) ExtendedGamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// SetExtendedGamepad sets the value of the extendedGamepad property.
// The extended gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) SetExtendedGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}

// The gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// SetGamepad sets the value of the gamepad property.
// The gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}

// The micro gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) MicroGamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// SetMicroGamepad sets the value of the microGamepad property.
// The micro gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) SetMicroGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}

// The motion input profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) Motion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("motion"))
	return rv
}


// SetMotion sets the value of the motion property.
// The motion input profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) SetMotion(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}

// The physical input profile for the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) PhysicalInputProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// SetPhysicalInputProfile sets the value of the physicalInputProfile property.
// The physical input profile for the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) SetPhysicalInputProfile(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}

// The axes in the profile as key-value pairs for lookup by name.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/axes
func (g_ GCPhysicalInputProfile) Axes() string {
	rv := objc.Send[string](g_.ID, objc.Sel("axes"))
	return rv
}


// SetAxes sets the value of the axes property.
// The axes in the profile as key-value pairs for lookup by name.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/axes
func (g_ GCPhysicalInputProfile) SetAxes(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAxes:"), objc.String(value))
}

// The buttons in the profile as key-value pairs for lookup by name.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/buttons
func (g_ GCPhysicalInputProfile) Buttons() string {
	rv := objc.Send[string](g_.ID, objc.Sel("buttons"))
	return rv
}


// SetButtons sets the value of the buttons property.
// The buttons in the profile as key-value pairs for lookup by name.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/buttons
func (g_ GCPhysicalInputProfile) SetButtons(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtons:"), objc.String(value))
}

// The physical device that the profile represents.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/device
func (g_ GCPhysicalInputProfile) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The physical device that the profile represents.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/device
func (g_ GCPhysicalInputProfile) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDevice:"), value)
}

// The directional pads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/dpads
func (g_ GCPhysicalInputProfile) Dpads() string {
	rv := objc.Send[string](g_.ID, objc.Sel("dpads"))
	return rv
}


// SetDpads sets the value of the dpads property.
// The directional pads in the profile as key-value pairs for lookup by name.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/dpads
func (g_ GCPhysicalInputProfile) SetDpads(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDpads:"), objc.String(value))
}

// The elements in the profile as key-value pairs for lookup by name.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/elements
func (g_ GCPhysicalInputProfile) Elements() string {
	rv := objc.Send[string](g_.ID, objc.Sel("elements"))
	return rv
}


// SetElements sets the value of the elements property.
// The elements in the profile as key-value pairs for lookup by name.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/elements
func (g_ GCPhysicalInputProfile) SetElements(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setElements:"), objc.String(value))
}

// The time of the most recent change to an element’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/lasteventtimestamp
func (g_ GCPhysicalInputProfile) LastEventTimestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("lastEventTimestamp"))
	return rv
}


// SetLastEventTimestamp sets the value of the lastEventTimestamp property.
// The time of the most recent change to an element’s value.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/lasteventtimestamp
func (g_ GCPhysicalInputProfile) SetLastEventTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLastEventTimestamp:"), value)
}

// The touchpads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/touchpads
func (g_ GCPhysicalInputProfile) Touchpads() string {
	rv := objc.Send[string](g_.ID, objc.Sel("touchpads"))
	return rv
}


// SetTouchpads sets the value of the touchpads property.
// The touchpads in the profile as key-value pairs for lookup by name.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/touchpads
func (g_ GCPhysicalInputProfile) SetTouchpads(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchpads:"), objc.String(value))
}

// The block that the profile calls when an element’s value changes.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/valuedidchangehandler
func (g_ GCPhysicalInputProfile) ValueDidChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueDidChangeHandler"))
	return rv
}


// SetValueDidChangeHandler sets the value of the valueDidChangeHandler property.
// The block that the profile calls when an element’s value changes.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/valuedidchangehandler
func (g_ GCPhysicalInputProfile) SetValueDidChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueDidChangeHandler:"), value)
}



