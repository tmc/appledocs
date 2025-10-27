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
	ScopedTo(scope foundation.foundation.INSString) IParameterKey


}





// Alloc allocates a new instance without initialization.
func (pc _ParameterKeyClass) Alloc() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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















// The key you use to access the Adam optimizer’s first beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta1
func (pc _ParameterKeyClass) Beta1() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("beta1"))
	return rv
}

// The key you use to access the Adam optimizer’s second beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta2
func (pc _ParameterKeyClass) Beta2() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("beta2"))
	return rv
}

// The key you use to access the biases of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/biases
func (pc _ParameterKeyClass) Biases() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("biases"))
	return rv
}

// The key you use to access the optimizer’s epochs parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/epochs
func (pc _ParameterKeyClass) Epochs() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("epochs"))
	return rv
}

// The key you use to access the Adam optimizer’s epsilon parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/eps
func (pc _ParameterKeyClass) Eps() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("eps"))
	return rv
}

// The key you use to access the optimizer’s learning rate parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/learningRate
func (pc _ParameterKeyClass) LearningRate() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("learningRate"))
	return rv
}

// The key you use to access the linked model’s filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelFileName
func (pc _ParameterKeyClass) LinkedModelFileName() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("linkedModelFileName"))
	return rv
}

// The key you use to access the linked model’s search path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelSearchPath
func (pc _ParameterKeyClass) LinkedModelSearchPath() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("linkedModelSearchPath"))
	return rv
}

// The key you use to access the optimizer’s mini batch-size parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/miniBatchSize
func (pc _ParameterKeyClass) MiniBatchSize() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("miniBatchSize"))
	return rv
}

// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/momentum
func (pc _ParameterKeyClass) Momentum() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("momentum"))
	return rv
}

// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/numberOfNeighbors
func (pc _ParameterKeyClass) NumberOfNeighbors() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("numberOfNeighbors"))
	return rv
}

// The key you use to access the seed parameter that initializes the random number generator for the shuffle option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/seed
func (pc _ParameterKeyClass) Seed() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("seed"))
	return rv
}

// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/shuffle
func (pc _ParameterKeyClass) Shuffle() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("shuffle"))
	return rv
}

// The key you use to access the weights of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/weights
func (pc _ParameterKeyClass) Weights() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("weights"))
	return rv
}






// Creates a copy of a parameter key and adds the scope to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/scoped(to:)
func (p_ ParameterKey) ScopedTo(scope foundation.foundation.INSString) IParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("scopedTo:"), scope)
	return rv
}







// The key you use to access the Adam optimizer’s first beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta1
func (p_ ParameterKey) Beta1() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("beta1"))
	return rv
}


// The key you use to access the Adam optimizer’s second beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta2
func (p_ ParameterKey) Beta2() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("beta2"))
	return rv
}


// The key you use to access the biases of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/biases
func (p_ ParameterKey) Biases() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("biases"))
	return rv
}


// The key you use to access the optimizer’s epochs parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/epochs
func (p_ ParameterKey) Epochs() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("epochs"))
	return rv
}


// The key you use to access the Adam optimizer’s epsilon parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/eps
func (p_ ParameterKey) Eps() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("eps"))
	return rv
}


// The key you use to access the optimizer’s learning rate parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/learningRate
func (p_ ParameterKey) LearningRate() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("learningRate"))
	return rv
}


// The key you use to access the linked model’s filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelFileName
func (p_ ParameterKey) LinkedModelFileName() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("linkedModelFileName"))
	return rv
}


// The key you use to access the linked model’s search path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelSearchPath
func (p_ ParameterKey) LinkedModelSearchPath() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("linkedModelSearchPath"))
	return rv
}


// The key you use to access the optimizer’s mini batch-size parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/miniBatchSize
func (p_ ParameterKey) MiniBatchSize() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("miniBatchSize"))
	return rv
}


// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/momentum
func (p_ ParameterKey) Momentum() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("momentum"))
	return rv
}


// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/numberOfNeighbors
func (p_ ParameterKey) NumberOfNeighbors() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("numberOfNeighbors"))
	return rv
}


// The key you use to access the seed parameter that initializes the random number generator for the shuffle option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/seed
func (p_ ParameterKey) Seed() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("seed"))
	return rv
}


// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/shuffle
func (p_ ParameterKey) Shuffle() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("shuffle"))
	return rv
}


// The key you use to access the weights of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/weights
func (p_ ParameterKey) Weights() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("weights"))
	return rv
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








