
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [scrubberContentSize] class.
var scrubberContentSizeClass _scrubberContentSizeClass

func init() {
	scrubberContentSizeClass = _scrubberContentSizeClass{objc.GetClass("scrubberContentSize")}
}

type _scrubberContentSizeClass struct {
	objc.Class
}

// An interface definition for the [scrubberContentSize] class.
type IscrubberContentSize interface {
	ID() objc.ID
}

type scrubberContentSize struct {
	id objc.ID
}

func scrubberContentSizeFrom(ptr unsafe.Pointer) scrubberContentSize {
	return scrubberContentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ scrubberContentSize) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _scrubberContentSizeClass) Alloc() scrubberContentSize {
	rv := objc.Send[scrubberContentSize](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _scrubberContentSizeClass) New() scrubberContentSize {
	rv := objc.Send[scrubberContentSize](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscrubberContentSize creates and returns a new initialized instance.
func NewscrubberContentSize() scrubberContentSize {
	return scrubberContentSizeClass.New()
}

// Init initializes the instance.
func (s_ scrubberContentSize) Init() scrubberContentSize {
	rv := objc.Send[scrubberContentSize](s_.ID(), selInit)
	return rv
}
