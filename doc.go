// Package appledocs provides programmatic access to Apple's documentation.
//
// The package offers two complementary APIs:
//
//  1. Typed API (recommended) - compile-time safe, clean usage
//  2. Map API (for flexibility) - access any field, handles edge cases
//
// # Typed API
//
// The typed API provides structured types and query functions:
//
//	fsys, err := appledocs.Open("docs/tutorials/data/documentation")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// List all frameworks
//	frameworks, _ := appledocs.ListFrameworks(fsys)
//	for _, fw := range frameworks {
//		fmt.Println(fw)
//	}
//
//	// Load a symbol
//	doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
//	fmt.Println(doc.Metadata.Title)        // NSString
//	fmt.Println(doc.Metadata.SymbolKind)   // class
//
//	// Iterate methods (all typed, no casts!)
//	for id, ref := range doc.References {
//		if ref.Role == "symbol" && ref.SymbolKind == "method" {
//			fmt.Println(ref.Title)
//		}
//	}
//
// # Map API
//
// The map API provides flexibility for accessing any field:
//
//	fsys, _ := appledocs.Open("docs/")
//	raw, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
//
//	title := appledocs.Title(raw)
//	kind := appledocs.SymbolKind(raw)
//
//	// Access custom/undocumented fields
//	customField := appledocs.GetString(raw, "custom", "nested", "field")
//
// # When to Use Each API
//
// Use the typed API when:
//   - Writing production code
//   - You know what fields you need
//   - You want compile-time safety
//   - You value clean, maintainable code
//
// Use the map API when:
//   - Exploring unknown structure
//   - Accessing undocumented fields
//   - Handling polymorphic data
//   - Writing quick scripts
//
// # File System Structure
//
// The documentation follows this structure:
//
//	docs/
//	├── Foundation.json           # Framework root
//	├── Foundation/
//	│   ├── NSString.json        # Class documentation
//	│   ├── NSArray.json
//	│   └── ...
//	├── UIKit.json
//	├── UIKit/
//	│   └── ...
//	└── ...
//
// Framework files are at the root. Symbol documentation is nested
// within framework directories.
//
// # Performance
//
// The package uses lazy loading via fs.FS. Files are read on-demand,
// so you can work with large documentation sets without loading
// everything into memory.
package appledocs
