// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPEndpoint] class.
var (
	MIDIUMPEndpointClass     _MIDIUMPEndpointClass
	MIDIUMPEndpointClassOnce sync.Once
)

func getMIDIUMPEndpointClass() _MIDIUMPEndpointClass {
	MIDIUMPEndpointClassOnce.Do(func() {
		MIDIUMPEndpointClass = _MIDIUMPEndpointClass{objc.GetClass("MIDIUMPEndpoint")}
	})
	return MIDIUMPEndpointClass
}

type _MIDIUMPEndpointClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPEndpoint] class.
type IMIDIUMPEndpoint interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint
type MIDIUMPEndpoint struct {
	objectivec.Object
}

// MIDIUMPEndpointFrom constructs a [MIDIUMPEndpoint] from an unsafe.Pointer.
func MIDIUMPEndpointFrom(ptr unsafe.Pointer) MIDIUMPEndpoint {
	return MIDIUMPEndpoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPEndpointClass) Alloc() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPEndpointClass) New() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPEndpoint) Init() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPEndpoint) Autorelease() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpoint creates a new MIDIUMPEndpoint instance.
func NewMIDIUMPEndpoint() MIDIUMPEndpoint {
	return getMIDIUMPEndpointClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/deviceInfo
func (m_ MIDIUMPEndpoint) DeviceInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceInfo"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/endpointType
func (m_ MIDIUMPEndpoint) EndpointType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endpointType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/functionBlocks
func (m_ MIDIUMPEndpoint) FunctionBlocks() []MIDIUMPFunctionBlock {
	rv := objc.Send[[]MIDIUMPFunctionBlock](m_.ID, objc.Sel("functionBlocks"))
	return rv
}


// SetFunctionBlocks sets the value of the functionBlocks property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/functionBlocks
func (m_ MIDIUMPEndpoint) SetFunctionBlocks(value []MIDIUMPFunctionBlock) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionBlocks:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSReceiveCapability
func (m_ MIDIUMPEndpoint) HasJRTSReceiveCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasJRTSReceiveCapability"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSTransmitCapability
func (m_ MIDIUMPEndpoint) HasJRTSTransmitCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasJRTSTransmitCapability"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasStaticFunctionBlocks
func (m_ MIDIUMPEndpoint) HasStaticFunctionBlocks() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasStaticFunctionBlocks"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiDestination
func (m_ MIDIUMPEndpoint) MIDIDestination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("MIDIDestination"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiProtocol
func (m_ MIDIUMPEndpoint) MIDIProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("MIDIProtocol"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiSource
func (m_ MIDIUMPEndpoint) MIDISource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("MIDISource"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/name
func (m_ MIDIUMPEndpoint) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/productInstanceID
func (m_ MIDIUMPEndpoint) ProductInstanceID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("productInstanceID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/supportedMIDIProtocols
func (m_ MIDIUMPEndpoint) SupportedMIDIProtocols() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("supportedMIDIProtocols"))
	return rv
}



