// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpellChecker] class.
var (
	SpellCheckerClass     _SpellCheckerClass
	SpellCheckerClassOnce sync.Once
)

func getSpellCheckerClass() _SpellCheckerClass {
	SpellCheckerClassOnce.Do(func() {
		SpellCheckerClass = _SpellCheckerClass{objc.GetClass("NSSpellChecker")}
	})
	return SpellCheckerClass
}

type _SpellCheckerClass struct {
	class objc.Class
}

// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	objectivec.IObject
	CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.IRange, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, orthography unsafe.Pointer, wordCount unsafe.Pointer) []unsafe.Pointer
	CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details []foundation.IDictionary) foundation.Range
	CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.Range
	CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, wordCount unsafe.Pointer) foundation.Range
	CountWordsInStringLanguage(stringToCount string, language string) int
	GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.IRange, string_ string, language string, tag int) []string
	GuessesForWord(word string) foundation.Array
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.IRange, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int
	RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.IRange, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int
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


// Requests unified text checking for the given range of the given string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/check(_:range:types:options:inSpellDocumentWithTag:orthography:wordCount:)
func (s_ SpellChecker) CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.IRange, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, orthography unsafe.Pointer, wordCount unsafe.Pointer) []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](s_.ID, objc.Sel("checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, orthography, wordCount)
	return rv
}

// Initiates a grammatical analysis of a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkGrammar(of:startingAt:language:wrap:inSpellDocumentWithTag:details:)
func (s_ SpellChecker) CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details []foundation.IDictionary) foundation.Range {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkGrammarOfString:startingAt:language:wrap:inSpellDocumentWithTag:details:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, details)
	return rv
}

// Starts the search for a misspelled word in starting at within the string object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.Range {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:"), objc.String(stringToCheck), startingOffset)
	return rv
}

// Starts the search for a misspelled word in a string starting at specified offset within the string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, wordCount unsafe.Pointer) foundation.Range {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, wordCount)
	return rv
}

// Returns the number of words in the specified string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/countWords(in:language:)
func (s_ SpellChecker) CountWordsInStringLanguage(stringToCount string, language string) int {
	rv := objc.Send[int](s_.ID, objc.Sel("countWordsInString:language:"), objc.String(stringToCount), objc.String(language))
	return rv
}

// Returns an array of possible substitutions for the specified string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/guesses(forWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.IRange, string_ string, language string, tag int) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("guessesForWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return rv
}

// Returns an array of suggested spellings for the misspelled word.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/guessesForWord:
func (s_ SpellChecker) GuessesForWord(word string) foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("guessesForWord:"), objc.String(word))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.IRange, stringToCheck string, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, objc.String(stringToCheck), checkingTypes, options, tag, completionHandler)
	return rv
}

// Requests that the string be checked in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.IRange, checkingTypes unsafe.Pointer, options unsafe.Pointer, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, completionHandler)
	return rv
}

// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/accessoryview
func (s_ SpellChecker) AccessoryView() NSView {
	rv := objc.Send[NSView](s_.ID, objc.Sel("accessoryView"))
	return rv
}


// SetAccessoryView sets the value of the accessoryView property.
// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/accessoryview
func (s_ SpellChecker) SetAccessoryView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccessoryView:"), value)
}

// Sets whether the spell checker will automatically identify languages.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/automaticallyidentifieslanguages
func (s_ SpellChecker) AutomaticallyIdentifiesLanguages() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyIdentifiesLanguages"))
	return rv
}


// SetAutomaticallyIdentifiesLanguages sets the value of the automaticallyIdentifiesLanguages property.
// Sets whether the spell checker will automatically identify languages.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/automaticallyidentifieslanguages
func (s_ SpellChecker) SetAutomaticallyIdentifiesLanguages(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyIdentifiesLanguages:"), value)
}

// Provides a list of all available languages.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/availablelanguages
func (s_ SpellChecker) AvailableLanguages() string {
	rv := objc.Send[string](s_.ID, objc.Sel("availableLanguages"))
	return rv
}


// SetAvailableLanguages sets the value of the availableLanguages property.
// Provides a list of all available languages.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/availablelanguages
func (s_ SpellChecker) SetAvailableLanguages(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAvailableLanguages:"), objc.String(value))
}

// Returns the spell checker’s panel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/spellingpanel
func (s_ SpellChecker) SpellingPanel() NSPanel {
	rv := objc.Send[NSPanel](s_.ID, objc.Sel("spellingPanel"))
	return rv
}


// SetSpellingPanel sets the value of the spellingPanel property.
// Returns the spell checker’s panel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/spellingpanel
func (s_ SpellChecker) SetSpellingPanel(value IPanel) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpellingPanel:"), value)
}

// Returns the substitutions panel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanel
func (s_ SpellChecker) SubstitutionsPanel() NSPanel {
	rv := objc.Send[NSPanel](s_.ID, objc.Sel("substitutionsPanel"))
	return rv
}


// SetSubstitutionsPanel sets the value of the substitutionsPanel property.
// Returns the substitutions panel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanel
func (s_ SpellChecker) SetSubstitutionsPanel(value IPanel) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstitutionsPanel:"), value)
}

// Sets the substitutions panel’s accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanelaccessoryviewcontroller
func (s_ SpellChecker) SubstitutionsPanelAccessoryViewController() NSViewController {
	rv := objc.Send[NSViewController](s_.ID, objc.Sel("substitutionsPanelAccessoryViewController"))
	return rv
}


// SetSubstitutionsPanelAccessoryViewController sets the value of the substitutionsPanelAccessoryViewController property.
// Sets the substitutions panel’s accessory view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanelaccessoryviewcontroller
func (s_ SpellChecker) SetSubstitutionsPanelAccessoryViewController(value IViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstitutionsPanelAccessoryViewController:"), value)
}

// Provides a subset of the available languages to be used for spell checking.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/userpreferredlanguages
func (s_ SpellChecker) UserPreferredLanguages() string {
	rv := objc.Send[string](s_.ID, objc.Sel("userPreferredLanguages"))
	return rv
}


// SetUserPreferredLanguages sets the value of the userPreferredLanguages property.
// Provides a subset of the available languages to be used for spell checking.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/userpreferredlanguages
func (s_ SpellChecker) SetUserPreferredLanguages(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUserPreferredLanguages:"), objc.String(value))
}

// Returns the dictionary used when replacing words.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/userreplacementsdictionary
func (s_ SpellChecker) UserReplacementsDictionary() string {
	rv := objc.Send[string](s_.ID, objc.Sel("userReplacementsDictionary"))
	return rv
}


// SetUserReplacementsDictionary sets the value of the userReplacementsDictionary property.
// Returns the dictionary used when replacing words.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/userreplacementsdictionary
func (s_ SpellChecker) SetUserReplacementsDictionary(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUserReplacementsDictionary:"), objc.String(value))
}



