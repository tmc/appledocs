// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchRequest] class.
var (
	FetchRequestClass     _FetchRequestClass
	FetchRequestClassOnce sync.Once
)

func getFetchRequestClass() _FetchRequestClass {
	FetchRequestClassOnce.Do(func() {
		FetchRequestClass = _FetchRequestClass{objc.GetClass("NSFetchRequest")}
	})
	return FetchRequestClass
}

type _FetchRequestClass struct {
	class objc.Class
}

// An interface definition for the [FetchRequest] class.
type IFetchRequest interface {
	IPersistentStoreRequest
	Execute(error_ unsafe.Pointer) []objc.ID
}

// A description of search criteria used to retrieve data from a persistent store.
//
// An instance of collects the criteria needed to select and optionally to sort a group of managed objects held in an persistent store. A fetch request contains an or an entity name that specifies which entity to search. It frequently also contains: An predicate that specifies which properties to filter by and the constraints on selection, such as, . If you don’t specify a predicate, then the system fetches all instances of the entity that you specified, subject to other constraints. For more information, see . An array of sort descriptors that specify how to order the returned objects, such as ascending by last name and then by first name. You can also specify other aspects of a fetch request: Use to perform the fetch directly on the managed object context that’s associated with the current queue. Or use one of the methods such as to execute the fetch. In , you can use a property wrapper to execute the fetch and assign the results to a property. First, create the request: Then use a property wrapper with the request to declare a property that receives the objects that the fetch returns: You often predefine fetch requests in an managed object model to provide an API to retrieve a stored fetch request by name. Stored fetch requests can include placeholders for variable substitution, and serve as templates for later completion. Fetch request templates allow you to predefine queries with variables to substitute at runtime.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest
type FetchRequest struct {
	PersistentStoreRequest
}

// FetchRequestFrom constructs a [FetchRequest] from an unsafe.Pointer.
//
// A description of search criteria used to retrieve data from a persistent store.
func FetchRequestFrom(ptr unsafe.Pointer) FetchRequest {
	return FetchRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FetchRequestClass) Alloc() FetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FetchRequestClass) New() FetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FetchRequest) Init() FetchRequest {
	rv := objc.Send[FetchRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FetchRequest) Autorelease() FetchRequest {
	rv := objc.Send[FetchRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchRequest creates a new FetchRequest instance.
func NewFetchRequest() FetchRequest {
	return getFetchRequestClass().New()
}




// Initializes a fetch request configured with a given entity name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/init(entityName:)
func NewFetchRequestWithEntityName(entityName string) FetchRequest {
	instance := getFetchRequestClass().Alloc()
	rv := objc.Send[FetchRequest](instance.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	rv.Autorelease()
	return rv
}


// Returns a fetch request configured with a given entity name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchRequestWithEntityName:
func (fc _FetchRequestClass) FetchRequestWithEntityName(entityName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fetchRequestWithEntityName:"), objc.String(entityName))
	return rv
}

// Executes the fetch request against the managed object context that is associated with the current queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/execute()
func (f_ FetchRequest) Execute(error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("execute:"), error_)
	return rv
}

// An array of persistent stores specified for the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/affectedStores
func (f_ FetchRequest) AffectedStores() []PersistentStore {
	rv := objc.Send[[]PersistentStore](f_.ID, objc.Sel("affectedStores"))
	return rv
}


// SetAffectedStores sets the value of the affectedStores property.
// An array of persistent stores specified for the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/affectedStores
func (f_ FetchRequest) SetAffectedStores(value []PersistentStore) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setAffectedStores:"), nsArray)
}

// The entity specified for the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/entity
func (f_ FetchRequest) Entity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("entity"))
	return rv
}


// SetEntity sets the value of the entity property.
// The entity specified for the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/entity
func (f_ FetchRequest) SetEntity(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEntity:"), value)
}

// The name of the entity the request is configured to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/entityName
func (f_ FetchRequest) EntityName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("entityName"))
	return rv
}

// The batch size of the objects specified in the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchBatchSize
func (f_ FetchRequest) FetchBatchSize() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("fetchBatchSize"))
	return rv
}


// SetFetchBatchSize sets the value of the fetchBatchSize property.
// The batch size of the objects specified in the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchBatchSize
func (f_ FetchRequest) SetFetchBatchSize(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchBatchSize:"), value)
}

// The fetch limit of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchLimit
func (f_ FetchRequest) FetchLimit() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("fetchLimit"))
	return rv
}


// SetFetchLimit sets the value of the fetchLimit property.
// The fetch limit of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchLimit
func (f_ FetchRequest) SetFetchLimit(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchLimit:"), value)
}

// The fetch offset of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchOffset
func (f_ FetchRequest) FetchOffset() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("fetchOffset"))
	return rv
}


// SetFetchOffset sets the value of the fetchOffset property.
// The fetch offset of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchOffset
func (f_ FetchRequest) SetFetchOffset(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchOffset:"), value)
}

