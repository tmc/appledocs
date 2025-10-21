// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Orthography] class.
type IOrthography interface {
	objectivec.IObject
	DominantLanguageForScript(script string) string
	LanguagesForScript(script string) []string
}

// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
//
// Use objects to describe the linguistic content of a piece of text, including which scripts the text contains, a dominant language (and possibly other languages) for each script, and a dominant script and language for the text as a whole. Scripts are uniformly described by four-letter ISO 15924 script codes, such as , , and . The supertags and are typically used for Japanese and Korean text, and and are typically used for Chinese text. The tag is used if a specific script cannot be identified. See for more information. Languages are uniformly described by BCP-47 tags (preferably in canonical form). The tag is used if a specific language cannot be determined. You typically work with orthography objects returned from methods and properties for classes like and .
//
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

// Alloc allocates a new instance without initialization.
func (oc _OrthographyClass) Alloc() Orthography {
	rv := objc.Send[Orthography](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/init(coder:)
func NewOrthographyWithCoder(coder unsafe.Pointer) Orthography {
	instance := getOrthographyClass().Alloc()
	rv := objc.Send[Orthography](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Creates an orthography object with the specified dominant script and language map.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/init(dominantScript:languageMap:)
func NewOrthographyWithDominantScriptLanguageMap(script string, map_ unsafe.Pointer) Orthography {
	instance := getOrthographyClass().Alloc()
	rv := objc.Send[Orthography](instance.ID, objc.Sel("initWithDominantScript:languageMap:"), objc.String(script), map_)
	rv.Autorelease()
	return rv
}

// Creates and returns an orthography object with the default language map for the specified language.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/defaultOrthography(forLanguage:)
func (oc _OrthographyClass) DefaultOrthographyForLanguage(language string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("defaultOrthographyForLanguage:"), objc.String(language))
	return rv
}

// Creates and returns an orthography object with the specified dominant script and language map.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/orthographyWithDominantScript:languageMap:
func (oc _OrthographyClass) OrthographyWithDominantScriptLanguageMap(script string, map_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("orthographyWithDominantScript:languageMap:"), objc.String(script), map_)
	return rv
}

// Returns the dominant language for the specified script.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage(forScript:)
func (o_ Orthography) DominantLanguageForScript(script string) string {
	rv := objc.Send[string](o_.ID, objc.Sel("dominantLanguageForScript:"), objc.String(script))
	return rv
}

// Returns the list of languages for the specified script.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/languages(forScript:)
func (o_ Orthography) LanguagesForScript(script string) []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("languagesForScript:"), objc.String(script))
	return rv
}

// The languages appearing in values of the language map.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/allLanguages
func (o_ Orthography) AllLanguages() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("allLanguages"))
	return rv
}

// The scripts appearing as keys in the language map.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/allScripts
func (o_ Orthography) AllScripts() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("allScripts"))
	return rv
}

// The first language in the list of languages for the dominant script.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage
func (o_ Orthography) DominantLanguage() string {
	rv := objc.Send[string](o_.ID, objc.Sel("dominantLanguage"))
	return rv
}

// The dominant script for the text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantScript
func (o_ Orthography) DominantScript() string {
	rv := objc.Send[string](o_.ID, objc.Sel("dominantScript"))
	return rv
}

// A dictionary that maps script tags to arrays of language tags.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography/languageMap
func (o_ Orthography) LanguageMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("languageMap"))
	return rv
}
