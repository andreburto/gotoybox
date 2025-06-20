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

### whichos

TBD

## Updates

**2025-06-14:** Added `waiter` service, which is just sample server that servers a web page and has one endpoint.
Used Copilot to generate the HTML, but added a few tweaks.
This is the POC on how to extend [vid2mp3](https://github.com/andreburto/vid2mp3) to work with a bookmarklet.

**2025-06-06:** Added `whichos` program.
Need a way to reliably identify which OS the program runs on.
Not putting it in a container.

**2025-05-26:** Adding automated tests. 
Better error handling in the `counter` program.
Add tests for `counter`.

**2025-05-24:** Initial commit and `counter` program added.