// The predicate used to filter rows being returned by a query containing a GROUP BY directive.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/havingPredicate
func (f_ FetchRequest) HavingPredicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("havingPredicate"))
	return rv
}


// SetHavingPredicate sets the value of the havingPredicate property.
// The predicate used to filter rows being returned by a query containing a GROUP BY directive.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/havingPredicate
func (f_ FetchRequest) SetHavingPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHavingPredicate:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPendingChanges
func (f_ FetchRequest) IncludesPendingChanges() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("includesPendingChanges"))
	return rv
}


// SetIncludesPendingChanges sets the value of the includesPendingChanges property.
// A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPendingChanges
func (f_ FetchRequest) SetIncludesPendingChanges(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIncludesPendingChanges:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPropertyValues
func (f_ FetchRequest) IncludesPropertyValues() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("includesPropertyValues"))
	return rv
}


// SetIncludesPropertyValues sets the value of the includesPropertyValues property.
// A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPropertyValues
func (f_ FetchRequest) SetIncludesPropertyValues(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIncludesPropertyValues:"), value)
}

// A Boolean value that indicates whether the fetch request includes subentities in the results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesSubentities
func (f_ FetchRequest) IncludesSubentities() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("includesSubentities"))
	return rv
}


// SetIncludesSubentities sets the value of the includesSubentities property.
// A Boolean value that indicates whether the fetch request includes subentities in the results.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesSubentities
func (f_ FetchRequest) SetIncludesSubentities(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIncludesSubentities:"), value)
}

// The predicate of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/predicate
func (f_ FetchRequest) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// The predicate of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/predicate
func (f_ FetchRequest) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPredicate:"), value)
}

// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch
func (f_ FetchRequest) PropertiesToFetch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("propertiesToFetch"))
	return rv
}


// SetPropertiesToFetch sets the value of the propertiesToFetch property.
// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch
func (f_ FetchRequest) SetPropertiesToFetch(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPropertiesToFetch:"), value)
}

// An array of objects that indicates how data should be grouped before a select statement is run in a SQL database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToGroupBy
func (f_ FetchRequest) PropertiesToGroupBy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("propertiesToGroupBy"))
	return rv
}


// SetPropertiesToGroupBy sets the value of the propertiesToGroupBy property.
// An array of objects that indicates how data should be grouped before a select statement is run in a SQL database.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToGroupBy
func (f_ FetchRequest) SetPropertiesToGroupBy(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPropertiesToGroupBy:"), value)
}

// The relationship key paths to prefetch along with the entity for the request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/relationshipKeyPathsForPrefetching
func (f_ FetchRequest) RelationshipKeyPathsForPrefetching() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("relationshipKeyPathsForPrefetching"))
	return rv
}


// SetRelationshipKeyPathsForPrefetching sets the value of the relationshipKeyPathsForPrefetching property.
// The relationship key paths to prefetch along with the entity for the request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/relationshipKeyPathsForPrefetching
func (f_ FetchRequest) SetRelationshipKeyPathsForPrefetching(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setRelationshipKeyPathsForPrefetching:"), nsArray)
}

// The result type of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/resultType
func (f_ FetchRequest) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("resultType"))
	return rv
}


// SetResultType sets the value of the resultType property.
// The result type of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/resultType
func (f_ FetchRequest) SetResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultType:"), value)
}

// A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsDistinctResults
func (f_ FetchRequest) ReturnsDistinctResults() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("returnsDistinctResults"))
	return rv
}


// SetReturnsDistinctResults sets the value of the returnsDistinctResults property.
// A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by .

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsDistinctResults
func (f_ FetchRequest) SetReturnsDistinctResults(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReturnsDistinctResults:"), value)
}

// A Boolean value that indicates whether the objects resulting from a fetch request are faults.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsObjectsAsFaults
func (f_ FetchRequest) ReturnsObjectsAsFaults() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("returnsObjectsAsFaults"))
	return rv
}


// SetReturnsObjectsAsFaults sets the value of the returnsObjectsAsFaults property.
// A Boolean value that indicates whether the objects resulting from a fetch request are faults.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsObjectsAsFaults
func (f_ FetchRequest) SetReturnsObjectsAsFaults(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReturnsObjectsAsFaults:"), value)
}

// A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/shouldRefreshRefetchedObjects
func (f_ FetchRequest) ShouldRefreshRefetchedObjects() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("shouldRefreshRefetchedObjects"))
	return rv
}


// SetShouldRefreshRefetchedObjects sets the value of the shouldRefreshRefetchedObjects property.
// A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/shouldRefreshRefetchedObjects
func (f_ FetchRequest) SetShouldRefreshRefetchedObjects(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setShouldRefreshRefetchedObjects:"), value)
}

// The sort descriptors of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (f_ FetchRequest) SortDescriptors() []NSSortDescriptor {
	rv := objc.Send[[]NSSortDescriptor](f_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// The sort descriptors of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (f_ FetchRequest) SetSortDescriptors(value []NSSortDescriptor) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}


