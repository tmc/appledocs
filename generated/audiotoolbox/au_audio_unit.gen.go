// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/coremidi"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUAudioUnit */


/* debug [class_header]: Header for AUAudioUnit */
// The class instance for the [AudioUnit] class.
var (
	AudioUnitClass     _AudioUnitClass
	AudioUnitClassOnce sync.Once
)

func getAudioUnitClass() _AudioUnitClass {
	AudioUnitClassOnce.Do(func() {
		AudioUnitClass = _AudioUnitClass{objc.GetClass("AUAudioUnit")}
	})
	return AudioUnitClass
}

type _AudioUnitClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnit */
// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnit */
	// properties:
	AllParameterValues() bool
	AudioUnitMIDIProtocol() objectivec.IObject
	AudioUnitName() objc.IObject /* cross-framework: NSString */
	AudioUnitShortName() objc.IObject /* cross-framework: NSString */
	CanPerformInput() bool
	CanPerformOutput() bool
	CanProcessInPlace() bool
	ChannelCapabilities() []foundation.Number
	ChannelMap() []foundation.Number
	SetChannelMap(value []foundation.Number)
	Component() AudioComponent /* typedef */
	ComponentDescription() objc.IObject /* cross-framework: AudioComponentDescription */
	ComponentName() objc.IObject /* cross-framework: NSString */
	ComponentVersion() uint32 /* not a class type */
	ContextName() objc.IObject /* cross-framework: NSString */
	SetContextName(value objc.IObject /* cross-framework: NSString */)
	CurrentPreset() IAUAudioUnitPreset
	SetCurrentPreset(value IAUAudioUnitPreset)
	DeviceID() AudioObjectID /* typedef */
	DeviceInputLatency() float64
	DeviceOutputLatency() float64
	FactoryPresets() []AudioUnitPreset
	FullState() foundation.IDictionary
	SetFullState(value foundation.IDictionary)
	FullStateForDocument() foundation.IDictionary
	SetFullStateForDocument(value foundation.IDictionary)
	HostMIDIProtocol() objectivec.IObject
	SetHostMIDIProtocol(value objectivec.IObject)
	InputBusses() IAUAudioUnitBusArray
	InputHandler() InputHandler /* not a class type */
	SetInputHandler(value InputHandler /* not a class type */)
	InternalRenderBlock() InternalRenderBlock /* not a class type */
	InputEnabled() bool
	SetInputEnabled(value bool)
	IsLoadedInProcess() bool
	MusicDeviceOrEffect() bool
	OutputEnabled() bool
	SetOutputEnabled(value bool)
	RenderingOffline() bool
	SetRenderingOffline(value bool)
	Running() bool
	Latency() float64
	ManufacturerName() objc.IObject /* cross-framework: NSString */
	MaximumFramesToRender() AudioFrameCount /* typedef */
	SetMaximumFramesToRender(value AudioFrameCount /* typedef */)
	MIDIOutputBufferSizeHint() int
	SetMIDIOutputBufferSizeHint(value int)
	MIDIOutputEventBlock() MIDIOutputEventBlock /* not a class type */
	SetMIDIOutputEventBlock(value MIDIOutputEventBlock /* not a class type */)
	MIDIOutputEventListBlock() MIDIEventListBlock /* not a class type */
	SetMIDIOutputEventListBlock(value MIDIEventListBlock /* not a class type */)
	MIDIOutputNames() []string
	MigrateFromPlugin() objc.IObject /* cross-framework: NSArray */
	MusicalContextBlock() HostMusicalContextBlock /* not a class type */
	SetMusicalContextBlock(value HostMusicalContextBlock /* not a class type */)
	OsWorkgroup() objectivec.IObject
	OutputBusses() IAUAudioUnitBusArray
	OutputProvider() RenderPullInputBlock /* not a class type */
	SetOutputProvider(value RenderPullInputBlock /* not a class type */)
	ParameterTree() IAUParameterTree
	SetParameterTree(value IAUParameterTree)
	ProfileChangedBlock() MIDICIProfileChangedBlock /* not a class type */
	SetProfileChangedBlock(value MIDICIProfileChangedBlock /* not a class type */)
	ProvidesUserInterface() bool
	RenderBlock() RenderBlock /* not a class type */
	RenderContextObserver() RenderContextObserver /* not a class type */
	RenderQuality() int
	SetRenderQuality(value int)
	RenderResourcesAllocated() bool
	ScheduleMIDIEventBlock() ScheduleMIDIEventBlock /* not a class type */
	ScheduleMIDIEventListBlock() MIDIEventListBlock /* not a class type */
	ScheduleParameterBlock() ScheduleParameterBlock /* not a class type */
	ShouldBypassEffect() bool
	SetShouldBypassEffect(value bool)
	SupportsMPE() bool
	SupportsUserPresets() bool
	TailTime() float64
	TransportStateBlock() HostTransportStateBlock /* not a class type */
	SetTransportStateBlock(value HostTransportStateBlock /* not a class type */)
	UserPresets() []AudioUnitPreset
	VirtualMIDICableCount() int
	IsInputEnabled() bool
	SetIsInputEnabled(value bool)
	IsMusicDeviceOrEffect() bool
	SetIsMusicDeviceOrEffect(value bool)
	IsOutputEnabled() bool
	SetIsOutputEnabled(value bool)
	IsRenderingOffline() bool
	SetIsRenderingOffline(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
	KAUPresetCPULoadKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetCPULoadKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetDataKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetDataKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetElementNameKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetElementNameKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetExternalFileRefs() objc.IObject /* cross-framework: NSString */
	SetKAUPresetExternalFileRefs(value objc.IObject /* cross-framework: NSString */)
	KAUPresetMASDataKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetMASDataKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetManufacturerKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetManufacturerKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetNameKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetNameKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetNumberKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetNumberKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetPartKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetPartKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetRenderQualityKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetRenderQualityKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetSubtypeKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetSubtypeKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetTypeKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetTypeKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetVSTDataKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetVSTDataKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetVSTPresetKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetVSTPresetKey(value objc.IObject /* cross-framework: NSString */)
	KAUPresetVersionKey() objc.IObject /* cross-framework: NSString */
	SetKAUPresetVersionKey(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnit */
	// methods:
	AllocateRenderResourcesAndReturnError(outError objectivec.IObject) bool
	DeallocateRenderResources()
	DeleteUserPresetError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) bool
	DisableProfileCableOnChannelError(profile coremidi.MIDICIProfile, cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */, outError objectivec.IObject) bool
	EnableProfileCableOnChannelError(profile coremidi.MIDICIProfile, cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */, outError objectivec.IObject) bool
	MessageChannelFor(channelName objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	ParametersForOverviewWithCount(count int) []foundation.Number
	PresetStateForError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) foundation.IDictionary
	ProfileStateForCableChannel(cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */) coremidi.MIDICIProfileState
	RemoveRenderObserver(token int)
	RequestViewControllerWithCompletionHandler(completionHandler unsafe.Pointer)
	Reset()
	SaveUserPresetError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) bool
	SelectViewConfiguration(viewConfiguration objectivec.IObject)
	SetDeviceIDError(deviceID AudioObjectID /* typedef */, outError objectivec.IObject) bool
	ShouldChangeToFormatForBus(format avfaudio.AudioFormat, bus IAUAudioUnitBus) bool
	StartHardwareAndReturnError(outError objectivec.IObject) bool
	StopHardware()
	SupportedViewConfigurations(availableViewConfigurations []objc.IObject) foundation.IndexSet
	TokenByAddingRenderObserver(observer RenderObserver /* not a class type */) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnit */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitClass) New() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnit) Init() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnit) Autorelease() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnit creates a new AudioUnit instance.
