// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [passkey] class.
var (
	PasskeyClass     _passkeyClass
	PasskeyClassOnce sync.Once
)

func getpasskeyClass() _passkeyClass {
	PasskeyClassOnce.Do(func() {
		PasskeyClass = _passkeyClass{objc.GetClass("passkey")}
	})
	return PasskeyClass
}

type _passkeyClass struct {
	class objc.Class
}

// An interface definition for the [passkey] class.
type Ipasskey interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/passkey-c.ivar
type passkey struct {
	objectivec.Object
}

// passkeyFrom constructs a [passkey] from an unsafe.Pointer.
func passkeyFrom(ptr unsafe.Pointer) passkey {
	return passkey{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _passkeyClass) Alloc() passkey {
	rv := objc.Send[passkey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _passkeyClass) New() passkey {
	rv := objc.Send[passkey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ passkey) Init() passkey {
	rv := objc.Send[passkey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ passkey) Autorelease() passkey {
	rv := objc.Send[passkey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// Newpasskey creates a new passkey instance.
func Newpasskey() passkey {
	return getpasskeyClass().New()
}




