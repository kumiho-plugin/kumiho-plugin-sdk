# kumiho-plugin-sdk

SDK for building Kumiho metadata plugins.

---

## Package layout

```
capability/   Capability constants (metadata.search, metadata.fetch, …)
config/       ConfigSchema, MigrationDescriptor
errors/       ErrorCode, PluginError
healthcheck/  Status, Response, Checker interface, timeout/retry conventions
manifest/     Manifest, TrustLevel, RuntimeType, Platform, Artifact, IconSet
plugin/       MetadataPlugin interface
service/      ServiceRuntime HTTP contract (path/status constants)
state/        Plugin lifecycle state constants
types/        SourceRef, SearchRequest/Response, FetchRequest/Response
version/      Versioning policy + IsCompatible helper
```

---

## Quickstart: BinaryRuntime plugin

### 1. Implement the interface

```go
package main

import (
    "context"

    "github.com/kumiho-plugin/kumiho-plugin-sdk/capability"
    plerrors "github.com/kumiho-plugin/kumiho-plugin-sdk/errors"
    "github.com/kumiho-plugin/kumiho-plugin-sdk/healthcheck"
    "github.com/kumiho-plugin/kumiho-plugin-sdk/manifest"
    "github.com/kumiho-plugin/kumiho-plugin-sdk/types"
    "github.com/kumiho-plugin/kumiho-plugin-sdk/version"
)

type MyPlugin struct{ apiKey string }

func (p *MyPlugin) Search(ctx context.Context, req *types.SearchRequest) (*types.SearchResponse, error) {
    if p.apiKey == "" {
        return &types.SearchResponse{
            Error: plerrors.New(plerrors.ErrCodeUnauthorized, "api key not configured"),
        }, nil
    }
    // Call your provider API, build candidates...
    return &types.SearchResponse{
        Candidates: []types.SearchCandidate{
            {
                Source:     types.SourceRef{ID: "vol_abc123", Name: "myprovider"},
                Title:      "Example Book",
                Score:      0.95,
                Confidence: 0.90,
                Reason:     "title exact match",
            },
        },
    }, nil
}

func (p *MyPlugin) Fetch(ctx context.Context, req *types.FetchRequest) (*types.FetchResponse, error) {
    // Use req.Source.ID to retrieve full details.
    return &types.FetchResponse{
        Result: &types.MetadataResult{
            Source:      req.Source,
            Title:       "Example Book",
            Authors:     []string{"Author Name"},
            Description: "A book about things.",
            Identifiers: map[string]string{"isbn13": "9780000000000"},
            Characters: []types.MetadataCharacter{
                {
                    ID:   "char_1",
                    Name: "Example Hero",
                    Role: "main",
                },
            },
        },
    }, nil
}

func (p *MyPlugin) Healthcheck(ctx context.Context) (*healthcheck.Response, error) {
    return &healthcheck.Response{Status: healthcheck.StatusOK, Version: "1.0.0"}, nil
}

func (p *MyPlugin) Manifest() *manifest.Manifest {
    return &manifest.Manifest{
        ID:                  "kumiho-plugin-metadata-myprovider",
        Name:                "My Provider",
        Version:             "1.0.0",
        Author:              "you",
        Icons:               &manifest.IconSet{SVG: "https://example.com/icon.svg", PNG: "https://example.com/icon.png"},
        TrustLevel:          manifest.TrustLevelUnverified,
        RuntimeType:         manifest.RuntimeTypeBinary,
        SupportedPlatforms:  []manifest.Platform{manifest.PlatformLinuxBinary},
        Capabilities:        []capability.Capability{capability.MetadataSearch, capability.MetadataFetch},
        MinCoreVersion:      "0.1.0",
        ConfigSchemaVersion: "1",
        SDKVersion:          version.SDK,
    }
}
```

### 1.1 Visual metadata

Use `Manifest.Icons` to provide plugin card artwork for the core UI.

- Prefer `svg` for scalable brand icons.
- Provide `png` as a fallback for clients that do not render SVG.
- Keep `Author` filled; plugin cards are expected to show both icon and author.

The core may dim inactive plugins or add glow/pulse effects for active plugins, so
the icon should remain recognizable on dark backgrounds.

