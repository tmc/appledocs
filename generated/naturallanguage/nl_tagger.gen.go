// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLTagger */


/* debug [class_header]: Header for NLTagger */
// The class instance for the [Tagger] class.
var (
	TaggerClass     _TaggerClass
	TaggerClassOnce sync.Once
)

func getTaggerClass() _TaggerClass {
	TaggerClassOnce.Do(func() {
		TaggerClass = _TaggerClass{objc.GetClass("NLTagger")}
	})
	return TaggerClass
}

type _TaggerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Tagger */
// An interface definition for the [Tagger] class.
type ITagger interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Tagger */
	// properties:
	DominantLanguage() Language /* typedef */
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
	TagSchemes() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Tagger */
	// methods:
	EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ corefoundation.Range, unit TokenUnit, scheme TagScheme /* typedef */, options TaggerOptions, block bool)
	GazetteersForTagScheme(tagScheme TagScheme /* typedef */) []Gazetteer
	ModelsForTagScheme(tagScheme TagScheme /* typedef */) []Model
	SetGazetteersForTagScheme(gazetteers []Gazetteer, tagScheme TagScheme /* typedef */)
	SetLanguageRange(language Language /* typedef */, range_ corefoundation.Range)
	SetModelsForTagScheme(models []Model, tagScheme TagScheme /* typedef */)
	SetOrthographyRange(orthography foundation.Orthography, range_ corefoundation.Range)
	TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit TokenUnit, scheme TagScheme /* typedef */, tokenRange RangePointer /* not a class type */) Tag /* typedef */
	TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit TokenUnit, scheme TagScheme /* typedef */, maximumCount uint, tokenRange RangePointer /* not a class type */) foundation.IDictionary
	TagsInRangeUnitSchemeOptionsTokenRanges(range_ corefoundation.Range, unit TokenUnit, scheme TagScheme /* typedef */, options TaggerOptions, tokenRanges []foundation.Value) []string
	TokenRangeAtIndexUnit(characterIndex uint, unit TokenUnit) corefoundation.Range
	TokenRangeForRangeUnit(range_ corefoundation.Range, unit TokenUnit) corefoundation.Range
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Tagger */
// Alloc allocates a new instance without initialization.
func (tc _TaggerClass) Alloc() Tagger {
	rv := objc.Send[Tagger](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TaggerClass) New() Tagger {
	rv := objc.Send[Tagger](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Tagger) Init() Tagger {
	rv := objc.Send[Tagger](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Tagger) Autorelease() Tagger {
	rv := objc.Send[Tagger](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTagger creates a new Tagger instance.
func NewTagger() Tagger {
	return getTaggerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Tagger */
// A tagger that analyzes natural language text.
//
// supports many different languages and scripts. Use it to segment natural language text into paragraph, sentence, or word units and to tag each unit with information like part of speech, lexical class, lemma, script, and language. When you create a linguistic tagger, you specify what kind of information you’re interested in by passing one or more values. Set the property to the natural language text you want to analyze, and the linguistic tagger processes it according to the specified tag schemes. You can then enumerate over the tags in a specified range, using the methods described in Enumerating linguistic tags, to get the information requested for a given scheme and unit.


// A tagger that analyzes natural language text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger
type Tagger struct {
	objectivec.Object
}

// TaggerFrom constructs a [Tagger] from an unsafe.Pointer.
//
// A tagger that analyzes natural language text.
func TaggerFrom(ptr unsafe.Pointer) Tagger {
	return Tagger{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Tagger */

// Creates a linguistic tagger instance using the specified tag schemes and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/init(tagSchemes:)
func NewTaggerWithTagSchemes(tagSchemes []string) Tagger {
	instance := getTaggerClass().Alloc()
	rv := objc.Send[Tagger](instance.ID, objc.Sel("initWithTagSchemes:"), tagSchemes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTaggerWithTagSchemes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Tagger */

// Retrieves the tag schemes available for a particular unit (like word or sentence) and language on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/availableTagSchemes(for:language:)
func (tc _TaggerClass) AvailableTagSchemesForUnitLanguage(unit TokenUnit, language Language /* typedef */) []string {
	rv := objc.Send[[]string](objc.ID(tc.class), objc.Sel("availableTagSchemesForUnit:language:"), unit, language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AvailableTagSchemesForUnitLanguage) */


// Asks the Natural Language framework to load any missing assets for a tag scheme onto the device for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/requestAssets(for:tagScheme:completionHandler:)
func (tc _TaggerClass) RequestAssetsForLanguageTagSchemeCompletionHandler(language Language /* typedef */, tagScheme TagScheme /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("requestAssetsForLanguage:tagScheme:completionHandler:"), language, tagScheme, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestAssetsForLanguageTagSchemeCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Tagger */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Tagger */

// Enumerates a block over the tagger’s string, given a range, token unit, and tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/enumerateTagsInRange:unit:scheme:options:usingBlock:
func (t_ Tagger) EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ corefoundation.Range, unit TokenUnit, scheme TagScheme /* typedef */, options TaggerOptions, block bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTagsInRange:unit:scheme:options:usingBlock:"), range_, unit, scheme, options, block)
}/* debug [instance_methods/method]: EnumerateTagsInRangeUnitSchemeOptionsUsingBlock */


// Retrieves the gazetteers attached to a tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/gazetteers(for:)
func (t_ Tagger) GazetteersForTagScheme(tagScheme TagScheme /* typedef */) []Gazetteer {
	rv := objc.Send[[]Gazetteer](t_.ID, objc.Sel("gazetteersForTagScheme:"), tagScheme)
	return rv
}/* debug [instance_methods/method]: GazetteersForTagScheme */


// Returns the models that apply to the given tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/models(forTagScheme:)
func (t_ Tagger) ModelsForTagScheme(tagScheme TagScheme /* typedef */) []Model {
	rv := objc.Send[[]Model](t_.ID, objc.Sel("modelsForTagScheme:"), tagScheme)
	return rv
}/* debug [instance_methods/method]: ModelsForTagScheme */


// Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setGazetteers(_:for:)
func (t_ Tagger) SetGazetteersForTagScheme(gazetteers []Gazetteer, tagScheme TagScheme /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGazetteers:forTagScheme:"), gazetteers, tagScheme)
}/* debug [instance_methods/method]: SetGazetteersForTagScheme */


// Sets the language for a range of text within the tagger’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setLanguage:range:
func (t_ Tagger) SetLanguageRange(language Language /* typedef */, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLanguage:range:"), language, range_)
}/* debug [instance_methods/method]: SetLanguageRange */


// Assigns models for a tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setModels(_:forTagScheme:)
func (t_ Tagger) SetModelsForTagScheme(models []Model, tagScheme TagScheme /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setModels:forTagScheme:"), models, tagScheme)
}/* debug [instance_methods/method]: SetModelsForTagScheme */


// Sets the orthography for the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setOrthography:range:
func (t_ Tagger) SetOrthographyRange(orthography foundation.Orthography, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOrthography:range:"), orthography, range_)
}/* debug [instance_methods/method]: SetOrthographyRange */


// Finds a tag for a given linguistic unit, for a single scheme, at the specified character position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagAtIndex:unit:scheme:tokenRange:
func (t_ Tagger) TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit TokenUnit, scheme TagScheme /* typedef */, tokenRange RangePointer /* not a class type */) Tag /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("tagAtIndex:unit:scheme:tokenRange:"), characterIndex, unit, scheme, tokenRange)
	return rv
}/* debug [instance_methods/method]: TagAtIndexUnitSchemeTokenRange */


// Finds multiple possible tags for a given linguistic unit, for a single scheme, at the specified character position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:
func (t_ Tagger) TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit TokenUnit, scheme TagScheme /* typedef */, maximumCount uint, tokenRange RangePointer /* not a class type */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:"), characterIndex, unit, scheme, maximumCount, tokenRange)
	return rv
}/* debug [instance_methods/method]: TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange */


// Finds an array of linguistic tags and token ranges for a given string range and linguistic unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagsInRange:unit:scheme:options:tokenRanges:
func (t_ Tagger) TagsInRangeUnitSchemeOptionsTokenRanges(range_ corefoundation.Range, unit TokenUnit, scheme TagScheme /* typedef */, options TaggerOptions, tokenRanges []foundation.Value) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("tagsInRange:unit:scheme:options:tokenRanges:"), range_, unit, scheme, options, tokenRanges)
	return rv
}/* debug [instance_methods/method]: TagsInRangeUnitSchemeOptionsTokenRanges */


// Returns the range of the linguistic unit containing the specified character index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeAtIndex:unit:
func (t_ Tagger) TokenRangeAtIndexUnit(characterIndex uint, unit TokenUnit) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeAtIndex:unit:"), characterIndex, unit)
	return rv
}/* debug [instance_methods/method]: TokenRangeAtIndexUnit */


// Finds the entire range of all tokens of the specified linguistic unit contained completely or partially within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeForRange:unit:
func (t_ Tagger) TokenRangeForRangeUnit(range_ corefoundation.Range, unit TokenUnit) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeForRange:unit:"), range_, unit)
	return rv
}/* debug [instance_methods/method]: TokenRangeForRangeUnit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Tagger */

// The dominant language of the string set for the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/dominantLanguage
func (t_ Tagger) DominantLanguage() Language /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("dominantLanguage"))
	return rv
}/* debug [instance_properties/getter]: dominantLanguage */


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/string
func (t_ Tagger) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/string
func (t_ Tagger) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */


// The tag schemes configured for this linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagSchemes
func (t_ Tagger) TagSchemes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("tagSchemes"))
	return rv
}/* debug [instance_properties/getter]: tagSchemes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLTagger */


