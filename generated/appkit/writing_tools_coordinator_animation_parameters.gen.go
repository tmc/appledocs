
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WritingToolsCoordinatorAnimationParameters] class.
var WritingToolsCoordinatorAnimationParametersClass _WritingToolsCoordinatorAnimationParametersClass

func init() {
	WritingToolsCoordinatorAnimationParametersClass = _WritingToolsCoordinatorAnimationParametersClass{objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")}
}

type _WritingToolsCoordinatorAnimationParametersClass struct {
	objc.Class
}

// An interface definition for the [WritingToolsCoordinatorAnimationParameters] class.
type IWritingToolsCoordinatorAnimationParameters interface {
	ID() objc.ID
}

type WritingToolsCoordinatorAnimationParameters struct {
	id objc.ID
}

func WritingToolsCoordinatorAnimationParametersFrom(ptr unsafe.Pointer) WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParameters{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WritingToolsCoordinatorAnimationParameters) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorAnimationParametersClass) Alloc() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WritingToolsCoordinatorAnimationParametersClass) New() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWritingToolsCoordinatorAnimationParameters creates and returns a new initialized instance.
func NewWritingToolsCoordinatorAnimationParameters() WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParametersClass.New()
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorAnimationParameters) Init() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](w_.ID(), selInit)
	return rv
}
