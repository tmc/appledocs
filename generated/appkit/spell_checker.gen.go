// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpellChecker] class.
var spellCheckerClass = _SpellCheckerClass{objc.GetClass("NSSpellChecker")}

type _SpellCheckerClass struct {
	class objc.Class
}

// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	objectivec.IObject
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange unsafe.Pointer, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int
}

// An interface to the Cocoa spell-checking service. [Full Topic]
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

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange unsafe.Pointer, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, stringToCheck, checkingTypes, options, tag, completionHandler)
	return rv
}


