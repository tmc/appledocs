// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioUnitComponent] class.
type IAudioUnitComponent interface {
	objectivec.IObject
	AllTagNames() string
	SetAllTagNames(value string)
	AudioComponent() unsafe.Pointer
	SetAudioComponent(value unsafe.Pointer)
	AudioComponentDescription() unsafe.Pointer
	SetAudioComponentDescription(value unsafe.Pointer)
	AvailableArchitectures() foundation.Number
	SetAvailableArchitectures(value foundation.INumber)
	ComponentURL() foundation.URL
	SetComponentURL(value foundation.IURL)
	ConfigurationDictionary() string
	SetConfigurationDictionary(value string)
	HasCustomView() bool
	SetHasCustomView(value bool)
	HasMIDIInput() bool
	SetHasMIDIInput(value bool)
	HasMIDIOutput() bool
	SetHasMIDIOutput(value bool)
	Icon() appkit.Image
	SetIcon(value appkit.IImage)
	IconURL() foundation.URL
	SetIconURL(value foundation.IURL)
	IsSandboxSafe() bool
	SetIsSandboxSafe(value bool)
	LocalizedTypeName() string
	SetLocalizedTypeName(value string)
	ManufacturerName() string
	SetManufacturerName(value string)
	Name() string
	SetName(value string)
	PassesAUVal() bool
	SetPassesAUVal(value bool)
	TypeName() string
	SetTypeName(value string)
	UserTagNames() string
	SetUserTagNames(value string)
	Version() int
	SetVersion(value int)
	VersionString() string
	SetVersionString(value string)
	AVAudioUnitManufacturerNameApple() string
	AVAudioUnitTypeEffect() string
	AVAudioUnitTypeFormatConverter() string
	AVAudioUnitTypeGenerator() string
	AVAudioUnitTypeMIDIProcessor() string
	AVAudioUnitTypeMixer() string
	AVAudioUnitTypeMusicDevice() string
	AVAudioUnitTypeMusicEffect() string
	AVAudioUnitTypeOfflineEffect() string
	AVAudioUnitTypeOutput() string
	AVAudioUnitTypePanner() string
}

// An object that provides details about an audio unit.
//
// Details can include information such as type, subtype, manufacturer, and location. An can include user tags, which you can query later for display.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitComponentClass) Alloc() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An array of tag names for the audio unit component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/alltagnames
func (a_ AudioUnitComponent) AllTagNames() string {
	rv := objc.Send[string](a_.ID, objc.Sel("allTagNames"))
	return rv
}


// SetAllTagNames sets the value of the allTagNames property.
// An array of tag names for the audio unit component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/alltagnames
func (a_ AudioUnitComponent) SetAllTagNames(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllTagNames:"), objc.String(value))
}

// The underlying audio component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/audiocomponent
func (a_ AudioUnitComponent) AudioComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioComponent"))
	return rv
}


// SetAudioComponent sets the value of the audioComponent property.
// The underlying audio component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/audiocomponent
func (a_ AudioUnitComponent) SetAudioComponent(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioComponent:"), value)
}

// The audio component description.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/audiocomponentdescription
func (a_ AudioUnitComponent) AudioComponentDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioComponentDescription"))
	return rv
}


// SetAudioComponentDescription sets the value of the audioComponentDescription property.
// The audio component description.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/audiocomponentdescription
func (a_ AudioUnitComponent) SetAudioComponentDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioComponentDescription:"), value)
}

// An array of architectures that the audio unit supports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/availablearchitectures
func (a_ AudioUnitComponent) AvailableArchitectures() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableArchitectures"))
	return rv
}


// SetAvailableArchitectures sets the value of the availableArchitectures property.
// An array of architectures that the audio unit supports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/availablearchitectures
func (a_ AudioUnitComponent) SetAvailableArchitectures(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableArchitectures:"), value)
}

// The URL of the audio unit component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/componenturl
func (a_ AudioUnitComponent) ComponentURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("componentURL"))
	return rv
}


