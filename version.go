// Package sdk is the Hatamoto plugin contract: hook names and payload DTOs,
// the plugin.json manifest, and the export-name rules of the WASM ABI. It has
// no dependencies beyond the standard library, so hosts and guests share it.
// Subpackage guest is the Go guest layer over the Extism PDK.
//
// Versioning: semantic. Minor releases are additive only; breaking changes
// bump the major version (before 1.0, the minor). A manifest declares the
// range it was built against and the host refuses a plugin whose range does
// not contain Version.
package sdk

// Version is the SDK semantic version. 0.5.0 replaced native Go plugins with
// sandboxed WASM modules; 0.6.0 adds capabilities (permissions, host
// functions, routes, cron, template funcs) and the form.submitted and
// render.page hooks; 0.7.0 renames the HeadContext and SitemapURLs
// per-site field to `language`; 0.8.0 makes sitemap.urls additions-only,
// turns the denied-capability error into
// {"code":"capability_denied","capability":"<name>"} and replaces
// limits.timeout_ms with per-hook limits.hook_timeouts_ms (see
// docs/PLUGIN_API.md); 0.9.0 slims the render.page payload (blocks by id,
// the host re-attaches content; `render_page.full_blocks` opts back in);
// 0.10.0 grows the entities host functions: `parent_id` on create and update,
// `parent_id`/`path`/`url` on the Entity, and entities.get by `path` or by
// `parent_id` + `slug`.
//
// 0.11.0 adds the builder.data filter, the builder.data permission, the
// builder_sources manifest field (plugin data sources for the page builder)
// and the builder.island permission (a widget pack that ships island.js).
const Version = "0.11.0"

// CompatibleWith lists earlier versions whose plugins this SDK still serves.
// Caret ranges before 1.0 stop at the next minor, so a purely additive minor
// lists its predecessors here. 0.4 plugins were native Go.
var CompatibleWith = []string{"0.5.0", "0.6.0", "0.7.0", "0.8.0", "0.9.0", "0.10.0"}
