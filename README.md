# GoWeda

GoWeda is a fast, lightweight command-line weather app. Type a city, get the current weather right in your terminal.

## Features

- **Check any city:** Look up the weather anywhere in the world with the `-city` flag.
- **Multi-word names work:** Cities like "New York" are handled for you.
- **Just the essentials:** See the current temperature, what it feels like, humidity, and conditions, with no clutter.
- **No account or API key needed:** Install it and go.

## Requirements

- [Go](https://go.dev/dl/) installed on your computer. To check, run `go version`.
- An internet connection.

## Installation

1. Clone the repository:

```bash
   git clone https://github.com/ehigai/goweda.git
```

2. Move into the project folder:

```bash
   cd goweda
```

3. Choose how to run it:

   **Run it directly:**

```bash
   go run main.go
```

**Or build an executable** for faster, repeated use:

```bash
   go build -o goweda main.go
   ./goweda
```

## Usage

Run GoWeda with no options and it shows the weather for **Lagos**:

```bash
go run main.go
```

Example output:

```plaintext
Fetching weather details for: Lagos

Weather in Lagos:
Conditions:  Partly cloudy
Temperature: 28°C (Feels like 31°C)
Humidity:    75%
```

### Check a different city

Use the `-city` flag. Put the name in quotes if it has more than one word:

```bash
go run main.go -city="London"
go run main.go -city="New York"
```

If you built the executable, use it the same way:

```bash
./goweda -city="London"
```
