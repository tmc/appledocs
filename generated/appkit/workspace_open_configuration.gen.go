// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var (
	workspaceOpenConfigurationClass     _WorkspaceOpenConfigurationClass
	workspaceOpenConfigurationClassOnce sync.Once
)

func getWorkspaceOpenConfigurationClass() _WorkspaceOpenConfigurationClass {
	workspaceOpenConfigurationClassOnce.Do(func() {
		workspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}
	})
	return workspaceOpenConfigurationClass
}

type _WorkspaceOpenConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [WorkspaceOpenConfiguration] class.
type IWorkspaceOpenConfiguration interface {
	objectivec.IObject
}

// The configuration options for opening URLs or launching apps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration

type WorkspaceOpenConfiguration struct {
	objectivec.Object
}

// WorkspaceOpenConfigurationFrom constructs a [WorkspaceOpenConfiguration] from an unsafe.Pointer.
//
// The configuration options for opening URLs or launching apps.
func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (wc _WorkspaceOpenConfigurationClass) Alloc() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (wc _WorkspaceOpenConfigurationClass) New() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WorkspaceOpenConfiguration) Init() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WorkspaceOpenConfiguration) Autorelease() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWorkspaceOpenConfiguration creates a new WorkspaceOpenConfiguration instance.
func NewWorkspaceOpenConfiguration() WorkspaceOpenConfiguration {
	return getWorkspaceOpenConfigurationClass().New()
}




