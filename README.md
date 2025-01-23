## Run Locally

###### This project needs CGO to run because of mattn/go-sqlite3 package, meaning a c++ compiler. Was used mingw-64 (x86_64-release-posix-seh-ucrt-rt), but you can install one of your preference. In windows extract the mingw64 folder to `C:\mingw64` and add `C:\mingw64\bin` to path

###### this project uses [air](https://github.com/air-verse/air), install it if you want live reload

Clone the project

```bash
  git clone <project-link>
```

Go to the project directory

```bash
  cd <project-name>
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  air
```

### Deployment in railway (only tested this one until now)

Need to setup custom command
We need to instal gcc because of sqlite library being used

```bash
  apt-get update && apt-get install -y gcc && go build ./cmd/main.go
```

Custom start command

```bash
  ./main
```

Environment Variables
`CC = gcc`, `CGO_ENABLED = 1`
