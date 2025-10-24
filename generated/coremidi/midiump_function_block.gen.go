// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDIUMPFunctionBlock */


/* debug [class_header]: Header for MIDIUMPFunctionBlock */
// The class instance for the [MIDIUMPFunctionBlock] class.
var (
	MIDIUMPFunctionBlockClass     _MIDIUMPFunctionBlockClass
	MIDIUMPFunctionBlockClassOnce sync.Once
)

func getMIDIUMPFunctionBlockClass() _MIDIUMPFunctionBlockClass {
	MIDIUMPFunctionBlockClassOnce.Do(func() {
		MIDIUMPFunctionBlockClass = _MIDIUMPFunctionBlockClass{objc.GetClass("MIDIUMPFunctionBlock")}
	})
	return MIDIUMPFunctionBlockClass
}

type _MIDIUMPFunctionBlockClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIUMPFunctionBlock */
// An interface definition for the [MIDIUMPFunctionBlock] class.
type IMIDIUMPFunctionBlock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDIUMPFunctionBlock */
	// properties:
	Direction() MIDIUMPFunctionBlockDirection
	FirstGroup() MIDIUMPGroupNumber /* typedef */
	FunctionBlockID() MIDIUMPFunctionBlockID /* typedef */
	IsEnabled() bool
	MaxSysEx8Streams() objectivec.IObject
	MIDI1Info() MIDIUMPFunctionBlockMIDI1Info
	MidiCIDevice() IMIDICIDevice
	Name() objc.IObject /* cross-framework: NSString */
	TotalGroupsSpanned() MIDIUInteger7 /* typedef */
	UIHint() MIDIUMPFunctionBlockUIHint
	UMPEndpoint() IMIDIUMPEndpoint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIUMPFunctionBlock */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIUMPFunctionBlock */
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPFunctionBlockClass) Alloc() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIUMPFunctionBlockClass) New() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPFunctionBlock) Init() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPFunctionBlock) Autorelease() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPFunctionBlock creates a new MIDIUMPFunctionBlock instance.
func NewMIDIUMPFunctionBlock() MIDIUMPFunctionBlock {
	return getMIDIUMPFunctionBlockClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIUMPFunctionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock
type MIDIUMPFunctionBlock struct {
	objectivec.Object
}

// MIDIUMPFunctionBlockFrom constructs a [MIDIUMPFunctionBlock] from an unsafe.Pointer.
func MIDIUMPFunctionBlockFrom(ptr unsafe.Pointer) MIDIUMPFunctionBlock {
	return MIDIUMPFunctionBlock{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIUMPFunctionBlock *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIUMPFunctionBlock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIUMPFunctionBlock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIUMPFunctionBlock */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIUMPFunctionBlock */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/direction
func (m_ MIDIUMPFunctionBlock) Direction() MIDIUMPFunctionBlockDirection {
	rv := objc.Send[MIDIUMPFunctionBlockDirection](m_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/firstGroup
func (m_ MIDIUMPFunctionBlock) FirstGroup() MIDIUMPGroupNumber /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("firstGroup"))
	return rv
}/* debug [instance_properties/getter]: firstGroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/functionBlockID
func (m_ MIDIUMPFunctionBlock) FunctionBlockID() MIDIUMPFunctionBlockID /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("functionBlockID"))
	return rv
}/* debug [instance_properties/getter]: functionBlockID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/isEnabled
func (m_ MIDIUMPFunctionBlock) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/maxSysEx8Streams
func (m_ MIDIUMPFunctionBlock) MaxSysEx8Streams() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("maxSysEx8Streams"))
	return rv
}/* debug [instance_properties/getter]: maxSysEx8Streams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midi1Info
func (m_ MIDIUMPFunctionBlock) MIDI1Info() MIDIUMPFunctionBlockMIDI1Info {
	rv := objc.Send[MIDIUMPFunctionBlockMIDI1Info](m_.ID, objc.Sel("MIDI1Info"))
	return rv
}/* debug [instance_properties/getter]: MIDI1Info */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midiCIDevice
func (m_ MIDIUMPFunctionBlock) MidiCIDevice() IMIDICIDevice {
	rv := objc.Send[MIDICIDevice](m_.ID, objc.Sel("midiCIDevice"))
	return rv
}/* debug [instance_properties/getter]: midiCIDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/name
func (m_ MIDIUMPFunctionBlock) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/totalGroupsSpanned
func (m_ MIDIUMPFunctionBlock) TotalGroupsSpanned() MIDIUInteger7 /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("totalGroupsSpanned"))
	return rv
}/* debug [instance_properties/getter]: totalGroupsSpanned */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/uiHint
func (m_ MIDIUMPFunctionBlock) UIHint() MIDIUMPFunctionBlockUIHint {
	rv := objc.Send[MIDIUMPFunctionBlockUIHint](m_.ID, objc.Sel("UIHint"))
	return rv
}/* debug [instance_properties/getter]: UIHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/umpEndpoint
func (m_ MIDIUMPFunctionBlock) UMPEndpoint() IMIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m_.ID, objc.Sel("UMPEndpoint"))
	return rv
}/* debug [instance_properties/getter]: UMPEndpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDIUMPFunctionBlock */



