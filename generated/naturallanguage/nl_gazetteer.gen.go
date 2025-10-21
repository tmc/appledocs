// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Gazetteer] class.
var (
	GazetteerClass     _GazetteerClass
	GazetteerClassOnce sync.Once
)

func getGazetteerClass() _GazetteerClass {
	GazetteerClassOnce.Do(func() {
		GazetteerClass = _GazetteerClass{objc.GetClass("NLGazetteer")}
	})
	return GazetteerClass
}

type _GazetteerClass struct {
	class objc.Class
}

// An interface definition for the [Gazetteer] class.
type IGazetteer interface {
	objectivec.IObject
	LabelForString(string_ string) string
}

// A collection of terms and their labels, which take precedence over a word tagger.
//
// Use an to augment an when you need to tag a specific set of terms (single words or short phrases) with a label. Typically, you add one gazetteer per language, or one language-independent gazetteer, to an with its method. The tagger uses its gazetteers to look up each term it processes. If a gazetteer has a label for a term, the tagger uses that label to tag the term, instead of inferring a tag itself. Typically, you create a gazetteer at development time, such as in a macOS playground, with Create ML’s . Alternatively, you can create an at runtime by using .
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer
type Gazetteer struct {
	objectivec.Object
}

// GazetteerFrom constructs a [Gazetteer] from an unsafe.Pointer.
//
// A collection of terms and their labels, which take precedence over a word tagger.
func GazetteerFrom(ptr unsafe.Pointer) Gazetteer {
	return Gazetteer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GazetteerClass) Alloc() Gazetteer {
	rv := objc.Send[Gazetteer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GazetteerClass) New() Gazetteer {
	rv := objc.Send[Gazetteer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Gazetteer) Init() Gazetteer {
	rv := objc.Send[Gazetteer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Gazetteer) Autorelease() Gazetteer {
	rv := objc.Send[Gazetteer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGazetteer creates a new Gazetteer instance.
func NewGazetteer() Gazetteer {
	return getGazetteerClass().New()
}


// Creates a Natural Language gazetteer from a model created with the Create ML framework.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(contentsOf:)
func NewGazetteerWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}

// Creates a gazetteer from a data instance.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(data:)
func NewGazetteerWithDataError(data unsafe.Pointer, error_ unsafe.Pointer) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}

// Creates a gazetteer from a set of labels for terms represented by a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(dictionary:language:)
func NewGazetteerWithDictionaryLanguageError(dictionary unsafe.Pointer, language unsafe.Pointer, error_ unsafe.Pointer) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithDictionary:language:error:"), dictionary, language, error_)
	rv.Autorelease()
	return rv
}


// Creates a Natural Language gazetteer from a model created with the Create ML framework.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/gazetteerWithContentsOfURL:error:
func (gc _GazetteerClass) GazetteerWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("gazetteerWithContentsOfURL:error:"), url, error_)
	return rv
}

// Creates a gazetteer from a set of labels for terms represented by a dictionary and saves the gazetteer to a file.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/write(_:language:to:)
func (gc _GazetteerClass) WriteGazetteerForDictionaryLanguageToURLError(dictionary unsafe.Pointer, language unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("writeGazetteerForDictionary:language:toURL:error:"), dictionary, language, url, error_)
	return rv
}

// Retrieves the label for the given term.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/label(for:)
func (g_ Gazetteer) LabelForString(string_ string) string {
	rv := objc.Send[string](g_.ID, objc.Sel("labelForString:"), objc.String(string_))
	return rv
}

// The gazetteer represented as a data instance.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/data
func (g_ Gazetteer) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("data"))
	return rv
}

// The language of the gazetteer.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/language
func (g_ Gazetteer) Language() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("language"))
	return rv
}


