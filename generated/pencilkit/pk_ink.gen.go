// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKInk */


/* debug [class_header]: Header for PKInk */
// The class instance for the [Ink] class.
var (
	InkClass     _InkClass
	InkClassOnce sync.Once
)

func getInkClass() _InkClass {
	InkClassOnce.Do(func() {
		InkClass = _InkClass{objc.GetClass("PKInk")}
	})
	return InkClass
}

type _InkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Ink */
// An interface definition for the [Ink] class.
type IInk interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Ink */
	// properties:
	Color() appkit.Color
	InkType() InkType /* not a class type */
	RequiredContentVersion() ContentVersion
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Ink */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Ink */
// Alloc allocates a new instance without initialization.
func (ic _InkClass) Alloc() Ink {
	rv := objc.Send[Ink](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InkClass) New() Ink {
	rv := objc.Send[Ink](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Ink) Init() Ink {
	rv := objc.Send[Ink](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Ink) Autorelease() Ink {
	rv := objc.Send[Ink](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInk creates a new Ink instance.
func NewInk() Ink {
	return getInkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Ink */
// Provides a description of the creation and rendering of marks on a canvas.


// Provides a description of the creation and rendering of marks on a canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference
type Ink struct {
	objectivec.Object
}

// InkFrom constructs a [Ink] from an unsafe.Pointer.
//
// Provides a description of the creation and rendering of marks on a canvas.
func InkFrom(ptr unsafe.Pointer) Ink {
	return Ink{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Ink */

// Create a new ink, specifying its type, color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/init(inkType:color:)
func NewInkWithInkTypeColor(type_ InkType /* not a class type */, color appkit.Color) Ink {
	instance := getInkClass().Alloc()
	rv := objc.Send[Ink](instance.ID, objc.Sel("initWithInkType:color:"), type_, color)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInkWithInkTypeColor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Ink */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Ink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Ink */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Ink */

// The base color for this ink.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/color
func (i_ Ink) Color() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The type of ink, such as pen or pencil, as defined in the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/inkType
func (i_ Ink) InkType() InkType /* not a class type */ {
	rv := objc.Send[InkType](i_.ID, objc.Sel("inkType"))
	return rv
}/* debug [instance_properties/getter]: inkType */


// The version of PencilKit necessary to use the ink.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/requiredContentVersion
func (i_ Ink) RequiredContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](i_.ID, objc.Sel("requiredContentVersion"))
	return rv
}/* debug [instance_properties/getter]: requiredContentVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKInk */


