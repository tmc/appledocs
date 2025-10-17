// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpellServer] class.
var SpellServerClass objc.Class

func init() {
	SpellServerClass = objc.GetClass("NSSpellServer")
}

type SpellServer struct {
	objc.ID
}

func SpellServerFrom(ptr unsafe.Pointer) SpellServer {
	return SpellServer{
		ID: objc.ID(ptr),
	}
}




