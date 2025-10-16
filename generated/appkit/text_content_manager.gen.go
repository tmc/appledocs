
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContentManager] class.
var TextContentManagerClass _TextContentManagerClass

func init() {
	TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}
}

type _TextContentManagerClass struct {
	objc.Class
}

// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	ID() objc.ID
}

type TextContentManager struct {
	id objc.ID
}

func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextContentManager) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextContentManagerClass) New() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextContentManager creates and returns a new initialized instance.
func NewTextContentManager() TextContentManager {
	return TextContentManagerClass.New()
}

// Init initializes the instance.
func (t_ TextContentManager) Init() TextContentManager {
	rv := objc.Send[TextContentManager](t_.ID(), selInit)
	return rv
}
