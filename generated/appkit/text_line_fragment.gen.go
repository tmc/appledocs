// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLineFragment] class.
var textLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	class objc.Class
}

// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	objectivec.IObject
}

// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment

type TextLineFragment struct {
	objectivec.Object
}

// TextLineFragmentFrom constructs a [TextLineFragment] from an unsafe.Pointer.
//
// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLineFragment) Autorelease() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLineFragment creates a new TextLineFragment instance.
func NewTextLineFragment() TextLineFragment {
	return textLineFragmentClass.New()
}




