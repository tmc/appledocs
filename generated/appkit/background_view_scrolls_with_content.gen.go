
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundViewScrollsWithContent] class.
var backgroundViewScrollsWithContentClass _backgroundViewScrollsWithContentClass

func init() {
	backgroundViewScrollsWithContentClass = _backgroundViewScrollsWithContentClass{objc.GetClass("backgroundViewScrollsWithContent")}
}

type _backgroundViewScrollsWithContentClass struct {
	objc.Class
}

// An interface definition for the [backgroundViewScrollsWithContent] class.
type IbackgroundViewScrollsWithContent interface {
	ID() objc.ID
}

type backgroundViewScrollsWithContent struct {
	id objc.ID
}

func backgroundViewScrollsWithContentFrom(ptr unsafe.Pointer) backgroundViewScrollsWithContent {
	return backgroundViewScrollsWithContent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundViewScrollsWithContent) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundViewScrollsWithContentClass) Alloc() backgroundViewScrollsWithContent {
	rv := objc.Send[backgroundViewScrollsWithContent](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundViewScrollsWithContentClass) New() backgroundViewScrollsWithContent {
	rv := objc.Send[backgroundViewScrollsWithContent](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundViewScrollsWithContent creates and returns a new initialized instance.
func NewbackgroundViewScrollsWithContent() backgroundViewScrollsWithContent {
	return backgroundViewScrollsWithContentClass.New()
}

// Init initializes the instance.
func (b_ backgroundViewScrollsWithContent) Init() backgroundViewScrollsWithContent {
	rv := objc.Send[backgroundViewScrollsWithContent](b_.ID(), selInit)
	return rv
}
