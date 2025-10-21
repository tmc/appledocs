// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRDoorLockClusterDlCredential] class.
var (
	MTRDoorLockClusterDlCredentialClass     _MTRDoorLockClusterDlCredentialClass
	MTRDoorLockClusterDlCredentialClassOnce sync.Once
)

func getMTRDoorLockClusterDlCredentialClass() _MTRDoorLockClusterDlCredentialClass {
	MTRDoorLockClusterDlCredentialClassOnce.Do(func() {
		MTRDoorLockClusterDlCredentialClass = _MTRDoorLockClusterDlCredentialClass{objc.GetClass("MTRDoorLockClusterDlCredential")}
	})
	return MTRDoorLockClusterDlCredentialClass
}

type _MTRDoorLockClusterDlCredentialClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterDlCredential] class.
type IMTRDoorLockClusterDlCredential interface {
	IMTRDoorLockClusterCredentialStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDlCredential
type MTRDoorLockClusterDlCredential struct {
	MTRDoorLockClusterCredentialStruct
}

// MTRDoorLockClusterDlCredentialFrom constructs a [MTRDoorLockClusterDlCredential] from an unsafe.Pointer.
func MTRDoorLockClusterDlCredentialFrom(ptr unsafe.Pointer) MTRDoorLockClusterDlCredential {
	return MTRDoorLockClusterDlCredential{
		MTRDoorLockClusterCredentialStruct: MTRDoorLockClusterCredentialStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterDlCredentialClass) Alloc() MTRDoorLockClusterDlCredential {
	rv := objc.Send[MTRDoorLockClusterDlCredential](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterDlCredentialClass) New() MTRDoorLockClusterDlCredential {
	rv := objc.Send[MTRDoorLockClusterDlCredential](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterDlCredential) Init() MTRDoorLockClusterDlCredential {
	rv := objc.Send[MTRDoorLockClusterDlCredential](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterDlCredential) Autorelease() MTRDoorLockClusterDlCredential {
	rv := objc.Send[MTRDoorLockClusterDlCredential](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterDlCredential creates a new MTRDoorLockClusterDlCredential instance.
func NewMTRDoorLockClusterDlCredential() MTRDoorLockClusterDlCredential {
	return getMTRDoorLockClusterDlCredentialClass().New()
}




