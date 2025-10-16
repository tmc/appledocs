
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsFirstResponder] class.
var showsFirstResponderClass _showsFirstResponderClass

func init() {
	showsFirstResponderClass = _showsFirstResponderClass{objc.GetClass("showsFirstResponder")}
}

type _showsFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [showsFirstResponder] class.
type IshowsFirstResponder interface {
	ID() objc.ID
}

type showsFirstResponder struct {
	id objc.ID
}

func showsFirstResponderFrom(ptr unsafe.Pointer) showsFirstResponder {
	return showsFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsFirstResponder) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsFirstResponderClass) Alloc() showsFirstResponder {
	rv := objc.Send[showsFirstResponder](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsFirstResponderClass) New() showsFirstResponder {
	rv := objc.Send[showsFirstResponder](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsFirstResponder creates and returns a new initialized instance.
func NewshowsFirstResponder() showsFirstResponder {
	return showsFirstResponderClass.New()
}

// Init initializes the instance.
func (s_ showsFirstResponder) Init() showsFirstResponder {
	rv := objc.Send[showsFirstResponder](s_.ID(), selInit)
	return rv
}
