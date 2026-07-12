# Configuration

profilegen reads a single `config.yml` file. Every field is optional except `name`, `username`, and `role`.

## Full Reference

```yaml
# ── Identity ──────────────────────────────────────────────────────────────────
name:     "Your Name"          # required
username: "yourusername"       # required — your GitHub username
role:     "Backend Engineer"   # required — your title/role
company:  "Acme Corp"          # optional
location: "City, Country"      # optional
website:  "https://you.dev"    # optional — must be a valid URL
bio:      "One sentence bio."  # optional

# ── Theme ─────────────────────────────────────────────────────────────────────
theme: terminal   # terminal | minimal | dashboard | docs | retro | hacker
                  # default: terminal

# ── Social Links ──────────────────────────────────────────────────────────────
socials:
  - platform: github      # used for badge logo
    handle:   "username"
    url:      "https://github.com/username"
  - platform: twitter
    handle:   "@username"
    url:      "https://twitter.com/username"
  - platform: linkedin
    handle:   "username"
    url:      "https://linkedin.com/in/username"

# ── Skills ────────────────────────────────────────────────────────────────────
skills:
  - category: Languages
    items: [Go, Python, TypeScript, SQL]
  - category: Frameworks
    items: [gin, FastAPI, React]
  - category: Databases
    items: [PostgreSQL, Redis]
  - category: DevOps
    items: [Docker, Kubernetes, GitHub Actions]

# ── Projects ──────────────────────────────────────────────────────────────────
projects:
  - name:        "project-name"
    description: "What it does in one sentence"
    url:         "https://github.com/you/project"
    language:    "Go"          # primary language
    stars:       42            # optional — star count
    featured:    true          # featured projects are highlighted
    tags:                      # optional
      - tag1
      - tag2

# ── Career Timeline ───────────────────────────────────────────────────────────
timeline:
  - year:        "2024"
    title:       "What happened"
    description: "Optional — one sentence of context"

# ── Contact ───────────────────────────────────────────────────────────────────
contact:
  email:   "you@example.com"   # optional
  twitter: "@handle"           # optional
  discord: "user#0000"         # optional

# ── Plugins ───────────────────────────────────────────────────────────────────
plugins:
  - name:    github-activity  # see: profilegen plugin list
    enabled: true
    options:
      limit: 5
```

## Validation

Run `profilegen validate` to check your config before building:

```bash
profilegen validate
profilegen validate --config path/to/config.yml
```

## Defaults

| Field   | Default    |
|---------|------------|
| `theme` | `terminal` |

## Tips

- **Featured projects** appear prominently; non-featured projects collapse into a `<details>` block.
- **Timeline** renders in the order you define it — most recent first is recommended.
- **Skills** categories are printed in definition order.
- Leave `contact.email` empty if you don't want to publish your email publicly.
