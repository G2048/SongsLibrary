## For developming

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