func NewAudioUnit() AudioUnit {
	return getAudioUnitClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnit */
// A class that defines a host’s interface to an audio unit.
//
// Hosts can instantiate either version 3 or version 2 audio units with this class, and to some extent control whether an audio unit is instantiated in-process or in a separate extension process. Version 3 audio units should subclass the class. Version 3 audio unit components can be registered in the following ways: Package the component into an app extension containing an entry. The principal class must conform to the protocol, which will typically instantiate an subclass. Call the method to associate a component description with an subclass. Use the convention when naming your audio unit component. Version 2 audio units should subclass the class instead. Version 2 audio unit components can be registered in the following ways: Package the component into a component bundle containing an entry, referring to an function. Call the function to associate a component description with an function. A host does not need to be aware of the concrete subclass that is being instantiated. The method ensures that the proper subclass is used.


// A class that defines a host’s interface to an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit
type AudioUnit struct {
	objectivec.Object
}

// AudioUnitFrom constructs a [AudioUnit] from an unsafe.Pointer.
//
// A class that defines a host’s interface to an audio unit.
func AudioUnitFrom(ptr unsafe.Pointer) AudioUnit {
	return AudioUnit{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnit */

// Synchronously initializes a new audio unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:)
func NewAudioUnitWithComponentDescriptionError(componentDescription objc.IObject /* cross-framework: AudioComponentDescription */, outError objectivec.IObject) AudioUnit {
	instance := getAudioUnitClass().Alloc()
	rv := objc.Send[AudioUnit](instance.ID, objc.Sel("initWithComponentDescription:error:"), componentDescription, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitWithComponentDescriptionError */


// Synchronously initializes a new audio unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:options:)
func NewAudioUnitWithComponentDescriptionOptionsError(componentDescription objc.IObject /* cross-framework: AudioComponentDescription */, options AudioComponentInstantiationOptions, outError objectivec.IObject) AudioUnit {
	instance := getAudioUnitClass().Alloc()
	rv := objc.Send[AudioUnit](instance.ID, objc.Sel("initWithComponentDescription:options:error:"), componentDescription, options, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitWithComponentDescriptionOptionsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnit */

// Asynchronously creates an audio unit instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/instantiate(with:options:completionHandler:)
func (ac _AudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(componentDescription objc.IObject /* cross-framework: AudioComponentDescription */, options AudioComponentInstantiationOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), componentDescription, options, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InstantiateWithComponentDescriptionOptionsCompletionHandler) */


// Registers an audio unit subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/registerSubclass(_:as:name:version:)
func (ac _AudioUnitClass) RegisterSubclassAsComponentDescriptionNameVersion(cls objc.Class, componentDescription objc.IObject /* cross-framework: AudioComponentDescription */, name objc.IObject /* cross-framework: NSString */, version objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("registerSubclass:asComponentDescription:name:version:"), cls, componentDescription, name, version)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterSubclassAsComponentDescriptionNameVersion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnit */

// Allocates resources required to render audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allocateRenderResources()
func (a_ AudioUnit) AllocateRenderResourcesAndReturnError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allocateRenderResourcesAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: AllocateRenderResourcesAndReturnError */


// Deallocates resources required to render audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deallocateRenderResources()
func (a_ AudioUnit) DeallocateRenderResources() {
	objc.Send[objc.ID](a_.ID, objc.Sel("deallocateRenderResources"))
}/* debug [instance_methods/method]: DeallocateRenderResources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deleteUserPreset(_:)
func (a_ AudioUnit) DeleteUserPresetError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("deleteUserPreset:error:"), userPreset, outError)
	return rv
}/* debug [instance_methods/method]: DeleteUserPresetError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/disableProfile(_:cable:onChannel:)
func (a_ AudioUnit) DisableProfileCableOnChannelError(profile coremidi.MIDICIProfile, cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("disableProfile:cable:onChannel:error:"), profile, cable, channel, outError)
	return rv
}/* debug [instance_methods/method]: DisableProfileCableOnChannelError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/enable(_:cable:onChannel:)
func (a_ AudioUnit) EnableProfileCableOnChannelError(profile coremidi.MIDICIProfile, cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableProfile:cable:onChannel:error:"), profile, cable, channel, outError)
	return rv
}/* debug [instance_methods/method]: EnableProfileCableOnChannelError */


// Returns an object for bidirectional communication between an audio unit and its host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/messageChannel(for:)
func (a_ AudioUnit) MessageChannelFor(channelName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("messageChannelFor:"), channelName)
	return rv
}/* debug [instance_methods/method]: MessageChannelFor */


// Returns the audio unit’s most important parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parametersForOverview(withCount:)
func (a_ AudioUnit) ParametersForOverviewWithCount(count int) []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("parametersForOverviewWithCount:"), count)
	return rv
}/* debug [instance_methods/method]: ParametersForOverviewWithCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/presetState(for:)
func (a_ AudioUnit) PresetStateForError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("presetStateFor:error:"), userPreset, outError)
	return rv
}/* debug [instance_methods/method]: PresetStateForError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileState(forCable:channel:)
func (a_ AudioUnit) ProfileStateForCableChannel(cable uint8 /* not a class type */, channel MIDIChannelNumber /* typedef */) coremidi.MIDICIProfileState {
	rv := objc.Send[coremidi.MIDICIProfileState](a_.ID, objc.Sel("profileStateForCable:channel:"), cable, channel)
	return rv
}/* debug [instance_methods/method]: ProfileStateForCableChannel */


// Removes an observer block previously added to the render cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/removeRenderObserver(_:)
func (a_ AudioUnit) RemoveRenderObserver(token int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeRenderObserver:"), token)
}/* debug [instance_methods/method]: RemoveRenderObserver */


// Requests an audio unit’s custom view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/requestViewController(completionHandler:)
func (a_ AudioUnit) RequestViewControllerWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestViewControllerWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RequestViewControllerWithCompletionHandler */


// Resets transitory rendering state to its initial state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/reset()
func (a_ AudioUnit) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/saveUserPreset(_:)
func (a_ AudioUnit) SaveUserPresetError(userPreset IAUAudioUnitPreset, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveUserPreset:error:"), userPreset, outError)
	return rv
}/* debug [instance_methods/method]: SaveUserPresetError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/select(_:)
func (a_ AudioUnit) SelectViewConfiguration(viewConfiguration objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectViewConfiguration:"), viewConfiguration)
}/* debug [instance_methods/method]: SelectViewConfiguration */


// Sets the I/O hardware device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/setDeviceID(_:)
func (a_ AudioUnit) SetDeviceIDError(deviceID AudioObjectID /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setDeviceID:error:"), deviceID, outError)
	return rv
}/* debug [instance_methods/method]: SetDeviceIDError */


// This is called when you set the format on a bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldChange(to:for:)
func (a_ AudioUnit) ShouldChangeToFormatForBus(format avfaudio.AudioFormat, bus IAUAudioUnitBus) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldChangeToFormat:forBus:"), format, bus)
	return rv
}/* debug [instance_methods/method]: ShouldChangeToFormatForBus */


// Starts the audio hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/startHardware()
func (a_ AudioUnit) StartHardwareAndReturnError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startHardwareAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: StartHardwareAndReturnError */


// Stops the audio hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/stopHardware()
func (a_ AudioUnit) StopHardware() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopHardware"))
}/* debug [instance_methods/method]: StopHardware */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportedViewConfigurations(_:)
func (a_ AudioUnit) SupportedViewConfigurations(availableViewConfigurations []objc.IObject) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](a_.ID, objc.Sel("supportedViewConfigurations:"), availableViewConfigurations)
	return rv
}/* debug [instance_methods/method]: SupportedViewConfigurations */


// Adds a block to be called on each render cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/token(byAddingRenderObserver:)
func (a_ AudioUnit) TokenByAddingRenderObserver(observer RenderObserver /* not a class type */) int {
	rv := objc.Send[int](a_.ID, objc.Sel("tokenByAddingRenderObserver:"), observer)
	return rv
}/* debug [instance_methods/method]: TokenByAddingRenderObserver */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnit */

// Special read-only property for KVO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allParameterValues
func (a_ AudioUnit) AllParameterValues() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allParameterValues"))
	return rv
}/* debug [instance_properties/getter]: allParameterValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitMIDIProtocol
func (a_ AudioUnit) AudioUnitMIDIProtocol() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("AudioUnitMIDIProtocol"))
	return rv
}/* debug [instance_properties/getter]: AudioUnitMIDIProtocol */


// The audio unit’s name, derived from the component’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitName
func (a_ AudioUnit) AudioUnitName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioUnitName"))
	return rv
}/* debug [instance_properties/getter]: audioUnitName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitShortName
func (a_ AudioUnit) AudioUnitShortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioUnitShortName"))
	return rv
}/* debug [instance_properties/getter]: audioUnitShortName */


// Determines whether the I/O device can perform input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformInput
func (a_ AudioUnit) CanPerformInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformInput"))
	return rv
}/* debug [instance_properties/getter]: canPerformInput */


