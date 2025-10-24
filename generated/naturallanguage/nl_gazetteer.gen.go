// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLGazetteer */


/* debug [class_header]: Header for NLGazetteer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Gazetteer */
// An interface definition for the [Gazetteer] class.
type IGazetteer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Gazetteer */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Language() Language /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Gazetteer */
	// methods:
	LabelForString(string_ objc.IObject /* cross-framework: NSString */) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Gazetteer */
// Alloc allocates a new instance without initialization.
func (gc _GazetteerClass) Alloc() Gazetteer {
	rv := objc.Send[Gazetteer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Gazetteer */
// A collection of terms and their labels, which take precedence over a word tagger.
//
// Use an to augment an when you need to tag a specific set of terms (single words or short phrases) with a label. Typically, you add one gazetteer per language, or one language-independent gazetteer, to an with its method. The tagger uses its gazetteers to look up each term it processes. If a gazetteer has a label for a term, the tagger uses that label to tag the term, instead of inferring a tag itself. Typically, you create a gazetteer at development time, such as in a macOS playground, with Create ML’s . Alternatively, you can create an at runtime by using .


// A collection of terms and their labels, which take precedence over a word tagger.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Gazetteer */

// Creates a Natural Language gazetteer from a model created with the Create ML framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(contentsOf:)
func NewGazetteerWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGazetteerWithContentsOfURLError */


// Creates a gazetteer from a data instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(data:)
func NewGazetteerWithDataError(data objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGazetteerWithDataError */


// Creates a gazetteer from a set of labels for terms represented by a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(dictionary:language:)
func NewGazetteerWithDictionaryLanguageError(dictionary foundation.IDictionary, language Language /* typedef */, error_ objectivec.IObject) Gazetteer {
	instance := getGazetteerClass().Alloc()
	rv := objc.Send[Gazetteer](instance.ID, objc.Sel("initWithDictionary:language:error:"), dictionary, language, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGazetteerWithDictionaryLanguageError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Gazetteer */

// Creates a Natural Language gazetteer from a model created with the Create ML framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/gazetteerWithContentsOfURL:error:
func (gc _GazetteerClass) GazetteerWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("gazetteerWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GazetteerWithContentsOfURLError) */


// Creates a gazetteer from a set of labels for terms represented by a dictionary and saves the gazetteer to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/write(_:language:to:)
func (gc _GazetteerClass) WriteGazetteerForDictionaryLanguageToURLError(dictionary foundation.IDictionary, language Language /* typedef */, url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("writeGazetteerForDictionary:language:toURL:error:"), dictionary, language, url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WriteGazetteerForDictionaryLanguageToURLError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Gazetteer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Gazetteer */

// Retrieves the label for the given term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/label(for:)
func (g_ Gazetteer) LabelForString(string_ objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](g_.ID, objc.Sel("labelForString:"), string_)
	return rv
}/* debug [instance_methods/method]: LabelForString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Gazetteer */

// The gazetteer represented as a data instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/data
func (g_ Gazetteer) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](g_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The language of the gazetteer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/language
func (g_ Gazetteer) Language() Language /* typedef */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLGazetteer */


