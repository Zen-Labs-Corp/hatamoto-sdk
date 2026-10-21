# Hatamoto Plugin SDK

`hatamoto.dev/sdk` is the plugin contract: hook names and
payload types (`events.go`, `hooks.go`), the `plugin.json` manifest
(`manifest.go`) and the export-name rules of the WASM ABI. The root package depends
on nothing but the standard library, never imports Hatamoto core packages, and the
module has its own `go.mod`.

```sh
go get hatamoto.dev/sdk@v0.11.0
```

Releases are plain `vX.Y.Z` tags of this repository, where `X.Y.Z` is the `Version`
constant in `version.go`. This repository is a one-way export of the `backend/pkg/sdk`
directory of the Hatamoto core repository, which stays the source of truth, so
changes are made upstream in Hatamoto core and exported here. Report SDK bugs and
send proposals as issues or pull requests on this repository
([github.com/zenlabscorp/hatamoto-sdk](https://github.com/zenlabscorp/hatamoto-sdk));
a pull request is applied upstream and arrives back in the next release.

`hatamoto.dev/sdk/guest` is the Go guest layer: a thin wrapper
over the Extism Go PDK (`github.com/extism/go-pdk`, the module's only dependency) that provides
`OnAction` / `OnFilter` registration and the JSON plumbing. It only builds the
PDK side for `GOOS=wasip1`.

Everything about writing a plugin (export names, JSON schemas of every payload,
manifest, limits, the Go quick start) is in the plugin documentation on
<https://hatamoto.dev/docs> (the plugin API reference ships with Hatamoto core). The first-party plugins
[gtm](https://github.com/zenlabscorp/hatamoto-plugin-gtm) and
[searchconsole](https://github.com/zenlabscorp/hatamoto-plugin-searchconsole) are real
examples built on this module.

## Versioning

Semantic. Minor releases are additive only: new hooks, new optional payload
fields, new manifest fields. Removing or renaming a hook, an export name or a
field is a breaking change and bumps the major version (before 1.0, the minor).
`Version` is the current version; `CompatibleWith` lists earlier versions whose
plugins this SDK still serves.

| SDK | Change |
|---|---|
| 0.11.0 | `builder.data` filter + permission and `builder_sources` manifest field: plugin data sources (values and item lists) for the page builder |
| 0.10.0 | Entity host functions: `parent_id` on create/update, `parent_id`/`path`/`url` on Entity, `entities.get` by `path` or `parent_id`+`slug`(+`schema_slug`) |
| 0.5.0 | Native Go plugins removed. Plugins are sandboxed WASM modules; payloads are JSON with snake_case fields; `guest` package added; no `pgx` or router types |
| 0.4.0 and earlier | Native, compiled-in plugins (gone) |

## Test

```sh
go vet ./... && go test ./...
GOOS=wasip1 GOARCH=wasm go vet ./guest
```
