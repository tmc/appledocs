// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [NowPlayingInfoLanguageOption] class.
type INowPlayingInfoLanguageOption interface {
	objectivec.IObject
	IsAutomaticAudibleLanguageOption() bool
	IsAutomaticLegibleLanguageOption() bool
}

// A set of interfaces for setting the language option for the Now Playing item.
//
// The and classes provide interfaces for setting information about language options, for example, audio and subtitles, in the Now Playing information area.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoLanguageOptionClass) Alloc() NowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a single language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/init(type:languageTag:characteristics:displayName:identifier:)
func NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(languageOptionType unsafe.Pointer, languageTag string, languageOptionCharacteristics unsafe.Pointer, displayName string, identifier string) NowPlayingInfoLanguageOption {
	instance := getNowPlayingInfoLanguageOptionClass().Alloc()
	rv := objc.Send[NowPlayingInfoLanguageOption](instance.ID, objc.Sel("initWithType:languageTag:characteristics:displayName:identifier:"), languageOptionType, objc.String(languageTag), languageOptionCharacteristics, objc.String(displayName), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// Returns a Boolean value that determines whether to use the best audible language option based on the system preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/isAutomaticAudibleLanguageOption()
func (n_ NowPlayingInfoLanguageOption) IsAutomaticAudibleLanguageOption() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isAutomaticAudibleLanguageOption"))
	return rv
}

// Returns a Boolean value that determines whether to use the best legible language option based on the system preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/isAutomaticLegibleLanguageOption()
func (n_ NowPlayingInfoLanguageOption) IsAutomaticLegibleLanguageOption() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isAutomaticLegibleLanguageOption"))
	return rv
}

// The display name for a language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/displayName
func (n_ NowPlayingInfoLanguageOption) DisplayName() string {
	rv := objc.Send[string](n_.ID, objc.Sel("displayName"))
	return rv
}

// The unique identifier for the language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/identifier
func (n_ NowPlayingInfoLanguageOption) Identifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("identifier"))
	return rv
}

// The characteristics that describe the content of the language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageOptionCharacteristics
func (n_ NowPlayingInfoLanguageOption) LanguageOptionCharacteristics() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("languageOptionCharacteristics"))
	return rv
}

// The type of language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageOptionType
func (n_ NowPlayingInfoLanguageOption) LanguageOptionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("languageOptionType"))
	return rv
}

// The abbreviated language code for the language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOption/languageTag
func (n_ NowPlayingInfoLanguageOption) LanguageTag() string {
	rv := objc.Send[string](n_.ID, objc.Sel("languageTag"))
	return rv
}


