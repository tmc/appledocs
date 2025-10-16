
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canChooseFiles] class.
var canChooseFilesClass _canChooseFilesClass

func init() {
	canChooseFilesClass = _canChooseFilesClass{objc.GetClass("canChooseFiles")}
}

type _canChooseFilesClass struct {
	objc.Class
}

// An interface definition for the [canChooseFiles] class.
type IcanChooseFiles interface {
	ID() objc.ID
}

type canChooseFiles struct {
	id objc.ID
}

func canChooseFilesFrom(ptr unsafe.Pointer) canChooseFiles {
	return canChooseFiles{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canChooseFiles) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canChooseFilesClass) Alloc() canChooseFiles {
	rv := objc.Send[canChooseFiles](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canChooseFilesClass) New() canChooseFiles {
	rv := objc.Send[canChooseFiles](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanChooseFiles creates and returns a new initialized instance.
func NewcanChooseFiles() canChooseFiles {
	return canChooseFilesClass.New()
}

// Init initializes the instance.
func (c_ canChooseFiles) Init() canChooseFiles {
	rv := objc.Send[canChooseFiles](c_.ID(), selInit)
	return rv
}
