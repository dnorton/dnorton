---
layout: post
title: "rabbit: a pace calculator for the terminal"
date: 2026-10-29
---

[rabbit](https://github.com/dnorton/rabbit-pace) is a small command-line running pace calculator I wrote in Go. It's named after the "rabbits": the pacers who set the tempo in a race.

I built it because I kept looking up pace tables online or asking an LLM to work out paces for workouts. This scratches that itch: the answer is one command away in the terminal I already have open.

Give it any two of time, pace, and distance and it works out the third, then prints a split table. Here's my Charlotte Marathon goal with a slight negative split:

```
$ rabbit -time 3:22:30 -dist marathon -unit mi -every 5mi -halfway -neg 2
Distance           42.2 km · 26.22 mi
Time               3:22:30
Pace               4:48 /km · 7:43 /mi
Negative split 2%  4:53 → 4:41 /km · 7:51 → 7:33 /mi

╭───────┬───────────────┬──────────┬──────────┬────────────┬────────────╮
│ Split │ Distance (mi) │ Pace /km │ Pace /mi │ Split Time │ Total Time │
├───────┼───────────────┼──────────┼──────────┼────────────┼────────────┤
│     1 │             5 │     4:53 │     7:51 │      39:17 │      39:17 │
│     2 │            10 │     4:51 │     7:48 │      38:58 │    1:18:15 │
│  Half │         13.11 │     4:48 │     7:44 │            │    1:42:17 │
│     3 │            15 │     4:48 │     7:44 │      38:40 │    1:56:54 │
│     4 │            20 │     4:46 │     7:40 │      38:21 │    2:35:15 │
│     5 │            25 │     4:44 │     7:37 │      38:03 │    3:13:18 │
│     6 │         26.22 │     4:41 │     7:33 │       9:12 │    3:22:30 │
╰───────┴───────────────┴──────────┴──────────┴────────────┴────────────╯
```

In a terminal, splits faster than target pace show in green, slower ones in yellow, and the halfway row is highlighted. Colour turns off automatically when output is piped or `NO_COLOR` is set.

### Install

```bash
# Homebrew (macOS and Linux)
brew install dnorton/tap/rabbit

# Install script (Linux and macOS) — installs to ~/.local/bin
curl -fsSL https://raw.githubusercontent.com/dnorton/rabbit-pace/main/install.sh | sh

# Anywhere with Go
go install github.com/dnorton/rabbit-pace/cmd/rabbit@latest

# Docker
docker run --rm ghcr.io/dnorton/rabbit-pace -pace 4:30 -dist 10k
```

`.deb`, `.rpm`, `.apk`, Arch packages, and Windows zips are on the [releases page](https://github.com/dnorton/rabbit-pace/releases).

### Things it's handy for

```bash
rabbit -pace 4:30 -dist half                 # finish time + splits
rabbit -time 3:00:00 -dist marathon          # pace for a goal time
rabbit -pace 5:00 -time 30:00                # how far you'll get
rabbit -pace 4:30 -dist marathon -every 5k   # custom split length
rabbit -pace 4:30 -dist half -neg 5          # negative split (-pos for positive)
rabbit -pace 7:00 -unit mi -dist 5k -mph     # treadmill speed (-kph too)
```

`-only` skips the table when you just want the numbers:

```
$ rabbit -pace 7:00 -unit mi -dist half -only -mph
Distance  21.1 km · 13.11 mi
Time      1:31:46
Pace      4:21 /km · 7:00 /mi
Speed     8.6 mph
```

`-md` prints Markdown instead, which pastes cleanly into Obsidian or a README:

```bash
rabbit -pace 7:00 -unit mi -dist 5k -halfway -md | pbcopy
```

Times and paces accept `hh:mm:ss`, `mm:ss`, or Go durations like `45m30s`. Distances accept `10k`, `13.1mi`, `800m`, `half`, `marathon`, or a bare number in `-unit` (km by default). `rabbit -h` lists every flag.

### Under the hood

- Written in Go, with tables styled by [Lip Gloss](https://github.com/charmbracelet/lipgloss).
- Releases are tag-driven: pushing a `v*` tag runs [GoReleaser](https://goreleaser.com) in GitHub Actions, which builds binaries and Linux packages, publishes a multi-arch image to `ghcr.io`, and updates the Homebrew tap.
- The Linux package is `rabbit-pace` because Debian and Ubuntu already ship an unrelated `rabbit` package. The command is still `rabbit`.

Source and full docs: [github.com/dnorton/rabbit-pace](https://github.com/dnorton/rabbit-pace) (MIT).
