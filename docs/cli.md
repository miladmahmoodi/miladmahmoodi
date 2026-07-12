# CLI Reference

## profilegen init

Scaffold a `config.yml` interactively.

```bash
profilegen init
profilegen init --force   # overwrite existing config.yml
```

Creates a starter `config.yml` by asking a few questions. Edit the file to add your projects, skills, and timeline.

---

## profilegen build

Generate `README.md` from `config.yml`.

```bash
profilegen build
profilegen build --config profile.yml
profilegen build --output path/to/README.md
profilegen build --dry-run    # print output without writing to disk
```

| Flag        | Short | Default      | Description                     |
|-------------|-------|--------------|---------------------------------|
| `--config`  | `-c`  | `config.yml` | Path to config file             |
| `--output`  | `-o`  | `README.md`  | Output path                     |
| `--dry-run` |       | false        | Print instead of writing        |

---

## profilegen preview

Serve a live preview at `localhost:4000`.

```bash
profilegen preview
profilegen preview --port 8080
```

Watches `config.yml` for changes and rebuilds automatically.

| Flag      | Short | Default      | Description               |
|-----------|-------|--------------|---------------------------|
| `--config`| `-c`  | `config.yml` | Path to config file       |
| `--port`  | `-p`  | `4000`       | Local port to serve on    |

---

## profilegen validate

Validate `config.yml` against the schema.

```bash
profilegen validate
profilegen validate --config path/to/config.yml
```

Reports all validation errors. Exits with code 1 if invalid.

---

## profilegen doctor

Check your environment for common issues.

```bash
profilegen doctor
```

Checks:
- `config.yml` exists
- `git` is in PATH
- `go` is in PATH
- Themes directory integrity

---

## profilegen theme list

List all available themes.

```bash
profilegen theme list
```

---

## profilegen theme install

Install a remote theme.

```bash
profilegen theme install minimal
profilegen theme install github.com/user/profilegen-theme-retro
```

> Remote theme installation coming in v0.2.0. Until then, place custom themes in `./themes/<name>/`.

---

## profilegen plugin list

List all available plugins.

```bash
profilegen plugin list
```

---

## profilegen plugin install

Install a community plugin.

```bash
profilegen plugin install github-activity
```

> Community plugin installation coming in v0.2.0.

---

## profilegen publish

Build and push to GitHub in one command.

```bash
profilegen publish
profilegen publish --message "update profile"
profilegen publish --dry-run
```

Runs `profilegen build`, then:

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
