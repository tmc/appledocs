// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [authenticationModuleEntries] class.
var (
	AuthenticationModuleEntriesClass     _authenticationModuleEntriesClass
	AuthenticationModuleEntriesClassOnce sync.Once
)

func getauthenticationModuleEntriesClass() _authenticationModuleEntriesClass {
	AuthenticationModuleEntriesClassOnce.Do(func() {
		AuthenticationModuleEntriesClass = _authenticationModuleEntriesClass{objc.GetClass("authenticationModuleEntries")}
	})
	return AuthenticationModuleEntriesClass
}

type _authenticationModuleEntriesClass struct {
	class objc.Class
}

// An interface definition for the [authenticationModuleEntries] class.
type IauthenticationModuleEntries interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-c.ivar
type authenticationModuleEntries struct {
	objectivec.Object
}

// authenticationModuleEntriesFrom constructs a [authenticationModuleEntries] from an unsafe.Pointer.
func authenticationModuleEntriesFrom(ptr unsafe.Pointer) authenticationModuleEntries {
	return authenticationModuleEntries{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _authenticationModuleEntriesClass) Alloc() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _authenticationModuleEntriesClass) New() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authenticationModuleEntries) Init() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authenticationModuleEntries) Autorelease() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthenticationModuleEntries creates a new authenticationModuleEntries instance.
func NewauthenticationModuleEntries() authenticationModuleEntries {
	return getauthenticationModuleEntriesClass().New()
}




