// PhotosSwift.swift
// Swift wrappers for Photos framework extensions
// Generated from Photos.swiftinterface

import Photos
import Foundation

// MARK: - PHProjectChangeRequest Extensions

/// Wrapper for PHProjectChangeRequest.removeAssets(_: PHFetchResult<PHAsset>)
/// Non-generic version - simplest to wrap
@_cdecl("photos_project_change_request_remove_assets_fetch_result")
public func photos_project_change_request_remove_assets_fetch_result(
    requestPtr: OpaquePointer,
    assetsPtr: OpaquePointer
) {
    let request = Unmanaged<PHProjectChangeRequest>
        .fromOpaque(UnsafeRawPointer(requestPtr))
        .takeUnretainedValue()

    let fetchResult = Unmanaged<PHFetchResult<PHAsset>>
        .fromOpaque(UnsafeRawPointer(assetsPtr))
        .takeUnretainedValue()

    request.removeAssets(fetchResult)
}

// MARK: - PHPersistentChangeFetchResult Extensions

/// Wrapper for PHPersistentChangeFetchResult.makeIterator()
/// Returns an iterator for the fetch result
@_cdecl("photos_persistent_change_fetch_result_make_iterator")
public func photos_persistent_change_fetch_result_make_iterator(
    fetchResultPtr: OpaquePointer
) -> OpaquePointer {
    let fetchResult = Unmanaged<PHPersistentChangeFetchResult>
        .fromOpaque(UnsafeRawPointer(fetchResultPtr))
        .takeUnretainedValue()

    let iterator = fetchResult.makeIterator()

    return OpaquePointer(Unmanaged.passRetained(iterator as AnyObject).toOpaque())
}

/// Wrapper for Iterator.next()
/// Returns the next PHPersistentChange or nil
@_cdecl("photos_persistent_change_iterator_next")
public func photos_persistent_change_iterator_next(
    iteratorPtr: OpaquePointer
) -> OpaquePointer? {
    let iterator = Unmanaged<PHPersistentChangeFetchResult.Iterator>
        .fromOpaque(UnsafeRawPointer(iteratorPtr))
        .takeUnretainedValue()

    if let change = iterator.next() {
        return OpaquePointer(Unmanaged.passRetained(change as AnyObject).toOpaque())
    }

    return nil
}

// MARK: - Helper Functions

/// Get shared PHPhotoLibrary instance
@_cdecl("photos_shared_library")
public func photos_shared_library() -> OpaquePointer {
    let library = PHPhotoLibrary.shared()
    return OpaquePointer(Unmanaged.passRetained(library).toOpaque())
}

/// Release an object created by these wrappers
@_cdecl("photos_release")
public func photos_release(ptr: OpaquePointer) {
    Unmanaged<AnyObject>.fromOpaque(UnsafeRawPointer(ptr)).release()
}

// MARK: - Info Functions

/// Get the count of items in a PHFetchResult
@_cdecl("photos_fetch_result_count")
public func photos_fetch_result_count(fetchResultPtr: OpaquePointer) -> Int {
    let fetchResult = Unmanaged<PHFetchResult<PHAsset>>
        .fromOpaque(UnsafeRawPointer(fetchResultPtr))
        .takeUnretainedValue()

    return fetchResult.count
}

/// Test function to verify the library works
@_cdecl("photos_test_hello")
public func photos_test_hello() {
    print("Hello from Photos Swift wrapper!")
    print("PHPhotoLibrary is available: \(PHPhotoLibrary.self)")
}
