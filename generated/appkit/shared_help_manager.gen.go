
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sharedHelpManager] class.
var sharedHelpManagerClass _sharedHelpManagerClass

func init() {
	sharedHelpManagerClass = _sharedHelpManagerClass{objc.GetClass("sharedHelpManager")}
}

type _sharedHelpManagerClass struct {
	objc.Class
}

// An interface definition for the [sharedHelpManager] class.
type IsharedHelpManager interface {
	ID() objc.ID
}

type sharedHelpManager struct {
	id objc.ID
}

func sharedHelpManagerFrom(ptr unsafe.Pointer) sharedHelpManager {
	return sharedHelpManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sharedHelpManager) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sharedHelpManagerClass) Alloc() sharedHelpManager {
	rv := objc.Send[sharedHelpManager](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sharedHelpManagerClass) New() sharedHelpManager {
	rv := objc.Send[sharedHelpManager](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsharedHelpManager creates and returns a new initialized instance.
func NewsharedHelpManager() sharedHelpManager {
	return sharedHelpManagerClass.New()
}

// Init initializes the instance.
func (s_ sharedHelpManager) Init() sharedHelpManager {
	rv := objc.Send[sharedHelpManager](s_.ID(), selInit)
	return rv
}
