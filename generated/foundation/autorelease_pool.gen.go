// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AutoreleasePool] class.
var (
	autoreleasePoolClass     _AutoreleasePoolClass
	autoreleasePoolClassOnce sync.Once
)

func getAutoreleasePoolClass() _AutoreleasePoolClass {
	autoreleasePoolClassOnce.Do(func() {
		autoreleasePoolClass = _AutoreleasePoolClass{objc.GetClass("NSAutoreleasePool")}
	})
	return autoreleasePoolClass
}

type _AutoreleasePoolClass struct {
	class objc.Class
}

// An interface definition for the [AutoreleasePool] class.
type IAutoreleasePool interface {
	objectivec.IObject
}

// An object that supports Cocoa’s reference-counted memory management system.
//
// An autorelease pool stores objects that are sent a message when the pool itself is drained. In a reference-counted environment (as opposed to one which uses garbage collection), an object contains objects that have received an message and when drained it sends a message to each of those objects. Thus, sending instead of to an object extends the lifetime of that object at least until the pool itself is drained (it may be longer if the object is subsequently retained). An object can be put into the same pool several times, in which case it receives a message for each time it was put into the pool. In a reference counted environment, Cocoa expects there to be an autorelease pool always available. If a pool is not available, autoreleased objects do not get released and you leak memory. In this situation, your program will typically log suitable warning messages. The Application Kit creates an autorelease pool on the main thread at the beginning of every cycle of the event loop, and drains it at the end, thereby releasing any autoreleased objects generated while processing an event. If you use the Application Kit, you therefore typically don’t have to create your own pools. If your application creates a lot of temporary autoreleased objects within the event loop, however, it may be beneficial to create “local” autorelease pools to help to minimize the peak memory footprint. You create an object with the usual and messages and dispose of it with (or —to understand the difference, see ). Since you cannot retain an autorelease pool (or autorelease it—see and ), draining a pool ultimately has the effect of deallocating it. You should always drain an autorelease pool in the same context (invocation of a method or function, or body of a loop) that it was created. See for more details. Each thread (including the main thread) maintains its own stack of objects (see ). As new pools are created, they get added to the top of the stack. When pools are deallocated, they are removed from the stack. Autoreleased objects are placed into the top autorelease pool for the current thread. When a thread terminates, it automatically drains all of the autorelease pools associated with itself.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool
type AutoreleasePool struct {
	objectivec.Object
}

// AutoreleasePoolFrom constructs a [AutoreleasePool] from an unsafe.Pointer.
//
// An object that supports Cocoa’s reference-counted memory management system.
func AutoreleasePoolFrom(ptr unsafe.Pointer) AutoreleasePool {
	return AutoreleasePool{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AutoreleasePoolClass) Alloc() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AutoreleasePoolClass) New() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutoreleasePool) Init() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutoreleasePool) Autorelease() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutoreleasePool creates a new AutoreleasePool instance.
func NewAutoreleasePool() AutoreleasePool {
	return getAutoreleasePoolClass().New()
}




