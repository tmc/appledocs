// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TermOfAddress] class.
var (
	TermOfAddressClass     _TermOfAddressClass
	TermOfAddressClassOnce sync.Once
)

func getTermOfAddressClass() _TermOfAddressClass {
	TermOfAddressClassOnce.Do(func() {
		TermOfAddressClass = _TermOfAddressClass{objc.GetClass("NSTermOfAddress")}
	})
	return TermOfAddressClass
}

type _TermOfAddressClass struct {
	class objc.Class
}

// An interface definition for the [TermOfAddress] class.
type ITermOfAddress interface {
	objectivec.IObject
	// properties:
	LanguageIdentifier() IString
	Pronouns() []IMorphologyPronoun
	// methods:
}

// The type for representing grammatical gender in localized text.
//
// Many languages rely on gender for their grammar. Without knowing the subject’s gender or pronoun preferences, some localized strings may have grammatical errors, resulting in a poor user experience. is a type that enables the system to make pronoun substitutions in localized text based on gender. You don’t create instances of this type directly. Instead, use the predefined types to specify the gender to use when referring to people in translated text. Or define your own pronoun terms for a specific language when the predefined types are insufficient. For example, to substitute the masculine pronoun , for the neutral pronoun , do the following: If the , , and terms of address are insufficient, create your own term of address specifying the pronouns and language. For examples of how to use terms of address, see:


// The type for representing grammatical gender in localized text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress
type TermOfAddress struct {
	objectivec.Object
}

// TermOfAddressFrom constructs a [TermOfAddress] from an unsafe.Pointer.
//
// The type for representing grammatical gender in localized text.
func TermOfAddressFrom(ptr unsafe.Pointer) TermOfAddress {
	return TermOfAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TermOfAddressClass) Alloc() TermOfAddress {
	rv := objc.Send[TermOfAddress](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TermOfAddressClass) New() TermOfAddress {
	rv := objc.Send[TermOfAddress](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TermOfAddress) Init() TermOfAddress {
	rv := objc.Send[TermOfAddress](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TermOfAddress) Autorelease() TermOfAddress {
	rv := objc.Send[TermOfAddress](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTermOfAddress creates a new TermOfAddress instance.
func NewTermOfAddress() TermOfAddress {
	return getTermOfAddressClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/currentUser
func (tc _TermOfAddressClass) CurrentUser() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("currentUser"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/feminine
func (tc _TermOfAddressClass) Feminine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("feminine"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/localizedForLanguageIdentifier:withPronouns:
func (tc _TermOfAddressClass) LocalizedForLanguageIdentifierWithPronouns(language IString, pronouns []IMorphologyPronoun) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("localizedForLanguageIdentifier:withPronouns:"), language, pronouns)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/masculine
func (tc _TermOfAddressClass) Masculine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("masculine"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/neutral
func (tc _TermOfAddressClass) Neutral() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("neutral"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/languageIdentifier
func (t_ TermOfAddress) LanguageIdentifier() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("languageIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/pronouns
func (t_ TermOfAddress) Pronouns() []IMorphologyPronoun {
	rv := objc.Send[[]MorphologyPronoun](t_.ID, objc.Sel("pronouns"))
	return rv
}



