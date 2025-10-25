// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPersonNameComponents */


/* debug [class_header]: Header for NSPersonNameComponents */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PersonNameComponents */
// An interface definition for the [PersonNameComponents] class.
type IPersonNameComponents interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PersonNameComponents */
	// properties:
	FamilyName() IString
	SetFamilyName(value IString)
	GivenName() IString
	SetGivenName(value IString)
	MiddleName() IString
	SetMiddleName(value IString)
	NamePrefix() IString
	SetNamePrefix(value IString)
	NameSuffix() IString
	SetNameSuffix(value IString)
	Nickname() IString
	SetNickname(value IString)
	PhoneticRepresentation() IPersonNameComponents
	SetPhoneticRepresentation(value IPersonNameComponents)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PersonNameComponents */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PersonNameComponents */
// Alloc allocates a new instance without initialization.
func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PersonNameComponents */
// An object that manages the separate parts of a person’s name to allow locale-aware formatting.
//
// You can use this type in Swift when you need reference semantics or other Foundation-specific behavior. An object encapsulates the components of a person’s name in an extendable, object-oriented manner. It is used to specify a person’s name by providing the components comprising a full name: given name, middle name, family name, prefix, suffix, nickname, and phonetic representation. objects can be used by an instance of to create string representations suitable for display in the current locale.


// An object that manages the separate parts of a person’s name to allow locale-aware formatting.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PersonNameComponents *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PersonNameComponents */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PersonNameComponents */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PersonNameComponents */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PersonNameComponents */

// Name bestowed upon an individual to denote membership in a group or family. .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/familyName
func (p_ PersonNameComponents) FamilyName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("familyName"))
	return rv
}/* debug [instance_properties/getter]: familyName */


// Name bestowed upon an individual to denote membership in a group or family. .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/familyName
func (p_ PersonNameComponents) SetFamilyName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFamilyName:"), value)
}/* debug [instance_properties/setter]: familyName */


// Name bestowed upon an individual to differentiate them from other members of a group that share a family name .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/givenName
func (p_ PersonNameComponents) GivenName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("givenName"))
	return rv
}/* debug [instance_properties/getter]: givenName */


// Name bestowed upon an individual to differentiate them from other members of a group that share a family name .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/givenName
func (p_ PersonNameComponents) SetGivenName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGivenName:"), value)
}/* debug [instance_properties/setter]: givenName */


// Secondary name bestowed upon an individual to differentiate them from others that have the same given name .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/middleName
func (p_ PersonNameComponents) MiddleName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("middleName"))
	return rv
}/* debug [instance_properties/getter]: middleName */


// Secondary name bestowed upon an individual to differentiate them from others that have the same given name .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/middleName
func (p_ PersonNameComponents) SetMiddleName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMiddleName:"), value)
}/* debug [instance_properties/setter]: middleName */


// The portion of a name’s full form of address that precedes the name itself .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/namePrefix
func (p_ PersonNameComponents) NamePrefix() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("namePrefix"))
	return rv
}/* debug [instance_properties/getter]: namePrefix */


// The portion of a name’s full form of address that precedes the name itself .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/namePrefix
func (p_ PersonNameComponents) SetNamePrefix(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNamePrefix:"), value)
}/* debug [instance_properties/setter]: namePrefix */


// The portion of a name’s full form of address that follows the name itself .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nameSuffix
func (p_ PersonNameComponents) NameSuffix() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("nameSuffix"))
	return rv
}/* debug [instance_properties/getter]: nameSuffix */


// The portion of a name’s full form of address that follows the name itself .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nameSuffix
func (p_ PersonNameComponents) SetNameSuffix(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNameSuffix:"), value)
}/* debug [instance_properties/setter]: nameSuffix */


// Name substituted for the purposes of familiarity .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nickname
func (p_ PersonNameComponents) Nickname() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("nickname"))
	return rv
}/* debug [instance_properties/getter]: nickname */


// Name substituted for the purposes of familiarity .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/nickname
func (p_ PersonNameComponents) SetNickname(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNickname:"), value)
}/* debug [instance_properties/setter]: nickname */


// The phonetic representation name components of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/phoneticRepresentation
func (p_ PersonNameComponents) PhoneticRepresentation() IPersonNameComponents {
	rv := objc.Send[PersonNameComponents](p_.ID, objc.Sel("phoneticRepresentation"))
	return rv
}/* debug [instance_properties/getter]: phoneticRepresentation */


// The phonetic representation name components of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents/phoneticRepresentation
func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPhoneticRepresentation:"), value)
}/* debug [instance_properties/setter]: phoneticRepresentation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPersonNameComponents */



