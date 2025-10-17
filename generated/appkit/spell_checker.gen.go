
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [SpellChecker] class.
var SpellCheckerClass _SpellCheckerClass

func init() {
	SpellCheckerClass = _SpellCheckerClass{objc.GetClass("NSSpellChecker")}
}

type _SpellCheckerClass struct {
	objc.Class
}

// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	ID() objc.ID
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int
}

type SpellChecker struct {
	id objc.ID
}

func SpellCheckerFrom(ptr unsafe.Pointer) SpellChecker {
	return SpellChecker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SpellChecker) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SpellCheckerClass) Alloc() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SpellCheckerClass) New() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSpellChecker creates and returns a new initialized instance.
func NewSpellChecker() SpellChecker {
	return SpellCheckerClass.New()
}

// Init initializes the instance.
func (s_ SpellChecker) Init() SpellChecker {
	rv := objc.Send[SpellChecker](s_.ID(), selInit)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID(), objc.RegisterName("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, stringToCheck, checkingTypes, options, tag, completionHandler)
	return rv
}
