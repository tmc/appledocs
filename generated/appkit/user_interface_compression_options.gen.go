// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserInterfaceCompressionOptions] class.
var (
	userInterfaceCompressionOptionsClass     _UserInterfaceCompressionOptionsClass
	userInterfaceCompressionOptionsClassOnce sync.Once
)

func getUserInterfaceCompressionOptionsClass() _UserInterfaceCompressionOptionsClass {
	userInterfaceCompressionOptionsClassOnce.Do(func() {
		userInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}
	})
	return userInterfaceCompressionOptionsClass
}

type _UserInterfaceCompressionOptionsClass struct {
	class objc.Class
}

// An interface definition for the [UserInterfaceCompressionOptions] class.
type IUserInterfaceCompressionOptions interface {
	objectivec.IObject
}

// An object that specifies how user interface elements resize themselves when space is constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions

type UserInterfaceCompressionOptions struct {
	objectivec.Object
}

// UserInterfaceCompressionOptionsFrom constructs a [UserInterfaceCompressionOptions] from an unsafe.Pointer.
//
// An object that specifies how user interface elements resize themselves when space is constrained.
func UserInterfaceCompressionOptionsFrom(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserInterfaceCompressionOptions) Autorelease() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserInterfaceCompressionOptions creates a new UserInterfaceCompressionOptions instance.
func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return getUserInterfaceCompressionOptionsClass().New()
}




