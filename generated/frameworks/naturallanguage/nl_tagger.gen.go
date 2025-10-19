// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLTagger] class.
var nLTaggerClass = _NLTaggerClass{objc.GetClass("NLTagger")}

type _NLTaggerClass struct {
	class objc.Class
}

// An interface definition for the [NLTagger] class.
type INLTagger interface {
	objectivec.IObject
	EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ unsafe.Pointer, unit unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer)
	GazetteersForTagScheme(tagScheme unsafe.Pointer) unsafe.Pointer
	ModelsForTagScheme(tagScheme unsafe.Pointer) unsafe.Pointer
	SetGazetteersForTagScheme(gazetteers unsafe.Pointer, tagScheme unsafe.Pointer)
	SetLanguageRange(language unsafe.Pointer, range_ unsafe.Pointer)
	SetModelsForTagScheme(models unsafe.Pointer, tagScheme unsafe.Pointer)
	SetOrthographyRange(orthography unsafe.Pointer, range_ unsafe.Pointer)
	TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit unsafe.Pointer, scheme unsafe.Pointer, tokenRange unsafe.Pointer) unsafe.Pointer
	TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit unsafe.Pointer, scheme unsafe.Pointer, maximumCount uint, tokenRange unsafe.Pointer) unsafe.Pointer
	TagsInRangeUnitSchemeOptionsTokenRanges(range_ unsafe.Pointer, unit unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, tokenRanges unsafe.Pointer) unsafe.Pointer
	TokenRangeAtIndexUnit(characterIndex uint, unit unsafe.Pointer) unsafe.Pointer
	TokenRangeForRangeUnit(range_ unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer
}

// A tagger that analyzes natural language text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger

type NLTagger struct {
	objectivec.Object
}

// NLTaggerFrom constructs a [NLTagger] from an unsafe.Pointer.
//
// A tagger that analyzes natural language text.
func NLTaggerFrom(ptr unsafe.Pointer) NLTagger {
	return NLTagger{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLTaggerClass) Alloc() NLTagger {
	rv := objc.Send[NLTagger](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLTaggerClass) New() NLTagger {
	rv := objc.Send[NLTagger](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLTagger) Init() NLTagger {
	rv := objc.Send[NLTagger](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLTagger) Autorelease() NLTagger {
	rv := objc.Send[NLTagger](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLTagger creates a new NLTagger instance.
func NewNLTagger() NLTagger {
	return nLTaggerClass.New()
}


// Creates a linguistic tagger instance using the specified tag schemes and options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/init(tagSchemes:)
func NewNLTaggerWithTagSchemes(tagSchemes unsafe.Pointer) NLTagger {
	instance := nLTaggerClass.Alloc()
	rv := objc.Send[NLTagger](instance.ID, objc.Sel("initWithTagSchemes:"), tagSchemes)
	rv.Autorelease()
	return rv
}


// Retrieves the tag schemes available for a particular unit (like word or sentence) and language on the current device. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/availableTagSchemes(for:language:)
func (nc _NLTaggerClass) AvailableTagSchemesForUnitLanguage(unit unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("availableTagSchemesForUnit:language:"), unit, language)
	return rv
}
// Asks the Natural Language framework to load any missing assets for a tag scheme onto the device for the given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/requestAssets(for:tagScheme:completionHandler:)
func (nc _NLTaggerClass) RequestAssetsForLanguageTagSchemeCompletionHandler(language unsafe.Pointer, tagScheme unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("requestAssetsForLanguage:tagScheme:completionHandler:"), language, tagScheme, completionHandler)
}
// Enumerates a block over the tagger’s string, given a range, token unit, and tag scheme. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/enumerateTagsInRange:unit:scheme:options:usingBlock:
func (n_ NLTagger) EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ unsafe.Pointer, unit unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateTagsInRange:unit:scheme:options:usingBlock:"), range_, unit, scheme, options, block)
}
// Retrieves the gazetteers attached to a tag scheme. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/gazetteers(for:)
func (n_ NLTagger) GazetteersForTagScheme(tagScheme unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("gazetteersForTagScheme:"), tagScheme)
	return rv
}
// Returns the models that apply to the given tag scheme. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/models(forTagScheme:)
func (n_ NLTagger) ModelsForTagScheme(tagScheme unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("modelsForTagScheme:"), tagScheme)
	return rv
}
// Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setGazetteers(_:for:)
func (n_ NLTagger) SetGazetteersForTagScheme(gazetteers unsafe.Pointer, tagScheme unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGazetteers:forTagScheme:"), gazetteers, tagScheme)
}
// Sets the language for a range of text within the tagger’s string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setLanguage:range:
func (n_ NLTagger) SetLanguageRange(language unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLanguage:range:"), language, range_)
}
// Assigns models for a tag scheme. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setModels(_:forTagScheme:)
func (n_ NLTagger) SetModelsForTagScheme(models unsafe.Pointer, tagScheme unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setModels:forTagScheme:"), models, tagScheme)
}
// Sets the orthography for the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setOrthography:range:
func (n_ NLTagger) SetOrthographyRange(orthography unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOrthography:range:"), orthography, range_)
}
// Finds a tag for a given linguistic unit, for a single scheme, at the specified character position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagAtIndex:unit:scheme:tokenRange:
func (n_ NLTagger) TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit unsafe.Pointer, scheme unsafe.Pointer, tokenRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tagAtIndex:unit:scheme:tokenRange:"), characterIndex, unit, scheme, tokenRange)
	return rv
}
// Finds multiple possible tags for a given linguistic unit, for a single scheme, at the specified character position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:
func (n_ NLTagger) TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit unsafe.Pointer, scheme unsafe.Pointer, maximumCount uint, tokenRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:"), characterIndex, unit, scheme, maximumCount, tokenRange)
	return rv
}
// Finds an array of linguistic tags and token ranges for a given string range and linguistic unit. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagsInRange:unit:scheme:options:tokenRanges:
func (n_ NLTagger) TagsInRangeUnitSchemeOptionsTokenRanges(range_ unsafe.Pointer, unit unsafe.Pointer, scheme unsafe.Pointer, options unsafe.Pointer, tokenRanges unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tagsInRange:unit:scheme:options:tokenRanges:"), range_, unit, scheme, options, tokenRanges)
	return rv
}
// Returns the range of the linguistic unit containing the specified character index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeAtIndex:unit:
func (n_ NLTagger) TokenRangeAtIndexUnit(characterIndex uint, unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tokenRangeAtIndex:unit:"), characterIndex, unit)
	return rv
}
// Finds the entire range of all tokens of the specified linguistic unit contained completely or partially within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeForRange:unit:
func (n_ NLTagger) TokenRangeForRangeUnit(range_ unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tokenRangeForRange:unit:"), range_, unit)
	return rv
}

