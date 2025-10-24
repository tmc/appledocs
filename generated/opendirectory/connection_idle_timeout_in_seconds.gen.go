// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [connectionIdleTimeoutInSeconds] class.
var (
	ConnectionIdleTimeoutInSecondsClass     _connectionIdleTimeoutInSecondsClass
	ConnectionIdleTimeoutInSecondsClassOnce sync.Once
)

func getconnectionIdleTimeoutInSecondsClass() _connectionIdleTimeoutInSecondsClass {
	ConnectionIdleTimeoutInSecondsClassOnce.Do(func() {
		ConnectionIdleTimeoutInSecondsClass = _connectionIdleTimeoutInSecondsClass{objc.GetClass("connectionIdleTimeoutInSeconds")}
	})
	return ConnectionIdleTimeoutInSecondsClass
}

type _connectionIdleTimeoutInSecondsClass struct {
	class objc.Class
}

// An interface definition for the [connectionIdleTimeoutInSeconds] class.
type IconnectionIdleTimeoutInSeconds interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-c.ivar
type connectionIdleTimeoutInSeconds struct {
	objectivec.Object
}

// connectionIdleTimeoutInSecondsFrom constructs a [connectionIdleTimeoutInSeconds] from an unsafe.Pointer.
func connectionIdleTimeoutInSecondsFrom(ptr unsafe.Pointer) connectionIdleTimeoutInSeconds {
	return connectionIdleTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _connectionIdleTimeoutInSecondsClass) Alloc() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _connectionIdleTimeoutInSecondsClass) New() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ connectionIdleTimeoutInSeconds) Init() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ connectionIdleTimeoutInSeconds) Autorelease() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewconnectionIdleTimeoutInSeconds creates a new connectionIdleTimeoutInSeconds instance.
func NewconnectionIdleTimeoutInSeconds() connectionIdleTimeoutInSeconds {
	return getconnectionIdleTimeoutInSecondsClass().New()
}




