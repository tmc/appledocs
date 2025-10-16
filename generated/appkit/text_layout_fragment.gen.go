
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLayoutFragment] class.
var TextLayoutFragmentClass _TextLayoutFragmentClass

func init() {
	TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}
}

type _TextLayoutFragmentClass struct {
	objc.Class
}

// An interface definition for the [TextLayoutFragment] class.
type ITextLayoutFragment interface {
	ID() objc.ID
}

type TextLayoutFragment struct {
	id objc.ID
}

func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextLayoutFragment) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextLayoutFragment creates and returns a new initialized instance.
func NewTextLayoutFragment() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

// Init initializes the instance.
func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID(), selInit)
	return rv
}
