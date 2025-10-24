// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLType */


/* debug [class_header]: Header for MTLType */
// The class instance for the [Type] class.
var (
	TypeClass     _TypeClass
	TypeClassOnce sync.Once
)

func getTypeClass() _TypeClass {
	TypeClassOnce.Do(func() {
		TypeClass = _TypeClass{objc.GetClass("MTLType")}
	})
	return TypeClass
}

type _TypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Type */
// An interface definition for the [Type] class.
type IType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Type */
	// properties:
	DataType() DataType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Type */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Type */
// Alloc allocates a new instance without initialization.
func (tc _TypeClass) Alloc() Type {
	rv := objc.Send[Type](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TypeClass) New() Type {
	rv := objc.Send[Type](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Type) Init() Type {
	rv := objc.Send[Type](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Type) Autorelease() Type {
	rv := objc.Send[Type](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewType creates a new Type instance.
func NewType() Type {
	return getTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Type */
// A description of a data type.


// A description of a data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLType
type Type struct {
	objectivec.Object
}

// TypeFrom constructs a [Type] from an unsafe.Pointer.
//
// A description of a data type.
func TypeFrom(ptr unsafe.Pointer) Type {
	return Type{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Type *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Type */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Type */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Type */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Type */

// The data type of the function argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLType/dataType
func (t_ Type) DataType() DataType {
	rv := objc.Send[DataType](t_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLType */



