// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SpellChecker] class.
var (
	spellCheckerClass     _SpellCheckerClass
	spellCheckerClassOnce sync.Once
)

func getSpellCheckerClass() _SpellCheckerClass {
	spellCheckerClassOnce.Do(func() {
		spellCheckerClass = _SpellCheckerClass{objc.GetClass("NSSpellChecker")}
	})
	return spellCheckerClass
}

type _SpellCheckerClass struct {
	class objc.Class
}

// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	objectivec.IObject
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int
}

// An interface to the Cocoa spell-checking service.
//
// To handle all its spell checking, an app needs only one instance of , known as the spell checker. Using the spell checker you manage the Spelling panel, in which the user can specify decisions about words that are suspect. The spell checker also offers the ability to provide word completions to augment the text completion system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
type SpellChecker struct {
	objectivec.Object
}

// SpellCheckerFrom constructs a [SpellChecker] from an unsafe.Pointer.
//
// An interface to the Cocoa spell-checking service.
func SpellCheckerFrom(ptr unsafe.Pointer) SpellChecker {
	return SpellChecker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpellCheckerClass) Alloc() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpellCheckerClass) New() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpellChecker) Init() SpellChecker {
	rv := objc.Send[SpellChecker](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpellChecker) Autorelease() SpellChecker {
	rv := objc.Send[SpellChecker](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpellChecker creates a new SpellChecker instance.
func NewSpellChecker() SpellChecker {
	return getSpellCheckerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, objc.String(stringToCheck), checkingTypes, options, tag, completionHandler)
	return rv
}



