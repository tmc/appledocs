// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ProcessInfo] class.
var ProcessInfoClass objc.Class

func init() {
	ProcessInfoClass = objc.GetClass("NSProcessInfo")
}

type ProcessInfo struct {
	objc.ID
}

func ProcessInfoFrom(ptr unsafe.Pointer) ProcessInfo {
	return ProcessInfo{
		ID: objc.ID(ptr),
	}
}


// Performs the specified block asynchronously and notifies you if the process is about to be suspended. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p_ ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer) {
	sel := objc.RegisterName("performExpiringActivityWithReason:usingBlock:")
	p_.ID.Send(sel, reason, block)
}


