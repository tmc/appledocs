
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsTagField] class.
var showsTagFieldClass _showsTagFieldClass

func init() {
	showsTagFieldClass = _showsTagFieldClass{objc.GetClass("showsTagField")}
}

type _showsTagFieldClass struct {
	objc.Class
}

// An interface definition for the [showsTagField] class.
type IshowsTagField interface {
	ID() objc.ID
}

type showsTagField struct {
	id objc.ID
}

func showsTagFieldFrom(ptr unsafe.Pointer) showsTagField {
	return showsTagField{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsTagField) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsTagFieldClass) Alloc() showsTagField {
	rv := objc.Send[showsTagField](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsTagFieldClass) New() showsTagField {
	rv := objc.Send[showsTagField](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsTagField creates and returns a new initialized instance.
func NewshowsTagField() showsTagField {
	return showsTagFieldClass.New()
}

// Init initializes the instance.
func (s_ showsTagField) Init() showsTagField {
	rv := objc.Send[showsTagField](s_.ID(), selInit)
	return rv
}
