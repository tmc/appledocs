// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpellChecker] class.
var SpellCheckerClass objc.Class

func init() {
	SpellCheckerClass = objc.GetClass("NSSpellChecker")
}

type SpellChecker struct {
	objc.ID
}

func SpellCheckerFrom(ptr unsafe.Pointer) SpellChecker {
	return SpellChecker{
		ID: objc.ID(ptr),
	}
}


//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	sel := objc.RegisterName("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:")
	ret := s_.ID.Send(sel, selectedRange, stringToCheck, checkingTypes, options, tag, completionHandler)
	return int(ret)
}

