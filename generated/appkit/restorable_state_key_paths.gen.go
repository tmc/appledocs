
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [restorableStateKeyPaths] class.
var restorableStateKeyPathsClass _restorableStateKeyPathsClass

func init() {
	restorableStateKeyPathsClass = _restorableStateKeyPathsClass{objc.GetClass("restorableStateKeyPaths")}
}

type _restorableStateKeyPathsClass struct {
	objc.Class
}

// An interface definition for the [restorableStateKeyPaths] class.
type IrestorableStateKeyPaths interface {
	ID() objc.ID
}

type restorableStateKeyPaths struct {
	id objc.ID
}

func restorableStateKeyPathsFrom(ptr unsafe.Pointer) restorableStateKeyPaths {
	return restorableStateKeyPaths{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ restorableStateKeyPaths) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _restorableStateKeyPathsClass) Alloc() restorableStateKeyPaths {
	rv := objc.Send[restorableStateKeyPaths](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _restorableStateKeyPathsClass) New() restorableStateKeyPaths {
	rv := objc.Send[restorableStateKeyPaths](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrestorableStateKeyPaths creates and returns a new initialized instance.
func NewrestorableStateKeyPaths() restorableStateKeyPaths {
	return restorableStateKeyPathsClass.New()
}

// Init initializes the instance.
func (r_ restorableStateKeyPaths) Init() restorableStateKeyPaths {
	rv := objc.Send[restorableStateKeyPaths](r_.ID(), selInit)
	return rv
}
