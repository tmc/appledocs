// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var localeClass _LocaleClass

func init() {
	localeClass = _LocaleClass{objc.GetClass("NSLocale")}
}

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