// SetComponentURL sets the value of the componentURL property.
// The URL of the audio unit component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/componenturl
func (a_ AudioUnitComponent) SetComponentURL(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setComponentURL:"), value)
}

// The audio unit component’s configuration dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/configurationdictionary
func (a_ AudioUnitComponent) ConfigurationDictionary() string {
	rv := objc.Send[string](a_.ID, objc.Sel("configurationDictionary"))
	return rv
}


// SetConfigurationDictionary sets the value of the configurationDictionary property.
// The audio unit component’s configuration dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/configurationdictionary
func (a_ AudioUnitComponent) SetConfigurationDictionary(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationDictionary:"), objc.String(value))
}

// A Boolean value that indicates whether the audio unit component has a custom view.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hascustomview
func (a_ AudioUnitComponent) HasCustomView() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasCustomView"))
	return rv
}


// SetHasCustomView sets the value of the hasCustomView property.
// A Boolean value that indicates whether the audio unit component has a custom view.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hascustomview
func (a_ AudioUnitComponent) SetHasCustomView(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasCustomView:"), value)
}

// A Boolean value that indicates whether the audio unit component has MIDI input.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hasmidiinput
func (a_ AudioUnitComponent) HasMIDIInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasMIDIInput"))
	return rv
}


// SetHasMIDIInput sets the value of the hasMIDIInput property.
// A Boolean value that indicates whether the audio unit component has MIDI input.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hasmidiinput
func (a_ AudioUnitComponent) SetHasMIDIInput(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasMIDIInput:"), value)
}

// A Boolean value that indicates whether the audio unit component has MIDI output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hasmidioutput
func (a_ AudioUnitComponent) HasMIDIOutput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasMIDIOutput"))
	return rv
}


// SetHasMIDIOutput sets the value of the hasMIDIOutput property.
// A Boolean value that indicates whether the audio unit component has MIDI output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/hasmidioutput
func (a_ AudioUnitComponent) SetHasMIDIOutput(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasMIDIOutput:"), value)
}

// An icon that represents the component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/icon
func (a_ AudioUnitComponent) Icon() appkit.Image {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("icon"))
	return rv
}


// SetIcon sets the value of the icon property.
// An icon that represents the component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/icon
func (a_ AudioUnitComponent) SetIcon(value appkit.IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIcon:"), value)
}

// The URL of an icon that represents the audio unit component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/iconurl
func (a_ AudioUnitComponent) IconURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("iconURL"))
	return rv
}


// SetIconURL sets the value of the iconURL property.
// The URL of an icon that represents the audio unit component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/iconurl
func (a_ AudioUnitComponent) SetIconURL(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIconURL:"), value)
}

// A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/issandboxsafe
func (a_ AudioUnitComponent) IsSandboxSafe() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSandboxSafe"))
	return rv
}


// SetIsSandboxSafe sets the value of the isSandboxSafe property.
// A Boolean value that indicates whether the audio unit component is safe for sandboxing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/issandboxsafe
func (a_ AudioUnitComponent) SetIsSandboxSafe(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSandboxSafe:"), value)
}

// The localized type name of the component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/localizedtypename
func (a_ AudioUnitComponent) LocalizedTypeName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("localizedTypeName"))
	return rv
}


// SetLocalizedTypeName sets the value of the localizedTypeName property.
// The localized type name of the component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/localizedtypename
func (a_ AudioUnitComponent) SetLocalizedTypeName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocalizedTypeName:"), objc.String(value))
}

// The name of the manufacturer of the audio unit component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/manufacturername
func (a_ AudioUnitComponent) ManufacturerName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("manufacturerName"))
	return rv
}


// SetManufacturerName sets the value of the manufacturerName property.
// The name of the manufacturer of the audio unit component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/manufacturername
func (a_ AudioUnitComponent) SetManufacturerName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManufacturerName:"), objc.String(value))
}

