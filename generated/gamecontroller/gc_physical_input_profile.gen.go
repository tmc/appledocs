// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AllElements() unsafe.Pointer
	Dpads() foundation.IDictionary
	HasRemappedElements() bool
	LastEventTimestamp() float64
	ValueDidChangeHandler() unsafe.Pointer
	SetValueDidChangeHandler(value unsafe.Pointer)
	ExtendedGamepad() IGCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	Gamepad() unsafe.Pointer
	SetGamepad(value unsafe.Pointer)
	MicroGamepad() objc.IObject /* cross-framework: GCMicroGamepad */
	SetMicroGamepad(value objc.IObject /* cross-framework: GCMicroGamepad */)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	PhysicalInputProfile() IGCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
	AllAxes() objc.IObject /* cross-framework: GCControllerAxisInput */
	SetAllAxes(value objc.IObject /* cross-framework: GCControllerAxisInput */)
	AllButtons() objc.IObject /* cross-framework: GCControllerButtonInput */
	SetAllButtons(value objc.IObject /* cross-framework: GCControllerButtonInput */)
	AllDpads() objc.IObject /* cross-framework: GCControllerDirectionPad */
	SetAllDpads(value objc.IObject /* cross-framework: GCControllerDirectionPad */)
	AllTouchpads() IGCControllerTouchpad
	SetAllTouchpads(value IGCControllerTouchpad)
	Axes() objc.IObject /* cross-framework: GCControllerAxisInput */
	SetAxes(value objc.IObject /* cross-framework: GCControllerAxisInput */)
	Buttons() objc.IObject /* cross-framework: GCControllerButtonInput */
	SetButtons(value objc.IObject /* cross-framework: GCControllerButtonInput */)
	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	Elements() IGCControllerElement
	SetElements(value IGCControllerElement)
	Touchpads() IGCControllerTouchpad
	SetTouchpads(value IGCControllerTouchpad)
	// methods:
	MappedElementAliasForPhysicalInputName(inputName objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: String */
	MappedPhysicalInputNamesForElementAlias(elementAlias objc.IObject /* cross-framework: NSString */) unsafe.Pointer
}

// The base class for controller profiles that support physical buttons, thumbsticks, and directional pads.
//
// This class provides properties and methods for accessing common elements of controllers, and for creating snapshots of profiles.


// The base class for controller profiles that support physical buttons, thumbsticks, and directional pads.
//
// [Full Topic]
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



// Returns the name of the input element to which the user remaps the given physical element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedElementAlias(forPhysicalInputName:)
func (g_ GCPhysicalInputProfile) MappedElementAliasForPhysicalInputName(inputName objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](g_.ID, objc.Sel("mappedElementAliasForPhysicalInputName:"), inputName)
	return rv
}


// Returns the physical input elements to which the user remaps the given input element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedPhysicalInputNames(forElementAlias:)
func (g_ GCPhysicalInputProfile) MappedPhysicalInputNamesForElementAlias(elementAlias objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mappedPhysicalInputNamesForElementAlias:"), elementAlias)
	return rv
}


// The elements in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allElements
func (g_ GCPhysicalInputProfile) AllElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allElements"))
	return rv
}


// The directional pads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/dpads
func (g_ GCPhysicalInputProfile) Dpads() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("dpads"))
	return rv
}


// A Boolean value that indicates whether the user remaps elements in this profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/hasRemappedElements
func (g_ GCPhysicalInputProfile) HasRemappedElements() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasRemappedElements"))
	return rv
}


// The time of the most recent change to an element’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/lastEventTimestamp
func (g_ GCPhysicalInputProfile) LastEventTimestamp() float64 {
	rv := objc.Send[TimeInterval](g_.ID, objc.Sel("lastEventTimestamp"))
	return rv
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/valueDidChangeHandler
func (g_ GCPhysicalInputProfile) ValueDidChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueDidChangeHandler"))
	return rv
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/valueDidChangeHandler
func (g_ GCPhysicalInputProfile) SetValueDidChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueDidChangeHandler:"), value)
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) MicroGamepad() objc.IObject /* cross-framework: GCMicroGamepad */ {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) SetMicroGamepad(value objc.IObject /* cross-framework: GCMicroGamepad */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}


// The axes in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/allaxes
func (g_ GCPhysicalInputProfile) AllAxes() objc.IObject /* cross-framework: GCControllerAxisInput */ {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("allAxes"))
	return rv
}


// The axes in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/allaxes
func (g_ GCPhysicalInputProfile) SetAllAxes(value objc.IObject /* cross-framework: GCControllerAxisInput */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllAxes:"), value)
}


// The buttons in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/allbuttons
func (g_ GCPhysicalInputProfile) AllButtons() objc.IObject /* cross-framework: GCControllerButtonInput */ {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("allButtons"))
	return rv
}


// The buttons in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/allbuttons
func (g_ GCPhysicalInputProfile) SetAllButtons(value objc.IObject /* cross-framework: GCControllerButtonInput */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllButtons:"), value)
}


// The directional pads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/alldpads
func (g_ GCPhysicalInputProfile) AllDpads() objc.IObject /* cross-framework: GCControllerDirectionPad */ {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("allDpads"))
	return rv
}


// The directional pads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/alldpads
func (g_ GCPhysicalInputProfile) SetAllDpads(value objc.IObject /* cross-framework: GCControllerDirectionPad */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllDpads:"), value)
}


// The touchpads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/alltouchpads
func (g_ GCPhysicalInputProfile) AllTouchpads() IGCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("allTouchpads"))
	return rv
}


// The touchpads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/alltouchpads
func (g_ GCPhysicalInputProfile) SetAllTouchpads(value IGCControllerTouchpad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllTouchpads:"), value)
}


// The axes in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/axes
func (g_ GCPhysicalInputProfile) Axes() objc.IObject /* cross-framework: GCControllerAxisInput */ {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("axes"))
	return rv
}


// The axes in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/axes
func (g_ GCPhysicalInputProfile) SetAxes(value objc.IObject /* cross-framework: GCControllerAxisInput */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAxes:"), value)
}


// The buttons in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/buttons
func (g_ GCPhysicalInputProfile) Buttons() objc.IObject /* cross-framework: GCControllerButtonInput */ {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttons"))
	return rv
}


// The buttons in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/buttons
func (g_ GCPhysicalInputProfile) SetButtons(value objc.IObject /* cross-framework: GCControllerButtonInput */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtons:"), value)
}


// The physical device that the profile represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/device
func (g_ GCPhysicalInputProfile) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("device"))
	return rv
}


// The physical device that the profile represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/device
func (g_ GCPhysicalInputProfile) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDevice:"), value)
}


// The elements in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/elements
func (g_ GCPhysicalInputProfile) Elements() IGCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("elements"))
	return rv
}


// The elements in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/elements
func (g_ GCPhysicalInputProfile) SetElements(value IGCControllerElement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setElements:"), value)
}


// The touchpads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/touchpads
func (g_ GCPhysicalInputProfile) Touchpads() IGCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("touchpads"))
	return rv
}


// The touchpads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcphysicalinputprofile/touchpads
func (g_ GCPhysicalInputProfile) SetTouchpads(value IGCControllerTouchpad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchpads:"), value)
}



