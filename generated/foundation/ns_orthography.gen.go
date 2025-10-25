// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOrthography */


/* debug [class_header]: Header for NSOrthography */
// The class instance for the [Orthography] class.
var (
	OrthographyClass     _OrthographyClass
	OrthographyClassOnce sync.Once
)

func getOrthographyClass() _OrthographyClass {
	OrthographyClassOnce.Do(func() {
		OrthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}
	})
	return OrthographyClass
}

type _OrthographyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Orthography */
// An interface definition for the [Orthography] class.
type IOrthography interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Orthography */
	// properties:
	AllLanguages() []string
	AllScripts() []string
	DominantLanguage() IString
	DominantScript() IString
	LanguageMap() IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Orthography */
	// methods:
	DominantLanguageForScript(script IString) IString
	LanguagesForScript(script IString) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Orthography */
// Alloc allocates a new instance without initialization.
func (oc _OrthographyClass) Alloc() Orthography {
	rv := objc.Send[Orthography](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OrthographyClass) New() Orthography {
	rv := objc.Send[Orthography](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Orthography) Init() Orthography {
	rv := objc.Send[Orthography](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Orthography) Autorelease() Orthography {
	rv := objc.Send[Orthography](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrthography creates a new Orthography instance.
func NewOrthography() Orthography {
	return getOrthographyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Orthography */
// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
//
// Use objects to describe the linguistic content of a piece of text, including which scripts the text contains, a dominant language (and possibly other languages) for each script, and a dominant script and language for the text as a whole. Scripts are uniformly described by four-letter ISO 15924 script codes, such as , , and . The supertags and are typically used for Japanese and Korean text, and and are typically used for Chinese text. The tag is used if a specific script cannot be identified. See for more information. Languages are uniformly described by BCP-47 tags (preferably in canonical form). The tag is used if a specific language cannot be determined. You typically work with orthography objects returned from methods and properties for classes like and .


// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography
type Orthography struct {
	objectivec.Object
}

// OrthographyFrom constructs a [Orthography] from an unsafe.Pointer.
//
// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Orthography */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/init(coder:)
func NewOrthographyWithCoder(coder ICoder) Orthography {
	instance := getOrthographyClass().Alloc()
	rv := objc.Send[Orthography](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOrthographyWithCoder */


// Creates an orthography object with the specified dominant script and language map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/init(dominantScript:languageMap:)
func NewOrthographyWithDominantScriptLanguageMap(script IString, map_ IDictionary) Orthography {
	instance := getOrthographyClass().Alloc()
	rv := objc.Send[Orthography](instance.ID, objc.Sel("initWithDominantScript:languageMap:"), script, map_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOrthographyWithDominantScriptLanguageMap */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Orthography */

// Creates and returns an orthography object with the default language map for the specified language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/defaultOrthography(forLanguage:)
func (oc _OrthographyClass) DefaultOrthographyForLanguage(language IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("defaultOrthographyForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultOrthographyForLanguage) */


// Creates and returns an orthography object with the specified dominant script and language map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/orthographyWithDominantScript:languageMap:
func (oc _OrthographyClass) OrthographyWithDominantScriptLanguageMap(script IString, map_ IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orthographyWithDominantScript:languageMap:"), script, map_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OrthographyWithDominantScriptLanguageMap) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Orthography */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Orthography */

// Returns the dominant language for the specified script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage(forScript:)
func (o_ Orthography) DominantLanguageForScript(script IString) IString {
	rv := objc.Send[String](o_.ID, objc.Sel("dominantLanguageForScript:"), script)
	return rv
}/* debug [instance_methods/method]: DominantLanguageForScript */


// Returns the list of languages for the specified script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/languages(forScript:)
func (o_ Orthography) LanguagesForScript(script IString) []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("languagesForScript:"), script)
	return rv
}/* debug [instance_methods/method]: LanguagesForScript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Orthography */

// The languages appearing in values of the language map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/allLanguages
func (o_ Orthography) AllLanguages() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("allLanguages"))
	return rv
}/* debug [instance_properties/getter]: allLanguages */


// The scripts appearing as keys in the language map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/allScripts
func (o_ Orthography) AllScripts() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("allScripts"))
	return rv
}/* debug [instance_properties/getter]: allScripts */


// The first language in the list of languages for the dominant script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage
func (o_ Orthography) DominantLanguage() IString {
	rv := objc.Send[String](o_.ID, objc.Sel("dominantLanguage"))
	return rv
}/* debug [instance_properties/getter]: dominantLanguage */


// The dominant script for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantScript
func (o_ Orthography) DominantScript() IString {
	rv := objc.Send[String](o_.ID, objc.Sel("dominantScript"))
	return rv
}/* debug [instance_properties/getter]: dominantScript */


// A dictionary that maps script tags to arrays of language tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/languageMap
func (o_ Orthography) LanguageMap() IDictionary {
	rv := objc.Send[Dictionary](o_.ID, objc.Sel("languageMap"))
	return rv
}/* debug [instance_properties/getter]: languageMap */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOrthography */


