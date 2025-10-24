// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [session] class.
var (
	SessionClass     _sessionClass
	SessionClassOnce sync.Once
)

func getsessionClass() _sessionClass {
	SessionClassOnce.Do(func() {
		SessionClass = _sessionClass{objc.GetClass("session")}
	})
	return SessionClass
}

type _sessionClass struct {
	class objc.Class
}

// An interface definition for the [session] class.
type Isession interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/session
type session struct {
	objectivec.Object
}

// sessionFrom constructs a [session] from an unsafe.Pointer.
func sessionFrom(ptr unsafe.Pointer) session {
	return session{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _sessionClass) Alloc() session {
	rv := objc.Send[session](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _sessionClass) New() session {
	rv := objc.Send[session](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ session) Init() session {
	rv := objc.Send[session](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ session) Autorelease() session {
	rv := objc.Send[session](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newsession creates a new session instance.
func Newsession() session {
	return getsessionClass().New()
}




