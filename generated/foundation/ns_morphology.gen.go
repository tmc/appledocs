// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMorphology */


/* debug [class_header]: Header for NSMorphology */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Morphology */
// An interface definition for the [Morphology] class.
type IMorphology interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Morphology */
	// properties:
	Definiteness() GrammaticalDefiniteness
	SetDefiniteness(value GrammaticalDefiniteness)
	Determination() GrammaticalDetermination
	SetDetermination(value GrammaticalDetermination)
	GrammaticalCase() GrammaticalCase
	SetGrammaticalCase(value GrammaticalCase)
	GrammaticalGender() GrammaticalGender
	SetGrammaticalGender(value GrammaticalGender)
	GrammaticalPerson() GrammaticalPerson
	SetGrammaticalPerson(value GrammaticalPerson)
	Number() GrammaticalNumber
	SetNumber(value GrammaticalNumber)
	PartOfSpeech() GrammaticalPartOfSpeech
	SetPartOfSpeech(value GrammaticalPartOfSpeech)
	PronounType() GrammaticalPronounType
	SetPronounType(value GrammaticalPronounType)
	Unspecified() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Morphology */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Morphology */
// Alloc allocates a new instance without initialization.
func (mc _MorphologyClass) Alloc() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Morphology */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Morphology *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Morphology */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Morphology */

// The addressing preferences of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/userMorphology
func (mc _MorphologyClass) UserMorphology() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("userMorphology"))
	return rv
}/* debug [class_properties_class/property]: userMorphology */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Morphology */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Morphology */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/definiteness
func (m_ Morphology) Definiteness() GrammaticalDefiniteness {
	rv := objc.Send[GrammaticalDefiniteness](m_.ID, objc.Sel("definiteness"))
	return rv
}/* debug [instance_properties/getter]: definiteness */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/definiteness
func (m_ Morphology) SetDefiniteness(value GrammaticalDefiniteness) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefiniteness:"), value)
}/* debug [instance_properties/setter]: definiteness */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/determination
func (m_ Morphology) Determination() GrammaticalDetermination {
	rv := objc.Send[GrammaticalDetermination](m_.ID, objc.Sel("determination"))
	return rv
}/* debug [instance_properties/getter]: determination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/determination
func (m_ Morphology) SetDetermination(value GrammaticalDetermination) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDetermination:"), value)
}/* debug [instance_properties/setter]: determination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalCase
func (m_ Morphology) GrammaticalCase() GrammaticalCase {
	rv := objc.Send[GrammaticalCase](m_.ID, objc.Sel("grammaticalCase"))
	return rv
}/* debug [instance_properties/getter]: grammaticalCase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalCase
func (m_ Morphology) SetGrammaticalCase(value GrammaticalCase) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalCase:"), value)
}/* debug [instance_properties/setter]: grammaticalCase */


// The grammatical gender used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalGender
func (m_ Morphology) GrammaticalGender() GrammaticalGender {
	rv := objc.Send[GrammaticalGender](m_.ID, objc.Sel("grammaticalGender"))
	return rv
}/* debug [instance_properties/getter]: grammaticalGender */


// The grammatical gender used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalGender
func (m_ Morphology) SetGrammaticalGender(value GrammaticalGender) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalGender:"), value)
}/* debug [instance_properties/setter]: grammaticalGender */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalPerson
func (m_ Morphology) GrammaticalPerson() GrammaticalPerson {
	rv := objc.Send[GrammaticalPerson](m_.ID, objc.Sel("grammaticalPerson"))
	return rv
}/* debug [instance_properties/getter]: grammaticalPerson */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalPerson
func (m_ Morphology) SetGrammaticalPerson(value GrammaticalPerson) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrammaticalPerson:"), value)
}/* debug [instance_properties/setter]: grammaticalPerson */


// The grammatical number used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/number
func (m_ Morphology) Number() GrammaticalNumber {
	rv := objc.Send[GrammaticalNumber](m_.ID, objc.Sel("number"))
	return rv
}/* debug [instance_properties/getter]: number */


// The grammatical number used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/number
func (m_ Morphology) SetNumber(value GrammaticalNumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumber:"), value)
}/* debug [instance_properties/setter]: number */


// The grammatical part of speech used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/partOfSpeech
func (m_ Morphology) PartOfSpeech() GrammaticalPartOfSpeech {
	rv := objc.Send[GrammaticalPartOfSpeech](m_.ID, objc.Sel("partOfSpeech"))
	return rv
}/* debug [instance_properties/getter]: partOfSpeech */


// The grammatical part of speech used for inflecting strings with this morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/partOfSpeech
func (m_ Morphology) SetPartOfSpeech(value GrammaticalPartOfSpeech) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPartOfSpeech:"), value)
}/* debug [instance_properties/setter]: partOfSpeech */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/pronounType
func (m_ Morphology) PronounType() GrammaticalPronounType {
	rv := objc.Send[GrammaticalPronounType](m_.ID, objc.Sel("pronounType"))
	return rv
}/* debug [instance_properties/getter]: pronounType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/pronounType
func (m_ Morphology) SetPronounType(value GrammaticalPronounType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPronounType:"), value)
}/* debug [instance_properties/setter]: pronounType */


// A Boolean value that indicates whether this instance specifies no particular grammar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/unspecified
func (m_ Morphology) Unspecified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("unspecified"))
	return rv
}/* debug [instance_properties/getter]: unspecified */


// The addressing preferences of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology/userMorphology
func (m_ Morphology) UserMorphology() IMorphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("userMorphology"))
	return rv
}/* debug [instance_properties/getter]: userMorphology */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMorphology */



