
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLineFragment] class.
var TextLineFragmentClass _TextLineFragmentClass

func init() {
	TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}
}

type _TextLineFragmentClass struct {
	objc.Class
}

// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	ID() objc.ID
}

type TextLineFragment struct {
	id objc.ID
}

func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextLineFragment) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextLineFragment creates and returns a new initialized instance.
func NewTextLineFragment() TextLineFragment {
	return TextLineFragmentClass.New()
}

// Init initializes the instance.
func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID(), selInit)
	return rv
}
