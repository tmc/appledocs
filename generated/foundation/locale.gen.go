// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Locale] class.
var (
	localeClass     _LocaleClass
	localeClassOnce sync.Once
)

func getLocaleClass() _LocaleClass {
	localeClassOnce.Do(func() {
		localeClass = _LocaleClass{objc.GetClass("NSLocale")}
	})
	return localeClass
}

type _LocaleClass struct {
	class objc.Class
}

// An interface definition for the [Locale] class.
type ILocale interface {
	objectivec.IObject
}

// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale
type Locale struct {
	objectivec.Object
}

// LocaleFrom constructs a [Locale] from an unsafe.Pointer.
//
// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
func LocaleFrom(ptr unsafe.Pointer) Locale {
	return Locale{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LocaleClass) Alloc() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocaleClass) New() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Locale) Init() Locale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Locale) Autorelease() Locale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocale creates a new Locale instance.
func NewLocale() Locale {
	return getLocaleClass().New()
}




