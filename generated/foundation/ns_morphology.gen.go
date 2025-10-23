// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Morphology] class.
var (
	MorphologyClass     _MorphologyClass
	MorphologyClassOnce sync.Once
)

func getMorphologyClass() _MorphologyClass {
	MorphologyClassOnce.Do(func() {
		MorphologyClass = _MorphologyClass{objc.GetClass("NSMorphology")}
	})
	return MorphologyClass
}

type _MorphologyClass struct {
	class objc.Class
}

// An interface definition for the [Morphology] class.
type IMorphology interface {
	objectivec.IObject
	Definiteness() NSGrammaticalDefiniteness
	SetDefiniteness(value NSGrammaticalDefiniteness)
	Determination() NSGrammaticalDetermination
	SetDetermination(value NSGrammaticalDetermination)
	GrammaticalCase() NSGrammaticalCase
	SetGrammaticalCase(value NSGrammaticalCase)
	GrammaticalGender() NSGrammaticalGender
	SetGrammaticalGender(value NSGrammaticalGender)
	GrammaticalPerson() NSGrammaticalPerson
	SetGrammaticalPerson(value NSGrammaticalPerson)
	Number() NSGrammaticalNumber
	SetNumber(value NSGrammaticalNumber)
	PartOfSpeech() NSGrammaticalPartOfSpeech
	SetPartOfSpeech(value NSGrammaticalPartOfSpeech)
	PronounType() NSGrammaticalPronounType
	SetPronounType(value NSGrammaticalPronounType)
	Unspecified() bool
}

// A description of the grammatical properties of a string.
//
// Use a morphology with an to specify how to interpret a specific word when inflecting an . This affects grammatical agreement with traits like number and gender, as well as declaring the word’s part of speech. The type’s design is language-independent; the concepts it can specify encompass the spectrum of what languages can do. Even for languages that don’t have one or more of those properties benefit the system as hints to make appropriate choices even when an exact inflection isn’t possible. Examples of properties absent from languages include Spanish’s lack of a grammatical gender of neuter, or the nonexistence of a paucal (plural few) pronoun in English.


// A description of the grammatical properties of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology
type Morphology struct {
	objectivec.Object
}

// MorphologyFrom constructs a [Morphology] from an unsafe.Pointer.
//
// A description of the grammatical properties of a string.
func MorphologyFrom(ptr unsafe.Pointer) Morphology {
	return Morphology{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MorphologyClass) Alloc() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MorphologyClass) New() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Morphology) Init() Morphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Morphology) Autorelease() Morphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphology creates a new Morphology instance.
func NewMorphology() Morphology {
	return getMorphologyClass().New()
}



// The addressing preferences of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/userMorphology
func (mc _MorphologyClass) UserMorphology() Morphology {
	rv := objc.Send[NSMorphology](objc.ID(mc.class), objc.Sel("userMorphology"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/definiteness
func (m_ Morphology) Definiteness() NSGrammaticalDefiniteness {
	rv := objc.Send[GrammaticalDefiniteness](m_.ID, objc.Sel("definiteness"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/definiteness
func (m_ Morphology) SetDefiniteness(value NSGrammaticalDefiniteness) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefiniteness:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/determination
func (m_ Morphology) Determination() NSGrammaticalDetermination {
	rv := objc.Send[GrammaticalDetermination](m_.ID, objc.Sel("determination"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/determination
func (m_ Morphology) SetDetermination(value NSGrammaticalDetermination) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDetermination:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalCase
func (m_ Morphology) GrammaticalCase() NSGrammaticalCase {
	rv := objc.Send[GrammaticalCase](m_.ID, objc.Sel("grammaticalCase"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalCase
func (m_ Morphology) SetGrammaticalCase(value NSGrammaticalCase) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalCase:"), value)
}


// The grammatical gender used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalGender
func (m_ Morphology) GrammaticalGender() NSGrammaticalGender {
	rv := objc.Send[GrammaticalGender](m_.ID, objc.Sel("grammaticalGender"))
	return rv
}


// The grammatical gender used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalGender
func (m_ Morphology) SetGrammaticalGender(value NSGrammaticalGender) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalGender:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalPerson
func (m_ Morphology) GrammaticalPerson() NSGrammaticalPerson {
	rv := objc.Send[GrammaticalPerson](m_.ID, objc.Sel("grammaticalPerson"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalPerson
func (m_ Morphology) SetGrammaticalPerson(value NSGrammaticalPerson) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalPerson:"), value)
}


// The grammatical number used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/number
func (m_ Morphology) Number() NSGrammaticalNumber {
	rv := objc.Send[GrammaticalNumber](m_.ID, objc.Sel("number"))
	return rv
}


// The grammatical number used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/number
func (m_ Morphology) SetNumber(value NSGrammaticalNumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumber:"), value)
}


// The grammatical part of speech used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/partOfSpeech
func (m_ Morphology) PartOfSpeech() NSGrammaticalPartOfSpeech {
	rv := objc.Send[GrammaticalPartOfSpeech](m_.ID, objc.Sel("partOfSpeech"))
	return rv
}


// The grammatical part of speech used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/partOfSpeech
func (m_ Morphology) SetPartOfSpeech(value NSGrammaticalPartOfSpeech) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPartOfSpeech:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/pronounType
func (m_ Morphology) PronounType() NSGrammaticalPronounType {
	rv := objc.Send[GrammaticalPronounType](m_.ID, objc.Sel("pronounType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/pronounType
func (m_ Morphology) SetPronounType(value NSGrammaticalPronounType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPronounType:"), value)
}


// A Boolean value that indicates whether this instance specifies no particular grammar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/unspecified
func (m_ Morphology) Unspecified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("unspecified"))
	return rv
}


// The addressing preferences of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/userMorphology
func (m_ Morphology) UserMorphology() IMorphology {
	rv := objc.Send[NSMorphology](m_.ID, objc.Sel("userMorphology"))
	return rv
}



