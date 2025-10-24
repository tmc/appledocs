// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCPhysicalInputProfile */


/* debug [class_header]: Header for GCPhysicalInputProfile */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCPhysicalInputProfile */
// An interface definition for the [GCPhysicalInputProfile] class.
type IGCPhysicalInputProfile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCPhysicalInputProfile */
	// properties:
	AllAxes() unsafe.Pointer
	AllButtons() unsafe.Pointer
	AllDpads() unsafe.Pointer
	AllElements() unsafe.Pointer
	AllTouchpads() unsafe.Pointer
	Axes() foundation.IDictionary
	Buttons() foundation.IDictionary
	Device() unsafe.Pointer
	Dpads() foundation.IDictionary
	Elements() foundation.IDictionary
	HasRemappedElements() bool
	LastEventTimestamp() float64
	Touchpads() foundation.IDictionary
	ValueDidChangeHandler() unsafe.Pointer
	SetValueDidChangeHandler(value unsafe.Pointer)
	ExtendedGamepad() IGCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	Gamepad() IGCGamepad
	SetGamepad(value IGCGamepad)
	MicroGamepad() IGCMicroGamepad
	SetMicroGamepad(value IGCMicroGamepad)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	PhysicalInputProfile() IGCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCPhysicalInputProfile */
	// methods:
	Capture() unsafe.Pointer
	MappedElementAliasForPhysicalInputName(inputName objc.IObject /* cross-framework: NSString */) foundation.String
	MappedPhysicalInputNamesForElementAlias(elementAlias objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	SetStateFromPhysicalInput(physicalInput IGCPhysicalInputProfile)
	ObjectForKeyedSubscript(key objc.IObject /* cross-framework: NSString */) IGCControllerElement
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCPhysicalInputProfile */
// Alloc allocates a new instance without initialization.
func (gc _GCPhysicalInputProfileClass) Alloc() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCPhysicalInputProfile */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCPhysicalInputProfile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCPhysicalInputProfile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCPhysicalInputProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCPhysicalInputProfile */

// Returns a snapshot of the profile with its current element values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/capture()
func (g_ GCPhysicalInputProfile) Capture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("capture"))
	return rv
}/* debug [instance_methods/method]: Capture */


// Returns the name of the input element to which the user remaps the given physical element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedElementAlias(forPhysicalInputName:)
func (g_ GCPhysicalInputProfile) MappedElementAliasForPhysicalInputName(inputName objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](g_.ID, objc.Sel("mappedElementAliasForPhysicalInputName:"), inputName)
	return rv
}/* debug [instance_methods/method]: MappedElementAliasForPhysicalInputName */


// Returns the physical input elements to which the user remaps the given input element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedPhysicalInputNames(forElementAlias:)
func (g_ GCPhysicalInputProfile) MappedPhysicalInputNamesForElementAlias(elementAlias objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mappedPhysicalInputNamesForElementAlias:"), elementAlias)
	return rv
}/* debug [instance_methods/method]: MappedPhysicalInputNamesForElementAlias */


// Copies the input values from a specified physical input profile to a snapshot of the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/setStateFromPhysicalInput(_:)
func (g_ GCPhysicalInputProfile) SetStateFromPhysicalInput(physicalInput IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromPhysicalInput:"), physicalInput)
}/* debug [instance_methods/method]: SetStateFromPhysicalInput */


// Returns the element that the key specifies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/subscript(_:)
func (g_ GCPhysicalInputProfile) ObjectForKeyedSubscript(key objc.IObject /* cross-framework: NSString */) IGCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCPhysicalInputProfile */

// The axes in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allAxes
func (g_ GCPhysicalInputProfile) AllAxes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allAxes"))
	return rv
}/* debug [instance_properties/getter]: allAxes */


// The buttons in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allButtons
func (g_ GCPhysicalInputProfile) AllButtons() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allButtons"))
	return rv
}/* debug [instance_properties/getter]: allButtons */


// The directional pads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allDpads
func (g_ GCPhysicalInputProfile) AllDpads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allDpads"))
	return rv
}/* debug [instance_properties/getter]: allDpads */


// The elements in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allElements
func (g_ GCPhysicalInputProfile) AllElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allElements"))
	return rv
}/* debug [instance_properties/getter]: allElements */


// The touchpads in the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allTouchpads
func (g_ GCPhysicalInputProfile) AllTouchpads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allTouchpads"))
	return rv
}/* debug [instance_properties/getter]: allTouchpads */


// The axes in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/axes
func (g_ GCPhysicalInputProfile) Axes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("axes"))
	return rv
}/* debug [instance_properties/getter]: axes */


// The buttons in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/buttons
func (g_ GCPhysicalInputProfile) Buttons() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("buttons"))
	return rv
}/* debug [instance_properties/getter]: buttons */


// The physical device that the profile represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/device
func (g_ GCPhysicalInputProfile) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The directional pads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/dpads
func (g_ GCPhysicalInputProfile) Dpads() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("dpads"))
	return rv
}/* debug [instance_properties/getter]: dpads */


// The elements in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/elements
func (g_ GCPhysicalInputProfile) Elements() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("elements"))
	return rv
}/* debug [instance_properties/getter]: elements */


// A Boolean value that indicates whether the user remaps elements in this profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/hasRemappedElements
func (g_ GCPhysicalInputProfile) HasRemappedElements() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasRemappedElements"))
	return rv
}/* debug [instance_properties/getter]: hasRemappedElements */


// The time of the most recent change to an element’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/lastEventTimestamp
func (g_ GCPhysicalInputProfile) LastEventTimestamp() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("lastEventTimestamp"))
	return rv
}/* debug [instance_properties/getter]: lastEventTimestamp */


// The touchpads in the profile as key-value pairs for lookup by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/touchpads
func (g_ GCPhysicalInputProfile) Touchpads() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("touchpads"))
	return rv
}/* debug [instance_properties/getter]: touchpads */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/valueDidChangeHandler
func (g_ GCPhysicalInputProfile) ValueDidChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueDidChangeHandler"))
	return rv
}/* debug [instance_properties/getter]: valueDidChangeHandler */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/valueDidChangeHandler
func (g_ GCPhysicalInputProfile) SetValueDidChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueDidChangeHandler:"), value)
}/* debug [instance_properties/setter]: valueDidChangeHandler */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCPhysicalInputProfile) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCPhysicalInputProfile) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCPhysicalInputProfile) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCPhysicalInputProfile) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCPhysicalInputProfile) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCPhysicalInputProfile */



