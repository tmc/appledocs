
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var WorkspaceOpenConfigurationClass _WorkspaceOpenConfigurationClass

func init() {
	WorkspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}
}

type _WorkspaceOpenConfigurationClass struct {
	objc.Class
}

// An interface definition for the [WorkspaceOpenConfiguration] class.
type IWorkspaceOpenConfiguration interface {
	ID() objc.ID
}

type WorkspaceOpenConfiguration struct {
	id objc.ID
}

func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WorkspaceOpenConfiguration) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WorkspaceOpenConfigurationClass) Alloc() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WorkspaceOpenConfigurationClass) New() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWorkspaceOpenConfiguration creates and returns a new initialized instance.
func NewWorkspaceOpenConfiguration() WorkspaceOpenConfiguration {
	return WorkspaceOpenConfigurationClass.New()
}

// Init initializes the instance.
func (w_ WorkspaceOpenConfiguration) Init() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID(), selInit)
	return rv
}