// Determines whether the I/O device can perform output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformOutput
func (a_ AudioUnit) CanPerformOutput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformOutput"))
	return rv
}/* debug [instance_properties/getter]: canPerformOutput */


// Determines whether an audio unit can process in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canProcessInPlace
func (a_ AudioUnit) CanProcessInPlace() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProcessInPlace"))
	return rv
}/* debug [instance_properties/getter]: canProcessInPlace */


// Expresses valid combinations of input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelCapabilities
func (a_ AudioUnit) ChannelCapabilities() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("channelCapabilities"))
	return rv
}/* debug [instance_properties/getter]: channelCapabilities */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelMap
func (a_ AudioUnit) ChannelMap() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("channelMap"))
	return rv
}/* debug [instance_properties/getter]: channelMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelMap
func (a_ AudioUnit) SetChannelMap(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelMap:"), nsArray)
}/* debug [instance_properties/setter]: channelMap */


// The component found in the component description with which the audio unit was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/component
func (a_ AudioUnit) Component() AudioComponent /* typedef */ {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("component"))
	return rv
}/* debug [instance_properties/getter]: component */


// The component description with which the audio unit was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentDescription
func (a_ AudioUnit) ComponentDescription() objc.IObject /* cross-framework: AudioComponentDescription */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("componentDescription"))
	return rv
}/* debug [instance_properties/getter]: componentDescription */


// The audio unit’s component’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentName
func (a_ AudioUnit) ComponentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("componentName"))
	return rv
}/* debug [instance_properties/getter]: componentName */


// The audio unit’s component’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentVersion
func (a_ AudioUnit) ComponentVersion() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("componentVersion"))
	return rv
}/* debug [instance_properties/getter]: componentVersion */


// Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/contextName
func (a_ AudioUnit) ContextName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contextName"))
	return rv
}/* debug [instance_properties/getter]: contextName */


// Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/contextName
func (a_ AudioUnit) SetContextName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContextName:"), value)
}/* debug [instance_properties/setter]: contextName */


// The audio unit’s last-selected preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/currentPreset
func (a_ AudioUnit) CurrentPreset() IAUAudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](a_.ID, objc.Sel("currentPreset"))
	return rv
}/* debug [instance_properties/getter]: currentPreset */


// The audio unit’s last-selected preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/currentPreset
func (a_ AudioUnit) SetCurrentPreset(value IAUAudioUnitPreset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPreset:"), value)
}/* debug [instance_properties/setter]: currentPreset */


// Gets the I/O hardware device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceID
func (a_ AudioUnit) DeviceID() AudioObjectID /* typedef */ {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("deviceID"))
	return rv
}/* debug [instance_properties/getter]: deviceID */


// The audio device’s input latency, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceInputLatency
func (a_ AudioUnit) DeviceInputLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("deviceInputLatency"))
	return rv
}/* debug [instance_properties/getter]: deviceInputLatency */


// The audio devic’s output latency, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceOutputLatency
func (a_ AudioUnit) DeviceOutputLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("deviceOutputLatency"))
	return rv
}/* debug [instance_properties/getter]: deviceOutputLatency */


