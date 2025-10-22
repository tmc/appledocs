// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ParameterGroup] class.
type IParameterGroup interface {
	IParameterNode
	AllParameters() []Parameter
	Children() []ParameterNode
	Identifier() string
	SetIdentifier(value string)
}

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

// Alloc allocates a new instance without initialization.
func (pc _ParameterGroupClass) Alloc() ParameterGroup {
	rv := objc.Send[ParameterGroup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a flat array of all parameters in the group, including those in child groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/allParameters

func (p_ ParameterGroup) AllParameters() []Parameter {
	rv := objc.Send[[]Parameter](p_.ID, objc.Sel("allParameters"))
	return rv
}


// The group’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/children

func (p_ ParameterGroup) Children() []ParameterNode {
	rv := objc.Send[[]ParameterNode](p_.ID, objc.Sel("children"))
	return rv
}


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/identifier

func (p_ ParameterGroup) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/identifier

func (p_ ParameterGroup) SetIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}



