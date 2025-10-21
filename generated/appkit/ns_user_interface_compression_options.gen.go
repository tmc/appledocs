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
	UserInterfaceCompressionOptionsClass     _UserInterfaceCompressionOptionsClass
	UserInterfaceCompressionOptionsClassOnce sync.Once
)

func getUserInterfaceCompressionOptionsClass() _UserInterfaceCompressionOptionsClass {
	UserInterfaceCompressionOptionsClassOnce.Do(func() {
		UserInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}
	})
	return UserInterfaceCompressionOptionsClass
}

type _UserInterfaceCompressionOptionsClass struct {
	class objc.Class
}

// An interface definition for the [UserInterfaceCompressionOptions] class.
type IUserInterfaceCompressionOptions interface {
	objectivec.IObject
}

// An object that specifies how user interface elements resize themselves when space is constrained.
//
// An instance of contains zero or more options. Because a compression options object behaves like a set, you can use common operations like intersection, union and subtraction to interact with instances and their members. You can access system-defined options through the class methods detailed in Creating standard options, or you can create your own custom options with the initializer. To compare two different compression options objects, use the methods described in the Comparing compression options section.
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




// Creates an option object with the given identifier string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(identifier:)
func NewUserInterfaceCompressionOptionsWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEmpty"))
	return rv
}


// SetIsEmpty sets the value of the isEmpty property.
// A Boolean value that denotes whether the option is empty.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) SetIsEmpty(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEmpty:"), value)
}


