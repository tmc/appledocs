// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPNowPlayingInfoLanguageOptionGroup */


/* debug [class_header]: Header for MPNowPlayingInfoLanguageOptionGroup */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NowPlayingInfoLanguageOptionGroup */
// An interface definition for the [NowPlayingInfoLanguageOptionGroup] class.
type INowPlayingInfoLanguageOptionGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NowPlayingInfoLanguageOptionGroup */
	// properties:
	AllowEmptySelection() bool
	DefaultLanguageOption() IMPNowPlayingInfoLanguageOption
	LanguageOptions() []NowPlayingInfoLanguageOption
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NowPlayingInfoLanguageOptionGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NowPlayingInfoLanguageOptionGroup */
// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoLanguageOptionGroupClass) Alloc() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NowPlayingInfoLanguageOptionGroup */
// A grouped set of language options where only a single language option can be active at a time.
//
// The and classes provide interfaces for setting information about language options, for example, audio and subtitles, in the Now Playing information area.


// A grouped set of language options where only a single language option can be active at a time.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NowPlayingInfoLanguageOptionGroup */

// Creates a new language option group with the supplied language options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/init(languageOptions:defaultLanguageOption:allowEmptySelection:)
func NewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection(languageOptions []NowPlayingInfoLanguageOption, defaultLanguageOption IMPNowPlayingInfoLanguageOption, allowEmptySelection bool) NowPlayingInfoLanguageOptionGroup {
	instance := getNowPlayingInfoLanguageOptionGroupClass().Alloc()
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](instance.ID, objc.Sel("initWithLanguageOptions:defaultLanguageOption:allowEmptySelection:"), languageOptions, defaultLanguageOption, allowEmptySelection)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NowPlayingInfoLanguageOptionGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NowPlayingInfoLanguageOptionGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NowPlayingInfoLanguageOptionGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NowPlayingInfoLanguageOptionGroup */

// A Boolean that indicates whether the system requires a selection for the language option group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/allowEmptySelection
func (n_ NowPlayingInfoLanguageOptionGroup) AllowEmptySelection() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowEmptySelection"))
	return rv
}/* debug [instance_properties/getter]: allowEmptySelection */


// The default language option for the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/defaultLanguageOption
func (n_ NowPlayingInfoLanguageOptionGroup) DefaultLanguageOption() IMPNowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](n_.ID, objc.Sel("defaultLanguageOption"))
	return rv
}/* debug [instance_properties/getter]: defaultLanguageOption */


// The available language options for the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionGroup/languageOptions
func (n_ NowPlayingInfoLanguageOptionGroup) LanguageOptions() []NowPlayingInfoLanguageOption {
	rv := objc.Send[[]NowPlayingInfoLanguageOption](n_.ID, objc.Sel("languageOptions"))
	return rv
}/* debug [instance_properties/getter]: languageOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPNowPlayingInfoLanguageOptionGroup */


