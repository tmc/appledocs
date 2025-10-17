// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var workspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}

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



