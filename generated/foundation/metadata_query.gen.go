// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataQuery] class.
var (
	metadataQueryClass     _MetadataQueryClass
	metadataQueryClassOnce sync.Once
)

func getMetadataQueryClass() _MetadataQueryClass {
	metadataQueryClassOnce.Do(func() {
		metadataQueryClass = _MetadataQueryClass{objc.GetClass("NSMetadataQuery")}
	})
	return metadataQueryClass
}

type _MetadataQueryClass struct {
	class objc.Class
}

// An interface definition for the [MetadataQuery] class.
type IMetadataQuery interface {
	objectivec.IObject
}

// A query that you perform against Spotlight metadata.
//
// The class encapsulates the functionality provided by the opaque type for querying the Spotlight metadata. objects provide metadata query results in several ways: As individual attribute values for requested attributes. As value lists that contain the distinct values for given attributes in the query results. As a result array proxy, containing all the query results. This is suitable for use with Cocoa bindings. As a hierarchical collection of results, grouping together items with the same values for specified grouping attributes. This is also suitable for use with Cocoa bindings. Queries have two phases: the initial gathering phase that collects all currently matching results and a second live-update phase. By default, the receiver has no limitation on its search scope. Use the property to customize. By default, notification of updated results occurs at 1.0 seconds. Use the property to customize. You must set a predicate with the property before starting a query.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery
type MetadataQuery struct {
	objectivec.Object
}

// MetadataQueryFrom constructs a [MetadataQuery] from an unsafe.Pointer.
//
// A query that you perform against Spotlight metadata.
func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataQueryClass) Alloc() MetadataQuery {
	rv := objc.Send[MetadataQuery](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataQueryClass) New() MetadataQuery {
	rv := objc.Send[MetadataQuery](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataQuery) Init() MetadataQuery {
	rv := objc.Send[MetadataQuery](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataQuery) Autorelease() MetadataQuery {
	rv := objc.Send[MetadataQuery](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataQuery creates a new MetadataQuery instance.
func NewMetadataQuery() MetadataQuery {
	return getMetadataQueryClass().New()
}




