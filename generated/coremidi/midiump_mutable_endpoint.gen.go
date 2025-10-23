// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MIDIUMPMutableEndpoint] class.
var (
	MIDIUMPMutableEndpointClass     _MIDIUMPMutableEndpointClass
	MIDIUMPMutableEndpointClassOnce sync.Once
)

func getMIDIUMPMutableEndpointClass() _MIDIUMPMutableEndpointClass {
	MIDIUMPMutableEndpointClassOnce.Do(func() {
		MIDIUMPMutableEndpointClass = _MIDIUMPMutableEndpointClass{objc.GetClass("MIDIUMPMutableEndpoint")}
	})
	return MIDIUMPMutableEndpointClass
}

type _MIDIUMPMutableEndpointClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPMutableEndpoint] class.
type IMIDIUMPMutableEndpoint interface {
	IMIDIUMPEndpoint
	IsEnabled() bool
	MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock
	SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock)
	RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool, error_ unsafe.Pointer) bool
	SetEnabledError(isEnabled bool, error_ unsafe.Pointer) bool
	SetNameError(name string, error_ unsafe.Pointer) bool
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/init(name:deviceInfo:productInstanceID:midiProtocol:destinationCallback:)
func NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name string, deviceInfo IMIDI2DeviceInfo, productInstanceID string, MIDIProtocol MIDIProtocolID, destinationCallback unsafe.Pointer) MIDIUMPMutableEndpoint {
	instance := getMIDIUMPMutableEndpointClass().Alloc()
	rv := objc.Send[MIDIUMPMutableEndpoint](instance.ID, objc.Sel("initWithName:deviceInfo:productInstanceID:MIDIProtocol:destinationCallback:"), objc.String(name), deviceInfo, objc.String(productInstanceID), MIDIProtocol, destinationCallback)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/registerFunctionBlocks(_:markAsStatic:)
func (m_ MIDIUMPMutableEndpoint) RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("registerFunctionBlocks:markAsStatic:error:"), functionBlocks, markAsStatic, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setEnabled(_:)
func (m_ MIDIUMPMutableEndpoint) SetEnabledError(isEnabled bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setEnabled:error:"), isEnabled, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setName(_:)
func (m_ MIDIUMPMutableEndpoint) SetNameError(name string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setName:error:"), objc.String(name), error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/isEnabled
func (m_ MIDIUMPMutableEndpoint) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/mutableFunctionBlocks
func (m_ MIDIUMPMutableEndpoint) MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock {
	rv := objc.Send[[]MIDIUMPMutableFunctionBlock](m_.ID, objc.Sel("mutableFunctionBlocks"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/mutableFunctionBlocks
func (m_ MIDIUMPMutableEndpoint) SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock) {
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
	objc.Send[objc.ID](m_.ID, objc.Sel("setMutableFunctionBlocks:"), nsArray)
}


