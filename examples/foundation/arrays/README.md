# Foundation NSArray Examples

This example demonstrates working with NSArray and NSMutableArray collections in Go.

## What it demonstrates

- Creating NSArray and NSMutableArray objects
- Adding and accessing array elements
- Array operations (contains, indexOf, first, last)
- Mutable array modifications (add, insert, remove)
- Array enumeration
- Common use cases (file extensions, collections)

## Running the example

```bash
go run main.go
```

## Key Concepts

### NSArray vs NSMutableArray

**NSArray** (Immutable):
- Cannot be modified after creation
- Thread-safe for reading
- More efficient for fixed collections
- Use when array contents won't change

**NSMutableArray** (Mutable):
- Can add, remove, and reorder elements
- Subclass of NSArray
- Not thread-safe without synchronization
- Use when array needs to change

### Creating Arrays

```go
// Empty array (immutable)
array := foundation.NewArray()

// Array with single object
array := foundation.NewArrayWithObject(unsafe.Pointer(obj))

// Mutable array (can be modified)
mutableArrayClass := objc.GetClass("NSMutableArray")
newSel := objc.RegisterName("new")
mutableArray := objc.ID(mutableArrayClass).Send(newSel)
```

### Adding Elements (NSMutableArray)

```go
addObjectSel := objc.RegisterName("addObject:")
objc.ID(mutableArray).Send(addObjectSel, object.ID)

// Insert at specific index
insertObjectSel := objc.RegisterName("insertObject:atIndex:")
objc.ID(mutableArray).Send(insertObjectSel, object.ID, uint64(index))
```

### Accessing Elements

```go
// Get count
countSel := objc.RegisterName("count")
count := array.ID.Send(countSel)

// Get object at index
objectAtIndexSel := objc.RegisterName("objectAtIndex:")
obj := array.ID.Send(objectAtIndexSel, uint64(index))

// First and last objects
firstObjectSel := objc.RegisterName("firstObject")
lastObjectSel := objc.RegisterName("lastObject")

firstObj := array.ID.Send(firstObjectSel)
lastObj := array.ID.Send(lastObjectSel)
```

### Searching

```go
// Check if contains object
containsObjectSel := objc.RegisterName("containsObject:")
contains := array.ID.Send(containsObjectSel, searchObject.ID)
// contains != 0 means found

// Get index of object
indexOfObjectSel := objc.RegisterName("indexOfObject:")
index := array.ID.Send(indexOfObjectSel, searchObject.ID)
// index == NSNotFound if not found
```

### Removing Elements (NSMutableArray)

```go
// Remove at index
removeObjectAtIndexSel := objc.RegisterName("removeObjectAtIndex:")
objc.ID(mutableArray).Send(removeObjectAtIndexSel, uint64(index))

// Remove specific object
removeObjectSel := objc.RegisterName("removeObject:")
objc.ID(mutableArray).Send(removeObjectSel, object.ID)

// Remove all objects
removeAllObjectsSel := objc.RegisterName("removeAllObjects")
objc.ID(mutableArray).Send(removeAllObjectsSel)
```

## Use Cases

### File Path Collections

```go
// Array of file paths
paths := []string{"/Documents/file1.txt", "/Documents/file2.pdf"}
for _, path := range paths {
    pathStr := goStringToNS(path)
    objc.ID(pathsArray).Send(addObjectSel, pathStr.ID)
}
```

### Configuration Lists

```go
// Array of settings/options
settings := []string{"auto", "manual", "disabled"}
for _, setting := range settings {
    settingStr := goStringToNS(setting)
    objc.ID(settingsArray).Send(addObjectSel, settingStr.ID)
}
```

### Data Processing

```go
// Collect results
for item := range processItems() {
    result := goStringToNS(item.String())
    objc.ID(resultsArray).Send(addObjectSel, result.ID)
}
```

## NSArray vs Go Slices

| Feature | NSArray | Go []interface{} |
|---------|---------|-----------------|
| Mutability | Immutable | Mutable |
| Type Safety | Dynamic (objc.ID) | Static (with type) |
| Performance | Slower (ObjC) | Faster (native) |
| Memory | Reference counted | GC managed |
| API | Rich ObjC methods | Simple, functional |
| Interop | Required for Cocoa | N/A |

**When to use NSArray:**
- Interfacing with Foundation/AppKit/UIKit APIs
- Need NSArray-specific methods
- Working with Cocoa collections
- Plist serialization

**When to use Go slices:**
- Pure Go code
- Performance-critical paths
- Type safety is important
- No Foundation dependency

## Best Practices

1. **Choose mutability wisely**:
   - Use NSArray when collection is fixed
   - Use NSMutableArray only when needed

2. **Index bounds checking**:
   - Always check count before accessing by index
   - Handle NSNotFound return values

3. **Memory management**:
   - Arrays retain their objects
   - Objects released when array is deallocated
   - Use autorelease pools for intensive operations

4. **Type safety**:
   - Store homogeneous types when possible
   - Validate object types before casting

5. **Enumeration**:
   - Use modern enumeration when available
   - Cache selectors for performance

## Advanced Features

NSArray also supports:
- Fast enumeration (NSFastEnumeration protocol)
- Filtering (filteredArrayUsingPredicate:)
- Mapping (via NSPredicate or blocks)
- Sorting (sortedArrayUsingDescriptors:)
- Binary search (indexOfObject:inSortedRange:)
- Subarrays (subarrayWithRange:)

## Common Patterns

### Building from Go Slice

```go
func goSliceToNSArray(items []string) objc.ID {
    arrayClass := objc.GetClass("NSMutableArray")
    array := objc.ID(arrayClass).Send(objc.RegisterName("new"))
    addSel := objc.RegisterName("addObject:")

    for _, item := range items {
        str := goStringToNS(item)
        objc.ID(array).Send(addSel, str.ID)
    }

    return array
}
```

### Converting to Go Slice

```go
func nsArrayToGoSlice(array objc.ID) []string {
    countSel := objc.RegisterName("count")
    count := objc.ID(array).Send(countSel)

    result := make([]string, 0, int(count))
    objectAtSel := objc.RegisterName("objectAtIndex:")

    for i := uint64(0); i < uint64(count); i++ {
        obj := objc.ID(array).Send(objectAtSel, i)
        str := foundation.StringFrom(unsafe.Pointer(obj))
        result = append(result, nsStringToGo(str))
    }

    return result
}
```

## References

- [NSArray Documentation](https://developer.apple.com/documentation/foundation/nsarray)
- [NSMutableArray Documentation](https://developer.apple.com/documentation/foundation/nsmutablearray)
- [Collections Programming Topics](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Collections/Collections.html)
