
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FilePromiseProvider] class.
var FilePromiseProviderClass _FilePromiseProviderClass

func init() {
	FilePromiseProviderClass = _FilePromiseProviderClass{objc.GetClass("NSFilePromiseProvider")}
}

type _FilePromiseProviderClass struct {
	objc.Class
}

// An interface definition for the [FilePromiseProvider] class.
type IFilePromiseProvider interface {
	ID() objc.ID
}

type FilePromiseProvider struct {
	id objc.ID
}

func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FilePromiseProvider) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FilePromiseProviderClass) Alloc() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FilePromiseProviderClass) New() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFilePromiseProvider creates and returns a new initialized instance.
func NewFilePromiseProvider() FilePromiseProvider {
	return FilePromiseProviderClass.New()
}

// Init initializes the instance.
func (f_ FilePromiseProvider) Init() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](f_.ID(), selInit)
	return rv
}
