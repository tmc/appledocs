
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var MutableParagraphStyleClass _MutableParagraphStyleClass

func init() {
	MutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}
}

type _MutableParagraphStyleClass struct {
	objc.Class
}

// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	ID() objc.ID
}

type MutableParagraphStyle struct {
	id objc.ID
}

func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MutableParagraphStyle) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMutableParagraphStyle creates and returns a new initialized instance.
func NewMutableParagraphStyle() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

// Init initializes the instance.
func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](m_.ID(), selInit)
	return rv
}
