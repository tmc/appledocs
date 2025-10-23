// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WorkspaceAuthorization] class.
var (
	WorkspaceAuthorizationClass     _WorkspaceAuthorizationClass
	WorkspaceAuthorizationClassOnce sync.Once
)

func getWorkspaceAuthorizationClass() _WorkspaceAuthorizationClass {
	WorkspaceAuthorizationClassOnce.Do(func() {
		WorkspaceAuthorizationClass = _WorkspaceAuthorizationClass{objc.GetClass("NSWorkspaceAuthorization")}
	})
	return WorkspaceAuthorizationClass
}

type _WorkspaceAuthorizationClass struct {
	class objc.Class
}

// An interface definition for the [WorkspaceAuthorization] class.
type IWorkspaceAuthorization interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The authorization granted to the app by the user.
//
// To enable your app to prompt the user for these file permissions, you must have a Privileged File Operation entitlement. If you have an app on the Mac App Store or plan to submit your app for review, you can .


// The authorization granted to the app by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/Authorization
type WorkspaceAuthorization struct {
	objectivec.Object
}

// WorkspaceAuthorizationFrom constructs a [WorkspaceAuthorization] from an unsafe.Pointer.
//
// The authorization granted to the app by the user.
func WorkspaceAuthorizationFrom(ptr unsafe.Pointer) WorkspaceAuthorization {
	return WorkspaceAuthorization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WorkspaceAuthorizationClass) Alloc() WorkspaceAuthorization {
	rv := objc.Send[WorkspaceAuthorization](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WorkspaceAuthorizationClass) New() WorkspaceAuthorization {
	rv := objc.Send[WorkspaceAuthorization](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WorkspaceAuthorization) Init() WorkspaceAuthorization {
	rv := objc.Send[WorkspaceAuthorization](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WorkspaceAuthorization) Autorelease() WorkspaceAuthorization {
	rv := objc.Send[WorkspaceAuthorization](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWorkspaceAuthorization creates a new WorkspaceAuthorization instance.
func NewWorkspaceAuthorization() WorkspaceAuthorization {
	return getWorkspaceAuthorizationClass().New()
}




