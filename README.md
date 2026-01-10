# opfetch

![Latest Release](https://img.shields.io/github/v/release/parthivsaikia/opfetch?style=flat-square&color=orange)
[![Go Report Card](https://goreportcard.com/badge/github.com/parthivsaikia/opfetch)](https://goreportcard.com/report/github.com/parthivsaikia/opfetch)
![License](https://img.shields.io/github/license/parthivsaikia/opfetch?style=flat-square)

**A "One Piece" inspired system fetch tool tailored for Linux.**

Turn your boring terminal specs into a Wanted Poster. `opfetch` is written in Go, focusing on speed, aesthetics, and that pirate life.

<p align="center">
  <img src="assets/screenshot.png" alt="opfetch demo" width="400">
</p>

## Features

- **The Aesthetic:** Renders your system stats inside a classic Wanted Poster.
- **Fast:** Written in Go, it runs instantly.
- **Zero Config:** Automatically detects your Distro, Kernel, Uptime, and Shell.
- **Dynamic Bounty:** Generates a unique "Bounty" based on your system stats.
- **Paper Style:** Uses a parchment color scheme that looks great on any terminal background.

## Installation

### Option 1: Go Install (Recommended)

If you have Go installed, this is the fastest way:

```bash
go install github.com/parthivsaikia/opfetch@latest
```

### Option 2: Binary Download

Download the pre-compiled binary for your architecture from the [Releases Page](https://github.com/parthivsaikia/opfetch/releases).

1. Download the `tar.gz` file for your architecture (e.g., `Linux_x86_64`).
2. Extract it:

```bash
tar -xvf opfetch_*.tar.gz
```

3. Move the binary to your path:

```bash
sudo mv opfetch /usr/local/bin/
```

## Usage

Simply run the command in your terminal:

```bash
opfetch
```

**Note:** Currently, `opfetch` supports Linux only (as it reads directly from `/proc`).

## Building from Source

```bash
git clone https://github.com/parthivsaikia/opfetch.git
cd opfetch
go build
./opfetch
```

## Contributing

Pull requests are welcome! If you want to add support for MacOS or Windows, feel free to open a PR.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
