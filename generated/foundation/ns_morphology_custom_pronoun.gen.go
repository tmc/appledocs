// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// A custom pronoun behavior for use in a specific langauge.
//
// Set a instance on a instance when you want to provide a langauge-specific customization of pronoun use in that language. Different languages have different requirements for the grammatical information needed to apply a custom pronoun, so you set custom pronoun behavior on a per-language basis. The example below shows how to create English “ze” and “hir” custom pronouns: only supports third-person pronouns. Use this feature when your app needs to refer to third parties with a specific pronoun.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MorphologyCustomPronounClass) Alloc() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The reflexive pronoun form to apply when using this custom pronoun behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/reflexiveForm
func (m_ MorphologyCustomPronoun) ReflexiveForm() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("reflexiveForm"))
	return rv
}


// SetReflexiveForm sets the value of the reflexiveForm property.
// The reflexive pronoun form to apply when using this custom pronoun behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/reflexiveForm
func (m_ MorphologyCustomPronoun) SetReflexiveForm(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReflexiveForm:"), value)
}



