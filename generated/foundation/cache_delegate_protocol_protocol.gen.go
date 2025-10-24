// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCacheDelegate is the NSCacheDelegate protocol interface.
//
// The delegate of an   object implements this protocol to perform specialized actions when an object is about to be evicted or removed from the cache.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSCacheDelegate
type PCacheDelegate interface {
	// Optional methods
	CacheWillEvictObject(cache ICache, obj objc.IObject)
	HasCacheWillEvictObject() bool
}

// CacheDelegate is a delegate implementation builder for the PCacheDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CacheDelegate struct {
	_CacheWillEvictObject func(cache ICache, obj objc.IObject)
}

// SetCacheWillEvictObject sets the handler for the CacheWillEvictObject delegate method.
//
// Called when an object is about to be evicted or removed from the cache.
func (d *CacheDelegate) SetCacheWillEvictObject(f func(cache ICache, obj objc.IObject)) {
	d._CacheWillEvictObject = f
}

// CacheWillEvictObject implements the PCacheDelegate interface.
func (d *CacheDelegate) CacheWillEvictObject(cache ICache, obj objc.IObject) {
	if d._CacheWillEvictObject != nil {
		d._CacheWillEvictObject(cache, obj)
	}
}

// HasCacheWillEvictObject returns true if a handler for CacheWillEvictObject has been set.
func (d *CacheDelegate) HasCacheWillEvictObject() bool {
	return d._CacheWillEvictObject != nil
}
