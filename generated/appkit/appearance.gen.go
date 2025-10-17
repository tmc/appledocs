// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Appearance] class.
var appearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}

type _AppearanceClass struct {
	class objc.Class
}

// An object that manages standard appearance attributes for UI elements in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance

type Appearance struct {
	objectivec.Object
}

// AppearanceFrom constructs a [Appearance] from an unsafe.Pointer.
//
// An object that manages standard appearance attributes for UI elements in an app.
func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{objectivec.Object{objc.ID(ptr)}}
}



