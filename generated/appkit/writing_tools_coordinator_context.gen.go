
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WritingToolsCoordinatorContext] class.
var WritingToolsCoordinatorContextClass _WritingToolsCoordinatorContextClass

func init() {
	WritingToolsCoordinatorContextClass = _WritingToolsCoordinatorContextClass{objc.GetClass("NSWritingToolsCoordinatorContext")}
}

type _WritingToolsCoordinatorContextClass struct {
	objc.Class
}

// An interface definition for the [WritingToolsCoordinatorContext] class.
type IWritingToolsCoordinatorContext interface {
	ID() objc.ID
}

type WritingToolsCoordinatorContext struct {
	id objc.ID
}

func WritingToolsCoordinatorContextFrom(ptr unsafe.Pointer) WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WritingToolsCoordinatorContext) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorContextClass) Alloc() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WritingToolsCoordinatorContextClass) New() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWritingToolsCoordinatorContext creates and returns a new initialized instance.
func NewWritingToolsCoordinatorContext() WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContextClass.New()
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorContext) Init() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](w_.ID(), selInit)
	return rv
}
// The unique identifier of the context object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/Context/range
func (w_ WritingToolsCoordinatorContext) Range() foundation.Range {
	rv := objc.Send[foundation.Range](w_.ID(), objc.RegisterName("range"))
	return rv
}

