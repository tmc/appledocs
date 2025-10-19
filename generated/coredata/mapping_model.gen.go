// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MappingModel] class.
var mappingModelClass = _MappingModelClass{objc.GetClass("NSMappingModel")}

type _MappingModelClass struct {
	class objc.Class
}

// An interface definition for the [MappingModel] class.
type IMappingModel interface {
	objectivec.IObject
}

// A model instance that specifies how to map a model from a source to a destination managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel

type MappingModel struct {
	objectivec.Object
}

// MappingModelFrom constructs a [MappingModel] from an unsafe.Pointer.
//
// A model instance that specifies how to map a model from a source to a destination managed object model.
func MappingModelFrom(ptr unsafe.Pointer) MappingModel {
	return MappingModel{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MappingModelClass) Alloc() MappingModel {
	rv := objc.Send[MappingModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MappingModelClass) New() MappingModel {
	rv := objc.Send[MappingModel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MappingModel) Init() MappingModel {
	rv := objc.Send[MappingModel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MappingModel) Autorelease() MappingModel {
	rv := objc.Send[MappingModel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMappingModel creates a new MappingModel instance.
func NewMappingModel() MappingModel {
	return mappingModelClass.New()
}


// Returns the mapping model that will translate data from the source to the destination model. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(from:forSourceModel:destinationModel:)
func NewMappingModelFromBundlesForSourceModelDestinationModel(bundles unsafe.Pointer, sourceModel unsafe.Pointer, destinationModel unsafe.Pointer) MappingModel {
	rv := objc.Send[MappingModel](objc.ID(mappingModelClass.class), objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), bundles, sourceModel, destinationModel)
	rv.Autorelease()
	return rv
}


// Returns a newly created mapping model that will migrate data from the source to the destination model. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/inferredMappingModel(forSourceModel:destinationModel:)
func (mc _MappingModelClass) InferredMappingModelForSourceModelDestinationModelError(sourceModel unsafe.Pointer, destinationModel unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("inferredMappingModelForSourceModel:destinationModel:error:"), sourceModel, destinationModel, error)
	return rv
}
// Returns the mapping model that will translate data from the source to the destination model. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(from:forSourceModel:destinationModel:)
func (mc _MappingModelClass) MappingModelFromBundlesForSourceModelDestinationModel(bundles unsafe.Pointer, sourceModel unsafe.Pointer, destinationModel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), bundles, sourceModel, destinationModel)
	return rv
}

