// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSpellServerDelegate is the NSSpellServerDelegate protocol interface.
//
// The optional methods implemented by the delegate of a spell server.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSSpellServerDelegate
type PSpellServerDelegate interface {
	// Optional methods
	SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender ISpellServer, stringToCheck IString, offset uint, checkingTypes TextCheckingTypes, options IDictionary, orthography IOrthography, wordCount int) []TextCheckingResult
	HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool
	SpellServerCheckGrammarInStringLanguageDetails(sender ISpellServer, stringToCheck IString, language IString, details IDictionary) Range
	HasSpellServerCheckGrammarInStringLanguageDetails() bool
	SpellServerDidForgetWordInLanguage(sender ISpellServer, word IString, language IString)
	HasSpellServerDidForgetWordInLanguage() bool
	SpellServerDidLearnWordInLanguage(sender ISpellServer, word IString, language IString)
	HasSpellServerDidLearnWordInLanguage() bool
	SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender ISpellServer, stringToCheck IString, language IString, wordCount int, countOnly bool) Range
	HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool
	SpellServerRecordResponseToCorrectionForWordLanguage(sender ISpellServer, response uint, correction IString, word IString, language IString)
	HasSpellServerRecordResponseToCorrectionForWordLanguage() bool
	SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender ISpellServer, range_ Range, string_ IString, language IString) []string
	HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool
	SpellServerSuggestGuessesForWordInLanguage(sender ISpellServer, word IString, language IString) []string
	HasSpellServerSuggestGuessesForWordInLanguage() bool
}

// SpellServerDelegate is a delegate implementation builder for the PSpellServerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpellServerDelegate struct {
	_SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount func(sender ISpellServer, stringToCheck IString, offset uint, checkingTypes TextCheckingTypes, options IDictionary, orthography IOrthography, wordCount int) []TextCheckingResult
	_SpellServerCheckGrammarInStringLanguageDetails func(sender ISpellServer, stringToCheck IString, language IString, details IDictionary) Range
	_SpellServerDidForgetWordInLanguage func(sender ISpellServer, word IString, language IString)
	_SpellServerDidLearnWordInLanguage func(sender ISpellServer, word IString, language IString)
	_SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly func(sender ISpellServer, stringToCheck IString, language IString, wordCount int, countOnly bool) Range
	_SpellServerRecordResponseToCorrectionForWordLanguage func(sender ISpellServer, response uint, correction IString, word IString, language IString)
	_SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage func(sender ISpellServer, range_ Range, string_ IString, language IString) []string
	_SpellServerSuggestGuessesForWordInLanguage func(sender ISpellServer, word IString, language IString) []string
}

// SetSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount sets the handler for the SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount delegate method.
//
// Gives the delegate the opportunity to analyze both the spelling and grammar simultaneously, which is more efficient.
func (d *SpellServerDelegate) SetSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(f func(sender ISpellServer, stringToCheck IString, offset uint, checkingTypes TextCheckingTypes, options IDictionary, orthography IOrthography, wordCount int) []TextCheckingResult) {
	d._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount = f
}

// SetSpellServerCheckGrammarInStringLanguageDetails sets the handler for the SpellServerCheckGrammarInStringLanguageDetails delegate method.
//
// Gives the delegate the opportunity to customize the grammatical analysis of a given string.
func (d *SpellServerDelegate) SetSpellServerCheckGrammarInStringLanguageDetails(f func(sender ISpellServer, stringToCheck IString, language IString, details IDictionary) Range) {
	d._SpellServerCheckGrammarInStringLanguageDetails = f
}

// SetSpellServerDidForgetWordInLanguage sets the handler for the SpellServerDidForgetWordInLanguage delegate method.
//
// Notifies the delegate that the sender has removed the specified word from the user’s list of acceptable words in the specified language.
func (d *SpellServerDelegate) SetSpellServerDidForgetWordInLanguage(f func(sender ISpellServer, word IString, language IString)) {
	d._SpellServerDidForgetWordInLanguage = f
}

// SetSpellServerDidLearnWordInLanguage sets the handler for the SpellServerDidLearnWordInLanguage delegate method.
//
// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language.
func (d *SpellServerDelegate) SetSpellServerDidLearnWordInLanguage(f func(sender ISpellServer, word IString, language IString)) {
	d._SpellServerDidLearnWordInLanguage = f
}

// SetSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly sets the handler for the SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly delegate method.
//
// Asks the delegate to search for a misspelled word in a given string, using the specified language, and marking the first misspelled word found by returning its range within the string.
func (d *SpellServerDelegate) SetSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(f func(sender ISpellServer, stringToCheck IString, language IString, wordCount int, countOnly bool) Range) {
	d._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly = f
}

// SetSpellServerRecordResponseToCorrectionForWordLanguage sets the handler for the SpellServerRecordResponseToCorrectionForWordLanguage delegate method.
//
// Notifies the spell checker of the users’s response to a correction.
func (d *SpellServerDelegate) SetSpellServerRecordResponseToCorrectionForWordLanguage(f func(sender ISpellServer, response uint, correction IString, word IString, language IString)) {
	d._SpellServerRecordResponseToCorrectionForWordLanguage = f
}

// SetSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage sets the handler for the SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage delegate method.
//
// This delegate method returns an array of possible word completions from the spell checker, based on a partially completed string and a given range.
func (d *SpellServerDelegate) SetSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(f func(sender ISpellServer, range_ Range, string_ IString, language IString) []string) {
	d._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage = f
}

// SetSpellServerSuggestGuessesForWordInLanguage sets the handler for the SpellServerSuggestGuessesForWordInLanguage delegate method.
//
// Gives the delegate the opportunity to suggest guesses to the sender for the correct spelling of the given misspelled word in the specified language.
func (d *SpellServerDelegate) SetSpellServerSuggestGuessesForWordInLanguage(f func(sender ISpellServer, word IString, language IString) []string) {
	d._SpellServerSuggestGuessesForWordInLanguage = f
}

// SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender ISpellServer, stringToCheck IString, offset uint, checkingTypes TextCheckingTypes, options IDictionary, orthography IOrthography, wordCount int) []TextCheckingResult {
	if d._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount != nil {
		return d._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender, stringToCheck, offset, checkingTypes, options, orthography, wordCount)
	}
	var zero []TextCheckingResult
	return zero
}

// HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount returns true if a handler for SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount has been set.
func (d *SpellServerDelegate) HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool {
	return d._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount != nil
}

// SpellServerCheckGrammarInStringLanguageDetails implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerCheckGrammarInStringLanguageDetails(sender ISpellServer, stringToCheck IString, language IString, details IDictionary) Range {
	if d._SpellServerCheckGrammarInStringLanguageDetails != nil {
		return d._SpellServerCheckGrammarInStringLanguageDetails(sender, stringToCheck, language, details)
	}
	var zero Range
	return zero
}

// HasSpellServerCheckGrammarInStringLanguageDetails returns true if a handler for SpellServerCheckGrammarInStringLanguageDetails has been set.
func (d *SpellServerDelegate) HasSpellServerCheckGrammarInStringLanguageDetails() bool {
	return d._SpellServerCheckGrammarInStringLanguageDetails != nil
}

// SpellServerDidForgetWordInLanguage implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerDidForgetWordInLanguage(sender ISpellServer, word IString, language IString) {
	if d._SpellServerDidForgetWordInLanguage != nil {
		d._SpellServerDidForgetWordInLanguage(sender, word, language)
	}
}

// HasSpellServerDidForgetWordInLanguage returns true if a handler for SpellServerDidForgetWordInLanguage has been set.
func (d *SpellServerDelegate) HasSpellServerDidForgetWordInLanguage() bool {
	return d._SpellServerDidForgetWordInLanguage != nil
}

// SpellServerDidLearnWordInLanguage implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerDidLearnWordInLanguage(sender ISpellServer, word IString, language IString) {
	if d._SpellServerDidLearnWordInLanguage != nil {
		d._SpellServerDidLearnWordInLanguage(sender, word, language)
	}
}

// HasSpellServerDidLearnWordInLanguage returns true if a handler for SpellServerDidLearnWordInLanguage has been set.
func (d *SpellServerDelegate) HasSpellServerDidLearnWordInLanguage() bool {
	return d._SpellServerDidLearnWordInLanguage != nil
}

// SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender ISpellServer, stringToCheck IString, language IString, wordCount int, countOnly bool) Range {
	if d._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly != nil {
		return d._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender, stringToCheck, language, wordCount, countOnly)
	}
	var zero Range
	return zero
}

// HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly returns true if a handler for SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly has been set.
func (d *SpellServerDelegate) HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool {
	return d._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly != nil
}

// SpellServerRecordResponseToCorrectionForWordLanguage implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerRecordResponseToCorrectionForWordLanguage(sender ISpellServer, response uint, correction IString, word IString, language IString) {
	if d._SpellServerRecordResponseToCorrectionForWordLanguage != nil {
		d._SpellServerRecordResponseToCorrectionForWordLanguage(sender, response, correction, word, language)
	}
}

// HasSpellServerRecordResponseToCorrectionForWordLanguage returns true if a handler for SpellServerRecordResponseToCorrectionForWordLanguage has been set.
func (d *SpellServerDelegate) HasSpellServerRecordResponseToCorrectionForWordLanguage() bool {
	return d._SpellServerRecordResponseToCorrectionForWordLanguage != nil
}

// SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender ISpellServer, range_ Range, string_ IString, language IString) []string {
	if d._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage != nil {
		return d._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender, range_, string_, language)
	}
	var zero []string
	return zero
}

// HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage returns true if a handler for SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage has been set.
func (d *SpellServerDelegate) HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool {
	return d._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage != nil
}

// SpellServerSuggestGuessesForWordInLanguage implements the PSpellServerDelegate interface.
func (d *SpellServerDelegate) SpellServerSuggestGuessesForWordInLanguage(sender ISpellServer, word IString, language IString) []string {
	if d._SpellServerSuggestGuessesForWordInLanguage != nil {
		return d._SpellServerSuggestGuessesForWordInLanguage(sender, word, language)
	}
	var zero []string
	return zero
}

// HasSpellServerSuggestGuessesForWordInLanguage returns true if a handler for SpellServerSuggestGuessesForWordInLanguage has been set.
func (d *SpellServerDelegate) HasSpellServerSuggestGuessesForWordInLanguage() bool {
	return d._SpellServerSuggestGuessesForWordInLanguage != nil
}

// SpellServerDelegateObject wraps an existing Objective-C object that conforms to the PSpellServerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SpellServerDelegateObject struct {
	objectivec.Object
}

// NewSpellServerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSpellServerDelegate protocol.
func NewSpellServerDelegateObject(obj objectivec.Object) *SpellServerDelegateObject {
	return &SpellServerDelegateObject{obj}
}

// Make sure SpellServerDelegateObject implements PSpellServerDelegate.
var _ PSpellServerDelegate = (*SpellServerDelegateObject)(nil)

// SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender ISpellServer, stringToCheck IString, offset uint, checkingTypes TextCheckingTypes, options IDictionary, orthography IOrthography, wordCount int) []TextCheckingResult {
	return objc.Send[[]TextCheckingResult](o.ID, objc.Sel("spellServer:checkString:offset:types:options:orthography:wordCount:"), sender, stringToCheck, offset, checkingTypes, options, orthography, wordCount)
}

// HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerCheckGrammarInStringLanguageDetails implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerCheckGrammarInStringLanguageDetails(sender ISpellServer, stringToCheck IString, language IString, details IDictionary) Range {
	return objc.Send[Range](o.ID, objc.Sel("spellServer:checkGrammarInString:language:details:"), sender, stringToCheck, language, details)
}

// HasSpellServerCheckGrammarInStringLanguageDetails returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerCheckGrammarInStringLanguageDetails() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerDidForgetWordInLanguage implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerDidForgetWordInLanguage(sender ISpellServer, word IString, language IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("spellServer:didForgetWord:inLanguage:"), sender, word, language)
}

// HasSpellServerDidForgetWordInLanguage returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerDidForgetWordInLanguage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerDidLearnWordInLanguage implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerDidLearnWordInLanguage(sender ISpellServer, word IString, language IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("spellServer:didLearnWord:inLanguage:"), sender, word, language)
}

// HasSpellServerDidLearnWordInLanguage returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerDidLearnWordInLanguage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender ISpellServer, stringToCheck IString, language IString, wordCount int, countOnly bool) Range {
	return objc.Send[Range](o.ID, objc.Sel("spellServer:findMisspelledWordInString:language:wordCount:countOnly:"), sender, stringToCheck, language, wordCount, countOnly)
}

// HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerRecordResponseToCorrectionForWordLanguage implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerRecordResponseToCorrectionForWordLanguage(sender ISpellServer, response uint, correction IString, word IString, language IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("spellServer:recordResponse:toCorrection:forWord:language:"), sender, response, correction, word, language)
}

// HasSpellServerRecordResponseToCorrectionForWordLanguage returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerRecordResponseToCorrectionForWordLanguage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender ISpellServer, range_ Range, string_ IString, language IString) []string {
	return objc.Send[[]string](o.ID, objc.Sel("spellServer:suggestCompletionsForPartialWordRange:inString:language:"), sender, range_, string_, language)
}

// HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpellServerSuggestGuessesForWordInLanguage implements the PSpellServerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpellServerDelegateObject) SpellServerSuggestGuessesForWordInLanguage(sender ISpellServer, word IString, language IString) []string {
	return objc.Send[[]string](o.ID, objc.Sel("spellServer:suggestGuessesForWord:inLanguage:"), sender, word, language)
}

// HasSpellServerSuggestGuessesForWordInLanguage returns true; this is a placeholder for optional method checks.
func (o *SpellServerDelegateObject) HasSpellServerSuggestGuessesForWordInLanguage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
