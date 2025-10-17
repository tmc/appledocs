// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Thread] class.
var threadClass = _ThreadClass{objc.GetClass("NSThread")}

type _ThreadClass struct {
	class objc.Class
}

// A thread of execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread

type Thread struct {
	objectivec.Object
}

// ThreadFrom constructs a [Thread] from an unsafe.Pointer.
//
// A thread of execution.
func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{objectivec.Object{objc.ID(ptr)}}
}



