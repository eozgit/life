# life
Conway's game of life

```sh
docker pull eozgit/life
```

```sh
docker run --rm --tty --interactive --publish 8080:8080 --name life eozgit/life
```

### Hotkeys

|Key|Function|
|-|-|
|h|Display help|
|1 - 9|Set game speed|
|r + 1 - 9|Reset game with selected initial population density|
|t + 1|Select *black & white* theme|
|t + 2|Select *CGA* theme\*|
|t + 3|Select *earth* theme|
|t + 4|Select *dune* theme|
|LMB|Resurrect cell|
|z + LMB|Resurrect *block*|
|x + LMB|Create *glider*|
|c + LMB|Create *light-weight spaceship*|
|v + LMB|Create *middle-weight spaceship*|
|b + LMB|Create *heavy-weight spaceship*|

\* When CGA theme is selected cells updated on last iteration will be highlighted showing the active areas subject to potential update next turn

---

#### Docker repo
https://hub.docker.com/r/eozgit/life

---

### Docker

#### Build
```sh
docker build --tag eozgit/life:latest --tag eozgit/life:YYMMDD .
```

#### Push
```sh
docker image push eozgit/life --all-tags
```

### Go

#### Run
```sh
go run .
```

#### Wasm
```sh
wasmserve ./
```
