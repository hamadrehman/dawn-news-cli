# 📰 Dawn News CLI

A delightful command-line interface for reading Dawn News headlines, built with Go and powered by the charming [Bubble Tea TUI framework](https://github.com/charmbracelet/bubbletea).

![Static Badge](https://img.shields.io/badge/go-1.20%2B-blue)
![Static Badge](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)

## ✨ Features

- Browse Dawn News headlines in your terminal
- Clean, modern terminal user interface
- Keyboard-driven navigation
- Fast and lightweight
- Real-time RSS feed updates

## 🚀 Installation

```bash
go install github.com/hamadrehman/dawn-news-cli@latest
```

Or build from source:

```bash
git clone https://github.com/hamadrehman/dawn-news-cli.git
cd dawn-news-cli
go build
```

## 🎮 Usage

Simply run the program:

```bash
dawn-news-cli
```

### Navigation

- `↑/↓` - Navigate through headlines
- `enter` - Select a headline
- `ctrl+c` - Quit

## 🛠 Architecture

The project is structured into three main packages:

- `feed` - Handles RSS feed fetching and parsing
- `ui` - Manages the terminal user interface using Bubble Tea
- `main` - Ties everything together

## 📦 Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal applications

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- [Dawn News](https://www.dawn.com/) for providing the RSS feed
- The amazing [Charm](https://charm.sh/) team for their beautiful terminal UI libraries

---

Built with ❤️ by [Hamad Rehman](https://github.com/hamadrehman)
