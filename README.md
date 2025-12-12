
# SteamDL Client (Unofficial)

![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat&logo=go)![Platform](https://img.shields.io/badge/Platform-Windows%20|%20Linux-gray)

> **A reverse-engineered, lightweight version of SteamDL written in Go :)**

## Supported Services:
- Steam
- PlayStation
- Xbox
- Riot
- Epic

## Features

- TUI Interface
- Categorized Traffic Statistics
- Lightweight and low resource usage
- Cross Platform Support
  - *Note: Currently, this has only been tested on **Windows**.*

## Usage

If you are on **Windows**, you don't need to compile anything!

1. Go to the [Releases](https://github.com/nimazerobit/steamdl-client/releases) page.
2. Download the latest `.exe` file.
3. Open it ^_~

For **Linux** users, or if you want the latest development changes, please follow the [instructions](#compile) below.

## Compile

### Prerequisites
- **Go**: You need Go installed on your machine (version 1.25.5 is recommended)
  - [Download Go here](https://go.dev/dl/)

### 1. Clone the Repository
Open your terminal and run the following commands to download the source code:
```bash
git clone https://github.com/nimazerobit/steamdl-client.git
cd steamdl-client
```

### 2. Install Dependencies
Ensure all required Go modules are downloaded and tidy:

```bash
go mod tidy
```

### 3. Compile
Run the build command based on your operating system.

**For Windows:**
```bash
go build -o steamdl.exe .\cmd\app\main.go
```

**For Linux:**
```bash
go build -o steamdl .\cmd\app\main.go
```

### 4. Run
Once the build is complete, you can run the executable directly:

**Windows:**
```bash
.\steamdl.exe
```

**Linux:**
```bash
./steamdl
```

## Contributing

[Issues](https://github.com/nimazerobit/steamdl-client/issues) requests are welcome

## Disclaimer

This is an **unofficial** client for [SteamDL](https://steamdl.ir)