// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PageLayout] class.
var pageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}

type _PageLayoutClass struct {
	class objc.Class
}

// A panel that queries the user for information such as paper type and orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout

type PageLayout struct {
	objectivec.Object
}

// PageLayoutFrom constructs a [PageLayout] from an unsafe.Pointer.
//
// A panel that queries the user for information such as paper type and orientation.
func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{objectivec.Object{objc.ID(ptr)}}
}