// The name of the audio unit component.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/name
func (a_ AudioUnitComponent) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the audio unit component.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/name
func (a_ AudioUnitComponent) SetName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}

// A Boolean value that indicates whether the audio unit component passes the validation tests.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/passesauval
func (a_ AudioUnitComponent) PassesAUVal() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("passesAUVal"))
	return rv
}


// SetPassesAUVal sets the value of the passesAUVal property.
// A Boolean value that indicates whether the audio unit component passes the validation tests.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/passesauval
func (a_ AudioUnitComponent) SetPassesAUVal(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPassesAUVal:"), value)
}

// The audio unit component type.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/typename
func (a_ AudioUnitComponent) TypeName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("typeName"))
	return rv
}


// SetTypeName sets the value of the typeName property.
// The audio unit component type.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/typename
func (a_ AudioUnitComponent) SetTypeName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypeName:"), objc.String(value))
}

// An array of tags the user creates.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/usertagnames
func (a_ AudioUnitComponent) UserTagNames() string {
	rv := objc.Send[string](a_.ID, objc.Sel("userTagNames"))
	return rv
}


// SetUserTagNames sets the value of the userTagNames property.
// An array of tags the user creates.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/usertagnames
func (a_ AudioUnitComponent) SetUserTagNames(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserTagNames:"), objc.String(value))
}

// The audio unit component version number.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/version
func (a_ AudioUnitComponent) Version() int {
	rv := objc.Send[int](a_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The audio unit component version number.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/version
func (a_ AudioUnitComponent) SetVersion(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersion:"), value)
}

// A string that represents the audio unit component version number.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/versionstring
func (a_ AudioUnitComponent) VersionString() string {
	rv := objc.Send[string](a_.ID, objc.Sel("versionString"))
	return rv
}


// SetVersionString sets the value of the versionString property.
// A string that represents the audio unit component version number.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponent/versionstring
func (a_ AudioUnitComponent) SetVersionString(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersionString:"), objc.String(value))
}

// The audio unit manufacturer is Apple.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitmanufacturernameapple
func (a_ AudioUnitComponent) AVAudioUnitManufacturerNameApple() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitManufacturerNameApple"))
	return rv
}

// An audio unit type that represents an effect.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeeffect
func (a_ AudioUnitComponent) AVAudioUnitTypeEffect() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeEffect"))
	return rv
}

// An audio unit type that represents a format converter.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeformatconverter
func (a_ AudioUnitComponent) AVAudioUnitTypeFormatConverter() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeFormatConverter"))
	return rv
}

// An audio unit type that represents a generator.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypegenerator
func (a_ AudioUnitComponent) AVAudioUnitTypeGenerator() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeGenerator"))
	return rv
}

// An audio unit type that represents a MIDI processor.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemidiprocessor
func (a_ AudioUnitComponent) AVAudioUnitTypeMIDIProcessor() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeMIDIProcessor"))
	return rv
}

// An audio unit type that represents a mixer.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemixer
func (a_ AudioUnitComponent) AVAudioUnitTypeMixer() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeMixer"))
	return rv
}

// An audio unit type that represents a music device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusicdevice
func (a_ AudioUnitComponent) AVAudioUnitTypeMusicDevice() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeMusicDevice"))
	return rv
}

// An audio unit type that represents a music effect.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusiceffect
func (a_ AudioUnitComponent) AVAudioUnitTypeMusicEffect() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeMusicEffect"))
	return rv
}

// An audio unit type that represents an offline effect.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeofflineeffect
func (a_ AudioUnitComponent) AVAudioUnitTypeOfflineEffect() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeOfflineEffect"))
	return rv
}

// An audio unit type that represents an output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypeoutput
func (a_ AudioUnitComponent) AVAudioUnitTypeOutput() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypeOutput"))
	return rv
}

// An audio unit type that represents a panner.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittypepanner
func (a_ AudioUnitComponent) AVAudioUnitTypePanner() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVAudioUnitTypePanner"))
	return rv
}



