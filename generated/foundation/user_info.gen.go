// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [userInfo] class.
var (
	UserInfoClass     _userInfoClass
	UserInfoClassOnce sync.Once
)

func getuserInfoClass() _userInfoClass {
	UserInfoClassOnce.Do(func() {
		UserInfoClass = _userInfoClass{objc.GetClass("userInfo")}
	})
	return UserInfoClass
}

type _userInfoClass struct {
	class objc.Class
}

// An interface definition for the [userInfo] class.
type IuserInfo interface {
	objectivec.IObject
}



//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-c.ivar

type userInfo struct {
	objectivec.Object
}

// userInfoFrom constructs a [userInfo] from an unsafe.Pointer.
func userInfoFrom(ptr unsafe.Pointer) userInfo {
	return userInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _userInfoClass) Alloc() userInfo {
	rv := objc.Send[userInfo](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _userInfoClass) New() userInfo {
	rv := objc.Send[userInfo](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ userInfo) Init() userInfo {
	rv := objc.Send[userInfo](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ userInfo) Autorelease() userInfo {
	rv := objc.Send[userInfo](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewuserInfo creates a new userInfo instance.
func NewuserInfo() userInfo {
	return getuserInfoClass().New()
}




