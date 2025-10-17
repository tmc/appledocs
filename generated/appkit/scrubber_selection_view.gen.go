
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberSelectionView] class.
var ScrubberSelectionViewClass _ScrubberSelectionViewClass

func init() {
	ScrubberSelectionViewClass = _ScrubberSelectionViewClass{objc.GetClass("NSScrubberSelectionView")}
}

type _ScrubberSelectionViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberSelectionView] class.
type IScrubberSelectionView interface {
	ID() objc.ID
}

type ScrubberSelectionView struct {
	id objc.ID
}

func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberSelectionView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionViewClass) Alloc() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberSelectionViewClass) New() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberSelectionView creates and returns a new initialized instance.
func NewScrubberSelectionView() ScrubberSelectionView {
	return ScrubberSelectionViewClass.New()
}

// Init initializes the instance.
func (s_ ScrubberSelectionView) Init() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID(), selInit)
	return rv
}
