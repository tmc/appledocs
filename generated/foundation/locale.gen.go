// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Locale] class.
var LocaleClass = _LocaleClass{objc.GetClass("NSLocale")}

type _LocaleClass struct {
	class objc.Class
}

type Locale struct {
	objc.ID
}

func LocaleFrom(ptr unsafe.Pointer) Locale {
	return Locale{
		ID: objc.ID(ptr),
	}
}




