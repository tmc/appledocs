// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	objectivec.IObject
	AllocateRenderResourcesAndReturnError(outError unsafe.Pointer) bool
	DeallocateRenderResources()
	DeleteUserPresetError(userPreset unsafe.Pointer, outError unsafe.Pointer) bool
	DisableProfileCableOnChannelError(profile unsafe.Pointer, cable unsafe.Pointer, channel unsafe.Pointer, outError unsafe.Pointer) bool
	EnableProfileCableOnChannelError(profile unsafe.Pointer, cable unsafe.Pointer, channel unsafe.Pointer, outError unsafe.Pointer) bool
	MessageChannelFor(channelName string) objc.ID
	ParametersForOverviewWithCount(count int) []foundation.Number
	PresetStateForError(userPreset unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer
	ProfileStateForCableChannel(cable unsafe.Pointer, channel unsafe.Pointer) unsafe.Pointer
	RemoveRenderObserver(token int)
	RequestViewControllerWithCompletionHandler(completionHandler unsafe.Pointer)
	Reset()
	SaveUserPresetError(userPreset unsafe.Pointer, outError unsafe.Pointer) bool
	SelectViewConfiguration(viewConfiguration unsafe.Pointer)
	SetDeviceIDError(deviceID unsafe.Pointer, outError unsafe.Pointer) bool
	ShouldChangeToFormatForBus(format unsafe.Pointer, bus unsafe.Pointer) bool
	StartHardwareAndReturnError(outError unsafe.Pointer) bool
	StopHardware()
	SupportedViewConfigurations(availableViewConfigurations unsafe.Pointer) unsafe.Pointer
	TokenByAddingRenderObserver(observer unsafe.Pointer) int
}

// A class that defines a host’s interface to an audio unit.
//
// Hosts can instantiate either version 3 or version 2 audio units with this class, and to some extent control whether an audio unit is instantiated in-process or in a separate extension process. Version 3 audio units should subclass the class. Version 3 audio unit components can be registered in the following ways: Package the component into an app extension containing an entry. The principal class must conform to the protocol, which will typically instantiate an subclass. Call the method to associate a component description with an subclass. Use the convention when naming your audio unit component. Version 2 audio units should subclass the class instead. Version 2 audio unit components can be registered in the following ways: Package the component into a component bundle containing an entry, referring to an function. Call the function to associate a component description with an function. A host does not need to be aware of the concrete subclass that is being instantiated. The method ensures that the proper subclass is used.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Synchronously initializes a new audio unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:)
func NewAudioUnitWithComponentDescriptionError(componentDescription unsafe.Pointer, outError unsafe.Pointer) AudioUnit {
	instance := getAudioUnitClass().Alloc()
	rv := objc.Send[AudioUnit](instance.ID, objc.Sel("initWithComponentDescription:error:"), componentDescription, outError)
	rv.Autorelease()
	return rv
}



// Synchronously initializes a new audio unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:options:)
func NewAudioUnitWithComponentDescriptionOptionsError(componentDescription unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) AudioUnit {
	instance := getAudioUnitClass().Alloc()
	rv := objc.Send[AudioUnit](instance.ID, objc.Sel("initWithComponentDescription:options:error:"), componentDescription, options, outError)
	rv.Autorelease()
	return rv
}


// Asynchronously creates an audio unit instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/instantiate(with:options:completionHandler:)
func (ac _AudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(componentDescription unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), componentDescription, options, completionHandler)
}

// Registers an audio unit subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/registerSubclass(_:as:name:version:)
func (ac _AudioUnitClass) RegisterSubclassAsComponentDescriptionNameVersion(cls objc.Class, componentDescription unsafe.Pointer, name string, version unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("registerSubclass:asComponentDescription:name:version:"), cls, componentDescription, objc.String(name), version)
}

// Allocates resources required to render audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allocateRenderResources()
func (a_ AudioUnit) AllocateRenderResourcesAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allocateRenderResourcesAndReturnError:"), outError)
	return rv
}

