// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ParameterKey] class.
var (
	ParameterKeyClass     _ParameterKeyClass
	ParameterKeyClassOnce sync.Once
)

func getParameterKeyClass() _ParameterKeyClass {
	ParameterKeyClassOnce.Do(func() {
		ParameterKeyClass = _ParameterKeyClass{objc.GetClass("MLParameterKey")}
	})
	return ParameterKeyClass
}

type _ParameterKeyClass struct {
	class objc.Class
}

// An interface definition for the [ParameterKey] class.
type IParameterKey interface {
	IKey
	// properties:
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	Parameters() IMLParameterKey
	SetParameters(value IMLParameterKey)
	// methods:
}

// The keys for the parameter dictionary in a model configuration or a model update context.
//
// Use an to retrieve a model’s parameter value using: The model’s method The dictionary of an The dictionary of an


// The keys for the parameter dictionary in a model configuration or a model update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey
type ParameterKey struct {
	Key
}

// ParameterKeyFrom constructs a [ParameterKey] from an unsafe.Pointer.
//
// The keys for the parameter dictionary in a model configuration or a model update context.
func ParameterKeyFrom(ptr unsafe.Pointer) ParameterKey {
	return ParameterKey{
		Key: KeyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _ParameterKeyClass) Alloc() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ParameterKeyClass) New() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterKey) Init() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterKey) Autorelease() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterKey creates a new ParameterKey instance.
func NewParameterKey() ParameterKey {
	return getParameterKeyClass().New()
}



// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (p_ ParameterKey) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](p_.ID, objc.Sel("configuration"))
	return rv
}


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (p_ ParameterKey) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguration:"), value)
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (p_ ParameterKey) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](p_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (p_ ParameterKey) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModelDescription:"), value)
}


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (p_ ParameterKey) Parameters() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("parameters"))
	return rv
}


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (p_ ParameterKey) SetParameters(value IMLParameterKey) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameters:"), value)
}



