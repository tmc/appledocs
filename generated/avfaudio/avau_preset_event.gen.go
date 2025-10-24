// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [AUPresetEvent] class.
type IAUPresetEvent interface {
	IMusicEvent
	

	// properties:
	Element() objectivec.IObject
	SetElement(value objectivec.IObject)
	PresetDictionary() objc.IObject /* cross-framework: NSDictionary */
	Scope() objectivec.IObject
	SetScope(value objectivec.IObject)


	

	// methods:


}





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






// Creates an event with the scope, element, and dictionary for the preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/init(scope:element:dictionary:)
func NewAUPresetEventWithScopeElementDictionary(scope objectivec.IObject, element objectivec.IObject, presetDictionary objc.IObject /* cross-framework: NSDictionary */) AUPresetEvent {
	instance := getAUPresetEventClass().Alloc()
	rv := objc.Send[AUPresetEvent](instance.ID, objc.Sel("initWithScope:element:dictionary:"), scope, element, presetDictionary)
	rv.Autorelease()
	return rv
}






















// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/element
func (p_ AUPresetEvent) Element() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("element"))
	return rv
}


// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/element
func (p_ AUPresetEvent) SetElement(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setElement:"), value)
}


// The dictionary that contains the preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/presetDictionary
func (p_ AUPresetEvent) PresetDictionary() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("presetDictionary"))
	return rv
}


// The audio unit scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/scope
func (p_ AUPresetEvent) Scope() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("scope"))
	return rv
}


// The audio unit scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/scope
func (p_ AUPresetEvent) SetScope(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScope:"), value)
}







