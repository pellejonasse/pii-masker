# pii-masker

A reflection-based PII masking library for Go. Given any value — struct, map, slice, pointer, or interface — it produces a deep copy with sensitive fields masked, anonymized, or passed through unchanged based on struct field tags. The original value is never mutated.

## Quick Start

```go
masker := piimasker.NewMasker(piimasker.MaskerConfig{})

type User struct {
    ID    string `Pii:"show"`
    Email string `Pii:"mask"`
    Age   int    `Pii:"mask"`
}

user := User{ID: "u-123", Email: "alice@example.com", Age: 31}
safe := masker.Mask(user).(User)
// safe → {ID: "u-123", Email: "***************", Age: 0}
```

## Configuration

`NewMasker` accepts a `MaskerConfig`:

```go
type MaskerConfig struct {
    MaxPiiStringLength int // max length of masked/anonymized string output; defaults to 100
}
```

`MaxPiiStringLength` caps two things:
- the number of `*` characters written for a masked string
- the length of the random string generated for an anonymized string

A per-call override can be passed as an optional argument to `Mask`:

```go
// Use the masker's default config
result := masker.Mask(user)

// Override config for this call only
result := masker.Mask(user, piimasker.MaskerConfig{MaxPiiStringLength: 20})
```

## Field Tags

Annotate struct fields with the `Pii` tag:

| Tag | Behaviour |
|---|---|
| `Pii:"mask"` | Replaces the value with its zero representation. Strings become a sequence of `*` characters matching the original length (up to `MaxPiiStringLength`). Numbers become `0`. Booleans become `false`. |
| `Pii:"show"` | Copies the original value unchanged. |
| `Pii:"anonymize"` | Replaces the value with a random value of the same shape. Strings get a random alphanumeric string of the same length. Numbers get a random value with the same number of digits. Booleans get a random bool. |
| *(absent)* | Inherits the mode propagated from the parent field (see [Propagation](#propagation-and-tag-priority)). If no mode has been propagated, the field is **masked** — untagged fields are never shown by default. |

```go
type PaymentCard struct {
    CardNumber string `Pii:"mask"`      // → "****************"
    HolderName string `Pii:"anonymize"` // → "xK9mPqRtZvBl3nYw"
    Currency   string `Pii:"show"`      // → "GBP"
    IsDefault  bool   `Pii:"mask"`      // → false
}
```

## Propagation and Tag Priority

When the masker descends into a nested struct, the mode from the containing field's tag is **propagated** into all sub-fields that have no tag of their own. A sub-field that has an explicit tag **overrides** the propagated mode for its own subtree — lower-level (deeper) explicit tags always win.

```go
type Address struct {
    Street  string // no tag — inherits parent mode
    City    string `Pii:"show"` // explicit override: always shown regardless of parent
    Country string // no tag — inherits parent mode
}

type User struct {
    Name    string  `Pii:"mask"`
    Address Address `Pii:"mask"` // propagates "mask" into Address
}
```

With the struct above:
- `User.Name` → masked
- `Address.Street` → masked (inherited from `User.Address`)
- `Address.City` → shown (explicit `show` overrides inherited `mask`)
- `Address.Country` → masked (inherited from `User.Address`)

### Propagation into External Structs

This makes it straightforward to handle structs from third-party packages where you cannot add tags. Tag the field that holds the external struct and the mode floods into all of its fields:

```go
import "github.com/aws/aws-lambda-go/events"

type EventWrapper struct {
    Event  events.CloudWatchEvent `Pii:"mask"` // masks all string/numeric fields inside CloudWatchEvent
    UserID string                 `Pii:"mask"`
}
```

## Performance: Mask vs Anonymize

**Prefer `mask` over `anonymize` when you only need to prevent PII from appearing in logs.** Masking is a simple zero-fill operation; anonymization requires generating random values, which is measurably slower — approximately 3× in microbenchmarks.

Only reach for `anonymize` when the output needs to look structurally plausible (e.g. for realistic test data or analytics pipelines that require valid-looking values).

Struct type metadata is parsed via reflection only once per unique type and cached in a `sync.Map`, so the per-call overhead is minimal at steady state. The masker is safe for concurrent use.

### Benchmarks

All figures are for **1,000 invocations** on an Apple M-series chip (`-12` = 12 logical cores). The full-struct benchmarks (`BenchmarkMask`) operate on a deeply nested `Person` fixture with slices, maps, and multiple levels of nesting.

#### Full struct (BenchmarkMask)

| Mode | Time / 1k calls | Memory / 1k calls | Allocs / 1k calls |
|---|---|---|---|
| `mask` | ~2.24 ms | ~2.80 MB | 45,000 |
| `show` | ~1.21 ms | ~1.98 MB | 15,000 |
| `anonymize` | ~6.63 ms | ~2.80 MB | 45,000 |
| no tag (default mask) | ~1.63 ms | ~2.80 MB | 45,000 |

#### Individual value anonymization (BenchmarkAnonymize)

| Type | Time / 1k calls | Memory / 1k calls | Allocs / 1k calls |
|---|---|---|---|
| `string` | ~227 µs | 62 KB | 4,000 |
| `int` | ~182 µs | 31 KB | 4,000 |
| `uint` | ~167 µs | 31 KB | 4,000 |
| `float64` | ~215 µs | 23 KB | 3,000 |

## Integration with Uber Zap

`Mask` returns `any`, so it slots directly into `zap.Any`:

```go
import (
    "go.uber.org/zap"
    "piimasker"
)

var masker = piimasker.NewMasker(piimasker.MaskerConfig{})

func processOrder(logger *zap.Logger, order Order) {
    // Pass the masked copy to the logger — the original order is unchanged
    logger.Info("processing order",
        zap.Any("order", masker.Mask(order)),
    )
}
```

For high-throughput paths, use `zap.Sugar()` or `zap.Object` with a custom marshaller if you want to avoid the extra allocation from `Mask` on every log call. However, if the log line is guarded by a level check, consider skipping the `Mask` call entirely:

```go
if logger.Core().Enabled(zap.DebugLevel) {
    logger.Debug("user payload", zap.Any("user", masker.Mask(user)))
}
```

## Unexported Fields and Named Types

- **Unexported fields** are silently skipped; their zero values appear in the copy.
- **Named string types** (e.g. `type Status string`, `type Role string`) are treated identically to plain strings — they are masked, anonymized, or shown according to the active tag or propagated mode. This means enum-like constants will have their underlying string value replaced when masked or anonymized.
