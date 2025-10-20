package main

import (
	"flag"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

// Helper to convert Go string to NSString
func goStringToNS(s string) foundation.String {
	cStr := append([]byte(s), 0)
	sel := objc.RegisterName("stringWithUTF8String:")
	class := objc.GetClass("NSString")
	result := objc.ID(class).Send(sel, unsafe.Pointer(&cStr[0]))
	return foundation.StringFrom(unsafe.Pointer(result))
}

// Helper to convert NSString to Go string
func nsStringToGo(nsStr foundation.String) string {
	if nsStr.ID == 0 {
		return ""
	}
	cStrSel := objc.RegisterName("UTF8String")
	cStr := nsStr.ID.Send(cStrSel)
	if cStr == 0 {
		return ""
	}
	result := ""
	for i := 0; ; i++ {
		b := *(*byte)(unsafe.Pointer(uintptr(cStr) + uintptr(i)))
		if b == 0 {
			break
		}
		result += string(b)
	}
	return result
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("Foundation NSArray Examples")
	fmt.Println("===========================")

	// Example 1: Create empty array
	fmt.Println("\n1. Creating Arrays:")

	emptyArray := foundation.NewArray()
	fmt.Printf("   Empty array created: %v\n", emptyArray)

	// Example 2: Array with objects (using NSString objects)
	fmt.Println("\n2. Array with String Objects:")

	// Create some NSString objects
	str1 := goStringToNS("Apple")
	str2 := goStringToNS("Banana")
	str3 := goStringToNS("Cherry")

	// Create array with single object
	singleArray := foundation.NewArrayWithObject(unsafe.Pointer(str1.ID))
	fmt.Printf("   Array with single object created\n")

	// Get count
	countSel := objc.RegisterName("count")
	count := singleArray.ID.Send(countSel)
	fmt.Printf("   Array count: %d\n", count)

	// Example 3: Accessing array elements
	fmt.Println("\n3. Accessing Array Elements:")

	// Create mutable array and add objects
	// (Note: arrayWithObjects: is variadic and complex to call from Go)
	mutableArrayClass := objc.GetClass("NSMutableArray")
	newSel := objc.RegisterName("new")
	mutableArray := objc.ID(mutableArrayClass).Send(newSel)

	// Add objects to mutable array
	addObjectSel := objc.RegisterName("addObject:")
	objc.ID(mutableArray).Send(addObjectSel, str1.ID)
	objc.ID(mutableArray).Send(addObjectSel, str2.ID)
	objc.ID(mutableArray).Send(addObjectSel, str3.ID)

	// Get count
	arrayCount := objc.ID(mutableArray).Send(countSel)
	fmt.Printf("   Array count: %d\n", arrayCount)

	// Access elements
	objectAtIndexSel := objc.RegisterName("objectAtIndex:")
	for i := uint64(0); i < uint64(arrayCount); i++ {
		obj := objc.ID(mutableArray).Send(objectAtIndexSel, i)
		str := foundation.StringFrom(unsafe.Pointer(obj))
		fmt.Printf("   [%d]: %s\n", i, nsStringToGo(str))
	}

	// Example 4: Array operations
	fmt.Println("\n4. Array Operations:")

	// Contains object
	containsObjectSel := objc.RegisterName("containsObject:")
	contains := objc.ID(mutableArray).Send(containsObjectSel, str2.ID)
	fmt.Printf("   Contains 'Banana': %v\n", contains != 0)

	searchStr := goStringToNS("Grape")
	containsGrape := objc.ID(mutableArray).Send(containsObjectSel, searchStr.ID)
	fmt.Printf("   Contains 'Grape': %v\n", containsGrape != 0)

	// Index of object
	indexOfObjectSel := objc.RegisterName("indexOfObject:")
	index := objc.ID(mutableArray).Send(indexOfObjectSel, str2.ID)
	fmt.Printf("   Index of 'Banana': %d\n", index)

	// First and last object
	firstObjectSel := objc.RegisterName("firstObject")
	lastObjectSel := objc.RegisterName("lastObject")

	firstObj := objc.ID(mutableArray).Send(firstObjectSel)
	lastObj := objc.ID(mutableArray).Send(lastObjectSel)

	if firstObj != 0 {
		firstStr := foundation.StringFrom(unsafe.Pointer(firstObj))
		fmt.Printf("   First object: %s\n", nsStringToGo(firstStr))
	}
	if lastObj != 0 {
		lastStr := foundation.StringFrom(unsafe.Pointer(lastObj))
		fmt.Printf("   Last object: %s\n", nsStringToGo(lastStr))
	}

	// Example 5: Mutable array modifications
	fmt.Println("\n5. Mutable Array Modifications:")

	// Add object
	newFruit := goStringToNS("Dragonfruit")
	objc.ID(mutableArray).Send(addObjectSel, newFruit.ID)
	fmt.Printf("   Added 'Dragonfruit'\n")

	// Insert object at index
	insertObjectSel := objc.RegisterName("insertObject:atIndex:")
	elderberry := goStringToNS("Elderberry")
	objc.ID(mutableArray).Send(insertObjectSel, elderberry.ID, uint64(1))
	fmt.Printf("   Inserted 'Elderberry' at index 1\n")

	// Remove object at index
	removeObjectAtIndexSel := objc.RegisterName("removeObjectAtIndex:")
	objc.ID(mutableArray).Send(removeObjectAtIndexSel, uint64(0))
	fmt.Printf("   Removed object at index 0\n")

	// Print final array
	finalCount := objc.ID(mutableArray).Send(countSel)
	fmt.Printf("   Final array (%d items):\n", finalCount)
	for i := uint64(0); i < uint64(finalCount); i++ {
		obj := objc.ID(mutableArray).Send(objectAtIndexSel, i)
		str := foundation.StringFrom(unsafe.Pointer(obj))
		fmt.Printf("     [%d]: %s\n", i, nsStringToGo(str))
	}

	// Example 6: Array enumeration
	fmt.Println("\n6. Array Enumeration:")

	fmt.Printf("   Fruits in array:\n")
	for i := uint64(0); i < uint64(finalCount); i++ {
		obj := objc.ID(mutableArray).Send(objectAtIndexSel, i)
		str := foundation.StringFrom(unsafe.Pointer(obj))
		fmt.Printf("   - %s\n", nsStringToGo(str))
	}

	// Example 7: Common use cases
	fmt.Println("\n7. Common Array Use Cases:")

	// Create array of file extensions
	extensionsArray := objc.ID(mutableArrayClass).Send(newSel)
	extensions := []string{".txt", ".pdf", ".jpg", ".png", ".doc"}

	fmt.Printf("   File extensions:\n")
	for _, ext := range extensions {
		extStr := goStringToNS(ext)
		objc.ID(extensionsArray).Send(addObjectSel, extStr.ID)
		fmt.Printf("   - %s\n", ext)
	}

	extCount := objc.ID(extensionsArray).Send(countSel)
	fmt.Printf("   Total extensions: %d\n", extCount)

	fmt.Println("\n✓ All Foundation NSArray operations completed successfully!")
	fmt.Println("\nNote: NSArray is immutable. Use NSMutableArray for arrays that can be modified.")
	fmt.Println("In this example, we used NSMutableArray to demonstrate array operations.")
}
