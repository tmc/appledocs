// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MIDIUMPMutableEndpoint] class.
var (
	mIDIUMPMutableEndpointClass     _MIDIUMPMutableEndpointClass
	mIDIUMPMutableEndpointClassOnce sync.Once
)

func getMIDIUMPMutableEndpointClass() _MIDIUMPMutableEndpointClass {
	mIDIUMPMutableEndpointClassOnce.Do(func() {
		mIDIUMPMutableEndpointClass = _MIDIUMPMutableEndpointClass{objc.GetClass("MIDIUMPMutableEndpoint")}
	})
	return mIDIUMPMutableEndpointClass
}

type _MIDIUMPMutableEndpointClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPMutableEndpoint] class.
type IMIDIUMPMutableEndpoint interface {
	IMIDIUMPEndpoint
	RegisterFunctionBlocksMarkAsStaticError(functionBlocks unsafe.Pointer, markAsStatic bool, error unsafe.Pointer) bool
	SetEnabledError(isEnabled bool, error unsafe.Pointer) bool
	SetNameError(name string, error unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint
type MIDIUMPMutableEndpoint struct {
	MIDIUMPEndpoint
}

// MIDIUMPMutableEndpointFrom constructs a [MIDIUMPMutableEndpoint] from an unsafe.Pointer.
func MIDIUMPMutableEndpointFrom(ptr unsafe.Pointer) MIDIUMPMutableEndpoint {
	return MIDIUMPMutableEndpoint{
		MIDIUMPEndpoint: MIDIUMPEndpointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPMutableEndpointClass) Alloc() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPMutableEndpointClass) New() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPMutableEndpoint) Init() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPMutableEndpoint) Autorelease() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPMutableEndpoint creates a new MIDIUMPMutableEndpoint instance.
func NewMIDIUMPMutableEndpoint() MIDIUMPMutableEndpoint {
	return getMIDIUMPMutableEndpointClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/init(name:deviceInfo:productInstanceID:midiProtocol:destinationCallback:)
func NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name string, deviceInfo unsafe.Pointer, productInstanceID string, MIDIProtocol unsafe.Pointer, destinationCallback unsafe.Pointer) MIDIUMPMutableEndpoint {
	instance := getMIDIUMPMutableEndpointClass().Alloc()
	rv := objc.Send[MIDIUMPMutableEndpoint](instance.ID, objc.Sel("initWithName:deviceInfo:productInstanceID:MIDIProtocol:destinationCallback:"), objc.String(name), deviceInfo, objc.String(productInstanceID), MIDIProtocol, destinationCallback)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/registerFunctionBlocks(_:markAsStatic:)
func (m_ MIDIUMPMutableEndpoint) RegisterFunctionBlocksMarkAsStaticError(functionBlocks unsafe.Pointer, markAsStatic bool, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("registerFunctionBlocks:markAsStatic:error:"), functionBlocks, markAsStatic, error)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setEnabled(_:)
func (m_ MIDIUMPMutableEndpoint) SetEnabledError(isEnabled bool, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setEnabled:error:"), isEnabled, error)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setName(_:)
func (m_ MIDIUMPMutableEndpoint) SetNameError(name string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setName:error:"), objc.String(name), error)
	return rv
}

