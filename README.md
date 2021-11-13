# Go2048
Yet another 2048 game implemented in golang with [go-sdl2](https://github.com/veandco/go-sdl2).

---

## Screenshots
![screenshot_1](./img/screenshot_2020-11-09_173912.png)

## Build
Install [go-sdl2](https://github.com/veandco/go-sdl2). After that simply type
```
go build main.go
```
and you'll get your game binary.

If you need the game to be built statically for redistribution, use the following line:
```
go build -tags static -ldflags "-s -w -H=windowsgui"
```
This will also get rid of the debug console. Remove "-H=windowsgui" if you need it.


## Disclaimer
All in-game assets used are legally obtained and can be redistributed freely.


---

2021, net2cn, proudly coded in boring classes.