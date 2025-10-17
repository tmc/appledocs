// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var WorkspaceOpenConfigurationClass objc.Class

func init() {
	WorkspaceOpenConfigurationClass = objc.GetClass("NSWorkspaceOpenConfiguration")
}

type WorkspaceOpenConfiguration struct {
	objc.ID
}

func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{
		ID: objc.ID(ptr),
	}
}




