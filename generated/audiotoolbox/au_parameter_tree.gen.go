// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ParameterTree] class.
type IParameterTree interface {
	IParameterGroup
	ParameterWithAddress(address IParameterAddress) Parameter
	ParameterWithIDScopeElement(paramID IAudioUnitParameterID, scope IAudioUnitScope, element IAudioUnitElement) Parameter
	ParameterTree() AUParameterTree
	SetParameterTree(value IAUParameterTree)
}

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

// Alloc allocates a new instance without initialization.
func (pc _ParameterTreeClass) Alloc() ParameterTree {
	rv := objc.Send[ParameterTree](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a group as a copied instance of a template group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(fromTemplate:identifier:name:addressOffset:)

func (pc _ParameterTreeClass) CreateGroupFromTemplateIdentifierNameAddressOffset(templateGroup IAUParameterGroup, identifier string, name string, addressOffset IParameterAddress) ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupFromTemplate:identifier:name:addressOffset:"), templateGroup, objc.String(identifier), objc.String(name), addressOffset)
	return rv
}


// Creates a parameter group object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(withIdentifier:name:children:)

func (pc _ParameterTreeClass) CreateGroupWithIdentifierNameChildren(identifier string, name string, children []ParameterNode) ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupWithIdentifier:name:children:"), objc.String(identifier), objc.String(name), children)
	return rv
}


// Creates a template group which may be used as a prototype for further group instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroupTemplate(_:)

func (pc _ParameterTreeClass) CreateGroupTemplate(children []ParameterNode) ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("createGroupTemplate:"), children)
	return rv
}


// Creates a single parameter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createParameter(withIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:)

func (pc _ParameterTreeClass) CreateParameterWithIdentifierNameAddressMinMaxUnitUnitNameFlagsValueStringsDependentParameters(identifier string, name string, address IParameterAddress, min IValue, max IValue, unit IAudioUnitParameterUnit, unitName string, flags AudioUnitParameterOptions, valueStrings []string, dependentParameters []foundation.INumber) Parameter {
	rv := objc.Send[Parameter](objc.ID(pc.class), objc.Sel("createParameterWithIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:"), objc.String(identifier), objc.String(name), address, min, max, unit, objc.String(unitName), flags, valueStrings, dependentParameters)
	return rv
}


// Creates a parameter tree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createTree(withChildren:)

func (pc _ParameterTreeClass) CreateTreeWithChildren(children []ParameterNode) ParameterTree {
	rv := objc.Send[ParameterTree](objc.ID(pc.class), objc.Sel("createTreeWithChildren:"), children)
	return rv
}



// Searches the tree for a parameter with a specific address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withAddress:)

func (p_ ParameterTree) ParameterWithAddress(address IParameterAddress) Parameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("parameterWithAddress:"), address)
	return rv
}



// Searches the tree for a specific version 2 audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withID:scope:element:)

func (p_ ParameterTree) ParameterWithIDScopeElement(paramID IAudioUnitParameterID, scope IAudioUnitScope, element IAudioUnitElement) Parameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("parameterWithID:scope:element:"), paramID, scope, element)
	return rv
}


// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/parametertree

func (p_ ParameterTree) ParameterTree() AUParameterTree {
	rv := objc.Send[AUParameterTree](p_.ID, objc.Sel("parameterTree"))
	return rv
}


// An audio unit’s parameters, organized in a tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/parametertree

func (p_ ParameterTree) SetParameterTree(value IAUParameterTree) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameterTree:"), value)
}



