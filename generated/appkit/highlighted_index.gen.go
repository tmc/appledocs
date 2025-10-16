
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [highlightedIndex] class.
var highlightedIndexClass _highlightedIndexClass

func init() {
	highlightedIndexClass = _highlightedIndexClass{objc.GetClass("highlightedIndex")}
}

type _highlightedIndexClass struct {
	objc.Class
}

// An interface definition for the [highlightedIndex] class.
type IhighlightedIndex interface {
	ID() objc.ID
}

type highlightedIndex struct {
	id objc.ID
}

func highlightedIndexFrom(ptr unsafe.Pointer) highlightedIndex {
	return highlightedIndex{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ highlightedIndex) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _highlightedIndexClass) Alloc() highlightedIndex {
	rv := objc.Send[highlightedIndex](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _highlightedIndexClass) New() highlightedIndex {
	rv := objc.Send[highlightedIndex](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhighlightedIndex creates and returns a new initialized instance.
func NewhighlightedIndex() highlightedIndex {
	return highlightedIndexClass.New()
}

// Init initializes the instance.
func (h_ highlightedIndex) Init() highlightedIndex {
	rv := objc.Send[highlightedIndex](h_.ID(), selInit)
	return rv
}
