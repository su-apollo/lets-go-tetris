# Lets go tetris!

[![Build Status](https://travis-ci.org/zrma/lets-go-tetris.svg?branch=master)](https://travis-ci.org/zrma/lets-go-tetris)
[![Quality Gate](https://sonar.team504.com/api/badges/gate?key=lets-go-tetris)](https://sonar.team504.com/dashboard?id=lets-go-tetris)
[![Coverage Status](https://coveralls.io/repos/github/zrma/lets-go-tetris/badge.svg?branch=master)](https://coveralls.io/github/zrma/lets-go-tetris?branch=master)

#### Requirement
> [Go](https://golang.org) (v1.12+)  
> [SDL2](https://libsdl.org/) (latest)

#### Test
> $ make test

#### Setting Go-SDL2 for Windows  
* download [SDL2](https://libsdl.org/download-2.0.php)  
SDL2-devel-2.x.x-mingw.tar.gz
* install [scoop](https://scoop.sh/)
* install gcc by scoop
> $ scoop install gcc   
* copy all files from SDL2-2.x.x/x86_64-w64-mingw32/lib/ to $(scoop path)/apps/gcc/current/x86_64-w64-mingw32/lib
