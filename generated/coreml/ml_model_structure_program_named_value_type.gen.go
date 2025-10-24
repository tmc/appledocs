// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramNamedValueType */


/* debug [class_header]: Header for MLModelStructureProgramNamedValueType */
// The class instance for the [ModelStructureProgramNamedValueType] class.
var (
	ModelStructureProgramNamedValueTypeClass     _ModelStructureProgramNamedValueTypeClass
	ModelStructureProgramNamedValueTypeClassOnce sync.Once
)

func getModelStructureProgramNamedValueTypeClass() _ModelStructureProgramNamedValueTypeClass {
	ModelStructureProgramNamedValueTypeClassOnce.Do(func() {
		ModelStructureProgramNamedValueTypeClass = _ModelStructureProgramNamedValueTypeClass{objc.GetClass("MLModelStructureProgramNamedValueType")}
	})
	return ModelStructureProgramNamedValueTypeClass
}

type _ModelStructureProgramNamedValueTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramNamedValueType */
// An interface definition for the [ModelStructureProgramNamedValueType] class.
type IModelStructureProgramNamedValueType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramNamedValueType */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Type() IMLModelStructureProgramValueType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramNamedValueType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramNamedValueType */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramNamedValueTypeClass) Alloc() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureProgramNamedValueTypeClass) New() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramNamedValueType) Init() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramNamedValueType) Autorelease() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramNamedValueType creates a new ModelStructureProgramNamedValueType instance.
func NewModelStructureProgramNamedValueType() ModelStructureProgramNamedValueType {
	return getModelStructureProgramNamedValueTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramNamedValueType */
// A class representing a named value type in a Program.


// A class representing a named value type in a Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type ModelStructureProgramNamedValueType struct {
	objectivec.Object
}

// ModelStructureProgramNamedValueTypeFrom constructs a [ModelStructureProgramNamedValueType] from an unsafe.Pointer.
//
// A class representing a named value type in a Program.
func ModelStructureProgramNamedValueTypeFrom(ptr unsafe.Pointer) ModelStructureProgramNamedValueType {
	return ModelStructureProgramNamedValueType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramNamedValueType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramNamedValueType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramNamedValueType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramNamedValueType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramNamedValueType */

// The name of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/name
func (m_ ModelStructureProgramNamedValueType) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The type of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/type
func (m_ ModelStructureProgramNamedValueType) Type() IMLModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramNamedValueType */



