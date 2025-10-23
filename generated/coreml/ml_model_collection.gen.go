// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelCollection] class.
var (
	ModelCollectionClass     _ModelCollectionClass
	ModelCollectionClassOnce sync.Once
)

func getModelCollectionClass() _ModelCollectionClass {
	ModelCollectionClassOnce.Do(func() {
		ModelCollectionClass = _ModelCollectionClass{objc.GetClass("MLModelCollection")}
	})
	return ModelCollectionClass
}

type _ModelCollectionClass struct {
	class objc.Class
}

// An interface definition for the [ModelCollection] class.
type IModelCollection interface {
	objectivec.IObject
	DeploymentID() string
	Entries() foundation.IDictionary
	Identifier() string
}

// A set of Core ML models from a model deployment.
//
// Use a model collection to access the models from a Core ML Model Deployment. For example, you can use a model collection to replace one or more of your app’s built-in models with a newer version. To access the newest model collection from a deployment, call the type method. Your app can also get a notification when Core ML receives an update to a model collection (see ).


// A set of Core ML models from a model deployment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection
type ModelCollection struct {
	objectivec.Object
}

// ModelCollectionFrom constructs a [ModelCollection] from an unsafe.Pointer.
//
// A set of Core ML models from a model deployment.
func ModelCollectionFrom(ptr unsafe.Pointer) ModelCollection {
	return ModelCollection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelCollectionClass) Alloc() ModelCollection {
	rv := objc.Send[ModelCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelCollectionClass) New() ModelCollection {
	rv := objc.Send[ModelCollection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelCollection) Init() ModelCollection {
	rv := objc.Send[ModelCollection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelCollection) Autorelease() ModelCollection {
	rv := objc.Send[ModelCollection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelCollection creates a new ModelCollection instance.
func NewModelCollection() ModelCollection {
	return getModelCollectionClass().New()
}



// Requests access to a model collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/beginAccessingModelCollectionWithIdentifier:completionHandler:
func (mc _ModelCollectionClass) BeginAccessingModelCollectionWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](objc.ID(mc.class), objc.Sel("beginAccessingModelCollectionWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
	return rv
}


// Terminates access to a model collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/endAccessing(identifier:)
func (mc _ModelCollectionClass) EndAccessingModelCollectionWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("endAccessingModelCollectionWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}


// The unique identifier of the model collection’s deployment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/deploymentID
func (m_ ModelCollection) DeploymentID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("deploymentID"))
	return rv
}


// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/entries
func (m_ ModelCollection) Entries() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("entries"))
	return rv
}


// The name of the model collection, unique to the development team.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/identifier
func (m_ ModelCollection) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}



