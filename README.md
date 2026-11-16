# Remote Outbound Subscriptions

An in-process Solovey Component for remote subscriptions, inline content and authorized local files. It owns source definitions, legacy provider representation fetch/merge, endpoint identity, groups, refresh/reconciliation, source checks and the Frontend contribution.

The Go module consumes typed application ports:

- `domain.Store` and `service.Store` expose finite source-table operations. Backend owns SQLite/GORM, admission, transactions, secrets and Backup/Restore.
- `domain.FetchTransport`, Host `parseport.Parser` and service security/fetcher ports use the existing Backend SSRF, redirect/DNS, SecretRef, file and parser owners. A missing transport or parser fails closed.
- The Host Outbounds binding owner admits and applies resource changes with Host, source, provider and resource revision fences. Outbounds remains the sole resource writer; Users/Subscription Publishing renders final subscriptions.
- `service.Scheduler` starts only through enabled Component composition. It owns one cancelable minute loop and drains admitted work on Stop. Importing the module starts no worker, parser registration, network request or timer.

`ManifestJSON()` returns an isolated copy of the canonical descriptor. Component ID, default-enabled state, four tables, settings and schema version 3 remain compatible. Relocation requires no Product schema migration.

Sources without an explicit User-Agent retain legacy multi-representation fetch and merge. Explicit User-Agent fetches the original resource once. Partial representations cannot authorize pruning accepted endpoints. Failure preserves accepted snapshots and working bindings. Inline content retains its 512 KiB limit; fetched snapshots retain the 4096 endpoint limit.

The Frontend preserves `/remote-subscriptions`, its existing alias, navigation, settings and client-source slots, six locales and Classic/Nexus behavior. It uses the first-party Component UI SDK and existing authenticated API client. `HOST_UI_PIN.json` pins the SDK/toolchain used for independent validation.

## Integration and verification

Require the accepted immutable Go module and consume its manifest/Frontend through the existing Component source-input lock. Implement the typed Backend ports; do not pass a global database or generic execution authority.

Use the Go version in `go.mod` and `GOWORK=off`:

```text
go test -race -count=1 ./...
go vet ./...
go mod tidy -diff
go test -tags minimal ./...
node scripts/check-source.mjs
node scripts/check-frontend.mjs /path/to/accepted/solovey-ui/frontend
```

The Frontend checker requires dependencies installed from the accepted SDK package lock. It checks types, contribution/locale fixtures and a library bundle with dynamic views and CSS. Linux and Windows CI run the same source checks.

The module grants no autonomous Node refresh authority. A headless integration may reuse the same source with explicitly supplied capabilities; production Node enrollment, transport and source storage remain outside this Component.

## Compatibility scope

This extraction preserves the existing source identity and representation model. The future explicit representation/EndpointIR migration remains separate R10 work. This Component does not claim that migration is complete.

The source retains GPL-3.0-only and filtered Git ancestry. See `NOTICE` and `SOURCE_PROVENANCE.json` for attribution and exact source mapping.
