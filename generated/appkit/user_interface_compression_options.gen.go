
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserInterfaceCompressionOptions] class.
var UserInterfaceCompressionOptionsClass _UserInterfaceCompressionOptionsClass

func init() {
	UserInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}
}

type _UserInterfaceCompressionOptionsClass struct {
	objc.Class
}

// An interface definition for the [UserInterfaceCompressionOptions] class.
type IUserInterfaceCompressionOptions interface {
	ID() objc.ID
}

type UserInterfaceCompressionOptions struct {
	id objc.ID
}

func UserInterfaceCompressionOptionsFrom(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ UserInterfaceCompressionOptions) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewUserInterfaceCompressionOptions creates and returns a new initialized instance.
func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

// Init initializes the instance.
func (u_ UserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID(), selInit)
	return rv
}
