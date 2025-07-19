# API Project Setup

0. Tips

Go doesn't provide a general project structure and a automatic project generation.

`Go tools` providses the central way to run, build, test and install packages/dependencies.

`go list -m all` list all dependencies of the module

`go.mod` containes the module name/path. Go version. And dependencies/packages. Kind of `pom.xml` or `requirements.txt`.

`go.sum` verifies dependencies/packages with hashes. Kind of `package-lock.json` or `poetry.lock`.

Check `go.sum`-integrity with `go mod verify`

`go build` builds the project/module and creates a binary. Won't install the binary on the system

`go run` temporary builds the binary to be able to run the module.

`go install` compiles and installs a go program onto the system, using the `$GOBIN` (default `~/go/bin`)

`go get` adds a dependency/module/package to the project (`go.mod`)

`go mod tidy` convenience to remove unused dependencies and add missing ones.

1. Install Go

Install `Go` on you local machine.

[Doc](https://go.dev/doc/install)

2. Init Project

First, move into the repository.

Second, init the project.

Follow the naming conventions for best practises, [see](https://go.dev/doc/modules/managing-dependencies#naming_module).

Common: `[repo]/[company name / project name]/[function like api/app]`

Like: `github.com/kprueger/circumitus-api`

Create the module (`go.mod` file):
```bash
go mod init [package/project name]
```

Create a `foo.go`:
```bash
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
```

Run the package:
```bash
go run .
```
or
```bash
go run cmd/circumitus-api/main.go
```

### Add a dependency

Use `go get/install [module-name@latest/version]`

or 

`go get .` to add all depenencies for a package in your module.

Search for packages on [pkg.go.dev](https://pkg.go.dev/)

### Project structure

Best practice to structure your module.

See [repo](https://github.com/golang-standards/project-layout) for best practice guide!!!

```plaintext
waste-api/
├── cmd/
│ └── server/
│    └── main.go    ← Entry point of the application
├── internal/       ← Private application packages (e.g. db, auth, logic)
├── graph/          ← gqlgen schema, resolvers, models
├── go.mod
└── go.sum
```

| Folder | Purpose |
|------------------|-------------------------------------------------------------------------|
| `cmd/` | Entry points for executables (API server, CLI tools, workers) |
| `internal/` | Private code only accessible within this module | 
| `pkg/` | Public, reusable packages (if needed outside your module) | 
| `graph/` | GraphQL schema, resolvers, and model definitions (for gqlgen) | 
| `config/` | Configuration loading (env files, YAML, etc.) | 
| `scripts/` | Shell scripts, Makefiles, or other automation | 
| `api/` (optional)| API documentation (GraphQL playground, OpenAPI specs) | 
| `migrations/` | SQL database migrations (e.g. for use with golang-migrate) |

### Endpoint

Check GraphQL [doc](https://graphql.org/learn/)

### Setup GraphQL

Used library: [repo](https://github.com/99designs/gqlgen?tab=readme-ov-file)

install the module
```bash
go install github.com/99designs/gqlgen@latest 
```

Update dependencies
```bash
go mod tidy
```

Init the module
```bash
gqlgen init
```
or
```bash
go run github.com/99designs/gqlgen init 
```

Modify of create `[Repository-Root]/graph/schema.graphqls`:
```yml
type Chip {
    id: ID!
    persons: [User]!
    createdAt: String!
    password: String!
}
```

Create `[Repository-Root]/gqlgen.yml`:
```yml
# gqlgen.yml
schema:
  - graph/schema.graphqls
exec:
  filename: graph/generated/generated.go
  package: generated
model:
  filename: graph/model/models_gen.go
  package: model
resolver:
  layout: follow-schema
  dir: graph
  package: graph
```

Generate the module structure. It is also used to encounter schema changes.
```bash
gqlgen generate
```

Next, implement the `graph/resolver.go`

## Setup Instructions (Development)

1. Copy the environment file:
```bash
   cp .env.example .env
```

2. Edit .env to match your local configuration.

3. Run the application:
```bash
go run cmd/circumitus-api/main.go
```

4. Test Query
```bash
mutation {
  createChip(id: "314242", password: "password") {
    id
    password
  }
}
```

5. In case the schema changes run
```bash
gqlgen generate
```