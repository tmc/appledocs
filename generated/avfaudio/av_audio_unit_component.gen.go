// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitComponent */


/* debug [class_header]: Header for AVAudioUnitComponent */
// The class instance for the [AudioUnitComponent] class.
var (
	AudioUnitComponentClass     _AudioUnitComponentClass
	AudioUnitComponentClassOnce sync.Once
)

func getAudioUnitComponentClass() _AudioUnitComponentClass {
	AudioUnitComponentClassOnce.Do(func() {
		AudioUnitComponentClass = _AudioUnitComponentClass{objc.GetClass("AVAudioUnitComponent")}
	})
	return AudioUnitComponentClass
}

type _AudioUnitComponentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitComponent */
// An interface definition for the [AudioUnitComponent] class.
type IAudioUnitComponent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnitComponent */
	// properties:
	AllTagNames() []string
	AudioComponent() objectivec.IObject
	AudioComponentDescription() audiotoolbox.AudioComponentDescription
	AvailableArchitectures() []foundation.Number
	ComponentURL() objc.IObject /* cross-framework: NSURL */
	ConfigurationDictionary() foundation.IDictionary
	HasCustomView() bool
	HasMIDIInput() bool
	HasMIDIOutput() bool
	Icon() appkit.Image
	IconURL() objc.IObject /* cross-framework: NSURL */
	SandboxSafe() bool
	LocalizedTypeName() objc.IObject /* cross-framework: NSString */
	ManufacturerName() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	PassesAUVal() bool
	TypeName() objc.IObject /* cross-framework: NSString */
	UserTagNames() []string
	SetUserTagNames(value []string)
	Version() uint
	VersionString() objc.IObject /* cross-framework: NSString */
	IsSandboxSafe() bool
	SetIsSandboxSafe(value bool)
	AVAudioUnitManufacturerNameApple() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeEffect() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeFormatConverter() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeGenerator() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeMIDIProcessor() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeMixer() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeMusicDevice() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeMusicEffect() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeOfflineEffect() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypeOutput() objc.IObject /* cross-framework: NSString */
	AVAudioUnitTypePanner() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitComponent */
	// methods:
	SupportsNumberInputChannelsOutputChannels(numInputChannels int, numOutputChannels int) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitComponent */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitComponentClass) Alloc() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitComponentClass) New() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitComponent) Init() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitComponent) Autorelease() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitComponent creates a new AudioUnitComponent instance.
func NewAudioUnitComponent() AudioUnitComponent {
	return getAudioUnitComponentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitComponent */
// An object that provides details about an audio unit.
//
// Details can include information such as type, subtype, manufacturer, and location. An can include user tags, which you can query later for display.


// An object that provides details about an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type AudioUnitComponent struct {
	objectivec.Object
}

// AudioUnitComponentFrom constructs a [AudioUnitComponent] from an unsafe.Pointer.
//
// An object that provides details about an audio unit.
func AudioUnitComponentFrom(ptr unsafe.Pointer) AudioUnitComponent {
	return AudioUnitComponent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitComponent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitComponent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitComponent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitComponent */

// Gets a Boolean value that indicates whether the audio unit component supports the specified number of input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/supportsNumberInputChannels(_:outputChannels:)
func (a_ AudioUnitComponent) SupportsNumberInputChannelsOutputChannels(numInputChannels int, numOutputChannels int) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsNumberInputChannels:outputChannels:"), numInputChannels, numOutputChannels)
	return rv
}/* debug [instance_methods/method]: SupportsNumberInputChannelsOutputChannels */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitComponent */

// An array of tag names for the audio unit component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/allTagNames
func (a_ AudioUnitComponent) AllTagNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("allTagNames"))
	return rv
}/* debug [instance_properties/getter]: allTagNames */


// The underlying audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/audioComponent
func (a_ AudioUnitComponent) AudioComponent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("audioComponent"))
	return rv
}/* debug [instance_properties/getter]: audioComponent */


// The audio component description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/audioComponentDescription
func (a_ AudioUnitComponent) AudioComponentDescription() audiotoolbox.AudioComponentDescription {
	rv := objc.Send[audiotoolbox.AudioComponentDescription](a_.ID, objc.Sel("audioComponentDescription"))
	return rv
}/* debug [instance_properties/getter]: audioComponentDescription */


// An array of architectures that the audio unit supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/availableArchitectures
func (a_ AudioUnitComponent) AvailableArchitectures() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("availableArchitectures"))
	return rv
}/* debug [instance_properties/getter]: availableArchitectures */


// The URL of the audio unit component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/componentURL
func (a_ AudioUnitComponent) ComponentURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("componentURL"))
	return rv
}/* debug [instance_properties/getter]: componentURL */


// The audio unit component’s configuration dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/configurationDictionary
func (a_ AudioUnitComponent) ConfigurationDictionary() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("configurationDictionary"))
	return rv
}/* debug [instance_properties/getter]: configurationDictionary */


// A Boolean value that indicates whether the audio unit component has a custom view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasCustomView
func (a_ AudioUnitComponent) HasCustomView() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasCustomView"))
	return rv
}/* debug [instance_properties/getter]: hasCustomView */


