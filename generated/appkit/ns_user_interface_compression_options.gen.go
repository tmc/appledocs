// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUserInterfaceCompressionOptions */


/* debug [class_header]: Header for NSUserInterfaceCompressionOptions */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserInterfaceCompressionOptions */
// An interface definition for the [UserInterfaceCompressionOptions] class.
type IUserInterfaceCompressionOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserInterfaceCompressionOptions */
	// properties:
	Empty() bool
	IsEmpty() bool
	SetIsEmpty(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserInterfaceCompressionOptions */
	// methods:
	ContainsOptions(options IUserInterfaceCompressionOptions) bool
	IntersectsOptions(options IUserInterfaceCompressionOptions) bool
	OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions
	OptionsByAddingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserInterfaceCompressionOptions */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserInterfaceCompressionOptions */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserInterfaceCompressionOptions */

// Creates an option object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(coder:)
func NewUserInterfaceCompressionOptionsWithCoder(coder foundation.Coder) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUserInterfaceCompressionOptionsWithCoder */


// Creates an option object that represents the union of the supplied options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(options:)
func NewUserInterfaceCompressionOptionsWithCompressionOptions(options unsafe.Pointer) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithCompressionOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUserInterfaceCompressionOptionsWithCompressionOptions */


// Creates an option object with the given identifier string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(identifier:)
func NewUserInterfaceCompressionOptionsWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UserInterfaceCompressionOptions {
	instance := getUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[UserInterfaceCompressionOptions](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUserInterfaceCompressionOptionsWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserInterfaceCompressionOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserInterfaceCompressionOptions */

// An option specifying that views should no longer maintain equal width constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/breakEqualWidths
func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("breakEqualWidthsOption"))
	return rv
}/* debug [class_properties_class/property]: breakEqualWidthsOption */

// An option specifying that views should hide their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideImages
func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("hideImagesOption"))
	return rv
}/* debug [class_properties_class/property]: hideImagesOption */

// An option specifying that views should hide their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideText
func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("hideTextOption"))
	return rv
}/* debug [class_properties_class/property]: hideTextOption */

// An option specifying that views should reduce their internal metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/reduceMetrics
func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("reduceMetricsOption"))
	return rv
}/* debug [class_properties_class/property]: reduceMetricsOption */

// An option that represents the union of all standard compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/standardOptions
func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](objc.ID(uc.class), objc.Sel("standardOptions"))
	return rv
}/* debug [class_properties_class/property]: standardOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserInterfaceCompressionOptions */

// Determines whether the supplied compression options are all present in the current instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/contains(_:)
func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("containsOptions:"), options)
	return rv
}/* debug [instance_methods/method]: ContainsOptions */


// Determines whether the supplied compression options intersect with the current instance’s options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/intersects(_:)
func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("intersectsOptions:"), options)
	return rv
}/* debug [instance_methods/method]: IntersectsOptions */


// Creates a new compression options object with the supplied options removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/subtracting(_:)
func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("optionsByRemovingOptions:"), options)
	return rv
}/* debug [instance_methods/method]: OptionsByRemovingOptions */


// Creates a new compression options object representing the union with the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/union(_:)
func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("optionsByAddingOptions:"), options)
	return rv
}/* debug [instance_methods/method]: OptionsByAddingOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserInterfaceCompressionOptions */

// An option specifying that views should no longer maintain equal width constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/breakEqualWidths
func (u_ UserInterfaceCompressionOptions) BreakEqualWidthsOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("breakEqualWidthsOption"))
	return rv
}/* debug [instance_properties/getter]: breakEqualWidthsOption */


// An option specifying that views should hide their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideImages
func (u_ UserInterfaceCompressionOptions) HideImagesOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("hideImagesOption"))
	return rv
}/* debug [instance_properties/getter]: hideImagesOption */


// An option specifying that views should hide their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideText
func (u_ UserInterfaceCompressionOptions) HideTextOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("hideTextOption"))
	return rv
}/* debug [instance_properties/getter]: hideTextOption */


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/isEmpty
func (u_ UserInterfaceCompressionOptions) Empty() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("empty"))
	return rv
}/* debug [instance_properties/getter]: empty */


// An option specifying that views should reduce their internal metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/reduceMetrics
func (u_ UserInterfaceCompressionOptions) ReduceMetricsOption() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("reduceMetricsOption"))
	return rv
}/* debug [instance_properties/getter]: reduceMetricsOption */


// An option that represents the union of all standard compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/standardOptions
func (u_ UserInterfaceCompressionOptions) StandardOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](u_.ID, objc.Sel("standardOptions"))
	return rv
}/* debug [instance_properties/getter]: standardOptions */


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEmpty"))
	return rv
}/* debug [instance_properties/getter]: isEmpty */


// A Boolean value that denotes whether the option is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/isempty
func (u_ UserInterfaceCompressionOptions) SetIsEmpty(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEmpty:"), value)
}/* debug [instance_properties/setter]: isEmpty */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserInterfaceCompressionOptions */


