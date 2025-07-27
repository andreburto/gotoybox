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

### jquery

My [jQuery](https://jquery.com/) toy box.
Could probably do this with nginx, but Go is a little more flexible for the API experiments I want to do.

### math

Simple API to test out [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) and JSON parsing.

### people

Importing modules from subdirectories.
Learning how to organize code.

### pointer

Did you know [Go has pointers](https://go.dev/tour/moretypes/1)?

### waiter

Simple API to test out using multuple methods with one endpoint.
Planning to use this in another project.

### whichos

Simple thing to test which OS is being used. 

## Updates

**2025-07-20:** Added the `jquery` project.
It's been too long since I played with jQuery, and I think its simplicity blends well with Go.

**2025-07-20:** Added test and GHA test for `math`.

**2025-07-19:** Added `math` service, which let's me play around with the Fetch API and JSON parsing libraries.

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