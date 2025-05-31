<h1 align="center">typtea</h1>

<p align="center">Minimal terminal-based typing speed test with multiple programming language support</p>

<br>
<div align="center">
<img src="assets/example.gif" />
</div>

---

## Features

- **Terminal-based typing** with WPM and accuracy tracking
- **Multi-language support** including English and 14+ programming languages
- **Infinite word generation** with smooth 3-line scrolling display
- **Minimalist TUI** built with Bubble Tea and Lipgloss for responsive design
- **Embedded language data** for easy distribution and offline use
- **Accurate metrics** following standard typing test calculations

### Supported Languages

| | | |
|---|---|---|
| English 1k | Javascript | TypeScript |
| Python | Java | PHP |
| C | C++ | C# |
| Go |Rust | Bash |
| HTML | CSS | SQL |

---

## Installation

### Via `go install`

```bash
go install github.com/ashish0kumar/typtea@latest
```

### Build from Source

```bash
git clone https://github.com/ashish0kumar/typtea.git
cd typtea
go build
sudo mv typtea /usr/local/bin/
typtea -h
```

---

## Usage

### Basic Commands

```bash
# Start a 30-second English typing test (default)
typtea start

# Start a 60-second typing test
typtea start --duration 60

# Start a Rust keywords typing test
typtea start --lang rust

# Combine duration and language
typtea start --duration 45 --lang javascript

# List all available languages
typtea start --list-langs

# Get help
typtea --help
typtea start --help
```

### During the Test

- **The test starts** when you begin typing
- **Backspace** to correct mistakes
- **Enter** to restart after completion
- **Esc** to quit the application

---

## Development

### Prerequisites

[Go 1.19+](https://go.dev/doc/install)

### Setup

```bash
git clone https://github.com/ashish0kumar/typtea.git
cd typtea
go mod tidy
go build
./typtea start
```

### Adding New Languages

1. Create a `JSON` file in `internal/game/data/` with the format:

```json
{
  "name": "Language Name",
  "words": ["word1", "word2", "word3", ...]
}
```

2. Rebuild the application to embed the new language data

---

## Contributing

Contributions are always welcome! If you have ideas, bug reports, or want to submit code, please feel free to open an issue or a pull request.

## Dependencies

- [**Bubble Tea**](https://github.com/charmbracelet/bubbletea) - TUI framework
- [**Lipgloss**](https://github.com/charmbracelet/lipgloss) - Styling and layout
- [**Cobra**](https://github.com/spf13/cobra) - CLI framework

## License

[MIT License](LICENSE)
