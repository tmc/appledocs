// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INIntent] class.
var (
	INIntentClass     _INIntentClass
	INIntentClassOnce sync.Once
)

func getINIntentClass() _INIntentClass {
	INIntentClassOnce.Do(func() {
		INIntentClass = _INIntentClass{objc.GetClass("INIntent")}
	})
	return INIntentClass
}

type _INIntentClass struct {
	class objc.Class
}

// An interface definition for the [INIntent] class.
type IINIntent interface {
	objectivec.IObject
	SetImageForParameterNamed(image INImage, parameterName string)
	ImageForParameterNamed(parameterName string) INImage
	KeyImage() INImage
	DonationMetadata() INIntentDonationMetadata
	SetDonationMetadata(value INIntentDonationMetadata)
	Identifier() string
	IntentDescription() string
	ShortcutAvailability() unsafe.Pointer
	SetShortcutAvailability(value unsafe.Pointer)
	SuggestedInvocationPhrase() string
	SetSuggestedInvocationPhrase(value string)
}

// A request to fulfill in your app or Intents extension.
//
// The class is abstract and provides behaviors shared by all intent objects. You don’t create instances of this class directly or implement your own custom subclasses. For a list of intent types that SiriKit already handles, see the Standard Intents section of . You may also define custom intent types in an Intent Definition file. Each subclass of defines the properties needed to perform the corresponding action. You use instances of those classes when responding to a request sent to your app or Intents extension by SiriKit. For more information about a specific type of action, see the appropriate subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent
type INIntent struct {
	objectivec.Object
}

// INIntentFrom constructs a [INIntent] from an unsafe.Pointer.
//
// A request to fulfill in your app or Intents extension.
func INIntentFrom(ptr unsafe.Pointer) INIntent {
	return INIntent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INIntentClass) Alloc() INIntent {
	rv := objc.Send[INIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INIntentClass) New() INIntent {
	rv := objc.Send[INIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INIntent) Init() INIntent {
	rv := objc.Send[INIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INIntent) Autorelease() INIntent {
	rv := objc.Send[INIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINIntent creates a new INIntent instance.
func NewINIntent() INIntent {
	return getINIntentClass().New()
}


// Sets the image to use for the specified parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inintent/2976224-setimage
func (i_ INIntent) SetImageForParameterNamed(image INImage, parameterName string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:forParameterNamed:"), image, objc.String(parameterName))
}

// Returns the image associated with the specified parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/imageForParameterNamed:
func (i_ INIntent) ImageForParameterNamed(parameterName string) INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("imageForParameterNamed:"), objc.String(parameterName))
	return rv
}

// The most relevant image to display to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/keyImage()
func (i_ INIntent) KeyImage() INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("keyImage"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/donationMetadata
func (i_ INIntent) DonationMetadata() INIntentDonationMetadata {
	rv := objc.Send[INIntentDonationMetadata](i_.ID, objc.Sel("donationMetadata"))
	return rv
}


// SetDonationMetadata sets the value of the donationMetadata property.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/donationMetadata
func (i_ INIntent) SetDonationMetadata(value INIntentDonationMetadata) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDonationMetadata:"), value)
}

// The unique identifier for this intent object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/identifier
func (i_ INIntent) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}

// A string describing the content of the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/intentDescription
func (i_ INIntent) IntentDescription() string {
	rv := objc.Send[string](i_.ID, objc.Sel("intentDescription"))
	return rv
}

// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/shortcutAvailability
func (i_ INIntent) ShortcutAvailability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("shortcutAvailability"))
	return rv
}


// SetShortcutAvailability sets the value of the shortcutAvailability property.
// A set of defined contexts in which an intent or activity might be relevant to a user.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/shortcutAvailability
func (i_ INIntent) SetShortcutAvailability(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setShortcutAvailability:"), value)
}

// The intent’s display name.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/suggestedInvocationPhrase
func (i_ INIntent) SuggestedInvocationPhrase() string {
	rv := objc.Send[string](i_.ID, objc.Sel("suggestedInvocationPhrase"))
	return rv
}


// SetSuggestedInvocationPhrase sets the value of the suggestedInvocationPhrase property.
// The intent’s display name.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntent/suggestedInvocationPhrase
func (i_ INIntent) SetSuggestedInvocationPhrase(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSuggestedInvocationPhrase:"), objc.String(value))
}



