// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Locale] class.
var localeClass = _LocaleClass{objc.GetClass("NSLocale")}

type _LocaleClass struct {
	class objc.Class
}

// An interface definition for the [Locale] class.
type ILocale interface {
	objectivec.IObject
}

// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation. [Full Topic]
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



