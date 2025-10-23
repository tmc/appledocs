// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Connection] class.
var (
	ConnectionClass     _ConnectionClass
	ConnectionClassOnce sync.Once
)

func getConnectionClass() _ConnectionClass {
	ConnectionClassOnce.Do(func() {
		ConnectionClass = _ConnectionClass{objc.GetClass("NSConnection")}
	})
	return ConnectionClass
}

type _ConnectionClass struct {
	class objc.Class
}

// An interface definition for the [Connection] class.
type IConnection interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
//
// Connection objects form the backbone of the distributed objects mechanism and normally operate in the background. You use the methods of explicitly when vending an object to other applications, when accessing such a vended object through a proxy, and when altering default communication parameters. At other times, you simply interact with a vended object or its proxy. A single connection object may be shared by multiple threads and used to access a vended object.


// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection
type Connection struct {
	objectivec.Object
}

// ConnectionFrom constructs a [Connection] from an unsafe.Pointer.
//
// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
func ConnectionFrom(ptr unsafe.Pointer) Connection {
	return Connection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConnectionClass) Alloc() Connection {
	rv := objc.Send[Connection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConnectionClass) New() Connection {
	rv := objc.Send[Connection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Connection) Init() Connection {
	rv := objc.Send[Connection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Connection) Autorelease() Connection {
	rv := objc.Send[Connection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConnection creates a new Connection instance.
func NewConnection() Connection {
	return getConnectionClass().New()
}