// A collection of presets provided by the audio unit’s developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/factoryPresets
func (a_ AudioUnit) FactoryPresets() []AudioUnitPreset {
	rv := objc.Send[[]AudioUnitPreset](a_.ID, objc.Sel("factoryPresets"))
	return rv
}/* debug [instance_properties/getter]: factoryPresets */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullState
func (a_ AudioUnit) FullState() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("fullState"))
	return rv
}/* debug [instance_properties/getter]: fullState */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullState
func (a_ AudioUnit) SetFullState(value foundation.IDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullState:"), value)
}/* debug [instance_properties/setter]: fullState */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullStateForDocument
func (a_ AudioUnit) FullStateForDocument() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("fullStateForDocument"))
	return rv
}/* debug [instance_properties/getter]: fullStateForDocument */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullStateForDocument
func (a_ AudioUnit) SetFullStateForDocument(value foundation.IDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullStateForDocument:"), value)
}/* debug [instance_properties/setter]: fullStateForDocument */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/hostMIDIProtocol
func (a_ AudioUnit) HostMIDIProtocol() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("hostMIDIProtocol"))
	return rv
}/* debug [instance_properties/getter]: hostMIDIProtocol */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/hostMIDIProtocol
func (a_ AudioUnit) SetHostMIDIProtocol(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHostMIDIProtocol:"), value)
}/* debug [instance_properties/setter]: hostMIDIProtocol */


// An array containing the audio unit’s input connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputBusses
func (a_ AudioUnit) InputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("inputBusses"))
	return rv
}/* debug [instance_properties/getter]: inputBusses */


// The block that the output unit will call to notify when input is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputHandler
func (a_ AudioUnit) InputHandler() InputHandler /* not a class type */ {
	rv := objc.Send[InputHandler](a_.ID, objc.Sel("inputHandler"))
	return rv
}/* debug [instance_properties/getter]: inputHandler */


// The block that the output unit will call to notify when input is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputHandler
func (a_ AudioUnit) SetInputHandler(value InputHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputHandler:"), value)
}/* debug [instance_properties/setter]: inputHandler */


// The block which you must provide, via a getter, in order to implement rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/internalRenderBlock
func (a_ AudioUnit) InternalRenderBlock() InternalRenderBlock /* not a class type */ {
	rv := objc.Send[InternalRenderBlock](a_.ID, objc.Sel("internalRenderBlock"))
	return rv
}/* debug [instance_properties/getter]: internalRenderBlock */


// A flag enabling audio input from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isInputEnabled
func (a_ AudioUnit) InputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputEnabled"))
	return rv
}/* debug [instance_properties/getter]: inputEnabled */


// A flag enabling audio input from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isInputEnabled
func (a_ AudioUnit) SetInputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputEnabled:"), value)
}/* debug [instance_properties/setter]: inputEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isLoadedInProcess
func (a_ AudioUnit) IsLoadedInProcess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isLoadedInProcess"))
	return rv
}/* debug [instance_properties/getter]: isLoadedInProcess */


// Specifies whether an audio unit responds to MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isMusicDeviceOrEffect
func (a_ AudioUnit) MusicDeviceOrEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("musicDeviceOrEffect"))
	return rv
}/* debug [instance_properties/getter]: musicDeviceOrEffect */


// A flag enabling audio output from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isOutputEnabled
func (a_ AudioUnit) OutputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("outputEnabled"))
	return rv
}/* debug [instance_properties/getter]: outputEnabled */


// A flag enabling audio output from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isOutputEnabled
func (a_ AudioUnit) SetOutputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputEnabled:"), value)
}/* debug [instance_properties/setter]: outputEnabled */


// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRenderingOffline
func (a_ AudioUnit) RenderingOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("renderingOffline"))
	return rv
}/* debug [instance_properties/getter]: renderingOffline */


// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRenderingOffline
func (a_ AudioUnit) SetRenderingOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenderingOffline:"), value)
}/* debug [instance_properties/setter]: renderingOffline */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRunning
func (a_ AudioUnit) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// The audio unit’s processing latency, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/latency
func (a_ AudioUnit) Latency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("latency"))
	return rv
}/* debug [instance_properties/getter]: latency */


// The manufacturer’s name, derived from the component’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/manufacturerName
func (a_ AudioUnit) ManufacturerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("manufacturerName"))
	return rv
}/* debug [instance_properties/getter]: manufacturerName */


// The maximum number of frames that the audio unit can render at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/maximumFramesToRender
func (a_ AudioUnit) MaximumFramesToRender() AudioFrameCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("maximumFramesToRender"))
	return rv
}/* debug [instance_properties/getter]: maximumFramesToRender */


