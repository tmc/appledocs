
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LevelIndicator] class.
var LevelIndicatorClass _LevelIndicatorClass

func init() {
	LevelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}
}

type _LevelIndicatorClass struct {
	objc.Class
}

// An interface definition for the [LevelIndicator] class.
type ILevelIndicator interface {
	ID() objc.ID
}

type LevelIndicator struct {
	id objc.ID
}

func LevelIndicatorFrom(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LevelIndicator) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := objc.Send[LevelIndicator](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LevelIndicatorClass) New() LevelIndicator {
	rv := objc.Send[LevelIndicator](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLevelIndicator creates and returns a new initialized instance.
func NewLevelIndicator() LevelIndicator {
	return LevelIndicatorClass.New()
}

// Init initializes the instance.
func (l_ LevelIndicator) Init() LevelIndicator {
	rv := objc.Send[LevelIndicator](l_.ID(), selInit)
	return rv
}
