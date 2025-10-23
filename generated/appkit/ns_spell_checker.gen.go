// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	AutomaticallyIdentifiesLanguages() bool /* primitive/slice/pointer. */
	SetAutomaticallyIdentifiesLanguages(value bool /* primitive/slice/pointer. */)
	AvailableLanguages() []string /* primitive/slice/pointer. */
	SpellingPanel() IPanel
	SubstitutionsPanel() IPanel
	SubstitutionsPanelAccessoryViewController() IViewController
	SetSubstitutionsPanelAccessoryViewController(value IViewController)
	UserPreferredLanguages() []string /* primitive/slice/pointer. */
	UserReplacementsDictionary() foundation.IDictionary /* already interface */
	// methods:
	CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string /* primitive/slice/pointer. */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, orthography unsafe.Pointer, wordCount Integer /* not a class type */) []foundation.objc.IObject /* cross-framework: TextCheckingResult */
	CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, wrapFlag bool /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */, details foundation.IDictionary /* already interface */) foundation.objc.IObject /* cross-framework: Range */
	CheckSpellingOfStringStartingAt(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */) foundation.objc.IObject /* cross-framework: Range */
	CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, wrapFlag bool /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */, wordCount Integer /* not a class type */) foundation.objc.IObject /* cross-framework: Range */
	CloseSpellDocumentWithTag(tag int /* primitive/slice/pointer. */)
	CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
	CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	CountWordsInStringLanguage(stringToCount string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	DeletesAutospaceBetweenStringAndStringLanguage(precedingString string /* primitive/slice/pointer. */, followingString string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	DismissCorrectionIndicatorForView(view IView)
	GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
	HasLearnedWord(word string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	IgnoreWordInSpellDocumentWithTag(wordToIgnore string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */)
	IgnoredWordsInSpellDocumentWithTag(tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
	Language() objc.IObject /* cross-framework: String */
	LanguageForWordRangeInStringOrthography(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, orthography objc.IObject /* cross-framework Orthography */) objc.IObject /* cross-framework: String */
	LearnWord(word string /* primitive/slice/pointer. */)
	MenuForResultStringOptionsAtLocationInView(result objc.IObject /* cross-framework TextCheckingResult */, checkedString string /* primitive/slice/pointer. */, options foundation.IDictionary /* already interface */, location coregraphics.CGPoint, view IView) IMenu
	PreventsAutocorrectionBeforeStringLanguage(string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse, correction string /* primitive/slice/pointer. */, word string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */)
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.objc.IObject /* cross-framework Range */, stringToCheck string /* primitive/slice/pointer. */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) int /* primitive/slice/pointer. */
	RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string /* primitive/slice/pointer. */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) int /* primitive/slice/pointer. */
	SetIgnoredWordsInSpellDocumentWithTag(words []string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */)
	SetLanguage(language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SetWordFieldStringValue(string_ string /* primitive/slice/pointer. */)
	ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType, primaryString string /* primitive/slice/pointer. */, alternativeStrings []string /* primitive/slice/pointer. */, rectOfTypedString coregraphics.CGRect, view IView, completionBlock unsafe.Pointer)
	ShowInlinePredictionForCandidatesClient(candidates []foundation.objc.IObject /* cross-framework TextCheckingResult */, client objectivec.IObject)
	UnlearnWord(word string /* primitive/slice/pointer. */)
	UpdatePanels()
	UpdateSpellingPanelWithGrammarStringDetail(string_ string /* primitive/slice/pointer. */, detail foundation.IDictionary /* already interface */)
	UpdateSpellingPanelWithMisspelledWord(word string /* primitive/slice/pointer. */)
	UserQuotesArrayForLanguage(language string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
}

// An interface to the Cocoa spell-checking service.
//
// To handle all its spell checking, an app needs only one instance of , known as the spell checker. Using the spell checker you manage the Spelling panel, in which the user can specify decisions about words that are suspect. The spell checker also offers the ability to provide word completions to augment the text completion system.


// An interface to the Cocoa spell-checking service.
//
// [Full Topic]
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



// Returns a guaranteed unique tag to use as the spell-document tag for a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/uniqueSpellDocumentTag()
func (sc _SpellCheckerClass) UniqueSpellDocumentTag() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](objc.ID(sc.class), objc.Sel("uniqueSpellDocumentTag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticCapitalizationEnabled
func (sc _SpellCheckerClass) AutomaticCapitalizationEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticCapitalizationEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticDashSubstitutionEnabled
func (sc _SpellCheckerClass) AutomaticDashSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticDashSubstitutionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticInlinePredictionEnabled
func (sc _SpellCheckerClass) AutomaticInlinePredictionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticInlinePredictionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticPeriodSubstitutionEnabled
func (sc _SpellCheckerClass) AutomaticPeriodSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticPeriodSubstitutionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticQuoteSubstitutionEnabled
func (sc _SpellCheckerClass) AutomaticQuoteSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticQuoteSubstitutionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticSpellingCorrectionEnabled
func (sc _SpellCheckerClass) AutomaticSpellingCorrectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticSpellingCorrectionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextCompletionEnabled
func (sc _SpellCheckerClass) AutomaticTextCompletionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticTextCompletionEnabled"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextReplacementEnabled
func (sc _SpellCheckerClass) AutomaticTextReplacementEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticTextReplacementEnabled"))
	return rv
}

// Returns the NSSpellChecker (one per application).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/shared
func (sc _SpellCheckerClass) SharedSpellChecker() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.class), objc.Sel("sharedSpellChecker"))
	return rv
}

