// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Nib] class.
var nibClass = _NibClass{objc.GetClass("NSNib")}

type _NibClass struct {
	class objc.Class
}

// An interface definition for the [Nib] class.
type INib interface {
	objectivec.IObject
}

// An object wrapper, or container, for an Interface Builder nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib

type Nib struct {
	objectivec.Object
}

// NibFrom constructs a [Nib] from an unsafe.Pointer.
//
// An object wrapper, or container, for an Interface Builder nib file.
func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{objectivec.Object{objc.ID(ptr)}}
}



