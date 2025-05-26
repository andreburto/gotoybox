# Go ToyBox

## About

If you want real [Golang](https://go.dev/) projects, check out [collaginator2](https://github.com/andreburto/collaginator2) or [vtubers](https://github.com/andreburto/vtubers/tree/convert-to-mongodb). This repo is a snippet collection, to help me get started with other projects after I've been immersed in Python or other languages.

## Build

### counter

**Build**:
```
docker compose build counter
```

**Run**
```
docker compose run --remove-orphans counter 10
```
...or...
```
docker compose run --remove-orphans counter 1 10
```

## Updates

**2025-05-26:** Adding automated tests. 
Better error handling in the `counter` program.
Add tests for `counter`.

**2025-05-24:** Initial commit and `counter` program added.