// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ParameterEvent] class.
var (
	ParameterEventClass     _ParameterEventClass
	ParameterEventClassOnce sync.Once
)

func getParameterEventClass() _ParameterEventClass {
	ParameterEventClassOnce.Do(func() {
		ParameterEventClass = _ParameterEventClass{objc.GetClass("AVParameterEvent")}
	})
	return ParameterEventClass
}

type _ParameterEventClass struct {
	class objc.Class
}





// An interface definition for the [ParameterEvent] class.
type IParameterEvent interface {
	IMusicEvent
	

	// properties:
	Element() objectivec.IObject
	SetElement(value objectivec.IObject)
	ParameterID() objectivec.IObject
	SetParameterID(value objectivec.IObject)
	Scope() objectivec.IObject
	SetScope(value objectivec.IObject)
	Value() float32
	SetValue(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _ParameterEventClass) Alloc() ParameterEvent {
	rv := objc.Send[ParameterEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterEventClass) New() ParameterEvent {
	rv := objc.Send[ParameterEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterEvent) Init() ParameterEvent {
	rv := objc.Send[ParameterEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterEvent) Autorelease() ParameterEvent {
	rv := objc.Send[ParameterEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterEvent creates a new ParameterEvent instance.
func NewParameterEvent() ParameterEvent {
	return getParameterEventClass().New()
}





// An object that represents a parameter event on a music track’s destination.
//
// When you configure an audio unit as the destination for an that contains this event, you can schedule and automate parameter changes. When the track is playing as part of a sequence, the destination audio unit receives set-parameter messages whose values change smoothly along a linear ramp between each event’s beat location. If you add an event to an empty, non-automation track, the track becomes an automation track.


// An object that represents a parameter event on a music track’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent
type ParameterEvent struct {
	MusicEvent
}

// ParameterEventFrom constructs a [ParameterEvent] from an unsafe.Pointer.
//
// An object that represents a parameter event on a music track’s destination.
func ParameterEventFrom(ptr unsafe.Pointer) ParameterEvent {
	return ParameterEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates an event with a parameter identifier, scope, element, and value for the parameter to set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/init(parameterID:scope:element:value:)
func NewParameterEventWithParameterIDScopeElementValue(parameterID objectivec.IObject, scope objectivec.IObject, element objectivec.IObject, value float32) ParameterEvent {
	instance := getParameterEventClass().Alloc()
	rv := objc.Send[ParameterEvent](instance.ID, objc.Sel("initWithParameterID:scope:element:value:"), parameterID, scope, element, value)
	rv.Autorelease()
	return rv
}






















// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/element
func (p_ ParameterEvent) Element() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("element"))
	return rv
}


// The element index in the scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/element
func (p_ ParameterEvent) SetElement(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setElement:"), value)
}


// The identifier of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/parameterID
func (p_ ParameterEvent) ParameterID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("parameterID"))
	return rv
}


// The identifier of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/parameterID
func (p_ ParameterEvent) SetParameterID(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameterID:"), value)
}


// The audio unit scope for the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/scope
func (p_ ParameterEvent) Scope() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("scope"))
	return rv
}


// The audio unit scope for the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/scope
func (p_ ParameterEvent) SetScope(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScope:"), value)
}


// The value of the parameter to set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/value
func (p_ ParameterEvent) Value() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("value"))
	return rv
}


// The value of the parameter to set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/value
func (p_ ParameterEvent) SetValue(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:"), value)
}