// Deallocates resources required to render audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deallocateRenderResources()
func (a_ AudioUnit) DeallocateRenderResources() {
	objc.Send[objc.ID](a_.ID, objc.Sel("deallocateRenderResources"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deleteUserPreset(_:)
func (a_ AudioUnit) DeleteUserPresetError(userPreset unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("deleteUserPreset:error:"), userPreset, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/disableProfile(_:cable:onChannel:)
func (a_ AudioUnit) DisableProfileCableOnChannelError(profile unsafe.Pointer, cable unsafe.Pointer, channel unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("disableProfile:cable:onChannel:error:"), profile, cable, channel, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/enable(_:cable:onChannel:)
func (a_ AudioUnit) EnableProfileCableOnChannelError(profile unsafe.Pointer, cable unsafe.Pointer, channel unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableProfile:cable:onChannel:error:"), profile, cable, channel, outError)
	return rv
}

// Returns an object for bidirectional communication between an audio unit and its host.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/messageChannel(for:)
func (a_ AudioUnit) MessageChannelFor(channelName string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("messageChannelFor:"), objc.String(channelName))
	return rv
}

// Returns the audio unit’s most important parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parametersForOverview(withCount:)
func (a_ AudioUnit) ParametersForOverviewWithCount(count int) []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("parametersForOverviewWithCount:"), count)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/presetState(for:)
func (a_ AudioUnit) PresetStateForError(userPreset unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presetStateFor:error:"), userPreset, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileState(forCable:channel:)
func (a_ AudioUnit) ProfileStateForCableChannel(cable unsafe.Pointer, channel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("profileStateForCable:channel:"), cable, channel)
	return rv
}

// Removes an observer block previously added to the render cycle.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/removeRenderObserver(_:)
func (a_ AudioUnit) RemoveRenderObserver(token int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeRenderObserver:"), token)
}

// Requests an audio unit’s custom view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/requestViewController(completionHandler:)
func (a_ AudioUnit) RequestViewControllerWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestViewControllerWithCompletionHandler:"), completionHandler)
}

// Resets transitory rendering state to its initial state.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/reset()
func (a_ AudioUnit) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/saveUserPreset(_:)
func (a_ AudioUnit) SaveUserPresetError(userPreset unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveUserPreset:error:"), userPreset, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/select(_:)
func (a_ AudioUnit) SelectViewConfiguration(viewConfiguration unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectViewConfiguration:"), viewConfiguration)
}

// Sets the I/O hardware device.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/setDeviceID(_:)
func (a_ AudioUnit) SetDeviceIDError(deviceID unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setDeviceID:error:"), deviceID, outError)
	return rv
}

// This is called when you set the format on a bus.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldChange(to:for:)
func (a_ AudioUnit) ShouldChangeToFormatForBus(format unsafe.Pointer, bus unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldChangeToFormat:forBus:"), format, bus)
	return rv
}

// Starts the audio hardware.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/startHardware()
func (a_ AudioUnit) StartHardwareAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startHardwareAndReturnError:"), outError)
	return rv
}

// Stops the audio hardware.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/stopHardware()
func (a_ AudioUnit) StopHardware() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopHardware"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportedViewConfigurations(_:)
func (a_ AudioUnit) SupportedViewConfigurations(availableViewConfigurations unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("supportedViewConfigurations:"), availableViewConfigurations)
	return rv
}

// Adds a block to be called on each render cycle.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/token(byAddingRenderObserver:)
func (a_ AudioUnit) TokenByAddingRenderObserver(observer unsafe.Pointer) int {
	rv := objc.Send[int](a_.ID, objc.Sel("tokenByAddingRenderObserver:"), observer)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnamekey
func (a_ AudioUnit) KAUPresetNameKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetNameKey"))
	return rv
}


// SetKAUPresetNameKey sets the value of the kAUPresetNameKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnamekey
func (a_ AudioUnit) SetKAUPresetNameKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetNameKey:"), objc.String(value))
}

// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrenderingoffline
func (a_ AudioUnit) IsRenderingOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRenderingOffline"))
	return rv
}


