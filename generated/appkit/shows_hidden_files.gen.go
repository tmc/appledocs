
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsHiddenFiles] class.
var showsHiddenFilesClass _showsHiddenFilesClass

func init() {
	showsHiddenFilesClass = _showsHiddenFilesClass{objc.GetClass("showsHiddenFiles")}
}

type _showsHiddenFilesClass struct {
	objc.Class
}

// An interface definition for the [showsHiddenFiles] class.
type IshowsHiddenFiles interface {
	ID() objc.ID
}

type showsHiddenFiles struct {
	id objc.ID
}

func showsHiddenFilesFrom(ptr unsafe.Pointer) showsHiddenFiles {
	return showsHiddenFiles{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsHiddenFiles) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsHiddenFilesClass) Alloc() showsHiddenFiles {
	rv := objc.Send[showsHiddenFiles](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsHiddenFilesClass) New() showsHiddenFiles {
	rv := objc.Send[showsHiddenFiles](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsHiddenFiles creates and returns a new initialized instance.
func NewshowsHiddenFiles() showsHiddenFiles {
	return showsHiddenFilesClass.New()
}

// Init initializes the instance.
func (s_ showsHiddenFiles) Init() showsHiddenFiles {
	rv := objc.Send[showsHiddenFiles](s_.ID(), selInit)
	return rv
}
