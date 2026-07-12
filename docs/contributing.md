# Contributing

profilegen is open source. Contributions are welcome.

## Development Setup

```bash
git clone https://github.com/miladmahmoodi/profilegen
cd profilegen
go mod download
make build
```

Run the tests:

```bash
make test
```

Run the linter:

```bash
make lint
```

## Project Structure

```
profilegen/
├── cmd/               # CLI commands (Cobra)
├── internal/
│   ├── config/        # Config parsing and validation
│   ├── engine/        # Template execution
│   ├── theme/         # Theme loading and registry
│   ├── plugin/        # Plugin interface and registry
│   └── generator/     # Pipeline orchestrator
├── themes/            # Built-in themes
│   ├── terminal/      # Default theme
│   └── minimal/
├── examples/          # Example configs
├── docs/              # Documentation
└── tests/             # Integration tests
```

## What to Contribute

- **New themes** — add a directory to `themes/` following the theme spec
- **Plugin implementations** — implement the `Plugin` interface
- **Bug fixes** — check open issues
- **Documentation** — improve clarity, fix typos
- **Examples** — add config examples for different developer personas

## Adding a Theme

1. Create `themes/<name>/`
2. Add `theme.yml` with name, description, author, version
3. Add `templates/base.md.tmpl` with your base template
4. Add section templates if needed
5. Add `assets/` for SVGs or other static files
6. Update `internal/theme/registry.go` to list it
7. Open a PR

## Adding a Plugin

1. Create `internal/plugin/plugins/<name>/plugin.go`
2. Implement `Plugin` interface (`Name()` and `Render()`)
3. Register in `internal/plugin/registry.go`
4. Document in `docs/plugins.md`
5. Open a PR

## Pull Request Guidelines

- Keep PRs focused — one concern per PR
- Add tests for new behaviour
- Run `make test lint vet` before submitting
- Follow existing code style

## Code Style

- `gofmt -s` formatting (enforced by CI)
- Error messages lowercase, no trailing punctuation
- No unnecessary abstractions
- Keep the external dependency count minimal

## License

MIT — see [LICENSE](../LICENSE).
