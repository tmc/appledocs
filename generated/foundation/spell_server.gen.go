// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpellServer] class.
var SpellServerClass = _SpellServerClass{objc.GetClass("NSSpellServer")}

type _SpellServerClass struct {
	class objc.Class
}

type SpellServer struct {
	objc.ID
}

func SpellServerFrom(ptr unsafe.Pointer) SpellServer {
	return SpellServer{
		ID: objc.ID(ptr),
	}
}




