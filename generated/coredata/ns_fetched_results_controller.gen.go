// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FetchedResultsController] class.
var (
	FetchedResultsControllerClass     _FetchedResultsControllerClass
	FetchedResultsControllerClassOnce sync.Once
)

func getFetchedResultsControllerClass() _FetchedResultsControllerClass {
	FetchedResultsControllerClassOnce.Do(func() {
		FetchedResultsControllerClass = _FetchedResultsControllerClass{objc.GetClass("NSFetchedResultsController")}
	})
	return FetchedResultsControllerClass
}

type _FetchedResultsControllerClass struct {
	class objc.Class
}

// An interface definition for the [FetchedResultsController] class.
type IFetchedResultsController interface {
	objectivec.IObject
	IndexPathForObject(object unsafe.Pointer) unsafe.Pointer
	ObjectAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer
	PerformFetch(error_ unsafe.Pointer) bool
	SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int
	SectionIndexTitleForSectionName(sectionName string) unsafe.Pointer
}

// A controller that you use to manage the results of a Core Data fetch request and to display data to the user.
//
// While you can use table views can in several ways, fetched results controllers primarily assist you with a primary list view. expects its data source to provide cells as an array of sections made up of rows. You configure a fetched results controller using a — an object that specifies what type of entity to fetch and how to sort the results. You can also add criteria for when to include a specific instance of the entity. The fetched results controller efficiently analyzes the result of the fetch request and computes all the information about sections in the result set. It also computes all the information for the index based on the result set. In addition, fetched results controllers: Optionally monitor changes to objects in the associated managed object context, and report changes in the results set to its delegate (see ). Optionally cache the results of its computation to enable redisplaying the same data without repeating the work to fetch it. For more information, see . A controller thus effectively has three modes of operation, determined by whether it has a delegate and whether you set the cache file name. No tracking: The delegate is . The controller provides access to the data as it was when it fetched it. Memory-only tracking: the delegate is non- and the file cache name is . The controller monitors objects in its result set and updates section and ordering information in response to relevant changes. Full persistent tracking: the delegate and the file cache name are non- . The controller monitors objects in its result set and updates section and ordering information in response to relevant changes. The controller maintains a persistent cache of the results of its computation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
type FetchedResultsController struct {
	objectivec.Object
}

// FetchedResultsControllerFrom constructs a [FetchedResultsController] from an unsafe.Pointer.
//
// A controller that you use to manage the results of a Core Data fetch request and to display data to the user.
func FetchedResultsControllerFrom(ptr unsafe.Pointer) FetchedResultsController {
	return FetchedResultsController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FetchedResultsControllerClass) Alloc() FetchedResultsController {
	rv := objc.Send[FetchedResultsController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FetchedResultsControllerClass) New() FetchedResultsController {
	rv := objc.Send[FetchedResultsController](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FetchedResultsController) Init() FetchedResultsController {
	rv := objc.Send[FetchedResultsController](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FetchedResultsController) Autorelease() FetchedResultsController {
	rv := objc.Send[FetchedResultsController](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchedResultsController creates a new FetchedResultsController instance.
func NewFetchedResultsController() FetchedResultsController {
	return getFetchedResultsControllerClass().New()
}


// Returns a fetch request controller initialized using the given arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/init(fetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:)
func NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest unsafe.Pointer, context unsafe.Pointer, sectionNameKeyPath string, name string) FetchedResultsController {
	instance := getFetchedResultsControllerClass().Alloc()
	rv := objc.Send[FetchedResultsController](instance.ID, objc.Sel("initWithFetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:"), fetchRequest, context, objc.String(sectionNameKeyPath), objc.String(name))
	rv.Autorelease()
	return rv
}


// Deletes the cached section information with the given name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/deleteCache(withName:)
func (fc _FetchedResultsControllerClass) DeleteCacheWithName(name string) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("deleteCacheWithName:"), objc.String(name))
}

// Returns the index path of a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/indexPath(forObject:)
func (f_ FetchedResultsController) IndexPathForObject(object unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("indexPathForObject:"), object)
	return rv
}

// Returns the object at the given index path in the fetch results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/object(at:)
func (f_ FetchedResultsController) ObjectAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("objectAtIndexPath:"), indexPath)
	return rv
}

// Executes the controller’s fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/performFetch()
func (f_ FetchedResultsController) PerformFetch(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("performFetch:"), error_)
	return rv
}

// Returns the section number for a given section title and index in the section index.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/section(forSectionIndexTitle:at:)
func (f_ FetchedResultsController) SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int {
	rv := objc.Send[int](f_.ID, objc.Sel("sectionForSectionIndexTitle:atIndex:"), objc.String(title), sectionIndex)
	return rv
}

// Returns the corresponding section index entry for a given section name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionIndexTitle(forSectionName:)
func (f_ FetchedResultsController) SectionIndexTitleForSectionName(sectionName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("sectionIndexTitleForSectionName:"), objc.String(sectionName))
	return rv
}

// The name of the file used to cache section information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/cacheName
func (f_ FetchedResultsController) CacheName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("cacheName"))
	return rv
}

// The object that is notified when the fetched results changed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/delegate
func (f_ FetchedResultsController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that is notified when the fetched results changed.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/delegate
func (f_ FetchedResultsController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}
// The fetch request used to do the fetching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/fetchRequest
func (f_ FetchedResultsController) FetchRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fetchRequest"))
	return rv
}

// The results of the fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/fetchedObjects
func (f_ FetchedResultsController) FetchedObjects() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("fetchedObjects"))
	return rv
}

// The managed object context used to fetch objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/managedObjectContext
func (f_ FetchedResultsController) ManagedObjectContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("managedObjectContext"))
	return rv
}

// The array of section index titles.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionIndexTitles
func (f_ FetchedResultsController) SectionIndexTitles() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("sectionIndexTitles"))
	return rv
}

// The key path of the attribute that determines which section the fetched entity belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionNameKeyPath
func (f_ FetchedResultsController) SectionNameKeyPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("sectionNameKeyPath"))
	return rv
}

// The sections for the fetch results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sections
func (f_ FetchedResultsController) Sections() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("sections"))
	return rv
}


