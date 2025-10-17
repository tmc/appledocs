// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpellServer] class.
var spellServerClass = _SpellServerClass{objc.GetClass("NSSpellServer")}

type _SpellServerClass struct {
	class objc.Class
}

// An interface definition for the [SpellServer] class.
type ISpellServer interface {
	objectivec.IObject
}

// A server that your app uses to provide a spell checker service to other apps running in the system. [Full Topic]
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