// Returns whether the application’s NSSpellChecker has already been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/sharedSpellCheckerExists
func (sc _SpellCheckerClass) SharedSpellCheckerExists() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("sharedSpellCheckerExists"))
	return rv
}

// Requests unified text checking for the given range of the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/check(_:range:types:options:inSpellDocumentWithTag:orthography:wordCount:)
func (s_ SpellChecker) CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string /* primitive/slice/pointer. */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, orthography unsafe.Pointer, wordCount Integer /* not a class type */) []foundation.objc.IObject /* cross-framework: TextCheckingResult */ {
	rv := objc.Send[[]foundation.TextCheckingResult](s_.ID, objc.Sel("checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, orthography, wordCount)
	return rv
}


// Initiates a grammatical analysis of a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkGrammar(of:startingAt:language:wrap:inSpellDocumentWithTag:details:)
func (s_ SpellChecker) CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, wrapFlag bool /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */, details foundation.IDictionary /* already interface */) foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkGrammarOfString:startingAt:language:wrap:inSpellDocumentWithTag:details:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, details)
	return rv
}


// Starts the search for a misspelled word in starting at within the string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAt(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */) foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:"), objc.String(stringToCheck), startingOffset)
	return rv
}


// Starts the search for a misspelled word in a string starting at specified offset within the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string /* primitive/slice/pointer. */, startingOffset int /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, wrapFlag bool /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */, wordCount Integer /* not a class type */) foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, wordCount)
	return rv
}


// Notifies the receiver that the user has finished with the tagged document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/closeSpellDocument(withTag:)
func (s_ SpellChecker) CloseSpellDocumentWithTag(tag int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("closeSpellDocumentWithTag:"), tag)
}


// Provides a list of complete words that the user might be trying to type based on a partial word in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/completions(forPartialWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("completionsForPartialWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return rv
}


// Returns a single proposed correction if a word is mis-spelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/correction(forWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](s_.ID, objc.Sel("correctionForWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return rv
}


// Returns the number of words in the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/countWords(in:language:)
func (s_ SpellChecker) CountWordsInStringLanguage(stringToCount string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("countWordsInString:language:"), objc.String(stringToCount), objc.String(language))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/deletesAutospaceBetweenString(_:andString:language:)
func (s_ SpellChecker) DeletesAutospaceBetweenStringAndStringLanguage(precedingString string /* primitive/slice/pointer. */, followingString string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("deletesAutospaceBetweenString:andString:language:"), objc.String(precedingString), objc.String(followingString), objc.String(language))
	return rv
}


// Dismisses the correction indicator for the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/dismissCorrectionIndicator(for:)
func (s_ SpellChecker) DismissCorrectionIndicatorForView(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("dismissCorrectionIndicatorForView:"), view)
}


// Returns an array of possible substitutions for the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/guesses(forWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("guessesForWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return rv
}


// Indicates whether the spell checker has learned a given word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/hasLearnedWord(_:)
func (s_ SpellChecker) HasLearnedWord(word string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasLearnedWord:"), objc.String(word))
	return rv
}


// Instructs the spell checker to ignore all future occurrences of in the document identified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoreWord(_:inSpellDocumentWithTag:)
func (s_ SpellChecker) IgnoreWordInSpellDocumentWithTag(wordToIgnore string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("ignoreWord:inSpellDocumentWithTag:"), objc.String(wordToIgnore), tag)
}


// Returns the array of ignored words for a document identified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoredWords(inSpellDocumentWithTag:)
func (s_ SpellChecker) IgnoredWordsInSpellDocumentWithTag(tag int /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("ignoredWordsInSpellDocumentWithTag:"), tag)
	return rv
}


// Returns the current language used in spell checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language()
func (s_ SpellChecker) Language() objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](s_.ID, objc.Sel("language"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language(forWordRange:in:orthography:)
func (s_ SpellChecker) LanguageForWordRangeInStringOrthography(range_ foundation.objc.IObject /* cross-framework Range */, string_ string /* primitive/slice/pointer. */, orthography objc.IObject /* cross-framework Orthography */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](s_.ID, objc.Sel("languageForWordRange:inString:orthography:"), range_, objc.String(string_), orthography)
	return rv
}


// Adds the word to the spell checker dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/learnWord(_:)
func (s_ SpellChecker) LearnWord(word string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("learnWord:"), objc.String(word))
}


// Provides a menu containing contextual menu items suitable for certain kinds of detected results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/menu(for:string:options:atLocation:in:)
func (s_ SpellChecker) MenuForResultStringOptionsAtLocationInView(result objc.IObject /* cross-framework TextCheckingResult */, checkedString string /* primitive/slice/pointer. */, options foundation.IDictionary /* already interface */, location coregraphics.CGPoint, view IView) IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("menuForResult:string:options:atLocation:inView:"), result, objc.String(checkedString), options, location, view)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/preventsAutocorrection(before:language:)
func (s_ SpellChecker) PreventsAutocorrectionBeforeStringLanguage(string_ string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsAutocorrectionBeforeString:language:"), objc.String(string_), objc.String(language))
	return rv
}


// Records the user response to the correction indicator being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/record(_:toCorrection:forWord:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse, correction string /* primitive/slice/pointer. */, word string /* primitive/slice/pointer. */, language string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("recordResponse:toCorrection:forWord:language:inSpellDocumentWithTag:"), response, objc.String(correction), objc.String(word), objc.String(language), tag)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.objc.IObject /* cross-framework Range */, stringToCheck string /* primitive/slice/pointer. */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, objc.String(stringToCheck), checkingTypes, options, tag, completionHandler)
	return rv
}


// Requests that the string be checked in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string /* primitive/slice/pointer. */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, tag int /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, completionHandler)
	return rv
}


// Initializes the ignored-words document (a dictionary identified by with ), an array of words to ignore.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setIgnoredWords(_:inSpellDocumentWithTag:)
func (s_ SpellChecker) SetIgnoredWordsInSpellDocumentWithTag(words []string /* primitive/slice/pointer. */, tag int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoredWords:inSpellDocumentWithTag:"), words, tag)
}


// Returns whether the specified language is in the Spelling pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setLanguage(_:)
func (s_ SpellChecker) SetLanguage(language string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("setLanguage:"), objc.String(language))
	return rv
}


// Sets the string that appears in the misspelled word field, using the string object .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setWordFieldStringValue(_:)
func (s_ SpellChecker) SetWordFieldStringValue(string_ string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWordFieldStringValue:"), objc.String(string_))
}


// Display a suitable user interface to indicate a correction may need to be made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/showCorrectionIndicator(of:primaryString:alternativeStrings:forStringIn:view:completionHandler:)
func (s_ SpellChecker) ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType, primaryString string /* primitive/slice/pointer. */, alternativeStrings []string /* primitive/slice/pointer. */, rectOfTypedString coregraphics.CGRect, view IView, completionBlock unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showCorrectionIndicatorOfType:primaryString:alternativeStrings:forStringInRect:view:completionHandler:"), type_, objc.String(primaryString), alternativeStrings, rectOfTypedString, view, completionBlock)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/showInlinePrediction(forCandidates:client:)
func (s_ SpellChecker) ShowInlinePredictionForCandidatesClient(candidates []foundation.objc.IObject /* cross-framework TextCheckingResult */, client objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showInlinePredictionForCandidates:client:"), candidates, client)
}


// Tells the spell checker to unlearn a given word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/unlearnWord(_:)
func (s_ SpellChecker) UnlearnWord(word string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("unlearnWord:"), objc.String(word))
}


// Updates the available panels to account for user changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updatePanels()
func (s_ SpellChecker) UpdatePanels() {
	objc.Send[objc.ID](s_.ID, objc.Sel("updatePanels"))
}