// The maximum number of frames that the audio unit can render at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/maximumFramesToRender
func (a_ AudioUnit) SetMaximumFramesToRender(value AudioFrameCount /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumFramesToRender:"), value)
}/* debug [instance_properties/setter]: maximumFramesToRender */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputBufferSizeHint
func (a_ AudioUnit) MIDIOutputBufferSizeHint() int {
	rv := objc.Send[int](a_.ID, objc.Sel("MIDIOutputBufferSizeHint"))
	return rv
}/* debug [instance_properties/getter]: MIDIOutputBufferSizeHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputBufferSizeHint
func (a_ AudioUnit) SetMIDIOutputBufferSizeHint(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputBufferSizeHint:"), value)
}/* debug [instance_properties/setter]: MIDIOutputBufferSizeHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventBlock
func (a_ AudioUnit) MIDIOutputEventBlock() MIDIOutputEventBlock /* not a class type */ {
	rv := objc.Send[MIDIOutputEventBlock](a_.ID, objc.Sel("MIDIOutputEventBlock"))
	return rv
}/* debug [instance_properties/getter]: MIDIOutputEventBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventBlock
func (a_ AudioUnit) SetMIDIOutputEventBlock(value MIDIOutputEventBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputEventBlock:"), value)
}/* debug [instance_properties/setter]: MIDIOutputEventBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventListBlock
func (a_ AudioUnit) MIDIOutputEventListBlock() MIDIEventListBlock /* not a class type */ {
	rv := objc.Send[MIDIEventListBlock](a_.ID, objc.Sel("MIDIOutputEventListBlock"))
	return rv
}/* debug [instance_properties/getter]: MIDIOutputEventListBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventListBlock
func (a_ AudioUnit) SetMIDIOutputEventListBlock(value MIDIEventListBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputEventListBlock:"), value)
}/* debug [instance_properties/setter]: MIDIOutputEventListBlock */


// The names of the MIDI outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputNames
func (a_ AudioUnit) MIDIOutputNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("MIDIOutputNames"))
	return rv
}/* debug [instance_properties/getter]: MIDIOutputNames */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/migrateFromPlugin
func (a_ AudioUnit) MigrateFromPlugin() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("migrateFromPlugin"))
	return rv
}/* debug [instance_properties/getter]: migrateFromPlugin */


// A callback to the host for musical context information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/musicalContextBlock
func (a_ AudioUnit) MusicalContextBlock() HostMusicalContextBlock /* not a class type */ {
	rv := objc.Send[HostMusicalContextBlock](a_.ID, objc.Sel("musicalContextBlock"))
	return rv
}/* debug [instance_properties/getter]: musicalContextBlock */


// A callback to the host for musical context information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/musicalContextBlock
func (a_ AudioUnit) SetMusicalContextBlock(value HostMusicalContextBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMusicalContextBlock:"), value)
}/* debug [instance_properties/setter]: musicalContextBlock */


// The workgroup associated with the audio device underlying this Audio Unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/osWorkgroup
func (a_ AudioUnit) OsWorkgroup() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("osWorkgroup"))
	return rv
}/* debug [instance_properties/getter]: osWorkgroup */


// An array containing the audio unit’s output connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputBusses
func (a_ AudioUnit) OutputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("outputBusses"))
	return rv
}/* debug [instance_properties/getter]: outputBusses */


// The block that the output unit will call to get audio to send to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputProvider
func (a_ AudioUnit) OutputProvider() RenderPullInputBlock /* not a class type */ {
	rv := objc.Send[RenderPullInputBlock](a_.ID, objc.Sel("outputProvider"))
	return rv
}/* debug [instance_properties/getter]: outputProvider */


// The block that the output unit will call to get audio to send to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputProvider
func (a_ AudioUnit) SetOutputProvider(value RenderPullInputBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputProvider:"), value)
}/* debug [instance_properties/setter]: outputProvider */


// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parameterTree
func (a_ AudioUnit) ParameterTree() IAUParameterTree {
	rv := objc.Send[ParameterTree](a_.ID, objc.Sel("parameterTree"))
	return rv
}/* debug [instance_properties/getter]: parameterTree */


// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parameterTree
func (a_ AudioUnit) SetParameterTree(value IAUParameterTree) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParameterTree:"), value)
}/* debug [instance_properties/setter]: parameterTree */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileChangedBlock
func (a_ AudioUnit) ProfileChangedBlock() MIDICIProfileChangedBlock /* not a class type */ {
	rv := objc.Send[MIDICIProfileChangedBlock](a_.ID, objc.Sel("profileChangedBlock"))
	return rv
}/* debug [instance_properties/getter]: profileChangedBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileChangedBlock
func (a_ AudioUnit) SetProfileChangedBlock(value MIDICIProfileChangedBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProfileChangedBlock:"), value)
}/* debug [instance_properties/setter]: profileChangedBlock */


// A Boolean that indicates whether the audio unit provides a user interface, normally in the form of a view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/providesUserInterface
func (a_ AudioUnit) ProvidesUserInterface() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesUserInterface"))
	return rv
}/* debug [instance_properties/getter]: providesUserInterface */


// The block that hosts use to ask the audio unit to render audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderBlock
func (a_ AudioUnit) RenderBlock() RenderBlock /* not a class type */ {
	rv := objc.Send[RenderBlock](a_.ID, objc.Sel("renderBlock"))
	return rv
}/* debug [instance_properties/getter]: renderBlock */


