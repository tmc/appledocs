
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [dividerStyle] class.
var dividerStyleClass _dividerStyleClass

func init() {
	dividerStyleClass = _dividerStyleClass{objc.GetClass("dividerStyle")}
}

type _dividerStyleClass struct {
	objc.Class
}

// An interface definition for the [dividerStyle] class.
type IdividerStyle interface {
	ID() objc.ID
}

type dividerStyle struct {
	id objc.ID
}

func dividerStyleFrom(ptr unsafe.Pointer) dividerStyle {
	return dividerStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ dividerStyle) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _dividerStyleClass) Alloc() dividerStyle {
	rv := objc.Send[dividerStyle](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _dividerStyleClass) New() dividerStyle {
	rv := objc.Send[dividerStyle](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdividerStyle creates and returns a new initialized instance.
func NewdividerStyle() dividerStyle {
	return dividerStyleClass.New()
}

// Init initializes the instance.
func (d_ dividerStyle) Init() dividerStyle {
	rv := objc.Send[dividerStyle](d_.ID(), selInit)
	return rv
}
