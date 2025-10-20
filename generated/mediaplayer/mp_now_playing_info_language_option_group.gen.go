// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NowPlayingInfoLanguageOptionGroup] class.
var (
	NowPlayingInfoLanguageOptionGroupClass     _NowPlayingInfoLanguageOptionGroupClass
	NowPlayingInfoLanguageOptionGroupClassOnce sync.Once
)

func getNowPlayingInfoLanguageOptionGroupClass() _NowPlayingInfoLanguageOptionGroupClass {
	NowPlayingInfoLanguageOptionGroupClassOnce.Do(func() {
		NowPlayingInfoLanguageOptionGroupClass = _NowPlayingInfoLanguageOptionGroupClass{objc.GetClass("MPNowPlayingInfoLanguageOptionGroup")}
	})
	return NowPlayingInfoLanguageOptionGroupClass
}

type _NowPlayingInfoLanguageOptionGroupClass struct {
	class objc.Class
}

// An interface definition for the [NowPlayingInfoLanguageOptionGroup] class.
type INowPlayingInfoLanguageOptionGroup interface {
	objectivec.IObject
}

// A grouped set of language options where only a single language option can be active at a time.
//
// The and classes provide interfaces for setting information about language options, for example, audio and subtitles, in the Now Playing information area.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup
type NowPlayingInfoLanguageOptionGroup struct {
	objectivec.Object
}

// NowPlayingInfoLanguageOptionGroupFrom constructs a [NowPlayingInfoLanguageOptionGroup] from an unsafe.Pointer.
//
// A grouped set of language options where only a single language option can be active at a time.
func NowPlayingInfoLanguageOptionGroupFrom(ptr unsafe.Pointer) NowPlayingInfoLanguageOptionGroup {
	return NowPlayingInfoLanguageOptionGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoLanguageOptionGroupClass) Alloc() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NowPlayingInfoLanguageOptionGroupClass) New() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NowPlayingInfoLanguageOptionGroup) Init() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NowPlayingInfoLanguageOptionGroup) Autorelease() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNowPlayingInfoLanguageOptionGroup creates a new NowPlayingInfoLanguageOptionGroup instance.
func NewNowPlayingInfoLanguageOptionGroup() NowPlayingInfoLanguageOptionGroup {
	return getNowPlayingInfoLanguageOptionGroupClass().New()
}


// Creates a new language option group with the supplied language options.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/init(languageOptions:defaultLanguageOption:allowEmptySelection:)
func NewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection(languageOptions unsafe.Pointer, defaultLanguageOption unsafe.Pointer, allowEmptySelection bool) NowPlayingInfoLanguageOptionGroup {
	instance := getNowPlayingInfoLanguageOptionGroupClass().Alloc()
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](instance.ID, objc.Sel("initWithLanguageOptions:defaultLanguageOption:allowEmptySelection:"), languageOptions, defaultLanguageOption, allowEmptySelection)
	rv.Autorelease()
	return rv
}


// A Boolean that indicates whether the system requires a selection for the language option group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/allowEmptySelection
func (n_ NowPlayingInfoLanguageOptionGroup) AllowEmptySelection() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowEmptySelection"))
	return rv
}

// The default language option for the group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/defaultLanguageOption
func (n_ NowPlayingInfoLanguageOptionGroup) DefaultLanguageOption() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("defaultLanguageOption"))
	return rv
}

// The available language options for the group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/languageOptions
func (n_ NowPlayingInfoLanguageOptionGroup) LanguageOptions() []NowPlayingInfoLanguageOption {
	rv := objc.Send[[]NowPlayingInfoLanguageOption](n_.ID, objc.Sel("languageOptions"))
	return rv
}


