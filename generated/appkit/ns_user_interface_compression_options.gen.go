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
	

	// properties:
	Empty() bool
	IsEmpty() bool
	SetIsEmpty(value bool)


	

	// methods:
	ContainsOptions(options IUserInterfaceCompressionOptions) bool
	IntersectsOptions(options IUserInterfaceCompressionOptions) bool
	OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions
	OptionsByAddingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions


}





// Alloc allocates a new instance without initialization.
func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that specifies how user interface elements resize themselves when space is constrained.
//
// An instance of contains zero or more options. Because a compression options object behaves like a set, you can use common operations like intersection, union and subtraction to interact with instances and their members. You can access system-defined options through the class methods detailed in Creating standard options, or you can create your own custom options with the initializer. To compare two different compression options objects, use the methods described in the Comparing compression options section.


// An object that specifies how user interface elements resize themselves when space is constrained.
//
// [Full Topic]
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






// Creates an option object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(coder:)
func NewUserInterfaceCompressionOptionsWithCoder(coder foundation.foundation.INSCoder) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates an option object that represents the union of the supplied options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(options:)
func NewUserInterfaceCompressionOptionsWithCompressionOptions(options unsafe.Pointer) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithCompressionOptions:"), options)
	rv.Autorelease()
	return rv
}


// Creates an option object with the given identifier string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(identifier:)
func NewUserInterfaceCompressionOptionsWithIdentifier(identifier foundation.foundation.INSString) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}












// An option specifying that views should no longer maintain equal width constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/breakEqualWidths
func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("breakEqualWidthsOption"))
	return rv
}

// An option specifying that views should hide their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideImages
func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("hideImagesOption"))
	return rv
}

// An option specifying that views should hide their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideText
func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("hideTextOption"))
	return rv
}

// An option specifying that views should reduce their internal metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/reduceMetrics
func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("reduceMetricsOption"))
	return rv
}

// An option that represents the union of all standard compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/standardOptions
func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("standardOptions"))
	return rv
}






// Determines whether the supplied compression options are all present in the current instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/contains(_:)
func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("containsOptions:"), options)
	return rv
}


// Determines whether the supplied compression options intersect with the current instance’s options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/intersects(_:)
func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("intersectsOptions:"), options)
	return rv
}


// Creates a new compression options object with the supplied options removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/subtracting(_:)
func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("optionsByRemovingOptions:"), options)
	return rv
}


// Creates a new compression options object representing the union with the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/union(_:)
func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("optionsByAddingOptions:"), options)
	return rv
}







// An option specifying that views should no longer maintain equal width constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/breakEqualWidths
func (u_ UserInterfaceCompressionOptions) BreakEqualWidthsOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("breakEqualWidthsOption"))
	return rv
}


// An option specifying that views should hide their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideImages
func (u_ UserInterfaceCompressionOptions) HideImagesOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("hideImagesOption"))
	return rv
}


// An option specifying that views should hide their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideText
func (u_ UserInterfaceCompressionOptions) HideTextOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("hideTextOption"))
	return rv
}


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/isEmpty
func (u_ UserInterfaceCompressionOptions) Empty() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("empty"))
	return rv
}


// An option specifying that views should reduce their internal metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/reduceMetrics
func (u_ UserInterfaceCompressionOptions) ReduceMetricsOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("reduceMetricsOption"))
	return rv
}


// An option that represents the union of all standard compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/standardOptions
func (u_ UserInterfaceCompressionOptions) StandardOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("standardOptions"))
	return rv
}


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEmpty"))
	return rv
}


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) SetIsEmpty(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEmpty:"), value)
}







