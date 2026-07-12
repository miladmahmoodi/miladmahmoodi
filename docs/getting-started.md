# Getting Started

Build your GitHub profile in under a minute.

## Install

```bash
git clone https://github.com/miladmahmoodi/miladmahmoodi
cd miladmahmoodi
make install
```

Or install directly with Go:

```bash
go install github.com/miladmahmoodi/miladmahmoodi@latest
```

## Quick Start

```bash
# 1. Create a config.yml in your profile repo
profilegen init

# 2. Edit config.yml with your details
# (name, role, skills, projects, timeline)

# 3. Generate your README.md
profilegen build

# 4. Preview locally before pushing
profilegen preview
```

## Your First config.yml

After `profilegen init`, you'll have a `config.yml` like this:

```yaml
name:     "Your Name"
username: "yourusername"
role:     "Backend Engineer"
location: "City, Country"
website:  "https://yoursite.dev"
bio:      "I build things."

theme: terminal

skills:
  - category: Languages
    items: [Go, Python, TypeScript]

projects:
  - name: "my-project"
    description: "What it does"
    url: "https://github.com/you/my-project"
    language: Go
    featured: true

timeline:
  - year: "2024"
    title: "Something important happened"
```

Run `profilegen build` and your `README.md` is generated.

## Automatic Updates

Add the GitHub Action to auto-regenerate your profile on every push:

```yaml
# .github/workflows/generate-readme.yml
name: Generate README

on:
  push:
    paths: ["config.yml"]

permissions:
  contents: write

jobs:
  generate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: "1.26" }
      - run: go build -o profilegen .
      - run: ./profilegen build
      - run: |
          git config user.name  "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git diff --quiet README.md || (git add README.md && git commit -m "chore: regenerate profile [skip ci]" && git push)
```

See [Configuration](./configuration.md) for the full config reference.
