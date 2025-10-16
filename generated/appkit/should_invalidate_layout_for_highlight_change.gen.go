
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [shouldInvalidateLayoutForHighlightChange] class.
var shouldInvalidateLayoutForHighlightChangeClass _shouldInvalidateLayoutForHighlightChangeClass

func init() {
	shouldInvalidateLayoutForHighlightChangeClass = _shouldInvalidateLayoutForHighlightChangeClass{objc.GetClass("shouldInvalidateLayoutForHighlightChange")}
}

type _shouldInvalidateLayoutForHighlightChangeClass struct {
	objc.Class
}

// An interface definition for the [shouldInvalidateLayoutForHighlightChange] class.
type IshouldInvalidateLayoutForHighlightChange interface {
	ID() objc.ID
}

type shouldInvalidateLayoutForHighlightChange struct {
	id objc.ID
}

func shouldInvalidateLayoutForHighlightChangeFrom(ptr unsafe.Pointer) shouldInvalidateLayoutForHighlightChange {
	return shouldInvalidateLayoutForHighlightChange{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ shouldInvalidateLayoutForHighlightChange) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _shouldInvalidateLayoutForHighlightChangeClass) Alloc() shouldInvalidateLayoutForHighlightChange {
	rv := objc.Send[shouldInvalidateLayoutForHighlightChange](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _shouldInvalidateLayoutForHighlightChangeClass) New() shouldInvalidateLayoutForHighlightChange {
	rv := objc.Send[shouldInvalidateLayoutForHighlightChange](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshouldInvalidateLayoutForHighlightChange creates and returns a new initialized instance.
func NewshouldInvalidateLayoutForHighlightChange() shouldInvalidateLayoutForHighlightChange {
	return shouldInvalidateLayoutForHighlightChangeClass.New()
}

// Init initializes the instance.
func (s_ shouldInvalidateLayoutForHighlightChange) Init() shouldInvalidateLayoutForHighlightChange {
	rv := objc.Send[shouldInvalidateLayoutForHighlightChange](s_.ID(), selInit)
	return rv
}