// A Boolean value that indicates whether the audio unit component has MIDI input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasMIDIInput
func (a_ AudioUnitComponent) HasMIDIInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasMIDIInput"))
	return rv
}/* debug [instance_properties/getter]: hasMIDIInput */


// A Boolean value that indicates whether the audio unit component has MIDI output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasMIDIOutput
func (a_ AudioUnitComponent) HasMIDIOutput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasMIDIOutput"))
	return rv
}/* debug [instance_properties/getter]: hasMIDIOutput */


// An icon that represents the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/icon
func (a_ AudioUnitComponent) Icon() appkit.Image {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The URL of an icon that represents the audio unit component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/iconURL
func (a_ AudioUnitComponent) IconURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("iconURL"))
	return rv
}/* debug [instance_properties/getter]: iconURL */


// A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/isSandboxSafe
func (a_ AudioUnitComponent) SandboxSafe() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sandboxSafe"))
	return rv
}/* debug [instance_properties/getter]: sandboxSafe */


// The localized type name of the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/localizedTypeName
func (a_ AudioUnitComponent) LocalizedTypeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedTypeName"))
	return rv
}/* debug [instance_properties/getter]: localizedTypeName */


// The name of the manufacturer of the audio unit component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/manufacturerName
func (a_ AudioUnitComponent) ManufacturerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("manufacturerName"))
	return rv
}/* debug [instance_properties/getter]: manufacturerName */


// The name of the audio unit component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/name
func (a_ AudioUnitComponent) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A Boolean value that indicates whether the audio unit component passes the validation tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/passesAUVal
func (a_ AudioUnitComponent) PassesAUVal() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("passesAUVal"))
	return rv
}/* debug [instance_properties/getter]: passesAUVal */


// The audio unit component type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/typeName
func (a_ AudioUnitComponent) TypeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("typeName"))
	return rv
}/* debug [instance_properties/getter]: typeName */


// An array of tags the user creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/userTagNames
func (a_ AudioUnitComponent) UserTagNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("userTagNames"))
	return rv
}/* debug [instance_properties/getter]: userTagNames */


// An array of tags the user creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/userTagNames
func (a_ AudioUnitComponent) SetUserTagNames(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserTagNames:"), nsArray)
}/* debug [instance_properties/setter]: userTagNames */


// The audio unit component version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/version
func (a_ AudioUnitComponent) Version() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */


// A string that represents the audio unit component version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/versionString
func (a_ AudioUnitComponent) VersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("versionString"))
	return rv
}/* debug [instance_properties/getter]: versionString */


// A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/issandboxsafe
func (a_ AudioUnitComponent) IsSandboxSafe() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSandboxSafe"))
	return rv
}/* debug [instance_properties/getter]: isSandboxSafe */


// A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/issandboxsafe
func (a_ AudioUnitComponent) SetIsSandboxSafe(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSandboxSafe:"), value)
}/* debug [instance_properties/setter]: isSandboxSafe */


// The audio unit manufacturer is Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitmanufacturernameapple
func (a_ AudioUnitComponent) AVAudioUnitManufacturerNameApple() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitManufacturerNameApple"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitManufacturerNameApple */


// An audio unit type that represents an effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeeffect
func (a_ AudioUnitComponent) AVAudioUnitTypeEffect() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeEffect"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeEffect */


// An audio unit type that represents a format converter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeformatconverter
func (a_ AudioUnitComponent) AVAudioUnitTypeFormatConverter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeFormatConverter"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeFormatConverter */


// An audio unit type that represents a generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypegenerator
func (a_ AudioUnitComponent) AVAudioUnitTypeGenerator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeGenerator"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeGenerator */


// An audio unit type that represents a MIDI processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemidiprocessor
func (a_ AudioUnitComponent) AVAudioUnitTypeMIDIProcessor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeMIDIProcessor"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeMIDIProcessor */


// An audio unit type that represents a mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemixer
func (a_ AudioUnitComponent) AVAudioUnitTypeMixer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeMixer"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeMixer */


// An audio unit type that represents a music device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusicdevice
func (a_ AudioUnitComponent) AVAudioUnitTypeMusicDevice() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeMusicDevice"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeMusicDevice */


// An audio unit type that represents a music effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusiceffect
func (a_ AudioUnitComponent) AVAudioUnitTypeMusicEffect() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeMusicEffect"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeMusicEffect */


// An audio unit type that represents an offline effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeofflineeffect
func (a_ AudioUnitComponent) AVAudioUnitTypeOfflineEffect() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeOfflineEffect"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeOfflineEffect */


// An audio unit type that represents an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeoutput
func (a_ AudioUnitComponent) AVAudioUnitTypeOutput() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypeOutput"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypeOutput */


// An audio unit type that represents a panner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypepanner
func (a_ AudioUnitComponent) AVAudioUnitTypePanner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioUnitTypePanner"))
	return rv
}/* debug [instance_properties/getter]: AVAudioUnitTypePanner */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitComponent */



