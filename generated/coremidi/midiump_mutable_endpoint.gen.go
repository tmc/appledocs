// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDIUMPMutableEndpoint */


/* debug [class_header]: Header for MIDIUMPMutableEndpoint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIUMPMutableEndpoint */
// An interface definition for the [MIDIUMPMutableEndpoint] class.
type IMIDIUMPMutableEndpoint interface {
	IMIDIUMPEndpoint
	
/* debug [class_interface_properties]: Properties for MIDIUMPMutableEndpoint */
	// properties:
	IsEnabled() bool
	MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock
	SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIUMPMutableEndpoint */
	// methods:
	RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool, error_ objectivec.IObject) bool
	SetEnabledError(isEnabled bool, error_ objectivec.IObject) bool
	SetNameError(name objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIUMPMutableEndpoint */
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPMutableEndpointClass) Alloc() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIUMPMutableEndpoint */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIUMPMutableEndpoint */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/init(name:deviceInfo:productInstanceID:midiProtocol:destinationCallback:)
func NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(name objc.IObject /* cross-framework: NSString */, deviceInfo IMIDI2DeviceInfo, productInstanceID objc.IObject /* cross-framework: NSString */, MIDIProtocol MIDIProtocolID, destinationCallback objectivec.IObject) MIDIUMPMutableEndpoint {
	instance := getMIDIUMPMutableEndpointClass().Alloc()
	rv := objc.Send[MIDIUMPMutableEndpoint](instance.ID, objc.Sel("initWithName:deviceInfo:productInstanceID:MIDIProtocol:destinationCallback:"), name, deviceInfo, productInstanceID, MIDIProtocol, destinationCallback)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIUMPMutableEndpoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIUMPMutableEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIUMPMutableEndpoint */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/registerFunctionBlocks(_:markAsStatic:)
func (m_ MIDIUMPMutableEndpoint) RegisterFunctionBlocksMarkAsStaticError(functionBlocks []MIDIUMPMutableFunctionBlock, markAsStatic bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("registerFunctionBlocks:markAsStatic:error:"), functionBlocks, markAsStatic, error_)
	return rv
}/* debug [instance_methods/method]: RegisterFunctionBlocksMarkAsStaticError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setEnabled(_:)
func (m_ MIDIUMPMutableEndpoint) SetEnabledError(isEnabled bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setEnabled:error:"), isEnabled, error_)
	return rv
}/* debug [instance_methods/method]: SetEnabledError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/setName(_:)
func (m_ MIDIUMPMutableEndpoint) SetNameError(name objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setName:error:"), name, error_)
	return rv
}/* debug [instance_methods/method]: SetNameError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIUMPMutableEndpoint */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/isEnabled
func (m_ MIDIUMPMutableEndpoint) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/mutableFunctionBlocks
func (m_ MIDIUMPMutableEndpoint) MutableFunctionBlocks() []MIDIUMPMutableFunctionBlock {
	rv := objc.Send[[]MIDIUMPMutableFunctionBlock](m_.ID, objc.Sel("mutableFunctionBlocks"))
	return rv
}/* debug [instance_properties/getter]: mutableFunctionBlocks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableEndpoint/mutableFunctionBlocks
func (m_ MIDIUMPMutableEndpoint) SetMutableFunctionBlocks(value []MIDIUMPMutableFunctionBlock) {
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
}/* debug [instance_properties/setter]: mutableFunctionBlocks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDIUMPMutableEndpoint */


