// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersonNameComponents] class.
var (
	PersonNameComponentsClass     _PersonNameComponentsClass
	PersonNameComponentsClassOnce sync.Once
)

func getPersonNameComponentsClass() _PersonNameComponentsClass {
	PersonNameComponentsClassOnce.Do(func() {
		PersonNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}
	})
	return PersonNameComponentsClass
}

type _PersonNameComponentsClass struct {
	class objc.Class
}

// An interface definition for the [PersonNameComponents] class.
type IPersonNameComponents interface {
	objectivec.IObject
}

// An object that manages the separate parts of a person’s name to allow locale-aware formatting.
//
// You can use this type in Swift when you need reference semantics or other Foundation-specific behavior. An object encapsulates the components of a person’s name in an extendable, object-oriented manner. It is used to specify a person’s name by providing the components comprising a full name: given name, middle name, family name, prefix, suffix, nickname, and phonetic representation. objects can be used by an instance of to create string representations suitable for display in the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents
type PersonNameComponents struct {
	objectivec.Object
}

// PersonNameComponentsFrom constructs a [PersonNameComponents] from an unsafe.Pointer.
//
// An object that manages the separate parts of a person’s name to allow locale-aware formatting.
func PersonNameComponentsFrom(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersonNameComponents) Autorelease() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersonNameComponents creates a new PersonNameComponents instance.
func NewPersonNameComponents() PersonNameComponents {
	return getPersonNameComponentsClass().New()
}


// Name bestowed upon an individual to denote membership in a group or family. .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/familyName
func (p_ PersonNameComponents) FamilyName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("familyName"))
	return rv
}


// SetFamilyName sets the value of the familyName property.
// Name bestowed upon an individual to denote membership in a group or family. .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/familyName
func (p_ PersonNameComponents) SetFamilyName(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFamilyName:"), value)
}

// Name bestowed upon an individual to differentiate them from other members of a group that share a family name .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/givenName
func (p_ PersonNameComponents) GivenName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("givenName"))
	return rv
}


// SetGivenName sets the value of the givenName property.
// Name bestowed upon an individual to differentiate them from other members of a group that share a family name .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/givenName
func (p_ PersonNameComponents) SetGivenName(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGivenName:"), value)
}

// Secondary name bestowed upon an individual to differentiate them from others that have the same given name .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/middleName
func (p_ PersonNameComponents) MiddleName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("middleName"))
	return rv
}


// SetMiddleName sets the value of the middleName property.
// Secondary name bestowed upon an individual to differentiate them from others that have the same given name .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/middleName
func (p_ PersonNameComponents) SetMiddleName(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMiddleName:"), value)
}

// The portion of a name’s full form of address that precedes the name itself .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/namePrefix
func (p_ PersonNameComponents) NamePrefix() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("namePrefix"))
	return rv
}


// SetNamePrefix sets the value of the namePrefix property.
// The portion of a name’s full form of address that precedes the name itself .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/namePrefix
func (p_ PersonNameComponents) SetNamePrefix(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNamePrefix:"), value)
}

// The portion of a name’s full form of address that follows the name itself .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nameSuffix
func (p_ PersonNameComponents) NameSuffix() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("nameSuffix"))
	return rv
}


// SetNameSuffix sets the value of the nameSuffix property.
// The portion of a name’s full form of address that follows the name itself .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nameSuffix
func (p_ PersonNameComponents) SetNameSuffix(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNameSuffix:"), value)
}

// Name substituted for the purposes of familiarity .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nickname
func (p_ PersonNameComponents) Nickname() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("nickname"))
	return rv
}


// SetNickname sets the value of the nickname property.
// Name substituted for the purposes of familiarity .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nickname
func (p_ PersonNameComponents) SetNickname(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNickname:"), value)
}

// The phonetic representation name components of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/phoneticRepresentation
func (p_ PersonNameComponents) PhoneticRepresentation() NSPersonNameComponents {
	rv := objc.Send[NSPersonNameComponents](p_.ID, objc.Sel("phoneticRepresentation"))
	return rv
}


// SetPhoneticRepresentation sets the value of the phoneticRepresentation property.
// The phonetic representation name components of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/phoneticRepresentation
func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPhoneticRepresentation:"), value)
}



