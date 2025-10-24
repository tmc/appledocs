// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [sendPort] class.
var (
	SendPortClass     _sendPortClass
	SendPortClassOnce sync.Once
)

func getsendPortClass() _sendPortClass {
	SendPortClassOnce.Do(func() {
		SendPortClass = _sendPortClass{objc.GetClass("sendPort")}
	})
	return SendPortClass
}

type _sendPortClass struct {
	class objc.Class
}

// An interface definition for the [sendPort] class.
type IsendPort interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/sendPort-c.ivar
type sendPort struct {
	objectivec.Object
}

// sendPortFrom constructs a [sendPort] from an unsafe.Pointer.
func sendPortFrom(ptr unsafe.Pointer) sendPort {
	return sendPort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _sendPortClass) Alloc() sendPort {
	rv := objc.Send[sendPort](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _sendPortClass) New() sendPort {
	rv := objc.Send[sendPort](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ sendPort) Init() sendPort {
	rv := objc.Send[sendPort](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ sendPort) Autorelease() sendPort {
	rv := objc.Send[sendPort](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsendPort creates a new sendPort instance.
func NewsendPort() sendPort {
	return getsendPortClass().New()
}




