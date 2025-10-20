// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var SpellServerClass _SpellServerClass

func init() {
	SpellServerClass = _SpellServerClass{objc.GetClass("NSSpellServer")}
}

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




