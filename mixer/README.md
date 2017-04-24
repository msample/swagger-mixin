# mixer
--
    import "github.com/msample/swagger-mixin/mixer"

mixer provides functions to merge Swagger 2.0 specs parsed into
https://github.com/go-openapi/spec format into one spec

Use cases include adding independently versioned metadata APIs to application
APIs for microservices.

Typically, multiple APIs to the same service instance is not a problem for
client generation as you can create more than one client to the service from the
same calling process (one for each API).

Server skeleton generation, ie generating the model & marshaling code, http
server instance etc. from Swagger, becomes easier with a merged spec for some
tools & target-languages. Server code generation tools that natively support
hosting multiple specs in one server process will not need this tool.

## Usage

#### func  FixEmptyDesc

```go
func FixEmptyDesc(rs *spec.Response)
```
FixEmptyDesc adds "(empty)" as the description to the given Response object if
it doesn't already have one and isn't a ref. No-op on nil input.

#### func  FixEmptyDescs

```go
func FixEmptyDescs(rs *spec.Responses)
```
FixEmptyDescs adds "(empty)" as the description for any Response in the given
Responses object that doesn't already have one.

#### func  FixEmptyResponseDescriptions

```go
func FixEmptyResponseDescriptions(s *spec.Swagger)
```
FixEmptyResponseDescriptions replaces empty ("") response descriptions in the
input with "(empty)" to ensure that the resulting Swagger is stays valid. The
problem appears to arise from reading in valid specs that have a explicit
response description of "" (valid, response.description is required), but due to
zero values being omitted upon re-serializing (omitempty) we lose them unless we
stick some chars in there.

#### func  Mixin

```go
func Mixin(primary *spec.Swagger, mixins ...*spec.Swagger) uint
```
Mixin modifies the primary swagger spec by adding the paths and definitions from
the mixin specs. Top level parameters and responses from the mixins are also
carried over. Operation id collisions are avoided by appending "Mixin<N>" but
only if needed. No other parts of primary are modified. Consider calling
FixEmptyResponseDescriptions() on the modified primary if you read them from
storage and they are valid to start with.

Entries in "paths", "definitions", "parameters" and "responses" are added to the
primary in the order of the given mixins. If the entry already exists in primary
it is skipped with a warning message.

The count of skipped entries (from collisions) is returned so any deviation from
the number expected can flag warning in your build scripts. Carefully review the
collisions before accepting them; consider renaming things if possible.

No normalization of any keys takes place (paths, type defs, etc). Ensure they
are canonical if your downstream tools do key normalization of any form.

#### func  MixinFiles

```go
func MixinFiles(primaryFile string, mixinFiles []string, w io.Writer) (uint, error)
```
MixinFiles is a convenience function for Mixin that reads the given swagger
files, adds the mixins to primary, calls FixEmptyResponseDescriptions on the
primary, and writes the primary with mixins to the given writer in JSON. Returns
the number of collsions that occured from mixins and any error.
