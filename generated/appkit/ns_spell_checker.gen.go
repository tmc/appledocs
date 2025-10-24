// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSSpellChecker */


/* debug [class_header]: Header for NSSpellChecker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpellChecker */
// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpellChecker */
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	AutomaticallyIdentifiesLanguages() bool
	SetAutomaticallyIdentifiesLanguages(value bool)
	AvailableLanguages() []string
	SpellingPanel() IPanel
	UserPreferredLanguages() []string
	UserReplacementsDictionary() foundation.IDictionary
	SubstitutionsPanel() IPanel
	SetSubstitutionsPanel(value IPanel)
	SubstitutionsPanelAccessoryViewController() IViewController
	SetSubstitutionsPanelAccessoryViewController(value IViewController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpellChecker */
	// methods:
	CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck objc.IObject /* cross-framework: NSString */, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, orthography objectivec.IObject, wordCount int) []foundation.TextCheckingResult
	CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int, language objc.IObject /* cross-framework: NSString */, wrapFlag bool, tag int, details foundation.IDictionary) corefoundation.Range
	CheckSpellingOfStringStartingAt(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int) corefoundation.Range
	CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int, language objc.IObject /* cross-framework: NSString */, wrapFlag bool, tag int, wordCount int) corefoundation.Range
	CloseSpellDocumentWithTag(tag int)
	CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) []string
	CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) foundation.String
	CountWordsInStringLanguage(stringToCount objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */) int
	DeletesAutospaceBetweenStringAndStringLanguage(precedingString objc.IObject /* cross-framework: NSString */, followingString objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */) bool
	GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) []string
	HasLearnedWord(word objc.IObject /* cross-framework: NSString */) bool
	IgnoreWordInSpellDocumentWithTag(wordToIgnore objc.IObject /* cross-framework: NSString */, tag int)
	IgnoredWordsInSpellDocumentWithTag(tag int) []string
	Language() foundation.String
	LanguageForWordRangeInStringOrthography(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, orthography foundation.Orthography) foundation.String
	LearnWord(word objc.IObject /* cross-framework: NSString */)
	MenuForResultStringOptionsAtLocationInView(result foundation.TextCheckingResult, checkedString objc.IObject /* cross-framework: NSString */, options foundation.IDictionary, location vision.Point, view IView) IMenu
	RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse /* not a class type */, correction objc.IObject /* cross-framework: NSString */, word objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int)
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange corefoundation.Range, stringToCheck objc.IObject /* cross-framework: NSString */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, completionHandler unsafe.Pointer) int
	RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck objc.IObject /* cross-framework: NSString */, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, completionHandler unsafe.Pointer) int
	SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int)
	SetLanguage(language objc.IObject /* cross-framework: NSString */) bool
	SetWordFieldStringValue(string_ objc.IObject /* cross-framework: NSString */)
	ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType /* not a class type */, primaryString objc.IObject /* cross-framework: NSString */, alternativeStrings []string, rectOfTypedString Rect /* not a class type */, view IView, completionBlock unsafe.Pointer)
	UnlearnWord(word objc.IObject /* cross-framework: NSString */)
	UpdateSpellingPanelWithGrammarStringDetail(string_ objc.IObject /* cross-framework: NSString */, detail foundation.IDictionary)
	UpdateSpellingPanelWithMisspelledWord(word objc.IObject /* cross-framework: NSString */)
	UserQuotesArrayForLanguage(language objc.IObject /* cross-framework: NSString */) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpellChecker */
// Alloc allocates a new instance without initialization.
func (sc _SpellCheckerClass) Alloc() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpellChecker */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpellChecker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpellChecker */

// Returns a guaranteed unique tag to use as the spell-document tag for a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/uniqueSpellDocumentTag()
func (sc _SpellCheckerClass) UniqueSpellDocumentTag() int {
	rv := objc.Send[int](objc.ID(sc.class), objc.Sel("uniqueSpellDocumentTag"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UniqueSpellDocumentTag) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpellChecker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticCapitalizationEnabled
func (sc _SpellCheckerClass) AutomaticCapitalizationEnabled() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticCapitalizationEnabled"))
	return rv
}/* debug [class_properties_class/property]: automaticCapitalizationEnabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticPeriodSubstitutionEnabled
func (sc _SpellCheckerClass) AutomaticPeriodSubstitutionEnabled() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticPeriodSubstitutionEnabled"))
	return rv
}/* debug [class_properties_class/property]: automaticPeriodSubstitutionEnabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextCompletionEnabled
func (sc _SpellCheckerClass) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticTextCompletionEnabled"))
	return rv
}/* debug [class_properties_class/property]: automaticTextCompletionEnabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextReplacementEnabled
func (sc _SpellCheckerClass) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("automaticTextReplacementEnabled"))
	return rv
}/* debug [class_properties_class/property]: automaticTextReplacementEnabled */

// Returns the NSSpellChecker (one per application).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/shared
func (sc _SpellCheckerClass) SharedSpellChecker() SpellChecker {
	rv := objc.Send[SpellChecker](objc.ID(sc.class), objc.Sel("sharedSpellChecker"))
	return rv
}/* debug [class_properties_class/property]: sharedSpellChecker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpellChecker */

// Requests unified text checking for the given range of the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/check(_:range:types:options:inSpellDocumentWithTag:orthography:wordCount:)
func (s_ SpellChecker) CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck objc.IObject /* cross-framework: NSString */, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, orthography objectivec.IObject, wordCount int) []foundation.TextCheckingResult {
	rv := objc.Send[[]foundation.TextCheckingResult](s_.ID, objc.Sel("checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:"), stringToCheck, range_, checkingTypes, options, tag, orthography, wordCount)
	return rv
}/* debug [instance_methods/method]: CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount */


// Initiates a grammatical analysis of a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkGrammar(of:startingAt:language:wrap:inSpellDocumentWithTag:details:)
func (s_ SpellChecker) CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int, language objc.IObject /* cross-framework: NSString */, wrapFlag bool, tag int, details foundation.IDictionary) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("checkGrammarOfString:startingAt:language:wrap:inSpellDocumentWithTag:details:"), stringToCheck, startingOffset, language, wrapFlag, tag, details)
	return rv
}/* debug [instance_methods/method]: CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails */


// Starts the search for a misspelled word in starting at within the string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAt(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:"), stringToCheck, startingOffset)
	return rv
}/* debug [instance_methods/method]: CheckSpellingOfStringStartingAt */


// Starts the search for a misspelled word in a string starting at specified offset within the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:)
func (s_ SpellChecker) CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck objc.IObject /* cross-framework: NSString */, startingOffset int, language objc.IObject /* cross-framework: NSString */, wrapFlag bool, tag int, wordCount int) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("checkSpellingOfString:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:"), stringToCheck, startingOffset, language, wrapFlag, tag, wordCount)
	return rv
}/* debug [instance_methods/method]: CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount */


// Notifies the receiver that the user has finished with the tagged document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/closeSpellDocument(withTag:)
func (s_ SpellChecker) CloseSpellDocumentWithTag(tag int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("closeSpellDocumentWithTag:"), tag)
}/* debug [instance_methods/method]: CloseSpellDocumentWithTag */


// Provides a list of complete words that the user might be trying to type based on a partial word in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/completions(forPartialWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("completionsForPartialWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}/* debug [instance_methods/method]: CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag */


// Returns a single proposed correction if a word is mis-spelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/correction(forWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("correctionForWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}/* debug [instance_methods/method]: CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag */


// Returns the number of words in the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/countWords(in:language:)
func (s_ SpellChecker) CountWordsInStringLanguage(stringToCount objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](s_.ID, objc.Sel("countWordsInString:language:"), stringToCount, language)
	return rv
}/* debug [instance_methods/method]: CountWordsInStringLanguage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/deletesAutospaceBetweenString(_:andString:language:)
func (s_ SpellChecker) DeletesAutospaceBetweenStringAndStringLanguage(precedingString objc.IObject /* cross-framework: NSString */, followingString objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("deletesAutospaceBetweenString:andString:language:"), precedingString, followingString, language)
	return rv
}/* debug [instance_methods/method]: DeletesAutospaceBetweenStringAndStringLanguage */


// Returns an array of possible substitutions for the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/guesses(forWordRange:in:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("guessesForWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}/* debug [instance_methods/method]: GuessesForWordRangeInStringLanguageInSpellDocumentWithTag */


// Indicates whether the spell checker has learned a given word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/hasLearnedWord(_:)
func (s_ SpellChecker) HasLearnedWord(word objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasLearnedWord:"), word)
	return rv
}/* debug [instance_methods/method]: HasLearnedWord */


// Instructs the spell checker to ignore all future occurrences of in the document identified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoreWord(_:inSpellDocumentWithTag:)
func (s_ SpellChecker) IgnoreWordInSpellDocumentWithTag(wordToIgnore objc.IObject /* cross-framework: NSString */, tag int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("ignoreWord:inSpellDocumentWithTag:"), wordToIgnore, tag)
}/* debug [instance_methods/method]: IgnoreWordInSpellDocumentWithTag */