// SetIsRenderingOffline sets the value of the isRenderingOffline property.
// Communicates to an audio unit that it is rendering offline.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrenderingoffline
func (a_ AudioUnit) SetIsRenderingOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRenderingOffline:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetdatakey
func (a_ AudioUnit) KAUPresetDataKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetDataKey"))
	return rv
}


// SetKAUPresetDataKey sets the value of the kAUPresetDataKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetdatakey
func (a_ AudioUnit) SetKAUPresetDataKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetDataKey:"), objc.String(value))
}

// Specifies whether an audio unit responds to MIDI events.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/ismusicdeviceoreffect
func (a_ AudioUnit) IsMusicDeviceOrEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMusicDeviceOrEffect"))
	return rv
}


// SetIsMusicDeviceOrEffect sets the value of the isMusicDeviceOrEffect property.
// Specifies whether an audio unit responds to MIDI events.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/ismusicdeviceoreffect
func (a_ AudioUnit) SetIsMusicDeviceOrEffect(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMusicDeviceOrEffect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetexternalfilerefs
func (a_ AudioUnit) KAUPresetExternalFileRefs() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetExternalFileRefs"))
	return rv
}


// SetKAUPresetExternalFileRefs sets the value of the kAUPresetExternalFileRefs property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetexternalfilerefs
func (a_ AudioUnit) SetKAUPresetExternalFileRefs(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetExternalFileRefs:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrunning
func (a_ AudioUnit) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isrunning
func (a_ AudioUnit) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}

// A flag enabling audio output from the unit.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isoutputenabled
func (a_ AudioUnit) IsOutputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOutputEnabled"))
	return rv
}


// SetIsOutputEnabled sets the value of the isOutputEnabled property.
// A flag enabling audio output from the unit.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isoutputenabled
func (a_ AudioUnit) SetIsOutputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOutputEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetelementnamekey
func (a_ AudioUnit) KAUPresetElementNameKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetElementNameKey"))
	return rv
}


// SetKAUPresetElementNameKey sets the value of the kAUPresetElementNameKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetelementnamekey
func (a_ AudioUnit) SetKAUPresetElementNameKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetElementNameKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetsubtypekey
func (a_ AudioUnit) KAUPresetSubtypeKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetSubtypeKey"))
	return rv
}


// SetKAUPresetSubtypeKey sets the value of the kAUPresetSubtypeKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetsubtypekey
func (a_ AudioUnit) SetKAUPresetSubtypeKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetSubtypeKey:"), objc.String(value))
}

// If present, distinguishes a global preset that is set on the global scope from a part-based preset that is set on the part scope. The value of this key is defined by the audio unit it applies to.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetpartkey
func (a_ AudioUnit) KAUPresetPartKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetPartKey"))
	return rv
}


// SetKAUPresetPartKey sets the value of the kAUPresetPartKey property.
// If present, distinguishes a global preset that is set on the global scope from a part-based preset that is set on the part scope. The value of this key is defined by the audio unit it applies to.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetpartkey
func (a_ AudioUnit) SetKAUPresetPartKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetPartKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmasdatakey
func (a_ AudioUnit) KAUPresetMASDataKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetMASDataKey"))
	return rv
}


// SetKAUPresetMASDataKey sets the value of the kAUPresetMASDataKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmasdatakey
func (a_ AudioUnit) SetKAUPresetMASDataKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetMASDataKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetversionkey
func (a_ AudioUnit) KAUPresetVersionKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetVersionKey"))
	return rv
}


// SetKAUPresetVersionKey sets the value of the kAUPresetVersionKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetversionkey
func (a_ AudioUnit) SetKAUPresetVersionKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVersionKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetrenderqualitykey
func (a_ AudioUnit) KAUPresetRenderQualityKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetRenderQualityKey"))
	return rv
}


// SetKAUPresetRenderQualityKey sets the value of the kAUPresetRenderQualityKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetrenderqualitykey
func (a_ AudioUnit) SetKAUPresetRenderQualityKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetRenderQualityKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetcpuloadkey
func (a_ AudioUnit) KAUPresetCPULoadKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetCPULoadKey"))
	return rv
}


