// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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
	CacheWillEvictObject(cache ICache, obj objectivec.IObject)
	HasCacheWillEvictObject() bool
}

// CacheDelegate is a delegate implementation builder for the PCacheDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CacheDelegate struct {
	_CacheWillEvictObject func(cache ICache, obj objectivec.IObject)
}

// SetCacheWillEvictObject sets the handler for the CacheWillEvictObject delegate method.
//
// Called when an object is about to be evicted or removed from the cache.
func (d *CacheDelegate) SetCacheWillEvictObject(f func(cache ICache, obj objectivec.IObject)) {
	d._CacheWillEvictObject = f
}

// CacheWillEvictObject implements the PCacheDelegate interface.
func (d *CacheDelegate) CacheWillEvictObject(cache ICache, obj objectivec.IObject) {
	if d._CacheWillEvictObject != nil {
		d._CacheWillEvictObject(cache, obj)
	}
}

// HasCacheWillEvictObject returns true if a handler for CacheWillEvictObject has been set.
func (d *CacheDelegate) HasCacheWillEvictObject() bool {
	return d._CacheWillEvictObject != nil
}

// CacheDelegateObject wraps an existing Objective-C object that conforms to the PCacheDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type CacheDelegateObject struct {
	objectivec.Object
}

// NewCacheDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSCacheDelegate protocol.
func NewCacheDelegateObject(obj objectivec.Object) *CacheDelegateObject {
	return &CacheDelegateObject{obj}
}

// Make sure CacheDelegateObject implements PCacheDelegate.
var _ PCacheDelegate = (*CacheDelegateObject)(nil)

// CacheWillEvictObject implements the PCacheDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CacheDelegateObject) CacheWillEvictObject(cache ICache, obj objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("cache:willEvictObject:"), cache, obj)
}

// HasCacheWillEvictObject returns true; this is a placeholder for optional method checks.
func (o *CacheDelegateObject) HasCacheWillEvictObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