### 2. Error handling rules

Return a non-nil response with the `Error` field set — never a bare Go error — for provider-side failures.

```go
// Transient: core will retry
return &types.SearchResponse{
    Error: plerrors.NewRetryable(plerrors.ErrCodeRateLimited, "quota exceeded"),
}, nil

// Permanent
return &types.SearchResponse{
    Error: plerrors.New(plerrors.ErrCodeUnauthorized, "invalid api key"),
}, nil
```

### 3. Declare config schema

```go
import "github.com/kumiho-plugin/kumiho-plugin-sdk/config"

var schema = config.Schema{
    Version: "1",
    Fields: []config.ConfigField{
        {Key: "api_key",  Type: config.FieldTypeSecret,  Label: "API Key", Required: true},
        {Key: "language", Type: config.FieldTypeSelect,  Label: "Language",
         Options: []string{"ko", "en", "ja"}, Default: "ko"},
    },
}
```

---

## Quickstart: ServiceRuntime plugin (Docker / HTTP)

Expose the four endpoints defined in `service`. Use the package constants for paths and status codes.

```go
import "github.com/kumiho-plugin/kumiho-plugin-sdk/service"

// Paths:
//   service.PathSearch   = "/search"    POST  types.SearchRequest  → types.SearchResponse
//   service.PathFetch    = "/fetch"     POST  types.FetchRequest   → types.FetchResponse
//   service.PathHealth   = "/health"    GET   —                   → healthcheck.Response
//   service.PathManifest = "/manifest"  GET   —                   → manifest.Manifest

// All responses use service.StatusOK (200) — even when the body contains an error field.
// Use service.StatusBadRequest (400) for malformed JSON.
// Use service.StatusInternalError (500) for unexpected panics.
// Set Content-Type: service.ContentTypeJSON on all responses.
```

The plugin must be ready within `plugin.StartupTimeout` (15s) of process/container start.

---

## Healthcheck conventions

| Constant                          | Default |
|-----------------------------------|---------|
| `healthcheck.DefaultTimeout`      | 5s      |
| `healthcheck.DefaultInterval`     | 30s     |
| `healthcheck.DefaultMaxRetries`   | 3       |
| `healthcheck.DefaultRetryBackoff` | 2s      |

After `DefaultMaxRetries` consecutive failures the core transitions the plugin to `state.Unhealthy`.

---

## Versioning policy

Both core and plugins follow `vMAJOR.MINOR.PATCH` semver.

- Declare `MinCoreVersion` / `MaxCoreVersion` in your manifest.
- The core calls `version.IsCompatible` at activation.
- MAJOR bumps are breaking; update `MaxCoreVersion` explicitly.
- MINOR / PATCH within the same MAJOR are backward-compatible.

```go
ok, _ := version.IsCompatible("1.3.0", "1.0.0", "1.99.99") // ok == true
```

Current SDK version: `version.SDK`

---

## Config schema migration

When updating a plugin changes its schema, include a `config.MigrationDescriptor`.
The **core** executes migration; the plugin must never modify stored config directly.

```go
migration := config.MigrationDescriptor{
    FromVersion: "1",
    ToVersion:   "2",
    Ops: []config.MigrationOp{
        {Type: config.OpTypeRenameField, OldKey: "api_key",      NewKey: "provider_api_key"},
        {Type: config.OpTypeAddField,    NewKey: "timeout_secs", Default: 10},
        {Type: config.OpTypeRemoveField, OldKey: "legacy_flag"},
    },
}
```

Set `ManualReviewRequired: true` when automated migration is not possible.

Supported operations:

| Op              | Effect                          |
|-----------------|---------------------------------|
| `rename_field`  | Renames a key                   |
| `add_field`     | Adds a key with a default value |
| `remove_field`  | Removes a deprecated key        |
| `move_secret`   | Moves a secret to a new key     |

---

## Plugin state lifecycle

```
not_installed → downloading → installed → registered → activation_pending → active
                                                                          ↓
                                                               disabled / unhealthy / error / incompatible
```

`state.IsRunning(s)` — true for `active` and `unhealthy`.
`state.IsInstalled(s)` — true for all states where the artifact is on disk.
