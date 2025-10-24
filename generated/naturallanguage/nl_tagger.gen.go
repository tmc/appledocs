// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Tagger] class.
type ITagger interface {
	objectivec.IObject
	// properties:
	DominantLanguage() objc.IObject /* cross-framework: Language */
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
	TagSchemes() []string
	// methods:
	EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, options TaggerOptions, block bool)
	GazetteersForTagScheme(tagScheme objc.IObject /* cross-framework: TagScheme */) []IGazetteer
	ModelsForTagScheme(tagScheme objc.IObject /* cross-framework: TagScheme */) []IModel
	SetGazetteersForTagScheme(gazetteers []IGazetteer, tagScheme objc.IObject /* cross-framework: TagScheme */)
	SetLanguageRange(language objc.IObject /* cross-framework: Language */, range_ objc.IObject /* cross-framework: Range */)
	SetModelsForTagScheme(models []IModel, tagScheme objc.IObject /* cross-framework: TagScheme */)
	SetOrthographyRange(orthography objc.IObject /* cross-framework: Orthography */, range_ objc.IObject /* cross-framework: Range */)
	TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, tokenRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Tag */
	TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, maximumCount uint, tokenRange RangePointer /* not a class type */) foundation.IDictionary
	TagsInRangeUnitSchemeOptionsTokenRanges(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, options TaggerOptions, tokenRanges []objc.IObject /* cross-framework: Value */) []string
	TokenRangeAtIndexUnit(characterIndex uint, unit TokenUnit) objc.IObject /* cross-framework: Range */
	TokenRangeForRangeUnit(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit) objc.IObject /* cross-framework: Range */
}

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

// Alloc allocates a new instance without initialization.
func (tc _TaggerClass) Alloc() Tagger {
	rv := objc.Send[Tagger](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a linguistic tagger instance using the specified tag schemes and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/init(tagSchemes:)
func NewTaggerWithTagSchemes(tagSchemes []string) Tagger {
	instance := getTaggerClass().Alloc()
	rv := objc.Send[Tagger](instance.ID, objc.Sel("initWithTagSchemes:"), tagSchemes)
	rv.Autorelease()
	return rv
}



// Retrieves the tag schemes available for a particular unit (like word or sentence) and language on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/availableTagSchemes(for:language:)
func (tc _TaggerClass) AvailableTagSchemesForUnitLanguage(unit TokenUnit, language objc.IObject /* cross-framework: Language */) []string {
	rv := objc.Send[[]string](objc.ID(tc.class), objc.Sel("availableTagSchemesForUnit:language:"), unit, language)
	return rv
}


// Asks the Natural Language framework to load any missing assets for a tag scheme onto the device for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/requestAssets(for:tagScheme:completionHandler:)
func (tc _TaggerClass) RequestAssetsForLanguageTagSchemeCompletionHandler(language objc.IObject /* cross-framework: Language */, tagScheme objc.IObject /* cross-framework: TagScheme */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("requestAssetsForLanguage:tagScheme:completionHandler:"), language, tagScheme, completionHandler)
}


// Enumerates a block over the tagger’s string, given a range, token unit, and tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/enumerateTagsInRange:unit:scheme:options:usingBlock:
func (t_ Tagger) EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, options TaggerOptions, block bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTagsInRange:unit:scheme:options:usingBlock:"), range_, unit, scheme, options, block)
}


// Retrieves the gazetteers attached to a tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/gazetteers(for:)
func (t_ Tagger) GazetteersForTagScheme(tagScheme objc.IObject /* cross-framework: TagScheme */) []IGazetteer {
	rv := objc.Send[[]Gazetteer](t_.ID, objc.Sel("gazetteersForTagScheme:"), tagScheme)
	return rv
}


// Returns the models that apply to the given tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/models(forTagScheme:)
func (t_ Tagger) ModelsForTagScheme(tagScheme objc.IObject /* cross-framework: TagScheme */) []IModel {
	rv := objc.Send[[]Model](t_.ID, objc.Sel("modelsForTagScheme:"), tagScheme)
	return rv
}


// Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setGazetteers(_:for:)
func (t_ Tagger) SetGazetteersForTagScheme(gazetteers []IGazetteer, tagScheme objc.IObject /* cross-framework: TagScheme */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGazetteers:forTagScheme:"), gazetteers, tagScheme)
}


// Sets the language for a range of text within the tagger’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setLanguage:range:
func (t_ Tagger) SetLanguageRange(language objc.IObject /* cross-framework: Language */, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLanguage:range:"), language, range_)
}


// Assigns models for a tag scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setModels(_:forTagScheme:)
func (t_ Tagger) SetModelsForTagScheme(models []IModel, tagScheme objc.IObject /* cross-framework: TagScheme */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setModels:forTagScheme:"), models, tagScheme)
}


// Sets the orthography for the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setOrthography:range:
func (t_ Tagger) SetOrthographyRange(orthography objc.IObject /* cross-framework: Orthography */, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOrthography:range:"), orthography, range_)
}


// Finds a tag for a given linguistic unit, for a single scheme, at the specified character position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagAtIndex:unit:scheme:tokenRange:
func (t_ Tagger) TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, tokenRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Tag */ {
	rv := objc.Send[coremedia.Tag](t_.ID, objc.Sel("tagAtIndex:unit:scheme:tokenRange:"), characterIndex, unit, scheme, tokenRange)
	return rv
}


// Finds multiple possible tags for a given linguistic unit, for a single scheme, at the specified character position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:
func (t_ Tagger) TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, maximumCount uint, tokenRange RangePointer /* not a class type */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:"), characterIndex, unit, scheme, maximumCount, tokenRange)
	return rv
}


// Finds an array of linguistic tags and token ranges for a given string range and linguistic unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagsInRange:unit:scheme:options:tokenRanges:
func (t_ Tagger) TagsInRangeUnitSchemeOptionsTokenRanges(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit, scheme objc.IObject /* cross-framework: TagScheme */, options TaggerOptions, tokenRanges []objc.IObject /* cross-framework: Value */) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("tagsInRange:unit:scheme:options:tokenRanges:"), range_, unit, scheme, options, tokenRanges)
	return rv
}


// Returns the range of the linguistic unit containing the specified character index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeAtIndex:unit:
func (t_ Tagger) TokenRangeAtIndexUnit(characterIndex uint, unit TokenUnit) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeAtIndex:unit:"), characterIndex, unit)
	return rv
}


// Finds the entire range of all tokens of the specified linguistic unit contained completely or partially within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeForRange:unit:
func (t_ Tagger) TokenRangeForRangeUnit(range_ objc.IObject /* cross-framework: Range */, unit TokenUnit) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeForRange:unit:"), range_, unit)
	return rv
}


// The dominant language of the string set for the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/dominantLanguage
func (t_ Tagger) DominantLanguage() objc.IObject /* cross-framework: Language */ {
	rv := objc.Send[Language](t_.ID, objc.Sel("dominantLanguage"))
	return rv
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/string
func (t_ Tagger) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/string
func (t_ Tagger) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}


// The tag schemes configured for this linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagSchemes
func (t_ Tagger) TagSchemes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("tagSchemes"))
	return rv
}