// Specifies a grammar-analysis detail to highlight in the Spelling panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withGrammarString:detail:)
func (s_ SpellChecker) UpdateSpellingPanelWithGrammarStringDetail(string_ string /* primitive/slice/pointer. */, detail foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateSpellingPanelWithGrammarString:detail:"), objc.String(string_), detail)
}


// Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withMisspelledWord:)
func (s_ SpellChecker) UpdateSpellingPanelWithMisspelledWord(word string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateSpellingPanelWithMisspelledWord:"), objc.String(word))
}


// Returns the default values for quote replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userQuotesArray(forLanguage:)
func (s_ SpellChecker) UserQuotesArrayForLanguage(language string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("userQuotesArrayForLanguage:"), objc.String(language))
	return rv
}


// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/accessoryView
func (s_ SpellChecker) AccessoryView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("accessoryView"))
	return rv
}


// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/accessoryView
func (s_ SpellChecker) SetAccessoryView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccessoryView:"), value)
}


// Sets whether the spell checker will automatically identify languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/automaticallyIdentifiesLanguages
func (s_ SpellChecker) AutomaticallyIdentifiesLanguages() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyIdentifiesLanguages"))
	return rv
}


// Sets whether the spell checker will automatically identify languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/automaticallyIdentifiesLanguages
func (s_ SpellChecker) SetAutomaticallyIdentifiesLanguages(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyIdentifiesLanguages:"), value)
}


// Provides a list of all available languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/availableLanguages
func (s_ SpellChecker) AvailableLanguages() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("availableLanguages"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticCapitalizationEnabled
func (s_ SpellChecker) AutomaticCapitalizationEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticCapitalizationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticDashSubstitutionEnabled
func (s_ SpellChecker) AutomaticDashSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticDashSubstitutionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticInlinePredictionEnabled
func (s_ SpellChecker) AutomaticInlinePredictionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticInlinePredictionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticPeriodSubstitutionEnabled
func (s_ SpellChecker) AutomaticPeriodSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticPeriodSubstitutionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticQuoteSubstitutionEnabled
func (s_ SpellChecker) AutomaticQuoteSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticQuoteSubstitutionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticSpellingCorrectionEnabled
func (s_ SpellChecker) AutomaticSpellingCorrectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticSpellingCorrectionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextCompletionEnabled
func (s_ SpellChecker) AutomaticTextCompletionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextReplacementEnabled
func (s_ SpellChecker) AutomaticTextReplacementEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticTextReplacementEnabled"))
	return rv
}


// Returns the NSSpellChecker (one per application).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/shared
func (s_ SpellChecker) SharedSpellChecker() ISpellChecker {
	rv := objc.Send[SpellChecker](s_.ID, objc.Sel("sharedSpellChecker"))
	return rv
}


// Returns whether the application’s NSSpellChecker has already been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/sharedSpellCheckerExists
func (s_ SpellChecker) SharedSpellCheckerExists() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("sharedSpellCheckerExists"))
	return rv
}


// Returns the spell checker’s panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/spellingPanel
func (s_ SpellChecker) SpellingPanel() IPanel {
	rv := objc.Send[Panel](s_.ID, objc.Sel("spellingPanel"))
	return rv
}


// Returns the substitutions panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/substitutionsPanel
func (s_ SpellChecker) SubstitutionsPanel() IPanel {
	rv := objc.Send[Panel](s_.ID, objc.Sel("substitutionsPanel"))
	return rv
}


// Sets the substitutions panel’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/substitutionsPanelAccessoryViewController
func (s_ SpellChecker) SubstitutionsPanelAccessoryViewController() IViewController {
	rv := objc.Send[ViewController](s_.ID, objc.Sel("substitutionsPanelAccessoryViewController"))
	return rv
}


// Sets the substitutions panel’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/substitutionsPanelAccessoryViewController
func (s_ SpellChecker) SetSubstitutionsPanelAccessoryViewController(value IViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstitutionsPanelAccessoryViewController:"), value)
}


// Provides a subset of the available languages to be used for spell checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userPreferredLanguages
func (s_ SpellChecker) UserPreferredLanguages() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](s_.ID, objc.Sel("userPreferredLanguages"))
	return rv
}


// Returns the dictionary used when replacing words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userReplacementsDictionary
func (s_ SpellChecker) UserReplacementsDictionary() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("userReplacementsDictionary"))
	return rv
}



