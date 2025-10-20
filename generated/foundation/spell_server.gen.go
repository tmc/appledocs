// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpellServer] class.
var (
	spellServerClass     _SpellServerClass
	spellServerClassOnce sync.Once
)

func getSpellServerClass() _SpellServerClass {
	spellServerClassOnce.Do(func() {
		spellServerClass = _SpellServerClass{objc.GetClass("NSSpellServer")}
	})
	return spellServerClass
}

type _SpellServerClass struct {
	class objc.Class
}

// An interface definition for the [SpellServer] class.
type ISpellServer interface {
	objectivec.IObject
}

// A server that your app uses to provide a spell checker service to other apps running in the system.
//
// A is an application that declares its availability in a standard way, so that any other applications that wish to use it can do so. If you build a spelling checker that makes use of the class and list it as an available service, then users of any application that makes use of or includes a Services menu will see your spelling checker as one of the available dictionaries.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpellServer
type SpellServer struct {
	objectivec.Object
}

// SpellServerFrom constructs a [SpellServer] from an unsafe.Pointer.
//
// A server that your app uses to provide a spell checker service to other apps running in the system.
func SpellServerFrom(ptr unsafe.Pointer) SpellServer {
	return SpellServer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpellServerClass) Alloc() SpellServer {
	rv := objc.Send[SpellServer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpellServerClass) New() SpellServer {
	rv := objc.Send[SpellServer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpellServer) Init() SpellServer {
	rv := objc.Send[SpellServer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpellServer) Autorelease() SpellServer {
	rv := objc.Send[SpellServer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpellServer creates a new SpellServer instance.
func NewSpellServer() SpellServer {
	return getSpellServerClass().New()
}




