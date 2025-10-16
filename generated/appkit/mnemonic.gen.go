
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mnemonic] class.
var mnemonicClass _mnemonicClass

func init() {
	mnemonicClass = _mnemonicClass{objc.GetClass("mnemonic")}
}

type _mnemonicClass struct {
	objc.Class
}

// An interface definition for the [mnemonic] class.
type Imnemonic interface {
	ID() objc.ID
}

type mnemonic struct {
	id objc.ID
}

func mnemonicFrom(ptr unsafe.Pointer) mnemonic {
	return mnemonic{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mnemonic) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mnemonicClass) Alloc() mnemonic {
	rv := objc.Send[mnemonic](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mnemonicClass) New() mnemonic {
	rv := objc.Send[mnemonic](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmnemonic creates and returns a new initialized instance.
func Newmnemonic() mnemonic {
	return mnemonicClass.New()
}

// Init initializes the instance.
func (m_ mnemonic) Init() mnemonic {
	rv := objc.Send[mnemonic](m_.ID(), selInit)
	return rv
}
