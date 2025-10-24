// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDIUMPEndpoint */


/* debug [class_header]: Header for MIDIUMPEndpoint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIUMPEndpoint */
// An interface definition for the [MIDIUMPEndpoint] class.
type IMIDIUMPEndpoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDIUMPEndpoint */
	// properties:
	DeviceInfo() IMIDI2DeviceInfo
	EndpointType() MIDIUMPCIObjectBackingType
	FunctionBlocks() []MIDIUMPFunctionBlock
	SetFunctionBlocks(value []MIDIUMPFunctionBlock)
	HasJRTSReceiveCapability() bool
	HasJRTSTransmitCapability() bool
	HasStaticFunctionBlocks() bool
	MIDIDestination() MIDIEndpointRef /* typedef */
	MIDIProtocol() MIDIProtocolID
	MIDISource() MIDIEndpointRef /* typedef */
	Name() objc.IObject /* cross-framework: NSString */
	ProductInstanceID() objc.IObject /* cross-framework: NSString */
	SupportedMIDIProtocols() MIDIUMPProtocolOptions
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIUMPEndpoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIUMPEndpoint */
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPEndpointClass) Alloc() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIUMPEndpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint
type MIDIUMPEndpoint struct {
	objectivec.Object
}

// MIDIUMPEndpointFrom constructs a [MIDIUMPEndpoint] from an unsafe.Pointer.
func MIDIUMPEndpointFrom(ptr unsafe.Pointer) MIDIUMPEndpoint {
	return MIDIUMPEndpoint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIUMPEndpoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIUMPEndpoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIUMPEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIUMPEndpoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIUMPEndpoint */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/deviceInfo
func (m_ MIDIUMPEndpoint) DeviceInfo() IMIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/endpointType
func (m_ MIDIUMPEndpoint) EndpointType() MIDIUMPCIObjectBackingType {
	rv := objc.Send[MIDIUMPCIObjectBackingType](m_.ID, objc.Sel("endpointType"))
	return rv
}/* debug [instance_properties/getter]: endpointType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/functionBlocks
func (m_ MIDIUMPEndpoint) FunctionBlocks() []MIDIUMPFunctionBlock {
	rv := objc.Send[[]MIDIUMPFunctionBlock](m_.ID, objc.Sel("functionBlocks"))
	return rv
}/* debug [instance_properties/getter]: functionBlocks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/functionBlocks
func (m_ MIDIUMPEndpoint) SetFunctionBlocks(value []MIDIUMPFunctionBlock) {
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
}/* debug [instance_properties/setter]: functionBlocks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSReceiveCapability
func (m_ MIDIUMPEndpoint) HasJRTSReceiveCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasJRTSReceiveCapability"))
	return rv
}/* debug [instance_properties/getter]: hasJRTSReceiveCapability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasJRTSTransmitCapability
func (m_ MIDIUMPEndpoint) HasJRTSTransmitCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasJRTSTransmitCapability"))
	return rv
}/* debug [instance_properties/getter]: hasJRTSTransmitCapability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/hasStaticFunctionBlocks
func (m_ MIDIUMPEndpoint) HasStaticFunctionBlocks() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasStaticFunctionBlocks"))
	return rv
}/* debug [instance_properties/getter]: hasStaticFunctionBlocks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiDestination
func (m_ MIDIUMPEndpoint) MIDIDestination() MIDIEndpointRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("MIDIDestination"))
	return rv
}/* debug [instance_properties/getter]: MIDIDestination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiProtocol
func (m_ MIDIUMPEndpoint) MIDIProtocol() MIDIProtocolID {
	rv := objc.Send[MIDIProtocolID](m_.ID, objc.Sel("MIDIProtocol"))
	return rv
}/* debug [instance_properties/getter]: MIDIProtocol */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/midiSource
func (m_ MIDIUMPEndpoint) MIDISource() MIDIEndpointRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("MIDISource"))
	return rv
}/* debug [instance_properties/getter]: MIDISource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/name
func (m_ MIDIUMPEndpoint) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/productInstanceID
func (m_ MIDIUMPEndpoint) ProductInstanceID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("productInstanceID"))
	return rv
}/* debug [instance_properties/getter]: productInstanceID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint/supportedMIDIProtocols
func (m_ MIDIUMPEndpoint) SupportedMIDIProtocols() MIDIUMPProtocolOptions {
	rv := objc.Send[MIDIUMPProtocolOptions](m_.ID, objc.Sel("supportedMIDIProtocols"))
	return rv
}/* debug [instance_properties/getter]: supportedMIDIProtocols */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDIUMPEndpoint */



