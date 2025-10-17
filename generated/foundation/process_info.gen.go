// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ProcessInfo] class.
var ProcessInfoClass = _ProcessInfoClass{objc.GetClass("NSProcessInfo")}

type _ProcessInfoClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p_ ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performExpiringActivityWithReason:usingBlock:"), reason, block)
}


