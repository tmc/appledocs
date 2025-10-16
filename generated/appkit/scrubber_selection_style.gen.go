
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberSelectionStyle] class.
var ScrubberSelectionStyleClass _ScrubberSelectionStyleClass

func init() {
	ScrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}
}

type _ScrubberSelectionStyleClass struct {
	objc.Class
}

// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	ID() objc.ID
}

type ScrubberSelectionStyle struct {
	id objc.ID
}

func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberSelectionStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionStyleClass) Alloc() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberSelectionStyleClass) New() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberSelectionStyle creates and returns a new initialized instance.
func NewScrubberSelectionStyle() ScrubberSelectionStyle {
	return ScrubberSelectionStyleClass.New()
}

// Init initializes the instance.
func (s_ ScrubberSelectionStyle) Init() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID(), selInit)
	return rv
}