// Returns the array of ignored words for a document identified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoredWords(inSpellDocumentWithTag:)
func (s_ SpellChecker) IgnoredWordsInSpellDocumentWithTag(tag int) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("ignoredWordsInSpellDocumentWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: IgnoredWordsInSpellDocumentWithTag */


// Returns the current language used in spell checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language()
func (s_ SpellChecker) Language() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_methods/method]: Language */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language(forWordRange:in:orthography:)
func (s_ SpellChecker) LanguageForWordRangeInStringOrthography(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */, orthography foundation.Orthography) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("languageForWordRange:inString:orthography:"), range_, string_, orthography)
	return rv
}/* debug [instance_methods/method]: LanguageForWordRangeInStringOrthography */


// Adds the word to the spell checker dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/learnWord(_:)
func (s_ SpellChecker) LearnWord(word objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("learnWord:"), word)
}/* debug [instance_methods/method]: LearnWord */


// Provides a menu containing contextual menu items suitable for certain kinds of detected results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/menu(for:string:options:atLocation:in:)
func (s_ SpellChecker) MenuForResultStringOptionsAtLocationInView(result foundation.TextCheckingResult, checkedString objc.IObject /* cross-framework: NSString */, options foundation.IDictionary, location vision.Point, view IView) IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("menuForResult:string:options:atLocation:inView:"), result, checkedString, options, location, view)
	return rv
}/* debug [instance_methods/method]: MenuForResultStringOptionsAtLocationInView */


// Records the user response to the correction indicator being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/record(_:toCorrection:forWord:language:inSpellDocumentWithTag:)
func (s_ SpellChecker) RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse /* not a class type */, correction objc.IObject /* cross-framework: NSString */, word objc.IObject /* cross-framework: NSString */, language objc.IObject /* cross-framework: NSString */, tag int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("recordResponse:toCorrection:forWord:language:inSpellDocumentWithTag:"), response, correction, word, language, tag)
}/* debug [instance_methods/method]: RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange corefoundation.Range, stringToCheck objc.IObject /* cross-framework: NSString */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, stringToCheck, checkingTypes, options, tag, completionHandler)
	return rv
}/* debug [instance_methods/method]: RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler */


// Requests that the string be checked in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)
func (s_ SpellChecker) RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck objc.IObject /* cross-framework: NSString */, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, tag int, completionHandler unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:"), stringToCheck, range_, checkingTypes, options, tag, completionHandler)
	return rv
}/* debug [instance_methods/method]: RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler */


// Initializes the ignored-words document (a dictionary identified by with ), an array of words to ignore.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setIgnoredWords(_:inSpellDocumentWithTag:)
func (s_ SpellChecker) SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoredWords:inSpellDocumentWithTag:"), words, tag)
}/* debug [instance_methods/method]: SetIgnoredWordsInSpellDocumentWithTag */


// Returns whether the specified language is in the Spelling pop-up list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setLanguage(_:)
func (s_ SpellChecker) SetLanguage(language objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("setLanguage:"), language)
	return rv
}/* debug [instance_methods/method]: SetLanguage */


// Sets the string that appears in the misspelled word field, using the string object .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setWordFieldStringValue(_:)
func (s_ SpellChecker) SetWordFieldStringValue(string_ objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWordFieldStringValue:"), string_)
}/* debug [instance_methods/method]: SetWordFieldStringValue */


// Display a suitable user interface to indicate a correction may need to be made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/showCorrectionIndicator(of:primaryString:alternativeStrings:forStringIn:view:completionHandler:)
func (s_ SpellChecker) ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType /* not a class type */, primaryString objc.IObject /* cross-framework: NSString */, alternativeStrings []string, rectOfTypedString Rect /* not a class type */, view IView, completionBlock unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showCorrectionIndicatorOfType:primaryString:alternativeStrings:forStringInRect:view:completionHandler:"), type_, primaryString, alternativeStrings, rectOfTypedString, view, completionBlock)
}/* debug [instance_methods/method]: ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler */


// Tells the spell checker to unlearn a given word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/unlearnWord(_:)
func (s_ SpellChecker) UnlearnWord(word objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("unlearnWord:"), word)
}/* debug [instance_methods/method]: UnlearnWord */


// Specifies a grammar-analysis detail to highlight in the Spelling panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withGrammarString:detail:)
func (s_ SpellChecker) UpdateSpellingPanelWithGrammarStringDetail(string_ objc.IObject /* cross-framework: NSString */, detail foundation.IDictionary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateSpellingPanelWithGrammarString:detail:"), string_, detail)
}/* debug [instance_methods/method]: UpdateSpellingPanelWithGrammarStringDetail */


// Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withMisspelledWord:)
func (s_ SpellChecker) UpdateSpellingPanelWithMisspelledWord(word objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateSpellingPanelWithMisspelledWord:"), word)
}/* debug [instance_methods/method]: UpdateSpellingPanelWithMisspelledWord */


// Returns the default values for quote replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userQuotesArray(forLanguage:)
func (s_ SpellChecker) UserQuotesArrayForLanguage(language objc.IObject /* cross-framework: NSString */) []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("userQuotesArrayForLanguage:"), language)
	return rv
}/* debug [instance_methods/method]: UserQuotesArrayForLanguage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpellChecker */

// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/accessoryView
func (s_ SpellChecker) AccessoryView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/accessoryView
func (s_ SpellChecker) SetAccessoryView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// Sets whether the spell checker will automatically identify languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/automaticallyIdentifiesLanguages
func (s_ SpellChecker) AutomaticallyIdentifiesLanguages() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyIdentifiesLanguages"))
	return rv
}/* debug [instance_properties/getter]: automaticallyIdentifiesLanguages */


// Sets whether the spell checker will automatically identify languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/automaticallyIdentifiesLanguages
func (s_ SpellChecker) SetAutomaticallyIdentifiesLanguages(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyIdentifiesLanguages:"), value)
}/* debug [instance_properties/setter]: automaticallyIdentifiesLanguages */


// Provides a list of all available languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/availableLanguages
func (s_ SpellChecker) AvailableLanguages() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("availableLanguages"))
	return rv
}/* debug [instance_properties/getter]: availableLanguages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticCapitalizationEnabled
func (s_ SpellChecker) AutomaticCapitalizationEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticCapitalizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticCapitalizationEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticPeriodSubstitutionEnabled
func (s_ SpellChecker) AutomaticPeriodSubstitutionEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticPeriodSubstitutionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticPeriodSubstitutionEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextCompletionEnabled
func (s_ SpellChecker) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticTextCompletionEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextReplacementEnabled
func (s_ SpellChecker) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticTextReplacementEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticTextReplacementEnabled */


// Returns the NSSpellChecker (one per application).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/shared
func (s_ SpellChecker) SharedSpellChecker() ISpellChecker {
	rv := objc.Send[SpellChecker](s_.ID, objc.Sel("sharedSpellChecker"))
	return rv
}/* debug [instance_properties/getter]: sharedSpellChecker */


// Returns the spell checker’s panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/spellingPanel
func (s_ SpellChecker) SpellingPanel() IPanel {
	rv := objc.Send[Panel](s_.ID, objc.Sel("spellingPanel"))
	return rv
}/* debug [instance_properties/getter]: spellingPanel */


// Provides a subset of the available languages to be used for spell checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userPreferredLanguages
func (s_ SpellChecker) UserPreferredLanguages() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("userPreferredLanguages"))
	return rv
}/* debug [instance_properties/getter]: userPreferredLanguages */


// Returns the dictionary used when replacing words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userReplacementsDictionary
func (s_ SpellChecker) UserReplacementsDictionary() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("userReplacementsDictionary"))
	return rv
}/* debug [instance_properties/getter]: userReplacementsDictionary */


// Returns the substitutions panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanel
func (s_ SpellChecker) SubstitutionsPanel() IPanel {
	rv := objc.Send[Panel](s_.ID, objc.Sel("substitutionsPanel"))
	return rv
}/* debug [instance_properties/getter]: substitutionsPanel */


// Returns the substitutions panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanel
func (s_ SpellChecker) SetSubstitutionsPanel(value IPanel) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstitutionsPanel:"), value)
}/* debug [instance_properties/setter]: substitutionsPanel */


// Sets the substitutions panel’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanelaccessoryviewcontroller
func (s_ SpellChecker) SubstitutionsPanelAccessoryViewController() IViewController {
	rv := objc.Send[ViewController](s_.ID, objc.Sel("substitutionsPanelAccessoryViewController"))
	return rv
}/* debug [instance_properties/getter]: substitutionsPanelAccessoryViewController */


// Sets the substitutions panel’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/substitutionspanelaccessoryviewcontroller
func (s_ SpellChecker) SetSubstitutionsPanelAccessoryViewController(value IViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstitutionsPanelAccessoryViewController:"), value)
}/* debug [instance_properties/setter]: substitutionsPanelAccessoryViewController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSpellChecker */



