
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentView] class.
var contentViewClass _contentViewClass

func init() {
	contentViewClass = _contentViewClass{objc.GetClass("contentView")}
}

type _contentViewClass struct {
	objc.Class
}

// An interface definition for the [contentView] class.
type IcontentView interface {
	ID() objc.ID
}

type contentView struct {
	id objc.ID
}

func contentViewFrom(ptr unsafe.Pointer) contentView {
	return contentView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentViewClass) Alloc() contentView {
	rv := objc.Send[contentView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentViewClass) New() contentView {
	rv := objc.Send[contentView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentView creates and returns a new initialized instance.
func NewcontentView() contentView {
	return contentViewClass.New()
}

// Init initializes the instance.
func (c_ contentView) Init() contentView {
	rv := objc.Send[contentView](c_.ID(), selInit)
	return rv
}
