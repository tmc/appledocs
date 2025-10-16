
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titlebarSeparatorStyle] class.
var titlebarSeparatorStyleClass _titlebarSeparatorStyleClass

func init() {
	titlebarSeparatorStyleClass = _titlebarSeparatorStyleClass{objc.GetClass("titlebarSeparatorStyle")}
}

type _titlebarSeparatorStyleClass struct {
	objc.Class
}

// An interface definition for the [titlebarSeparatorStyle] class.
type ItitlebarSeparatorStyle interface {
	ID() objc.ID
}

type titlebarSeparatorStyle struct {
	id objc.ID
}

func titlebarSeparatorStyleFrom(ptr unsafe.Pointer) titlebarSeparatorStyle {
	return titlebarSeparatorStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titlebarSeparatorStyle) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titlebarSeparatorStyleClass) Alloc() titlebarSeparatorStyle {
	rv := objc.Send[titlebarSeparatorStyle](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titlebarSeparatorStyleClass) New() titlebarSeparatorStyle {
	rv := objc.Send[titlebarSeparatorStyle](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitlebarSeparatorStyle creates and returns a new initialized instance.
func NewtitlebarSeparatorStyle() titlebarSeparatorStyle {
	return titlebarSeparatorStyleClass.New()
}

// Init initializes the instance.
func (t_ titlebarSeparatorStyle) Init() titlebarSeparatorStyle {
	rv := objc.Send[titlebarSeparatorStyle](t_.ID(), selInit)
	return rv
}