// SetKAUPresetCPULoadKey sets the value of the kAUPresetCPULoadKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetcpuloadkey
func (a_ AudioUnit) SetKAUPresetCPULoadKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetCPULoadKey:"), objc.String(value))
}

// A flag enabling audio input from the unit.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isinputenabled
func (a_ AudioUnit) IsInputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInputEnabled"))
	return rv
}


// SetIsInputEnabled sets the value of the isInputEnabled property.
// A flag enabling audio input from the unit.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/isinputenabled
func (a_ AudioUnit) SetIsInputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInputEnabled:"), value)
}

// VST state from a VST “bank.”
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstdatakey
func (a_ AudioUnit) KAUPresetVSTDataKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetVSTDataKey"))
	return rv
}


// SetKAUPresetVSTDataKey sets the value of the kAUPresetVSTDataKey property.
// VST state from a VST “bank.”

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstdatakey
func (a_ AudioUnit) SetKAUPresetVSTDataKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVSTDataKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresettypekey
func (a_ AudioUnit) KAUPresetTypeKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetTypeKey"))
	return rv
}


// SetKAUPresetTypeKey sets the value of the kAUPresetTypeKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresettypekey
func (a_ AudioUnit) SetKAUPresetTypeKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetTypeKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmanufacturerkey
func (a_ AudioUnit) KAUPresetManufacturerKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetManufacturerKey"))
	return rv
}


// SetKAUPresetManufacturerKey sets the value of the kAUPresetManufacturerKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetmanufacturerkey
func (a_ AudioUnit) SetKAUPresetManufacturerKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetManufacturerKey:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnumberkey
func (a_ AudioUnit) KAUPresetNumberKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetNumberKey"))
	return rv
}


// SetKAUPresetNumberKey sets the value of the kAUPresetNumberKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetnumberkey
func (a_ AudioUnit) SetKAUPresetNumberKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetNumberKey:"), objc.String(value))
}

// VST state from a VST “preset.”
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstpresetkey
func (a_ AudioUnit) KAUPresetVSTPresetKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("kAUPresetVSTPresetKey"))
	return rv
}


// SetKAUPresetVSTPresetKey sets the value of the kAUPresetVSTPresetKey property.
// VST state from a VST “preset.”

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/kaupresetvstpresetkey
func (a_ AudioUnit) SetKAUPresetVSTPresetKey(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAUPresetVSTPresetKey:"), objc.String(value))
}

// Special read-only property for KVO.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allParameterValues
func (a_ AudioUnit) AllParameterValues() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allParameterValues"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitMIDIProtocol
func (a_ AudioUnit) AudioUnitMIDIProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("AudioUnitMIDIProtocol"))
	return rv
}

// The audio unit’s name, derived from the component’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitName
func (a_ AudioUnit) AudioUnitName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("audioUnitName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitShortName
func (a_ AudioUnit) AudioUnitShortName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("audioUnitShortName"))
	return rv
}

// Determines whether the I/O device can perform input.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformInput
func (a_ AudioUnit) CanPerformInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformInput"))
	return rv
}

// Determines whether the I/O device can perform output.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformOutput
func (a_ AudioUnit) CanPerformOutput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformOutput"))
	return rv
}

// Determines whether an audio unit can process in place.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canProcessInPlace
func (a_ AudioUnit) CanProcessInPlace() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProcessInPlace"))
	return rv
}

// Expresses valid combinations of input and output channels.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelCapabilities
func (a_ AudioUnit) ChannelCapabilities() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("channelCapabilities"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelMap
func (a_ AudioUnit) ChannelMap() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("channelMap"))
	return rv
}


// SetChannelMap sets the value of the channelMap property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelMap
func (a_ AudioUnit) SetChannelMap(value []foundation.Number) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelMap:"), nsArray)
}

// The component found in the component description with which the audio unit was created.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/component
func (a_ AudioUnit) Component() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("component"))
	return rv
}

