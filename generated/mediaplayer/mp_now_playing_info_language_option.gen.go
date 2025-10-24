// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPNowPlayingInfoLanguageOption */


/* debug [class_header]: Header for MPNowPlayingInfoLanguageOption */
// The class instance for the [NowPlayingInfoLanguageOption] class.
var (
	NowPlayingInfoLanguageOptionClass     _NowPlayingInfoLanguageOptionClass
	NowPlayingInfoLanguageOptionClassOnce sync.Once
)

func getNowPlayingInfoLanguageOptionClass() _NowPlayingInfoLanguageOptionClass {
	NowPlayingInfoLanguageOptionClassOnce.Do(func() {
		NowPlayingInfoLanguageOptionClass = _NowPlayingInfoLanguageOptionClass{objc.GetClass("MPNowPlayingInfoLanguageOption")}
	})
	return NowPlayingInfoLanguageOptionClass
}

type _NowPlayingInfoLanguageOptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NowPlayingInfoLanguageOption */
// An interface definition for the [NowPlayingInfoLanguageOption] class.
type INowPlayingInfoLanguageOption interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NowPlayingInfoLanguageOption */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	LanguageOptionCharacteristics() []string
	LanguageOptionType() NowPlayingInfoLanguageOptionType
	LanguageTag() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NowPlayingInfoLanguageOption */
	// methods:
	IsAutomaticAudibleLanguageOption() bool
	IsAutomaticLegibleLanguageOption() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NowPlayingInfoLanguageOption */
// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoLanguageOptionClass) Alloc() NowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NowPlayingInfoLanguageOptionClass) New() NowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NowPlayingInfoLanguageOption) Init() NowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NowPlayingInfoLanguageOption) Autorelease() NowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNowPlayingInfoLanguageOption creates a new NowPlayingInfoLanguageOption instance.
func NewNowPlayingInfoLanguageOption() NowPlayingInfoLanguageOption {
	return getNowPlayingInfoLanguageOptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NowPlayingInfoLanguageOption */
// A set of interfaces for setting the language option for the Now Playing item.
//
// The and classes provide interfaces for setting information about language options, for example, audio and subtitles, in the Now Playing information area.


// A set of interfaces for setting the language option for the Now Playing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption
type NowPlayingInfoLanguageOption struct {
	objectivec.Object
}

// NowPlayingInfoLanguageOptionFrom constructs a [NowPlayingInfoLanguageOption] from an unsafe.Pointer.
//
// A set of interfaces for setting the language option for the Now Playing item.
func NowPlayingInfoLanguageOptionFrom(ptr unsafe.Pointer) NowPlayingInfoLanguageOption {
	return NowPlayingInfoLanguageOption{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NowPlayingInfoLanguageOption */

// Creates a single language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/init(type:languageTag:characteristics:displayName:identifier:)
func NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(languageOptionType NowPlayingInfoLanguageOptionType, languageTag objc.IObject /* cross-framework: NSString */, languageOptionCharacteristics []string, displayName objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */) NowPlayingInfoLanguageOption {
	instance := getNowPlayingInfoLanguageOptionClass().Alloc()
	rv := objc.Send[NowPlayingInfoLanguageOption](instance.ID, objc.Sel("initWithType:languageTag:characteristics:displayName:identifier:"), languageOptionType, languageTag, languageOptionCharacteristics, displayName, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NowPlayingInfoLanguageOption */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NowPlayingInfoLanguageOption */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NowPlayingInfoLanguageOption */

// Returns a Boolean value that determines whether to use the best audible language option based on the system preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/isAutomaticAudibleLanguageOption()
func (n_ NowPlayingInfoLanguageOption) IsAutomaticAudibleLanguageOption() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isAutomaticAudibleLanguageOption"))
	return rv
}/* debug [instance_methods/method]: IsAutomaticAudibleLanguageOption */


// Returns a Boolean value that determines whether to use the best legible language option based on the system preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/isAutomaticLegibleLanguageOption()
func (n_ NowPlayingInfoLanguageOption) IsAutomaticLegibleLanguageOption() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isAutomaticLegibleLanguageOption"))
	return rv
}/* debug [instance_methods/method]: IsAutomaticLegibleLanguageOption */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NowPlayingInfoLanguageOption */

// The display name for a language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/displayName
func (n_ NowPlayingInfoLanguageOption) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// The unique identifier for the language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/identifier
func (n_ NowPlayingInfoLanguageOption) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The characteristics that describe the content of the language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageOptionCharacteristics
func (n_ NowPlayingInfoLanguageOption) LanguageOptionCharacteristics() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("languageOptionCharacteristics"))
	return rv
}/* debug [instance_properties/getter]: languageOptionCharacteristics */


// The type of language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageOptionType
func (n_ NowPlayingInfoLanguageOption) LanguageOptionType() NowPlayingInfoLanguageOptionType {
	rv := objc.Send[NowPlayingInfoLanguageOptionType](n_.ID, objc.Sel("languageOptionType"))
	return rv
}/* debug [instance_properties/getter]: languageOptionType */


// The abbreviated language code for the language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageTag
func (n_ NowPlayingInfoLanguageOption) LanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("languageTag"))
	return rv
}/* debug [instance_properties/getter]: languageTag */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPNowPlayingInfoLanguageOption */


