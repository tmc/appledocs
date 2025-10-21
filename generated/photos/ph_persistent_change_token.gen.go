// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHPersistentChangeToken] class.
var (
	PHPersistentChangeTokenClass     _PHPersistentChangeTokenClass
	PHPersistentChangeTokenClassOnce sync.Once
)

func getPHPersistentChangeTokenClass() _PHPersistentChangeTokenClass {
	PHPersistentChangeTokenClassOnce.Do(func() {
		PHPersistentChangeTokenClass = _PHPersistentChangeTokenClass{objc.GetClass("PHPersistentChangeToken")}
	})
	return PHPersistentChangeTokenClass
}

type _PHPersistentChangeTokenClass struct {
	class objc.Class
}

// An interface definition for the [PHPersistentChangeToken] class.
type IPHPersistentChangeToken interface {
	objectivec.IObject
}

// An opaque object that tracks the state of the Photos library between runs, and that you can copy and serialize for future use.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPersistentChangeToken
type PHPersistentChangeToken struct {
	objectivec.Object
}

// PHPersistentChangeTokenFrom constructs a [PHPersistentChangeToken] from an unsafe.Pointer.
//
// An opaque object that tracks the state of the Photos library between runs, and that you can copy and serialize for future use.
func PHPersistentChangeTokenFrom(ptr unsafe.Pointer) PHPersistentChangeToken {
	return PHPersistentChangeToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPersistentChangeTokenClass) Alloc() PHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPersistentChangeTokenClass) New() PHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPersistentChangeToken) Init() PHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPersistentChangeToken) Autorelease() PHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPersistentChangeToken creates a new PHPersistentChangeToken instance.
func NewPHPersistentChangeToken() PHPersistentChangeToken {
	return getPHPersistentChangeTokenClass().New()
}




