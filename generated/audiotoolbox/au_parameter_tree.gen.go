// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AUParameterTree */


/* debug [class_header]: Header for AUParameterTree */
// The class instance for the [ParameterTree] class.
var (
	ParameterTreeClass     _ParameterTreeClass
	ParameterTreeClassOnce sync.Once
)

func getParameterTreeClass() _ParameterTreeClass {
	ParameterTreeClassOnce.Do(func() {
		ParameterTreeClass = _ParameterTreeClass{objc.GetClass("AUParameterTree")}
	})
	return ParameterTreeClass
}

type _ParameterTreeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ParameterTree */
// An interface definition for the [ParameterTree] class.
type IParameterTree interface {
	IParameterGroup
	
/* debug [class_interface_properties]: Properties for ParameterTree */
	// properties:
	ParameterTree() IAUParameterTree
	SetParameterTree(value IAUParameterTree)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ParameterTree */
	// methods:
	ParameterWithAddress(address ParameterAddress /* typedef */) IParameter
	ParameterWithIDScopeElement(paramID AudioUnitParameterID /* typedef */, scope AudioUnitScope /* typedef */, element AudioUnitElement /* typedef */) IParameter
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ParameterTree */
// Alloc allocates a new instance without initialization.
func (pc _ParameterTreeClass) Alloc() ParameterTree {
	rv := objc.Send[ParameterTree](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterTreeClass) New() ParameterTree {
	rv := objc.Send[ParameterTree](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterTree) Init() ParameterTree {
	rv := objc.Send[ParameterTree](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterTree) Autorelease() ParameterTree {
	rv := objc.Send[ParameterTree](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterTree creates a new ParameterTree instance.
func NewParameterTree() ParameterTree {
	return getParameterTreeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ParameterTree */
// An object that represents a top-level group node that contains all of an audio unit’s parameters.
//
// An audio unit’s parameters are organized into a tree containing groups and parameters (groups may be nested). The parameter tree is KVO-compliant. An audio unit may choose to dynamically rearrange the tree; when doing so, it must issue a KVO notification on the audio unit’s property.


// An object that represents a top-level group node that contains all of an audio unit’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree
type ParameterTree struct {
	ParameterGroup
}

// ParameterTreeFrom constructs a [ParameterTree] from an unsafe.Pointer.
//
// An object that represents a top-level group node that contains all of an audio unit’s parameters.
func ParameterTreeFrom(ptr unsafe.Pointer) ParameterTree {
	return ParameterTree{
		ParameterGroup: ParameterGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ParameterTree *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ParameterTree */

// Initializes a group as a copied instance of a template group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(fromTemplate:identifier:name:addressOffset:)
func (pc _ParameterTreeClass) CreateGroupFromTemplateIdentifierNameAddressOffset(templateGroup IAUParameterGroup, identifier objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, addressOffset ParameterAddress /* typedef */) IParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupFromTemplate:identifier:name:addressOffset:"), templateGroup, identifier, name, addressOffset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateGroupFromTemplateIdentifierNameAddressOffset) */


// Creates a parameter group object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(withIdentifier:name:children:)
func (pc _ParameterTreeClass) CreateGroupWithIdentifierNameChildren(identifier objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, children []ParameterNode) IParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupWithIdentifier:name:children:"), identifier, name, children)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateGroupWithIdentifierNameChildren) */


// Creates a template group which may be used as a prototype for further group instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroupTemplate(_:)
func (pc _ParameterTreeClass) CreateGroupTemplate(children []ParameterNode) IParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupTemplate:"), children)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateGroupTemplate) */


// Creates a single parameter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createParameter(withIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:)
func (pc _ParameterTreeClass) CreateParameterWithIdentifierNameAddressMinMaxUnitUnitNameFlagsValueStringsDependentParameters(identifier objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, address ParameterAddress /* typedef */, min foundation.Value, max foundation.Value, unit AudioUnitParameterUnit, unitName objc.IObject /* cross-framework: NSString */, flags AudioUnitParameterOptions, valueStrings []string, dependentParameters []foundation.Number) IParameter {
	rv := objc.Send[Parameter](objc.ID(pc.class), objc.Sel("createParameterWithIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:"), identifier, name, address, min, max, unit, unitName, flags, valueStrings, dependentParameters)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateParameterWithIdentifierNameAddressMinMaxUnitUnitNameFlagsValueStringsDependentParameters) */


// Creates a parameter tree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createTree(withChildren:)
func (pc _ParameterTreeClass) CreateTreeWithChildren(children []ParameterNode) IParameterTree {
	rv := objc.Send[ParameterTree](objc.ID(pc.class), objc.Sel("createTreeWithChildren:"), children)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateTreeWithChildren) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ParameterTree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ParameterTree */

// Searches the tree for a parameter with a specific address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withAddress:)
func (p_ ParameterTree) ParameterWithAddress(address ParameterAddress /* typedef */) IParameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("parameterWithAddress:"), address)
	return rv
}/* debug [instance_methods/method]: ParameterWithAddress */


// Searches the tree for a specific version 2 audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withID:scope:element:)
func (p_ ParameterTree) ParameterWithIDScopeElement(paramID AudioUnitParameterID /* typedef */, scope AudioUnitScope /* typedef */, element AudioUnitElement /* typedef */) IParameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("parameterWithID:scope:element:"), paramID, scope, element)
	return rv
}/* debug [instance_methods/method]: ParameterWithIDScopeElement */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ParameterTree */

// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/parametertree
func (p_ ParameterTree) ParameterTree() IAUParameterTree {
	rv := objc.Send[ParameterTree](p_.ID, objc.Sel("parameterTree"))
	return rv
}/* debug [instance_properties/getter]: parameterTree */


// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/parametertree
func (p_ ParameterTree) SetParameterTree(value IAUParameterTree) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameterTree:"), value)
}/* debug [instance_properties/setter]: parameterTree */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUParameterTree */





