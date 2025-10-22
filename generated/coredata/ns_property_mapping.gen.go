// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyMapping] class.
var (
	PropertyMappingClass     _PropertyMappingClass
	PropertyMappingClassOnce sync.Once
)

func getPropertyMappingClass() _PropertyMappingClass {
	PropertyMappingClassOnce.Do(func() {
		PropertyMappingClass = _PropertyMappingClass{objc.GetClass("NSPropertyMapping")}
	})
	return PropertyMappingClass
}

type _PropertyMappingClass struct {
	class objc.Class
}

// An interface definition for the [PropertyMapping] class.
type IPropertyMapping interface {
	objectivec.IObject
	Name() string
	SetName(value string)
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
	ValueExpression() Expression
	SetValueExpression(value IExpression)
}

// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping
type PropertyMapping struct {
	objectivec.Object
}

// PropertyMappingFrom constructs a [PropertyMapping] from an unsafe.Pointer.
//
// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity.
func PropertyMappingFrom(ptr unsafe.Pointer) PropertyMapping {
	return PropertyMapping{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PropertyMappingClass) Alloc() PropertyMapping {
	rv := objc.Send[PropertyMapping](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyMappingClass) New() PropertyMapping {
	rv := objc.Send[PropertyMapping](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyMapping) Init() PropertyMapping {
	rv := objc.Send[PropertyMapping](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyMapping) Autorelease() PropertyMapping {
	rv := objc.Send[PropertyMapping](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyMapping creates a new PropertyMapping instance.
func NewPropertyMapping() PropertyMapping {
	return getPropertyMappingClass().New()
}


// The name of the property in the destination entity for the property mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/name
func (p_ PropertyMapping) Name() string {
	rv := objc.Send[string](p_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the property in the destination entity for the property mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/name
func (p_ PropertyMapping) SetName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), objc.String(value))
}

// The user info for the property mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/userInfo
func (p_ PropertyMapping) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// The user info for the property mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/userInfo
func (p_ PropertyMapping) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfo:"), value)
}

// The value expression for the property mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/valueExpression
func (p_ PropertyMapping) ValueExpression() Expression {
	rv := objc.Send[Expression](p_.ID, objc.Sel("valueExpression"))
	return rv
}


// SetValueExpression sets the value of the valueExpression property.
// The value expression for the property mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/valueExpression
func (p_ PropertyMapping) SetValueExpression(value IExpression) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValueExpression:"), value)
}