// The component description with which the audio unit was created.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentDescription
func (a_ AudioUnit) ComponentDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("componentDescription"))
	return rv
}

// The audio unit’s component’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentName
func (a_ AudioUnit) ComponentName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("componentName"))
	return rv
}

// The audio unit’s component’s version.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentVersion
func (a_ AudioUnit) ComponentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("componentVersion"))
	return rv
}

// Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/contextName
func (a_ AudioUnit) ContextName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("contextName"))
	return rv
}


// SetContextName sets the value of the contextName property.
// Information about the host context in which the audio unit is connected, for display in the audio unit’s view.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/contextName
func (a_ AudioUnit) SetContextName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContextName:"), objc.String(value))
}

// The audio unit’s last-selected preset.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/currentPreset
func (a_ AudioUnit) CurrentPreset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentPreset"))
	return rv
}


// SetCurrentPreset sets the value of the currentPreset property.
// The audio unit’s last-selected preset.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/currentPreset
func (a_ AudioUnit) SetCurrentPreset(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPreset:"), value)
}

// Gets the I/O hardware device.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceID
func (a_ AudioUnit) DeviceID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("deviceID"))
	return rv
}

// The audio device’s input latency, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceInputLatency
func (a_ AudioUnit) DeviceInputLatency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("deviceInputLatency"))
	return rv
}

// The audio devic’s output latency, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceOutputLatency
func (a_ AudioUnit) DeviceOutputLatency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("deviceOutputLatency"))
	return rv
}

// A collection of presets provided by the audio unit’s developer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/factoryPresets
func (a_ AudioUnit) FactoryPresets() []AudioUnitPreset {
	rv := objc.Send[[]AudioUnitPreset](a_.ID, objc.Sel("factoryPresets"))
	return rv
}

// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullState
func (a_ AudioUnit) FullState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fullState"))
	return rv
}


// SetFullState sets the value of the fullState property.
// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullState
func (a_ AudioUnit) SetFullState(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullState:"), value)
}

// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullStateForDocument
func (a_ AudioUnit) FullStateForDocument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fullStateForDocument"))
	return rv
}


// SetFullStateForDocument sets the value of the fullStateForDocument property.
// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullStateForDocument
func (a_ AudioUnit) SetFullStateForDocument(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullStateForDocument:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/hostMIDIProtocol
func (a_ AudioUnit) HostMIDIProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("hostMIDIProtocol"))
	return rv
}


// SetHostMIDIProtocol sets the value of the hostMIDIProtocol property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/hostMIDIProtocol
func (a_ AudioUnit) SetHostMIDIProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHostMIDIProtocol:"), value)
}

// An array containing the audio unit’s input connection points.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputBusses
func (a_ AudioUnit) InputBusses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputBusses"))
	return rv
}

// The block that the output unit will call to notify when input is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputHandler
func (a_ AudioUnit) InputHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputHandler"))
	return rv
}


// SetInputHandler sets the value of the inputHandler property.
// The block that the output unit will call to notify when input is available.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputHandler
func (a_ AudioUnit) SetInputHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputHandler:"), value)
}

// The AUAudioUnit’s intended spatial experience.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/intendedSpatialExperience-1dvhd
func (a_ AudioUnit) IntendedSpatialExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}


// SetIntendedSpatialExperience sets the value of the intendedSpatialExperience property.
// The AUAudioUnit’s intended spatial experience.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/intendedSpatialExperience-1dvhd
func (a_ AudioUnit) SetIntendedSpatialExperience(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}

// The block which you must provide, via a getter, in order to implement rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/internalRenderBlock
func (a_ AudioUnit) InternalRenderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("internalRenderBlock"))
	return rv
}

// A flag enabling audio input from the unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isInputEnabled
func (a_ AudioUnit) InputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputEnabled"))
	return rv
}


// SetInputEnabled sets the value of the inputEnabled property.
// A flag enabling audio input from the unit.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isInputEnabled
func (a_ AudioUnit) SetInputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isLoadedInProcess
func (a_ AudioUnit) IsLoadedInProcess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isLoadedInProcess"))
	return rv
}

// Specifies whether an audio unit responds to MIDI events.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isMusicDeviceOrEffect
func (a_ AudioUnit) MusicDeviceOrEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("musicDeviceOrEffect"))
	return rv
}

// A flag enabling audio output from the unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isOutputEnabled
func (a_ AudioUnit) OutputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("outputEnabled"))
	return rv
}


// SetOutputEnabled sets the value of the outputEnabled property.
// A flag enabling audio output from the unit.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isOutputEnabled
func (a_ AudioUnit) SetOutputEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputEnabled:"), value)
}

// Communicates to an audio unit that it is rendering offline.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRenderingOffline
func (a_ AudioUnit) RenderingOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("renderingOffline"))
	return rv
}


// SetRenderingOffline sets the value of the renderingOffline property.
// Communicates to an audio unit that it is rendering offline.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRenderingOffline
func (a_ AudioUnit) SetRenderingOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenderingOffline:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRunning
func (a_ AudioUnit) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}

// The audio unit’s processing latency, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/latency
func (a_ AudioUnit) Latency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("latency"))
	return rv
}

// The manufacturer’s name, derived from the component’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/manufacturerName
func (a_ AudioUnit) ManufacturerName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("manufacturerName"))
	return rv
}

// The maximum number of frames that the audio unit can render at once.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/maximumFramesToRender
func (a_ AudioUnit) MaximumFramesToRender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("maximumFramesToRender"))
	return rv
}


// SetMaximumFramesToRender sets the value of the maximumFramesToRender property.
// The maximum number of frames that the audio unit can render at once.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/maximumFramesToRender
func (a_ AudioUnit) SetMaximumFramesToRender(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumFramesToRender:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputBufferSizeHint
func (a_ AudioUnit) MIDIOutputBufferSizeHint() int {
	rv := objc.Send[int](a_.ID, objc.Sel("MIDIOutputBufferSizeHint"))
	return rv
}


// SetMIDIOutputBufferSizeHint sets the value of the MIDIOutputBufferSizeHint property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputBufferSizeHint
func (a_ AudioUnit) SetMIDIOutputBufferSizeHint(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputBufferSizeHint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventBlock
func (a_ AudioUnit) MIDIOutputEventBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("MIDIOutputEventBlock"))
	return rv
}


// SetMIDIOutputEventBlock sets the value of the MIDIOutputEventBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventBlock
func (a_ AudioUnit) SetMIDIOutputEventBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputEventBlock:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventListBlock
func (a_ AudioUnit) MIDIOutputEventListBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("MIDIOutputEventListBlock"))
	return rv
}


// SetMIDIOutputEventListBlock sets the value of the MIDIOutputEventListBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventListBlock
func (a_ AudioUnit) SetMIDIOutputEventListBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMIDIOutputEventListBlock:"), value)
}

// The names of the MIDI outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputNames
func (a_ AudioUnit) MIDIOutputNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("MIDIOutputNames"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/migrateFromPlugin
func (a_ AudioUnit) MigrateFromPlugin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("migrateFromPlugin"))
	return rv
}

// A callback to the host for musical context information.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/musicalContextBlock
func (a_ AudioUnit) MusicalContextBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("musicalContextBlock"))
	return rv
}


// SetMusicalContextBlock sets the value of the musicalContextBlock property.
// A callback to the host for musical context information.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/musicalContextBlock
func (a_ AudioUnit) SetMusicalContextBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMusicalContextBlock:"), value)
}

// The workgroup associated with the audio device underlying this Audio Unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/osWorkgroup
func (a_ AudioUnit) OsWorkgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("osWorkgroup"))
	return rv
}

// An array containing the audio unit’s output connection points.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputBusses
func (a_ AudioUnit) OutputBusses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputBusses"))
	return rv
}

// The block that the output unit will call to get audio to send to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputProvider
func (a_ AudioUnit) OutputProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputProvider"))
	return rv
}


// SetOutputProvider sets the value of the outputProvider property.
// The block that the output unit will call to get audio to send to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputProvider
func (a_ AudioUnit) SetOutputProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputProvider:"), value)
}

// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parameterTree
func (a_ AudioUnit) ParameterTree() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("parameterTree"))
	return rv
}


// SetParameterTree sets the value of the parameterTree property.
// An audio unit’s parameters, organized in a tree hierarchy.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parameterTree
func (a_ AudioUnit) SetParameterTree(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParameterTree:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileChangedBlock
func (a_ AudioUnit) ProfileChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("profileChangedBlock"))
	return rv
}


// SetProfileChangedBlock sets the value of the profileChangedBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileChangedBlock
func (a_ AudioUnit) SetProfileChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProfileChangedBlock:"), value)
}

// A Boolean that indicates whether the audio unit provides a user interface, normally in the form of a view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/providesUserInterface
func (a_ AudioUnit) ProvidesUserInterface() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesUserInterface"))
	return rv
}

// The block that hosts use to ask the audio unit to render audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderBlock
func (a_ AudioUnit) RenderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("renderBlock"))
	return rv
}

// The block that the system calls when the rendering context changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderContextObserver
func (a_ AudioUnit) RenderContextObserver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("renderContextObserver"))
	return rv
}

// Provides a trade-off between rendering quality and CPU load.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderQuality
func (a_ AudioUnit) RenderQuality() int {
	rv := objc.Send[int](a_.ID, objc.Sel("renderQuality"))
	return rv
}


// SetRenderQuality sets the value of the renderQuality property.
// Provides a trade-off between rendering quality and CPU load.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderQuality
func (a_ AudioUnit) SetRenderQuality(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenderQuality:"), value)
}

// Determines whether the audio unit has allocated render resources.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderResourcesAllocated
func (a_ AudioUnit) RenderResourcesAllocated() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("renderResourcesAllocated"))
	return rv
}

// A block used to schedule MIDI events.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventBlock
func (a_ AudioUnit) ScheduleMIDIEventBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("scheduleMIDIEventBlock"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventListBlock
func (a_ AudioUnit) ScheduleMIDIEventListBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("scheduleMIDIEventListBlock"))
	return rv
}

// The block that hosts use to schedule parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleParameterBlock
func (a_ AudioUnit) ScheduleParameterBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("scheduleParameterBlock"))
	return rv
}

// Determines whether an effect should route input directly to output, without any processing.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldBypassEffect
func (a_ AudioUnit) ShouldBypassEffect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldBypassEffect"))
	return rv
}


// SetShouldBypassEffect sets the value of the shouldBypassEffect property.
// Determines whether an effect should route input directly to output, without any processing.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldBypassEffect
func (a_ AudioUnit) SetShouldBypassEffect(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldBypassEffect:"), value)
}

// A Boolean value that indicates whether the audio unit supports multi-dimensional polyphonic expression.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsMPE
func (a_ AudioUnit) SupportsMPE() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsMPE"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsUserPresets
func (a_ AudioUnit) SupportsUserPresets() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsUserPresets"))
	return rv
}

// The audio unit’s tail time, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/tailTime
func (a_ AudioUnit) TailTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("tailTime"))
	return rv
}

// A callback to the host for transport state information.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/transportStateBlock
func (a_ AudioUnit) TransportStateBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transportStateBlock"))
	return rv
}


// SetTransportStateBlock sets the value of the transportStateBlock property.
// A callback to the host for transport state information.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/transportStateBlock
func (a_ AudioUnit) SetTransportStateBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransportStateBlock:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/userPresets
func (a_ AudioUnit) UserPresets() []AudioUnitPreset {
	rv := objc.Send[[]AudioUnitPreset](a_.ID, objc.Sel("userPresets"))
	return rv
}

// The number of virtual MIDI cables implemented by a music device or effect.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/virtualMIDICableCount
func (a_ AudioUnit) VirtualMIDICableCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("virtualMIDICableCount"))
	return rv
}


