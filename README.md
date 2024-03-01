# go-builder-generator <!-- omit in toc -->

- [How to use ?](#how-to-use-)
- [Commands](#commands)
  - [Generate](#generate)
    - [Tags](#tags)

## How to use ?

```sh
go install github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest
```

## Commands

CLI commands:

- `generate`: generates the builders for given go file with structs. Options are:
  - `--dest -d` (optional): destination package for the builders (if not provided, the builders will be generated in package where the structs are located). Path can be relative.
  - `--file -f` (required): filename where the structs to generate builders are located. Path can be relative.
  - `--no-notice` (optional): removes the `Code generated by` notice at the top of builders file.
  - `--use-validator` (optional): allows the use of https://github.com/go-playground/validator to validate structs in `Build` functions (see [tags](#tags) for an example).
  - `--structs -s` (required): structs name to generate builders on.

### Generate

Once installed, it's easy to use the builder generator. For it you just need to add a `//go:generate` section in your go files. You can also directly use the CLI inside a terminal but it's clearly easier to use directly inside go files.

```go
package pkg

//go:generate go-builder-generator generate --file/-f filename.go --structs/-s StructName,AnotherStructName --dest/-d <path/to/destination_package>

type StructName struct {}

type AnotherStructName struct {}
```

#### Tags

It's possible to tune builders generation with the struct tag `builder`. The available options are:

- `ignore`: when provided, the field will be ignored in builder generation.
- `append`: when provided on a slice (only), the generated builder will append instead of reaffecting.
- `pointer`: when provided on a pointer, the generated builder will keep it's pointer on the input parameter.
- `default_func`: when provided, an additional function will be generated in another file suffixed with `_impl.go` to allow manual affectation of field (or even other fields).

Additionally, having the `validate` tag on at least one of the struct (in addition to `--use-validator` CLI argument) fields will add https://github.com/go-playground/validator on the builder `Build` function.

Example:

```go
package pkg

type StructName struct {
	Pointer               *int64   `builder:"pointer"` // generated builder will be 'SetPointer(pointer *int64)'
	NoPointer             *int64   // generated builder will be 'SetNoPointer(noPointer int64)'
	ASlice                []string `builder:"append"` // generated builder will be 'SetASlice(aSlice ...string)', additionally the affectation will be `b.ASlice = append(b.ASlice, aSlice...)`
	NoAppend              []string // generated builder will be 'SetASlice(noAppend []string)', additionally the affectation will be `b.NoAppend = noAppend`
	Ignore                int64    `builder:"ignore"`                            // no builder will be generated on this field
	DefaultFunc           int64    `builder:"default_func=SomeFuncName"`         // an additional function named 'SomeFuncName' will be generated in target package file '_impl.go' and associated to builder struct
	IgnoreWithDefaultFunc int64    `builder:"ignore,default_func=SomeOtherFunc"` // no builder will be generated and the additional function will be generated
	UseValidator          *string  `validate:"required"`                         // validator from go-playground (https://github.com/go-playground/validator) will be added to Build function to validate the whole struct (only if --use-validator argument is given)
}
```

**Note:** `append` option and `pointer` option are exclusive with a priority for `append` if both provided. Also if `append` is provided on a field not being a slice, it will just be ignored.

For more examples, you can check in `examples` package at project root !