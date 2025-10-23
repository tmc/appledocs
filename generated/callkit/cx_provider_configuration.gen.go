// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXProviderConfiguration] class.
var (
	CXProviderConfigurationClass     _CXProviderConfigurationClass
	CXProviderConfigurationClassOnce sync.Once
)

func getCXProviderConfigurationClass() _CXProviderConfigurationClass {
	CXProviderConfigurationClassOnce.Do(func() {
		CXProviderConfigurationClass = _CXProviderConfigurationClass{objc.GetClass("CXProviderConfiguration")}
	})
	return CXProviderConfigurationClass
}

type _CXProviderConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [CXProviderConfiguration] class.
type ICXProviderConfiguration interface {
	objectivec.IObject
	IconTemplateImageData() foundation.NSData
	SetIconTemplateImageData(value foundation.NSData)
	IncludesCallsInRecents() bool
	SetIncludesCallsInRecents(value bool)
	LocalizedName() string
	MaximumCallGroups() uint
	SetMaximumCallGroups(value uint)
	MaximumCallsPerCallGroup() uint
	SetMaximumCallsPerCallGroup(value uint)
	RingtoneSound() string
	SetRingtoneSound(value string)
	SupportedHandleTypes() unsafe.Pointer
	SetSupportedHandleTypes(value unsafe.Pointer)
	SupportsAudioTranslation() bool
	SetSupportsAudioTranslation(value bool)
	SupportsVideo() bool
	SetSupportsVideo(value bool)
}

// An encapsulation of the configuration of a provider object.
//
// A object controls the native call UI for incoming and outgoing calls, including a localized name for the provider, the ringtone to play for incoming calls, and the icon to display during calls. A provider configuration can also set the maximum number of call groups and the number of calls in a single call group, determine whether to use emails and phone numbers as handles, and specify whether to support video.


// An encapsulation of the configuration of a provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration
type CXProviderConfiguration struct {
	objectivec.Object
}

// CXProviderConfigurationFrom constructs a [CXProviderConfiguration] from an unsafe.Pointer.
//
// An encapsulation of the configuration of a provider object.
func CXProviderConfigurationFrom(ptr unsafe.Pointer) CXProviderConfiguration {
	return CXProviderConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXProviderConfigurationClass) Alloc() CXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXProviderConfigurationClass) New() CXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXProviderConfiguration) Init() CXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXProviderConfiguration) Autorelease() CXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXProviderConfiguration creates a new CXProviderConfiguration instance.
func NewCXProviderConfiguration() CXProviderConfiguration {
	return getCXProviderConfigurationClass().New()
}



// Initializes a configuration with the specified localized name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/init(localizedName:)
func NewCXProviderConfigurationWithLocalizedName(localizedName string) CXProviderConfiguration {
	instance := getCXProviderConfigurationClass().Alloc()
	rv := objc.Send[CXProviderConfiguration](instance.ID, objc.Sel("initWithLocalizedName:"), objc.String(localizedName))
	rv.Autorelease()
	return rv
}



// The PNG data for the icon image to be displayed for the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/iconTemplateImageData
func (c_ CXProviderConfiguration) IconTemplateImageData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("iconTemplateImageData"))
	return rv
}


// The PNG data for the icon image to be displayed for the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/iconTemplateImageData
func (c_ CXProviderConfiguration) SetIconTemplateImageData(value foundation.NSData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIconTemplateImageData:"), value)
}


// A Boolean value that indicates whether the provider includes a call in the system’s Recents list after the call ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/includesCallsInRecents
func (c_ CXProviderConfiguration) IncludesCallsInRecents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includesCallsInRecents"))
	return rv
}


// A Boolean value that indicates whether the provider includes a call in the system’s Recents list after the call ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/includesCallsInRecents
func (c_ CXProviderConfiguration) SetIncludesCallsInRecents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludesCallsInRecents:"), value)
}


// The localized name of the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/localizedName
func (c_ CXProviderConfiguration) LocalizedName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedName"))
	return rv
}


// The maximum number of call groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallGroups
func (c_ CXProviderConfiguration) MaximumCallGroups() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumCallGroups"))
	return rv
}


// The maximum number of call groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallGroups
func (c_ CXProviderConfiguration) SetMaximumCallGroups(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumCallGroups:"), value)
}


// The maximum number of calls per call group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallsPerCallGroup
func (c_ CXProviderConfiguration) MaximumCallsPerCallGroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumCallsPerCallGroup"))
	return rv
}


// The maximum number of calls per call group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallsPerCallGroup
func (c_ CXProviderConfiguration) SetMaximumCallsPerCallGroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumCallsPerCallGroup:"), value)
}


// The name of the sound resource in the app bundle to be used for the provider ringtone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/ringtoneSound
func (c_ CXProviderConfiguration) RingtoneSound() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ringtoneSound"))
	return rv
}


// The name of the sound resource in the app bundle to be used for the provider ringtone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/ringtoneSound
func (c_ CXProviderConfiguration) SetRingtoneSound(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRingtoneSound:"), objc.String(value))
}


// The supported handle types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportedHandleTypes-995uh
func (c_ CXProviderConfiguration) SupportedHandleTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedHandleTypes"))
	return rv
}


// The supported handle types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportedHandleTypes-995uh
func (c_ CXProviderConfiguration) SetSupportedHandleTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedHandleTypes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsAudioTranslation
func (c_ CXProviderConfiguration) SupportsAudioTranslation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsAudioTranslation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsAudioTranslation
func (c_ CXProviderConfiguration) SetSupportsAudioTranslation(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsAudioTranslation:"), value)
}


// A Boolean value that indicates whether the provider supports video in addition to audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsVideo
func (c_ CXProviderConfiguration) SupportsVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideo"))
	return rv
}


// A Boolean value that indicates whether the provider supports video in addition to audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsVideo
func (c_ CXProviderConfiguration) SetSupportsVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsVideo:"), value)
}


