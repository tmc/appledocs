// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserScriptTask] class.
var userScriptTaskClass = _UserScriptTaskClass{objc.GetClass("NSUserScriptTask")}

type _UserScriptTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserScriptTask] class.
type IUserScriptTask interface {
	objectivec.IObject
}

// An object that executes scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask

type UserScriptTask struct {
	objectivec.Object
}

// UserScriptTaskFrom constructs a [UserScriptTask] from an unsafe.Pointer.
//
// An object that executes scripts.
func UserScriptTaskFrom(ptr unsafe.Pointer) UserScriptTask {
	return UserScriptTask{objectivec.Object{objc.ID(ptr)}}
}



