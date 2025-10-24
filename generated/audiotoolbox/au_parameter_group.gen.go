// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AUParameterGroup */


/* debug [class_header]: Header for AUParameterGroup */
// The class instance for the [ParameterGroup] class.
var (
	ParameterGroupClass     _ParameterGroupClass
	ParameterGroupClassOnce sync.Once
)

func getParameterGroupClass() _ParameterGroupClass {
	ParameterGroupClassOnce.Do(func() {
		ParameterGroupClass = _ParameterGroupClass{objc.GetClass("AUParameterGroup")}
	})
	return ParameterGroupClass
}

type _ParameterGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ParameterGroup */
// An interface definition for the [ParameterGroup] class.
type IParameterGroup interface {
	IParameterNode
	
/* debug [class_interface_properties]: Properties for ParameterGroup */
	// properties:
	AllParameters() []Parameter
	Children() []ParameterNode
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ParameterGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ParameterGroup */
// Alloc allocates a new instance without initialization.
func (pc _ParameterGroupClass) Alloc() ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterGroupClass) New() ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterGroup) Init() ParameterGroup {
	rv := objc.Send[ParameterGroup](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterGroup) Autorelease() ParameterGroup {
	rv := objc.Send[ParameterGroup](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterGroup creates a new ParameterGroup instance.
func NewParameterGroup() ParameterGroup {
	return getParameterGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ParameterGroup */
// A parameter group object represents a group of related audio unit parameters.
//
// A parameter group is KVC-compliant for its children. For example, calling the parameter group’s method, with a key value of , returns a child whose value matches that key.


// A parameter group object represents a group of related audio unit parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup
type ParameterGroup struct {
	ParameterNode
}

// ParameterGroupFrom constructs a [ParameterGroup] from an unsafe.Pointer.
//
// A parameter group object represents a group of related audio unit parameters.
func ParameterGroupFrom(ptr unsafe.Pointer) ParameterGroup {
	return ParameterGroup{
		ParameterNode: ParameterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ParameterGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ParameterGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ParameterGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ParameterGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ParameterGroup */

// Returns a flat array of all parameters in the group, including those in child groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/allParameters
func (p_ ParameterGroup) AllParameters() []Parameter {
	rv := objc.Send[[]Parameter](p_.ID, objc.Sel("allParameters"))
	return rv
}/* debug [instance_properties/getter]: allParameters */


// The group’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/children
func (p_ ParameterGroup) Children() []ParameterNode {
	rv := objc.Send[[]ParameterNode](p_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/identifier
func (p_ ParameterGroup) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/identifier
func (p_ ParameterGroup) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUParameterGroup */



