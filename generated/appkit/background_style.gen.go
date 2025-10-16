
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundStyle] class.
var backgroundStyleClass _backgroundStyleClass

func init() {
	backgroundStyleClass = _backgroundStyleClass{objc.GetClass("backgroundStyle")}
}

type _backgroundStyleClass struct {
	objc.Class
}

// An interface definition for the [backgroundStyle] class.
type IbackgroundStyle interface {
	ID() objc.ID
}

type backgroundStyle struct {
	id objc.ID
}

func backgroundStyleFrom(ptr unsafe.Pointer) backgroundStyle {
	return backgroundStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundStyle) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundStyleClass) Alloc() backgroundStyle {
	rv := objc.Send[backgroundStyle](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundStyleClass) New() backgroundStyle {
	rv := objc.Send[backgroundStyle](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundStyle creates and returns a new initialized instance.
func NewbackgroundStyle() backgroundStyle {
	return backgroundStyleClass.New()
}

// Init initializes the instance.
func (b_ backgroundStyle) Init() backgroundStyle {
	rv := objc.Send[backgroundStyle](b_.ID(), selInit)
	return rv
}
