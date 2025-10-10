// Package photosframework provides Swift-like Go bindings for Photos framework extensions.
//
// This package wraps Swift extensions from Photos.swiftinterface, providing idiomatic Go APIs
// for accessing photo library functionality that's only available in Swift.
package photosframework

import (
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	// Core functions from libPhotosSwift.dylib
	photosTestHello                        func()
	photosSharedLibrary                    func() unsafe.Pointer
	photosFetchResultCount                 func(unsafe.Pointer) int
	photosRelease                          func(unsafe.Pointer)
	photosProjectChangeRequestRemoveAssets func(unsafe.Pointer, unsafe.Pointer)
	photosPersistentChangeMakeIterator     func(unsafe.Pointer) unsafe.Pointer
	photosPersistentChangeIteratorNext     func(unsafe.Pointer) unsafe.Pointer
)

// init loads the Photos Swift wrapper library and registers all functions
func init() {
	// Load Swift wrapper library
	lib, err := purego.Dlopen("libPhotosSwift.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic("Failed to load libPhotosSwift.dylib: " + err.Error())
	}

	// Register all exported functions
	purego.RegisterLibFunc(&photosTestHello, lib, "photos_test_hello")
	purego.RegisterLibFunc(&photosSharedLibrary, lib, "photos_shared_library")
	purego.RegisterLibFunc(&photosFetchResultCount, lib, "photos_fetch_result_count")
	purego.RegisterLibFunc(&photosRelease, lib, "photos_release")
	purego.RegisterLibFunc(&photosProjectChangeRequestRemoveAssets, lib, "photos_project_change_request_remove_assets_fetch_result")
	purego.RegisterLibFunc(&photosPersistentChangeMakeIterator, lib, "photos_persistent_change_fetch_result_make_iterator")
	purego.RegisterLibFunc(&photosPersistentChangeIteratorNext, lib, "photos_persistent_change_iterator_next")
}

// TestHello is a diagnostic function that prints a message from the Swift wrapper.
// Useful for verifying the library is loaded correctly.
func TestHello() {
	photosTestHello()
}

// PhotoLibrary represents a PHPhotoLibrary instance.
type PhotoLibrary struct {
	ptr unsafe.Pointer
}

// SharedPhotoLibrary returns the shared photo library instance.
func SharedPhotoLibrary() *PhotoLibrary {
	ptr := photosSharedLibrary()
	if ptr == nil {
		return nil
	}

	lib := &PhotoLibrary{ptr: ptr}
	runtime.SetFinalizer(lib, (*PhotoLibrary).Release)
	return lib
}

// Release frees the photo library resources.
// This is called automatically by the garbage collector, but can be called manually.
func (pl *PhotoLibrary) Release() {
	if pl.ptr != nil {
		photosRelease(pl.ptr)
		pl.ptr = nil
	}
}

// FetchResult represents a PHFetchResult<PHAsset> instance.
type FetchResult struct {
	ptr unsafe.Pointer
}

// Count returns the number of assets in the fetch result.
func (fr *FetchResult) Count() int {
	if fr.ptr == nil {
		return 0
	}
	return photosFetchResultCount(fr.ptr)
}

// Release frees the fetch result resources.
func (fr *FetchResult) Release() {
	if fr.ptr != nil {
		photosRelease(fr.ptr)
		fr.ptr = nil
	}
}

// ProjectChangeRequest represents a PHProjectChangeRequest instance.
type ProjectChangeRequest struct {
	ptr unsafe.Pointer
}

// RemoveAssets removes assets from a project.
// This is a Swift extension method that's not available in the Objective-C API.
func (pcr *ProjectChangeRequest) RemoveAssets(assets *FetchResult) {
	if pcr.ptr == nil || assets.ptr == nil {
		return
	}
	photosProjectChangeRequestRemoveAssets(pcr.ptr, assets.ptr)
}

// PersistentChangeFetchResult represents a PHPersistentChangeFetchResult instance.
type PersistentChangeFetchResult struct {
	ptr unsafe.Pointer
}

// Iterator creates an iterator for the persistent change fetch result.
// This is a Swift extension that implements the Sequence protocol.
func (pcfr *PersistentChangeFetchResult) Iterator() *PersistentChangeIterator {
	if pcfr.ptr == nil {
		return nil
	}

	iterPtr := photosPersistentChangeMakeIterator(pcfr.ptr)
	if iterPtr == nil {
		return nil
	}

	iter := &PersistentChangeIterator{ptr: iterPtr}
	runtime.SetFinalizer(iter, (*PersistentChangeIterator).Release)
	return iter
}

// Release frees the persistent change fetch result resources.
func (pcfr *PersistentChangeFetchResult) Release() {
	if pcfr.ptr != nil {
		photosRelease(pcfr.ptr)
		pcfr.ptr = nil
	}
}

// PersistentChangeIterator represents a PHPersistentChangeFetchResult.Iterator instance.
type PersistentChangeIterator struct {
	ptr unsafe.Pointer
}

// Next returns the next persistent change, or nil if there are no more changes.
func (pci *PersistentChangeIterator) Next() *PersistentChange {
	if pci.ptr == nil {
		return nil
	}

	changePtr := photosPersistentChangeIteratorNext(pci.ptr)
	if changePtr == nil {
		return nil
	}

	change := &PersistentChange{ptr: changePtr}
	runtime.SetFinalizer(change, (*PersistentChange).Release)
	return change
}

// Release frees the iterator resources.
func (pci *PersistentChangeIterator) Release() {
	if pci.ptr != nil {
		photosRelease(pci.ptr)
		pci.ptr = nil
	}
}

// PersistentChange represents a PHPersistentChange instance.
type PersistentChange struct {
	ptr unsafe.Pointer
}

// Release frees the persistent change resources.
func (pc *PersistentChange) Release() {
	if pc.ptr != nil {
		photosRelease(pc.ptr)
		pc.ptr = nil
	}
}
