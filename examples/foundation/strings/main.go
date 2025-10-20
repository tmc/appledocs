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

// Helper to convert NSString to Go string
func nsStringToGo(nsStr foundation.String) string {
	if nsStr.ID == 0 {
		return ""
	}

	// Get UTF8 C string
	cStrSel := objc.RegisterName("UTF8String")
	cStr := nsStr.ID.Send(cStrSel)
	if cStr == 0 {
		return ""
	}

	// Convert C string to Go string
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

// Helper to convert Go string to NSString
func goStringToNS(s string) foundation.String {
	// Use stringWithUTF8String:
	cStr := append([]byte(s), 0) // null-terminated
	sel := objc.RegisterName("stringWithUTF8String:")
	class := objc.GetClass("NSString")
	result := objc.ID(class).Send(sel, unsafe.Pointer(&cStr[0]))
	return foundation.StringFrom(unsafe.Pointer(result))
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("Foundation NSString Examples")
	fmt.Println("============================")

	// Example 1: Create strings
	fmt.Println("\n1. Creating NSString Objects:")

	str1 := foundation.NewString()
	fmt.Printf("   Empty string: %v\n", str1)

	str2 := goStringToNS("Hello, World!")
	goStr := nsStringToGo(str2)
	fmt.Printf("   From Go string: \"%s\"\n", goStr)

	str3 := goStringToNS("Swift and Go together!")
	fmt.Printf("   Another string: \"%s\"\n", nsStringToGo(str3))

	// Example 2: String length
	fmt.Println("\n2. String Operations:")

	lengthSel := objc.RegisterName("length")
	length := str2.ID.Send(lengthSel)
	fmt.Printf("   Length of \"%s\": %d\n", goStr, length)

	// Example 3: String comparison
	fmt.Println("\n3. String Comparison:")

	compareStr1 := goStringToNS("Apple")
	compareStr2 := goStringToNS("Banana")
	compareStr3 := goStringToNS("Apple")

	compareSel := objc.RegisterName("compare:")
	result1 := compareStr1.ID.Send(compareSel, compareStr2.ID)
	result2 := compareStr1.ID.Send(compareSel, compareStr3.ID)

	fmt.Printf("   \"Apple\" compare \"Banana\": %d (negative = less than)\n", int64(result1))
	fmt.Printf("   \"Apple\" compare \"Apple\": %d (zero = equal)\n", int64(result2))

	// Example 4: String case conversion
	fmt.Println("\n4. Case Conversion:")

	original := goStringToNS("Hello World")

	uppercaseSel := objc.RegisterName("uppercaseString")
	uppercase := objc.ID(original.ID.Send(uppercaseSel))
	fmt.Printf("   Original: \"%s\"\n", nsStringToGo(original))
	fmt.Printf("   Uppercase: \"%s\"\n", nsStringToGo(foundation.StringFrom(unsafe.Pointer(uppercase))))

	lowercaseSel := objc.RegisterName("lowercaseString")
	lowercase := objc.ID(original.ID.Send(lowercaseSel))
	fmt.Printf("   Lowercase: \"%s\"\n", nsStringToGo(foundation.StringFrom(unsafe.Pointer(lowercase))))

	capitalizedSel := objc.RegisterName("capitalizedString")
	capitalized := objc.ID(original.ID.Send(capitalizedSel))
	fmt.Printf("   Capitalized: \"%s\"\n", nsStringToGo(foundation.StringFrom(unsafe.Pointer(capitalized))))

	// Example 5: String contains
	fmt.Println("\n5. Substring Search:")

	searchStr := goStringToNS("The quick brown fox jumps over the lazy dog")
	containsSel := objc.RegisterName("containsString:")

	search1 := goStringToNS("quick")
	contains1 := searchStr.ID.Send(containsSel, search1.ID)
	fmt.Printf("   Contains \"quick\": %v\n", contains1 != 0)

	search2 := goStringToNS("slow")
	contains2 := searchStr.ID.Send(containsSel, search2.ID)
	fmt.Printf("   Contains \"slow\": %v\n", contains2 != 0)

	// Example 6: String prefixes and suffixes
	fmt.Println("\n6. Prefix/Suffix Checking:")

	filename := goStringToNS("document.pdf")

	hasPrefix := objc.RegisterName("hasPrefix:")
	hasSuffix := objc.RegisterName("hasSuffix:")

	prefixDoc := goStringToNS("document")
	suffixPdf := goStringToNS(".pdf")
	suffixTxt := goStringToNS(".txt")

	fmt.Printf("   Filename: \"%s\"\n", nsStringToGo(filename))
	fmt.Printf("   Starts with \"document\": %v\n", filename.ID.Send(hasPrefix, prefixDoc.ID) != 0)
	fmt.Printf("   Ends with \".pdf\": %v\n", filename.ID.Send(hasSuffix, suffixPdf.ID) != 0)
	fmt.Printf("   Ends with \".txt\": %v\n", filename.ID.Send(hasSuffix, suffixTxt.ID) != 0)

	// Example 7: Common string patterns
	fmt.Println("\n7. Common String Patterns:")

	strings := []string{
		"com.apple.foundation",
		"/Users/username/Documents",
		"user@example.com",
		"192.168.1.1",
		"2025-10-19",
	}

	for _, s := range strings {
		nsStr := goStringToNS(s)
		length := nsStr.ID.Send(lengthSel)
		fmt.Printf("   %-30s (length: %2d)\n", s, length)
	}

	fmt.Println("\n✓ All Foundation NSString operations completed successfully!")
	fmt.Println("\nNote: NSString is the Foundation framework's immutable string class.")
	fmt.Println("For mutable strings, use NSMutableString.")
}
