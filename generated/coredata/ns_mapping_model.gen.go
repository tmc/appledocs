// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MappingModel] class.
var (
	MappingModelClass     _MappingModelClass
	MappingModelClassOnce sync.Once
)

func getMappingModelClass() _MappingModelClass {
	MappingModelClassOnce.Do(func() {
		MappingModelClass = _MappingModelClass{objc.GetClass("NSMappingModel")}
	})
	return MappingModelClass
}

type _MappingModelClass struct {
	class objc.Class
}

// An interface definition for the [MappingModel] class.
type IMappingModel interface {
	objectivec.IObject
	// properties:
	EntityMappings() IEntityMapping
	SetEntityMappings(value IEntityMapping)
	EntityMappingsByName() IEntityMapping
	SetEntityMappingsByName(value IEntityMapping)
	// methods:
}

// A model instance that specifies how to map a model from a source to a destination managed object model.


// A model instance that specifies how to map a model from a source to a destination managed object model.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getMappingModelClass().New()
}



// Returns the mapping model that will translate data from the source to the destination model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(from:forSourceModel:destinationModel:)
func NewMappingModelFromBundlesForSourceModelDestinationModel(bundles []foundation.objc.IObject /* cross-framework Bundle */, sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MappingModel {
	rv := objc.Send[MappingModel](objc.ID(getMappingModelClass().class), objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), bundles, sourceModel, destinationModel)
	return rv
}



// Returns a newly created mapping model that will migrate data from the source to the destination model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/inferredMappingModel(forSourceModel:destinationModel:)
func (mc _MappingModelClass) InferredMappingModelForSourceModelDestinationModelError(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel, error_ unsafe.Pointer) IMappingModel {
	rv := objc.Send[MappingModel](objc.ID(mc.class), objc.Sel("inferredMappingModelForSourceModel:destinationModel:error:"), sourceModel, destinationModel, error_)
	return rv
}


// Returns the mapping model that will translate data from the source to the destination model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(from:forSourceModel:destinationModel:)
func (mc _MappingModelClass) MappingModelFromBundlesForSourceModelDestinationModel(bundles []foundation.objc.IObject /* cross-framework Bundle */, sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) IMappingModel {
	rv := objc.Send[MappingModel](objc.ID(mc.class), objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), bundles, sourceModel, destinationModel)
	return rv
}


// The entity mappings for the mapping model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/entitymappings
func (m_ MappingModel) EntityMappings() IEntityMapping {
	rv := objc.Send[EntityMapping](m_.ID, objc.Sel("entityMappings"))
	return rv
}


// The entity mappings for the mapping model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/entitymappings
func (m_ MappingModel) SetEntityMappings(value IEntityMapping) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntityMappings:"), value)
}


// The entity mappings for the mapping model, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/entitymappingsbyname
func (m_ MappingModel) EntityMappingsByName() IEntityMapping {
	rv := objc.Send[EntityMapping](m_.ID, objc.Sel("entityMappingsByName"))
	return rv
}


// The entity mappings for the mapping model, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/entitymappingsbyname
func (m_ MappingModel) SetEntityMappingsByName(value IEntityMapping) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntityMappingsByName:"), value)
}


