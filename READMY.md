## Example the `.env` file:

```zsh
appname="SongsLibrary"
appversion="1.0.0"
loglevel="debug"
port="3333"
host="0.0.0.0"
dsn="postgres://librarian:librarian@0.0.0.0:5432/library?sslmode=disable"
```


## For developers

```zsh
go install github.com/air-verse/air@latest
```

```zsh
air init
```

**Change in `.air.toml`:**

```zsh
  cmd = "go build -o tmp\\main.exe src\\cmd\\main.go"
```

And run server:

```zsh
air -d
```
