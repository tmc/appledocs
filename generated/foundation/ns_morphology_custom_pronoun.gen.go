// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MorphologyCustomPronoun] class.
var (
	MorphologyCustomPronounClass     _MorphologyCustomPronounClass
	MorphologyCustomPronounClassOnce sync.Once
)

func getMorphologyCustomPronounClass() _MorphologyCustomPronounClass {
	MorphologyCustomPronounClassOnce.Do(func() {
		MorphologyCustomPronounClass = _MorphologyCustomPronounClass{objc.GetClass("NSMorphologyCustomPronoun")}
	})
	return MorphologyCustomPronounClass
}

type _MorphologyCustomPronounClass struct {
	class objc.Class
}





// An interface definition for the [MorphologyCustomPronoun] class.
type IMorphologyCustomPronoun interface {
	objectivec.IObject
	

	// properties:
	ObjectForm() IString
	SetObjectForm(value IString)
	PossessiveAdjectiveForm() IString
	SetPossessiveAdjectiveForm(value IString)
	PossessiveForm() IString
	SetPossessiveForm(value IString)
	ReflexiveForm() IString
	SetReflexiveForm(value IString)
	SubjectForm() IString
	SetSubjectForm(value IString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MorphologyCustomPronounClass) Alloc() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MorphologyCustomPronounClass) New() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MorphologyCustomPronoun) Init() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MorphologyCustomPronoun) Autorelease() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphologyCustomPronoun creates a new MorphologyCustomPronoun instance.
func NewMorphologyCustomPronoun() MorphologyCustomPronoun {
	return getMorphologyCustomPronounClass().New()
}





// A custom pronoun behavior for use in a specific langauge.
//
// Set a instance on a instance when you want to provide a langauge-specific customization of pronoun use in that language. Different languages have different requirements for the grammatical information needed to apply a custom pronoun, so you set custom pronoun behavior on a per-language basis. The example below shows how to create English “ze” and “hir” custom pronouns: only supports third-person pronouns. Use this feature when your app needs to refer to third parties with a specific pronoun.


// A custom pronoun behavior for use in a specific langauge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun
type MorphologyCustomPronoun struct {
	objectivec.Object
}

// MorphologyCustomPronounFrom constructs a [MorphologyCustomPronoun] from an unsafe.Pointer.
//
// A custom pronoun behavior for use in a specific langauge.
func MorphologyCustomPronounFrom(ptr unsafe.Pointer) MorphologyCustomPronoun {
	return MorphologyCustomPronoun{objectivec.Object{objc.ID(ptr)}}
}










// Returns a Boolean value that indicates whether the given language supports setting custom pronouns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/isSupportedForLanguage:
func (mc _MorphologyCustomPronounClass) IsSupportedForLanguage(language IString) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isSupportedForLanguage:"), language)
	return rv
}


// Returns a collection of the custom pronoun keys required by this language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/requiredKeysForLanguage:
func (mc _MorphologyCustomPronounClass) RequiredKeysForLanguage(language IString) []string {
	rv := objc.Send[[]string](objc.ID(mc.class), objc.Sel("requiredKeysForLanguage:"), language)
	return rv
}

















// The object pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/objectForm
func (m_ MorphologyCustomPronoun) ObjectForm() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("objectForm"))
	return rv
}


// The object pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/objectForm
func (m_ MorphologyCustomPronoun) SetObjectForm(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectForm:"), value)
}


// The posessive adjective pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveAdjectiveForm
func (m_ MorphologyCustomPronoun) PossessiveAdjectiveForm() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("possessiveAdjectiveForm"))
	return rv
}


// The posessive adjective pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveAdjectiveForm
func (m_ MorphologyCustomPronoun) SetPossessiveAdjectiveForm(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPossessiveAdjectiveForm:"), value)
}


// The posessive pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveForm
func (m_ MorphologyCustomPronoun) PossessiveForm() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("possessiveForm"))
	return rv
}


// The posessive pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveForm
func (m_ MorphologyCustomPronoun) SetPossessiveForm(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPossessiveForm:"), value)
}


// The reflexive pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/reflexiveForm
func (m_ MorphologyCustomPronoun) ReflexiveForm() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("reflexiveForm"))
	return rv
}


// The reflexive pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/reflexiveForm
func (m_ MorphologyCustomPronoun) SetReflexiveForm(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReflexiveForm:"), value)
}


// The subject pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/subjectForm
func (m_ MorphologyCustomPronoun) SubjectForm() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("subjectForm"))
	return rv
}


// The subject pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/subjectForm
func (m_ MorphologyCustomPronoun) SetSubjectForm(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubjectForm:"), value)
}








