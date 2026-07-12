# CLI Reference

## forge init

Scaffold a `config.yml` interactively.

```bash
forge init
forge init --force   # overwrite existing config.yml
```

Creates a starter `config.yml` by asking a few questions. Edit the file to add your projects, skills, and timeline.

---

## forge build

Generate `README.md` from `config.yml`.

```bash
forge build
forge build --config profile.yml
forge build --output path/to/README.md
forge build --dry-run    # print output without writing to disk
```

| Flag        | Short | Default      | Description                     |
|-------------|-------|--------------|---------------------------------|
| `--config`  | `-c`  | `config.yml` | Path to config file             |
| `--output`  | `-o`  | `README.md`  | Output path                     |
| `--dry-run` |       | false        | Print instead of writing        |

---

## forge preview

Serve a live preview at `localhost:4000`.

```bash
forge preview
forge preview --port 8080
```

Watches `config.yml` for changes and rebuilds automatically.

| Flag      | Short | Default      | Description               |
|-----------|-------|--------------|---------------------------|
| `--config`| `-c`  | `config.yml` | Path to config file       |
| `--port`  | `-p`  | `4000`       | Local port to serve on    |

---

## forge validate

Validate `config.yml` against the schema.

```bash
forge validate
forge validate --config path/to/config.yml
```

Reports all validation errors. Exits with code 1 if invalid.

---

## forge doctor

Check your environment for common issues.

```bash
forge doctor
```

Checks:
- `config.yml` exists
- `git` is in PATH
- `go` is in PATH
- Themes directory integrity

---

## forge theme list

List all available themes.

```bash
forge theme list
```

---

## forge theme install

Install a remote theme.

```bash
forge theme install minimal
forge theme install github.com/user/forge-theme-retro
```

> Remote theme installation coming in v0.2.0. Until then, place custom themes in `./themes/<name>/`.

---

## forge plugin list

List all available plugins.

```bash
forge plugin list
```

---

## forge plugin install

Install a community plugin.

```bash
forge plugin install github-activity
```

> Community plugin installation coming in v0.2.0.

---

## forge publish

Build and push to GitHub in one command.

```bash
forge publish
forge publish --message "update profile"
forge publish --dry-run
```

Runs `forge build`, then:

```bash
git add README.md
git commit -m "chore: regenerate profile [skip ci]"
git push
```

| Flag        | Short | Default       | Description              |
|-------------|-------|---------------|--------------------------|
| `--config`  | `-c`  | `config.yml`  | Path to config file      |
| `--message` | `-m`  | auto          | Git commit message       |
| `--dry-run` |       | false         | Build but don't push     |
