## Run Locally

###### This project needs CGO to run because of mattn/go-sqlite3 package, meaning a c++ compiler. Was used mingw-64, but you can install one of your preference

###### this project uses [air](github.com/air-verse/air 'air'), install it if you want live reload

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
