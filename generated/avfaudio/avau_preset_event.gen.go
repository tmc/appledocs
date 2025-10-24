// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAUPresetEvent */


/* debug [class_header]: Header for AVAUPresetEvent */
// The class instance for the [AUPresetEvent] class.
var (
	AUPresetEventClass     _AUPresetEventClass
	AUPresetEventClassOnce sync.Once
)

func getAUPresetEventClass() _AUPresetEventClass {
	AUPresetEventClassOnce.Do(func() {
		AUPresetEventClass = _AUPresetEventClass{objc.GetClass("AVAUPresetEvent")}
	})
	return AUPresetEventClass
}

type _AUPresetEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AUPresetEvent */
// An interface definition for the [AUPresetEvent] class.
type IAUPresetEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for AUPresetEvent */
	// properties:
	Element() objectivec.IObject
	SetElement(value objectivec.IObject)
	PresetDictionary() objc.IObject /* cross-framework: NSDictionary */
	Scope() objectivec.IObject
	SetScope(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AUPresetEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AUPresetEvent */
// Alloc allocates a new instance without initialization.
func (pc _AUPresetEventClass) Alloc() AUPresetEvent {
	rv := objc.Send[AUPresetEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _AUPresetEventClass) New() AUPresetEvent {
	rv := objc.Send[AUPresetEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ AUPresetEvent) Init() AUPresetEvent {
	rv := objc.Send[AUPresetEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ AUPresetEvent) Autorelease() AUPresetEvent {
	rv := objc.Send[AUPresetEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUPresetEvent creates a new AUPresetEvent instance.
func NewAUPresetEvent() AUPresetEvent {
	return getAUPresetEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AUPresetEvent */
// An object that represents a preset load and change on the music track’s destination audio unit.


// An object that represents a preset load and change on the music track’s destination audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent
type AUPresetEvent struct {
	MusicEvent
}

// AUPresetEventFrom constructs a [AUPresetEvent] from an unsafe.Pointer.
//
// An object that represents a preset load and change on the music track’s destination audio unit.
func AUPresetEventFrom(ptr unsafe.Pointer) AUPresetEvent {
	return AUPresetEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AUPresetEvent */

// Creates an event with the scope, element, and dictionary for the preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/init(scope:element:dictionary:)
func NewAUPresetEventWithScopeElementDictionary(scope objectivec.IObject, element objectivec.IObject, presetDictionary objc.IObject /* cross-framework: NSDictionary */) AUPresetEvent {
	instance := getAUPresetEventClass().Alloc()
	rv := objc.Send[AUPresetEvent](instance.ID, objc.Sel("initWithScope:element:dictionary:"), scope, element, presetDictionary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAUPresetEventWithScopeElementDictionary */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AUPresetEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AUPresetEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AUPresetEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AUPresetEvent */

// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/element
func (p_ AUPresetEvent) Element() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("element"))
	return rv
}/* debug [instance_properties/getter]: element */


// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/element
func (p_ AUPresetEvent) SetElement(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setElement:"), value)
}/* debug [instance_properties/setter]: element */


// The dictionary that contains the preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/presetDictionary
func (p_ AUPresetEvent) PresetDictionary() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("presetDictionary"))
	return rv
}/* debug [instance_properties/getter]: presetDictionary */


// The audio unit scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/scope
func (p_ AUPresetEvent) Scope() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("scope"))
	return rv
}/* debug [instance_properties/getter]: scope */


// The audio unit scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/scope
func (p_ AUPresetEvent) SetScope(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScope:"), value)
}/* debug [instance_properties/setter]: scope */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAUPresetEvent */


