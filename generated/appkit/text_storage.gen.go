
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextStorage] class.
var TextStorageClass _TextStorageClass

func init() {
	TextStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}
}

type _TextStorageClass struct {
	objc.Class
}

// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	ID() objc.ID
}

type TextStorage struct {
	id objc.ID
}

func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextStorage) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextStorageClass) New() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextStorage creates and returns a new initialized instance.
func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

// Init initializes the instance.
func (t_ TextStorage) Init() TextStorage {
	rv := objc.Send[TextStorage](t_.ID(), selInit)
	return rv
}
