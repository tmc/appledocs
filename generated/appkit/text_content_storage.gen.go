
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContentStorage] class.
var TextContentStorageClass _TextContentStorageClass

func init() {
	TextContentStorageClass = _TextContentStorageClass{objc.GetClass("NSTextContentStorage")}
}

type _TextContentStorageClass struct {
	objc.Class
}

// An interface definition for the [TextContentStorage] class.
type ITextContentStorage interface {
	ID() objc.ID
}

type TextContentStorage struct {
	id objc.ID
}

func TextContentStorageFrom(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextContentStorage) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextContentStorageClass) Alloc() TextContentStorage {
	rv := objc.Send[TextContentStorage](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextContentStorageClass) New() TextContentStorage {
	rv := objc.Send[TextContentStorage](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextContentStorage creates and returns a new initialized instance.
func NewTextContentStorage() TextContentStorage {
	return TextContentStorageClass.New()
}

// Init initializes the instance.
func (t_ TextContentStorage) Init() TextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID(), selInit)
	return rv
}
