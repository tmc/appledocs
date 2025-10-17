// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Locale] class.
var LocaleClass objc.Class

func init() {
	LocaleClass = objc.GetClass("NSLocale")
}

type Locale struct {
	objc.ID
}

func LocaleFrom(ptr unsafe.Pointer) Locale {
	return Locale{
		ID: objc.ID(ptr),
	}
}