// The block that the system calls when the rendering context changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderContextObserver
func (a_ AudioUnit) RenderContextObserver() RenderContextObserver /* not a class type */ {
	rv := objc.Send[RenderContextObserver](a_.ID, objc.Sel("renderContextObserver"))
	return rv
}/* debug [instance_properties/getter]: renderContextObserver */


// Provides a trade-off between rendering quality and CPU load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderQuality
func (a_ AudioUnit) RenderQuality() int {
	rv := objc.Send[int](a_.ID, objc.Sel("renderQuality"))
	return rv
}/* debug [instance_properties/getter]: renderQuality */


// Provides a trade-off between rendering quality and CPU load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderQuality
func (a_ AudioUnit) SetRenderQuality(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenderQuality:"), value)
}/* debug [instance_properties/setter]: renderQuality */


// Determines whether the audio unit has allocated render resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderResourcesAllocated
func (a_ AudioUnit) RenderResourcesAllocated() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("renderResourcesAllocated"))
	return rv
}/* debug [instance_properties/getter]: renderResourcesAllocated */


// A block used to schedule MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventBlock
func (a_ AudioUnit) ScheduleMIDIEventBlock() ScheduleMIDIEventBlock /* not a class type */ {
	rv := objc.Send[ScheduleMIDIEventBlock](a_.ID, objc.Sel("scheduleMIDIEventBlock"))
	return rv
}/* debug [instance_properties/getter]: scheduleMIDIEventBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventListBlock
func (a_ AudioUnit) ScheduleMIDIEventListBlock() MIDIEventListBlock /* not a class type */ {
	rv := objc.Send[MIDIEventListBlock](a_.ID, objc.Sel("scheduleMIDIEventListBlock"))
	return rv
}/* debug [instance_properties/getter]: scheduleMIDIEventListBlock */


// The block that hosts use to schedule parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleParameterBlock
func (a_ AudioUnit) ScheduleParameterBlock() ScheduleParameterBlock /* not a class type */ {
	rv := objc.Send[ScheduleParameterBlock](a_.ID, objc.Sel("scheduleParameterBlock"))
	return rv
}/* debug [instance_properties/getter]: scheduleParameterBlock */


// Determines whether an effect should route input directly to output, without any processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldBypassEffect
func (a_ AudioUnit) ShouldBypassEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldBypassEffect"))
	return rv
}/* debug [instance_properties/getter]: shouldBypassEffect */


// Determines whether an effect should route input directly to output, without any processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldBypassEffect
func (a_ AudioUnit) SetShouldBypassEffect(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldBypassEffect:"), value)
}/* debug [instance_properties/setter]: shouldBypassEffect */


// A Boolean value that indicates whether the audio unit supports multi-dimensional polyphonic expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsMPE
func (a_ AudioUnit) SupportsMPE() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsMPE"))
	return rv
}/* debug [instance_properties/getter]: supportsMPE */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsUserPresets
func (a_ AudioUnit) SupportsUserPresets() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsUserPresets"))
	return rv
}/* debug [instance_properties/getter]: supportsUserPresets */


// The audio unit’s tail time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/tailTime
func (a_ AudioUnit) TailTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("tailTime"))
	return rv
}/* debug [instance_properties/getter]: tailTime */


// A callback to the host for transport state information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/transportStateBlock
func (a_ AudioUnit) TransportStateBlock() HostTransportStateBlock /* not a class type */ {
	rv := objc.Send[HostTransportStateBlock](a_.ID, objc.Sel("transportStateBlock"))
	return rv
}/* debug [instance_properties/getter]: transportStateBlock */


// A callback to the host for transport state information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/transportStateBlock
func (a_ AudioUnit) SetTransportStateBlock(value HostTransportStateBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransportStateBlock:"), value)
}/* debug [instance_properties/setter]: transportStateBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/userPresets
func (a_ AudioUnit) UserPresets() []AudioUnitPreset {
	rv := objc.Send[[]AudioUnitPreset](a_.ID, objc.Sel("userPresets"))
	return rv
}/* debug [instance_properties/getter]: userPresets */


// The number of virtual MIDI cables implemented by a music device or effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/virtualMIDICableCount
func (a_ AudioUnit) VirtualMIDICableCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("virtualMIDICableCount"))
	return rv
}/* debug [instance_properties/getter]: virtualMIDICableCount */


// A flag enabling audio input from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isinputenabled
func (a_ AudioUnit) IsInputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInputEnabled"))
	return rv
}/* debug [instance_properties/getter]: isInputEnabled */


// A flag enabling audio input from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isinputenabled
func (a_ AudioUnit) SetIsInputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInputEnabled:"), value)
}/* debug [instance_properties/setter]: isInputEnabled */


// Specifies whether an audio unit responds to MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/ismusicdeviceoreffect
func (a_ AudioUnit) IsMusicDeviceOrEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMusicDeviceOrEffect"))
	return rv
}/* debug [instance_properties/getter]: isMusicDeviceOrEffect */


// Specifies whether an audio unit responds to MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/ismusicdeviceoreffect
func (a_ AudioUnit) SetIsMusicDeviceOrEffect(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMusicDeviceOrEffect:"), value)
}/* debug [instance_properties/setter]: isMusicDeviceOrEffect */


// A flag enabling audio output from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isoutputenabled
func (a_ AudioUnit) IsOutputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOutputEnabled"))
	return rv
}/* debug [instance_properties/getter]: isOutputEnabled */


// A flag enabling audio output from the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isoutputenabled
func (a_ AudioUnit) SetIsOutputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOutputEnabled:"), value)
}/* debug [instance_properties/setter]: isOutputEnabled */


// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrenderingoffline
func (a_ AudioUnit) IsRenderingOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRenderingOffline"))
	return rv
}/* debug [instance_properties/getter]: isRenderingOffline */


// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrenderingoffline
func (a_ AudioUnit) SetIsRenderingOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRenderingOffline:"), value)
}/* debug [instance_properties/setter]: isRenderingOffline */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrunning
func (a_ AudioUnit) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrunning
func (a_ AudioUnit) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetcpuloadkey
func (a_ AudioUnit) KAUPresetCPULoadKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetCPULoadKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetCPULoadKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetcpuloadkey
func (a_ AudioUnit) SetKAUPresetCPULoadKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetCPULoadKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetCPULoadKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetdatakey
func (a_ AudioUnit) KAUPresetDataKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetDataKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetDataKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetdatakey
func (a_ AudioUnit) SetKAUPresetDataKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetDataKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetDataKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetelementnamekey
func (a_ AudioUnit) KAUPresetElementNameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetElementNameKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetElementNameKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetelementnamekey
func (a_ AudioUnit) SetKAUPresetElementNameKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetElementNameKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetElementNameKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetexternalfilerefs
func (a_ AudioUnit) KAUPresetExternalFileRefs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetExternalFileRefs"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetExternalFileRefs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetexternalfilerefs
func (a_ AudioUnit) SetKAUPresetExternalFileRefs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetExternalFileRefs:"), value)
}/* debug [instance_properties/setter]: kAUPresetExternalFileRefs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmasdatakey
func (a_ AudioUnit) KAUPresetMASDataKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetMASDataKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetMASDataKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmasdatakey
func (a_ AudioUnit) SetKAUPresetMASDataKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetMASDataKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetMASDataKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmanufacturerkey
func (a_ AudioUnit) KAUPresetManufacturerKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetManufacturerKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetManufacturerKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmanufacturerkey
func (a_ AudioUnit) SetKAUPresetManufacturerKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetManufacturerKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetManufacturerKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnamekey
func (a_ AudioUnit) KAUPresetNameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetNameKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetNameKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnamekey
func (a_ AudioUnit) SetKAUPresetNameKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetNameKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetNameKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnumberkey
func (a_ AudioUnit) KAUPresetNumberKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetNumberKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetNumberKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnumberkey
func (a_ AudioUnit) SetKAUPresetNumberKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetNumberKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetNumberKey */


// If present, distinguishes a global preset that is set on the global scope from a part-based preset that is set on the part scope. The value of this key is defined by the audio unit it applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetpartkey
func (a_ AudioUnit) KAUPresetPartKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetPartKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetPartKey */


// If present, distinguishes a global preset that is set on the global scope from a part-based preset that is set on the part scope. The value of this key is defined by the audio unit it applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetpartkey
func (a_ AudioUnit) SetKAUPresetPartKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetPartKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetPartKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetrenderqualitykey
func (a_ AudioUnit) KAUPresetRenderQualityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetRenderQualityKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetRenderQualityKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetrenderqualitykey
func (a_ AudioUnit) SetKAUPresetRenderQualityKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetRenderQualityKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetRenderQualityKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetsubtypekey
func (a_ AudioUnit) KAUPresetSubtypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetSubtypeKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetSubtypeKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetsubtypekey
func (a_ AudioUnit) SetKAUPresetSubtypeKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetSubtypeKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetSubtypeKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresettypekey
func (a_ AudioUnit) KAUPresetTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetTypeKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetTypeKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresettypekey
func (a_ AudioUnit) SetKAUPresetTypeKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetTypeKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetTypeKey */


// VST state from a VST “bank.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstdatakey
func (a_ AudioUnit) KAUPresetVSTDataKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetVSTDataKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetVSTDataKey */


// VST state from a VST “bank.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstdatakey
func (a_ AudioUnit) SetKAUPresetVSTDataKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVSTDataKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetVSTDataKey */


// VST state from a VST “preset.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstpresetkey
func (a_ AudioUnit) KAUPresetVSTPresetKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetVSTPresetKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetVSTPresetKey */


// VST state from a VST “preset.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstpresetkey
func (a_ AudioUnit) SetKAUPresetVSTPresetKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVSTPresetKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetVSTPresetKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetversionkey
func (a_ AudioUnit) KAUPresetVersionKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("kAUPresetVersionKey"))
	return rv
}/* debug [instance_properties/getter]: kAUPresetVersionKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetversionkey
func (a_ AudioUnit) SetKAUPresetVersionKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVersionKey:"), value)
}/* debug [instance_properties/setter]: kAUPresetVersionKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUAudioUnit */


