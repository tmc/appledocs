# Apple Docs API Server

A REST API and GraphQL-like server that provides access to the Apple documentation that has been crawled and cached by the main appledocs tool.

## Requirements

- Go 1.18+ installed
- The `.cache` directory populated with Apple documentation (run `./appledocs -mode crawl` from the main directory first)

## Installation

```bash
# Build the server
go build -o appledocs-gql

# Run the server with default settings
./appledocs-gql
```

## Configuration

Command line flags:

- `-port`: HTTP server port (default: 8080)
- `-cache`: Directory containing cached documentation (default: ".cache")

Example:
```bash
./appledocs-gql -port 3000 -cache /path/to/cache/directory
```

## REST API Endpoints

### Get Document by Path

Retrieves a document by its path:

```
GET /api/document?path=tutorials/data/documentation/swiftui
```

### Search Documents

Searches for documents matching a query:

```
GET /api/search?q=SwiftUI&limit=5
```

### List Frameworks

Lists all top-level frameworks:

```
GET /api/frameworks
```

## GraphQL Interface

The server provides a full GraphQL-like interface that allows for more structured and flexible data querying.

You can access two different GraphQL interfaces:

- **GraphiQL Playground**: `http://localhost:8080/graphql` - The classic GraphQL IDE
- **Apollo Sandbox**: `http://localhost:8080/sandbox` - A more modern GraphQL development environment with additional features

### Schema

The GraphQL schema provides a rich set of types and queries:

- **Document**: Represents an Apple documentation item with fields like id, path, title, abstract, type, framework, platforms, etc.
- **Platform**: Represents platform availability information (iOS, macOS, etc.)
- **Section**: Represents a section of a document
- **Framework**: Represents framework information
- **Category**: Represents a document category
- **SearchResult**: Contains search results information

### Available Queries

- `document(path: String!)`: Get a document by its path
- `search(query: String!, limit: Int)`: Search for documents matching a query
- `frameworks`: List all top-level frameworks
- `framework(id: String!)`: Get a specific framework by its ID
- `documentsByType(type: String!, framework: String)`: Get documents of a specific type
- `documentsByPlatform(platform: String!, framework: String)`: Get documents for a specific platform
- `relatedDocuments(path: String!, limit: Int)`: Get related documents for a specific document

### Example Queries

```graphql
# Get a document by path with detailed information
{
  document(path: "tutorials/data/documentation/swiftui") {
    id
    title
    abstract
    type
    framework
    platforms {
      name
      introducedAt
    }
    isDeprecated
  }
}

# Search for documents with pagination
{
  search(query: "SwiftUI", limit: 5) {
    documents {
      id
      title
      abstract
    }
    count
    query
  }
}

# Get all frameworks with platform info
{
  frameworks {
    id
    title
    abstract
    platforms {
      name
    }
  }
}

# Get a specific framework
{
  framework(id: "swiftui") {
    id
    title
    abstract
  }
}

# Get class documents from SwiftUI
{
  documentsByType(type: "class", framework: "swiftui") {
    id
    title
    abstract
  }
}

# Get iOS-specific documents
{
  documentsByPlatform(platform: "ios") {
    id
    title
    abstract
  }
}

# Get documents related to SwiftUI
{
  relatedDocuments(path: "tutorials/data/documentation/swiftui", limit: 3) {
    id
    title
    abstract
  }
}
```

## JSON Schema

### Document

```json
{
  "id": "string",
  "path": "string",
  "title": "string",
  "abstract": "string",
  "metadata": {
    // Document metadata as a JSON object
  },
  "content": {
    // Full document content as a JSON object
  }
}
```

### Search Response

```json
{
  "results": [
    // Array of Document objects
  ],
  "count": 5,
  "query": "SwiftUI",
  "limit": 10
}
```

## Integration Examples

### Fetch data with JavaScript

```javascript
async function fetchFrameworks() {
  const response = await fetch('http://localhost:8080/api/frameworks');
  const frameworks = await response.json();
  console.log(frameworks);
}

async function searchDocuments(query, limit = 10) {
  const response = await fetch(`http://localhost:8080/api/search?q=${encodeURIComponent(query)}&limit=${limit}`);
  const results = await response.json();
  console.log(results);
}

async function getDocument(path) {
  const response = await fetch(`http://localhost:8080/api/document?path=${encodeURIComponent(path)}`);
  const document = await response.json();
  console.log(document);
}
```

### Fetch data with curl

```bash
# Get all frameworks
curl -X GET http://localhost:8080/api/frameworks

# Search for documents
curl -X GET "http://localhost:8080/api/search?q=SwiftUI&limit=5"

# Get a document by path
curl -X GET "http://localhost:8080/api/document?path=tutorials/data/documentation/swiftui"
```

## Architecture

The server follows a simple architecture:

1. REST API with three main endpoints:
   - `/api/document`: Retrieve a document by path
   - `/api/search`: Search for documents
   - `/api/frameworks`: List all frameworks

2. GraphQL-like interface:
   - Uses GraphiQL as the UI
   - Translates GraphQL queries to REST API calls
   - Provides a familiar GraphQL syntax for clients

3. Document service:
   - Reads cached JSON files from disk
   - Extracts relevant information for each document
   - Provides search functionality

## Performance Notes

- The server performs file system operations to retrieve documents, which may be slow for large documentation sets
- Search is implemented with a simple substring search, which may be inefficient for large queries
- Consider implementing a proper indexing solution for production use