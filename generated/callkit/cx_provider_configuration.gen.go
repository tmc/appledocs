// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXProviderConfiguration */


/* debug [class_header]: Header for CXProviderConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXProviderConfiguration */
// An interface definition for the [CXProviderConfiguration] class.
type ICXProviderConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXProviderConfiguration */
	// properties:
	IconTemplateImageData() objc.IObject /* cross-framework: NSData */
	SetIconTemplateImageData(value objc.IObject /* cross-framework: NSData */)
	IncludesCallsInRecents() bool
	SetIncludesCallsInRecents(value bool)
	LocalizedName() objc.IObject /* cross-framework: NSString */
	MaximumCallGroups() uint
	SetMaximumCallGroups(value uint)
	MaximumCallsPerCallGroup() uint
	SetMaximumCallsPerCallGroup(value uint)
	RingtoneSound() objc.IObject /* cross-framework: NSString */
	SetRingtoneSound(value objc.IObject /* cross-framework: NSString */)
	SupportedHandleTypes() unsafe.Pointer
	SetSupportedHandleTypes(value unsafe.Pointer)
	SupportsAudioTranslation() bool
	SetSupportsAudioTranslation(value bool)
	SupportsVideo() bool
	SetSupportsVideo(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXProviderConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXProviderConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CXProviderConfigurationClass) Alloc() CXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXProviderConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXProviderConfiguration */

// Initializes a configuration with the specified localized name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/init(localizedName:)
func NewCXProviderConfigurationWithLocalizedName(localizedName objc.IObject /* cross-framework: NSString */) CXProviderConfiguration {
	instance := getCXProviderConfigurationClass().Alloc()
	rv := objc.Send[CXProviderConfiguration](instance.ID, objc.Sel("initWithLocalizedName:"), localizedName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXProviderConfigurationWithLocalizedName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXProviderConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXProviderConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXProviderConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXProviderConfiguration */

// The PNG data for the icon image to be displayed for the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/iconTemplateImageData
func (c_ CXProviderConfiguration) IconTemplateImageData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("iconTemplateImageData"))
	return rv
}/* debug [instance_properties/getter]: iconTemplateImageData */


// The PNG data for the icon image to be displayed for the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/iconTemplateImageData
func (c_ CXProviderConfiguration) SetIconTemplateImageData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIconTemplateImageData:"), value)
}/* debug [instance_properties/setter]: iconTemplateImageData */


// A Boolean value that indicates whether the provider includes a call in the system’s Recents list after the call ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/includesCallsInRecents
func (c_ CXProviderConfiguration) IncludesCallsInRecents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includesCallsInRecents"))
	return rv
}/* debug [instance_properties/getter]: includesCallsInRecents */


// A Boolean value that indicates whether the provider includes a call in the system’s Recents list after the call ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/includesCallsInRecents
func (c_ CXProviderConfiguration) SetIncludesCallsInRecents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludesCallsInRecents:"), value)
}/* debug [instance_properties/setter]: includesCallsInRecents */


// The localized name of the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/localizedName
func (c_ CXProviderConfiguration) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */


// The maximum number of call groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallGroups
func (c_ CXProviderConfiguration) MaximumCallGroups() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumCallGroups"))
	return rv
}/* debug [instance_properties/getter]: maximumCallGroups */


// The maximum number of call groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallGroups
func (c_ CXProviderConfiguration) SetMaximumCallGroups(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumCallGroups:"), value)
}/* debug [instance_properties/setter]: maximumCallGroups */


// The maximum number of calls per call group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallsPerCallGroup
func (c_ CXProviderConfiguration) MaximumCallsPerCallGroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumCallsPerCallGroup"))
	return rv
}/* debug [instance_properties/getter]: maximumCallsPerCallGroup */


// The maximum number of calls per call group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/maximumCallsPerCallGroup
func (c_ CXProviderConfiguration) SetMaximumCallsPerCallGroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumCallsPerCallGroup:"), value)
}/* debug [instance_properties/setter]: maximumCallsPerCallGroup */


// The name of the sound resource in the app bundle to be used for the provider ringtone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/ringtoneSound
func (c_ CXProviderConfiguration) RingtoneSound() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ringtoneSound"))
	return rv
}/* debug [instance_properties/getter]: ringtoneSound */


// The name of the sound resource in the app bundle to be used for the provider ringtone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/ringtoneSound
func (c_ CXProviderConfiguration) SetRingtoneSound(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRingtoneSound:"), value)
}/* debug [instance_properties/setter]: ringtoneSound */


// The supported handle types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportedHandleTypes-995uh
func (c_ CXProviderConfiguration) SupportedHandleTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedHandleTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedHandleTypes */


// The supported handle types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportedHandleTypes-995uh
func (c_ CXProviderConfiguration) SetSupportedHandleTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedHandleTypes:"), value)
}/* debug [instance_properties/setter]: supportedHandleTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsAudioTranslation
func (c_ CXProviderConfiguration) SupportsAudioTranslation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsAudioTranslation"))
	return rv
}/* debug [instance_properties/getter]: supportsAudioTranslation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsAudioTranslation
func (c_ CXProviderConfiguration) SetSupportsAudioTranslation(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsAudioTranslation:"), value)
}/* debug [instance_properties/setter]: supportsAudioTranslation */


// A Boolean value that indicates whether the provider supports video in addition to audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsVideo
func (c_ CXProviderConfiguration) SupportsVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideo"))
	return rv
}/* debug [instance_properties/getter]: supportsVideo */


// A Boolean value that indicates whether the provider supports video in addition to audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProviderConfiguration/supportsVideo
func (c_ CXProviderConfiguration) SetSupportsVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsVideo:"), value)
}/* debug [instance_properties/setter]: supportsVideo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXProviderConfiguration */


