// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FetchedResultsController] class.
var fetchedResultsControllerClass = _FetchedResultsControllerClass{objc.GetClass("NSFetchedResultsController")}

type _FetchedResultsControllerClass struct {
	class objc.Class
}

// An interface definition for the [FetchedResultsController] class.
type IFetchedResultsController interface {
	objectivec.IObject
	IndexPathForObject(object unsafe.Pointer) unsafe.Pointer
	ObjectAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer
	PerformFetch(error unsafe.Pointer) bool
	SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int
	SectionIndexTitleForSectionName(sectionName string) unsafe.Pointer
}

// A controller that you use to manage the results of a Core Data fetch request and to display data to the user. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return fetchedResultsControllerClass.New()
}


// Returns a fetch request controller initialized using the given arguments. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/init(fetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:)
func NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest unsafe.Pointer, context unsafe.Pointer, sectionNameKeyPath string, name string) FetchedResultsController {
	instance := fetchedResultsControllerClass.Alloc()
	rv := objc.Send[FetchedResultsController](instance.ID, objc.Sel("initWithFetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:"), fetchRequest, context, objc.String(sectionNameKeyPath), objc.String(name))
	rv.Autorelease()
	return rv
}


// Deletes the cached section information with the given name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/deleteCache(withName:)
func (fc _FetchedResultsControllerClass) DeleteCacheWithName(name string) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("deleteCacheWithName:"), objc.String(name))
}
// Returns the index path of a given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/indexPath(forObject:)
func (f_ FetchedResultsController) IndexPathForObject(object unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("indexPathForObject:"), object)
	return rv
}
// Returns the object at the given index path in the fetch results. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/object(at:)
func (f_ FetchedResultsController) ObjectAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("objectAtIndexPath:"), indexPath)
	return rv
}
// Executes the controller’s fetch request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/performFetch()
func (f_ FetchedResultsController) PerformFetch(error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("performFetch:"), error)
	return rv
}
// Returns the section number for a given section title and index in the section index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/section(forSectionIndexTitle:at:)
func (f_ FetchedResultsController) SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int {
	rv := objc.Send[int](f_.ID, objc.Sel("sectionForSectionIndexTitle:atIndex:"), objc.String(title), sectionIndex)
	return rv
}
// Returns the corresponding section index entry for a given section name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionIndexTitle(forSectionName:)
func (f_ FetchedResultsController) SectionIndexTitleForSectionName(sectionName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("sectionIndexTitleForSectionName:"), objc.String(sectionName))
	return rv
}

