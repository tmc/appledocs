
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsBaselineSeparator] class.
var showsBaselineSeparatorClass _showsBaselineSeparatorClass

func init() {
	showsBaselineSeparatorClass = _showsBaselineSeparatorClass{objc.GetClass("showsBaselineSeparator")}
}

type _showsBaselineSeparatorClass struct {
	objc.Class
}

// An interface definition for the [showsBaselineSeparator] class.
type IshowsBaselineSeparator interface {
	ID() objc.ID
}

type showsBaselineSeparator struct {
	id objc.ID
}

func showsBaselineSeparatorFrom(ptr unsafe.Pointer) showsBaselineSeparator {
	return showsBaselineSeparator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsBaselineSeparator) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsBaselineSeparatorClass) Alloc() showsBaselineSeparator {
	rv := objc.Send[showsBaselineSeparator](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsBaselineSeparatorClass) New() showsBaselineSeparator {
	rv := objc.Send[showsBaselineSeparator](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsBaselineSeparator creates and returns a new initialized instance.
func NewshowsBaselineSeparator() showsBaselineSeparator {
	return showsBaselineSeparatorClass.New()
}

// Init initializes the instance.
func (s_ showsBaselineSeparator) Init() showsBaselineSeparator {
	rv := objc.Send[showsBaselineSeparator](s_.ID(), selInit)
	return rv
}
