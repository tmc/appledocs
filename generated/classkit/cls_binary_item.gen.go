// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CLSBinaryItem */


/* debug [class_header]: Header for CLSBinaryItem */
// The class instance for the [SBinaryItem] class.
var (
	SBinaryItemClass     _SBinaryItemClass
	SBinaryItemClassOnce sync.Once
)

func getSBinaryItemClass() _SBinaryItemClass {
	SBinaryItemClassOnce.Do(func() {
		SBinaryItemClass = _SBinaryItemClass{objc.GetClass("CLSBinaryItem")}
	})
	return SBinaryItemClass
}

type _SBinaryItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SBinaryItem */
// An interface definition for the [SBinaryItem] class.
type ISBinaryItem interface {
	ISActivityItem
	
/* debug [class_interface_properties]: Properties for SBinaryItem */
	// properties:
	Value() bool
	SetValue(value bool)
	ValueType() SBinaryValueType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SBinaryItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SBinaryItem */
// Alloc allocates a new instance without initialization.
func (sc _SBinaryItemClass) Alloc() SBinaryItem {
	rv := objc.Send[SBinaryItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SBinaryItemClass) New() SBinaryItem {
	rv := objc.Send[SBinaryItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SBinaryItem) Init() SBinaryItem {
	rv := objc.Send[SBinaryItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SBinaryItem) Autorelease() SBinaryItem {
	rv := objc.Send[SBinaryItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBinaryItem creates a new SBinaryItem instance.
func NewSBinaryItem() SBinaryItem {
	return getSBinaryItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SBinaryItem */
// Activity information that is true or false, pass or fail, yes or no.
//
// Use an activity item of this type to indicate a binary condition, such as whether a student passed a test or failed it. Set the property to specify how the binary condition should be reported to a teacher.


// Activity information that is true or false, pass or fail, yes or no.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryItem
type SBinaryItem struct {
	SActivityItem
}

// SBinaryItemFrom constructs a [SBinaryItem] from an unsafe.Pointer.
//
// Activity information that is true or false, pass or fail, yes or no.
func SBinaryItemFrom(ptr unsafe.Pointer) SBinaryItem {
	return SBinaryItem{
		SActivityItem: SActivityItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SBinaryItem */

// Initializes a new binary activity item of the given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryItem/init(identifier:title:type:)
func NewSBinaryItemWithIdentifierTitleType(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, valueType SBinaryValueType) SBinaryItem {
	instance := getSBinaryItemClass().Alloc()
	rv := objc.Send[SBinaryItem](instance.ID, objc.Sel("initWithIdentifier:title:type:"), identifier, title, valueType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSBinaryItemWithIdentifierTitleType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SBinaryItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SBinaryItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SBinaryItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SBinaryItem */

// The value that the binary activity item takes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryItem/value
func (s_ SBinaryItem) Value() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The value that the binary activity item takes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryItem/value
func (s_ SBinaryItem) SetValue(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// The kind of outcome that the binary activity item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryItem/valueType
func (s_ SBinaryItem) ValueType() SBinaryValueType {
	rv := objc.Send[SBinaryValueType](s_.ID, objc.Sel("valueType"))
	return rv
}/* debug [instance_properties/getter]: valueType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSBinaryItem */


