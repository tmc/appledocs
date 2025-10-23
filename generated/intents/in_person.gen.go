// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INPerson] class.
var (
	INPersonClass     _INPersonClass
	INPersonClassOnce sync.Once
)

func getINPersonClass() _INPersonClass {
	INPersonClassOnce.Do(func() {
		INPersonClass = _INPersonClass{objc.GetClass("INPerson")}
	})
	return INPersonClass
}

type _INPersonClass struct {
	class objc.Class
}

// An interface definition for the [INPerson] class.
type IINPerson interface {
	objectivec.IObject
	// properties:
	Aliases() INPersonHandle
	SetAliases(value INPersonHandle)
	ContactIdentifier() string
	SetContactIdentifier(value string)
	CustomIdentifier() string
	SetCustomIdentifier(value string)
	DisplayName() string
	SetDisplayName(value string)
	Handle() string
	SetHandle(value string)
	Image() INImage
	SetImage(value INImage)
	IsContactSuggestion() bool
	SetIsContactSuggestion(value bool)
	IsMe() bool
	SetIsMe(value bool)
	NameComponents() foundation.PersonNameComponents
	SetNameComponents(value foundation.PersonNameComponents)
	PersonHandle() INPersonHandle
	SetPersonHandle(value INPersonHandle)
	Relationship() unsafe.Pointer
	SetRelationship(value unsafe.Pointer)
	SiriMatches() INPerson
	SetSiriMatches(value INPerson)
	SuggestionType() INPersonSuggestionType
	SetSuggestionType(value INPersonSuggestionType)
	SpokenPhrase() string
	SetSpokenPhrase(value string)
	// methods:
}

// Information about a person participating in a SiriKit interaction.
//
// SiriKit uses objects to represent people with many different roles, including the sender or recipient of calls and messages, the payer or payee of a financial transaction, or the driver of a vehicle. You also use person objects to identify the corresponding contact in your app and to communicate information about that contact back to SiriKit. When resolving the parameters of an intent, use any provided objects to identify the corresponding contacts in your app. A person object contains information provided by the initial request, which could be as little as a single name spoken by the person interacting with Siri. After identifying the contact, create a new object and fill it with the information that you need to identify that contact again later. For example, you might specify a value for property that contains the information about how your app identifies that contact. When resolving the identities of contacts, SiriKit leverages the information in the device owner’s contacts database when that information is available. If the owner denies your app access to their contacts, SiriKit can’t use that information, which might cause many properties of a person object to be . Because the class conforms to the protocol, though, SiriKit still populates the property with what the person interacting with Siri said, and you can use that information to try to identify the contact. For more information about that protocol, see .


// Information about a person participating in a SiriKit interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPerson
type INPerson struct {
	objectivec.Object
}

// INPersonFrom constructs a [INPerson] from an unsafe.Pointer.
//
// Information about a person participating in a SiriKit interaction.
func INPersonFrom(ptr unsafe.Pointer) INPerson {
	return INPerson{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INPersonClass) Alloc() INPerson {
	rv := objc.Send[INPerson](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPersonClass) New() INPerson {
	rv := objc.Send[INPerson](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPerson) Init() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPerson) Autorelease() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPerson creates a new INPerson instance.
func NewINPerson() INPerson {
	return getINPersonClass().New()
}



// The additional handles that Siri may use to identify the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/aliases
func (i_ INPerson) Aliases() INPersonHandle {
	rv := objc.Send[INPersonHandle](i_.ID, objc.Sel("aliases"))
	return rv
}


// The additional handles that Siri may use to identify the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/aliases
func (i_ INPerson) SetAliases(value INPersonHandle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAliases:"), value)
}


// The Contacts database identifier for the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/contactidentifier
func (i_ INPerson) ContactIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("contactIdentifier"))
	return rv
}


// The Contacts database identifier for the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/contactidentifier
func (i_ INPerson) SetContactIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContactIdentifier:"), objc.String(value))
}


// The unique identifier that your app uses to identify the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/customidentifier
func (i_ INPerson) CustomIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("customIdentifier"))
	return rv
}


// The unique identifier that your app uses to identify the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/customidentifier
func (i_ INPerson) SetCustomIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomIdentifier:"), objc.String(value))
}


// The person’s formatted name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/displayname
func (i_ INPerson) DisplayName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("displayName"))
	return rv
}


// The person’s formatted name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/displayname
func (i_ INPerson) SetDisplayName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The unique identifier that your app assigned to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/handle
func (i_ INPerson) Handle() string {
	rv := objc.Send[string](i_.ID, objc.Sel("handle"))
	return rv
}


// The unique identifier that your app assigned to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/handle
func (i_ INPerson) SetHandle(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHandle:"), objc.String(value))
}


// An image of the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/image
func (i_ INPerson) Image() INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("image"))
	return rv
}


// An image of the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/image
func (i_ INPerson) SetImage(value INImage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:"), value)
}


// A Boolean value that indicates whether the person is a contact suggestion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/iscontactsuggestion
func (i_ INPerson) IsContactSuggestion() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isContactSuggestion"))
	return rv
}


// A Boolean value that indicates whether the person is a contact suggestion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/iscontactsuggestion
func (i_ INPerson) SetIsContactSuggestion(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsContactSuggestion:"), value)
}


// A Boolean value indicating whether the person is the user of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/isme
func (i_ INPerson) IsMe() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isMe"))
	return rv
}


// A Boolean value indicating whether the person is the user of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/isme
func (i_ INPerson) SetIsMe(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsMe:"), value)
}


// The individual components of the person’s full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/namecomponents
func (i_ INPerson) NameComponents() foundation.PersonNameComponents {
	rv := objc.Send[foundation.PersonNameComponents](i_.ID, objc.Sel("nameComponents"))
	return rv
}


// The individual components of the person’s full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/namecomponents
func (i_ INPerson) SetNameComponents(value foundation.PersonNameComponents) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNameComponents:"), value)
}


// The unique handle that your app assigns to the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/personhandle
func (i_ INPerson) PersonHandle() INPersonHandle {
	rv := objc.Send[INPersonHandle](i_.ID, objc.Sel("personHandle"))
	return rv
}


// The unique handle that your app assigns to the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/personhandle
func (i_ INPerson) SetPersonHandle(value INPersonHandle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPersonHandle:"), value)
}


// The relationship between this person and the person using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/relationship
func (i_ INPerson) Relationship() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relationship"))
	return rv
}


// The relationship between this person and the person using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/relationship
func (i_ INPerson) SetRelationship(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelationship:"), value)
}


// The list of matches Siri provides for you to resolve or disambiguate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/sirimatches
func (i_ INPerson) SiriMatches() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("siriMatches"))
	return rv
}


// The list of matches Siri provides for you to resolve or disambiguate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/sirimatches
func (i_ INPerson) SetSiriMatches(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSiriMatches:"), value)
}


// The type of contact information to donate with interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/suggestiontype
func (i_ INPerson) SuggestionType() INPersonSuggestionType {
	rv := objc.Send[INPersonSuggestionType](i_.ID, objc.Sel("suggestionType"))
	return rv
}


// The type of contact information to donate with interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inperson/suggestiontype
func (i_ INPerson) SetSuggestionType(value INPersonSuggestionType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSuggestionType:"), value)
}


// The phrase identified by Siri.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inspeakable/spokenphrase
func (i_ INPerson) SpokenPhrase() string {
	rv := objc.Send[string](i_.ID, objc.Sel("spokenPhrase"))
	return rv
}


// The phrase identified by Siri.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inspeakable/spokenphrase
func (i_ INPerson) SetSpokenPhrase(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSpokenPhrase:"), objc.String(value))
}



