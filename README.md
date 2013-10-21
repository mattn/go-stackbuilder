# go-stackbuilder

## Introduction

go-stackbuilder provides perl Plack::Builder like interface. Currently, it support constructor of handlers.

## Usage

    package main
    
    import (
    	"github.com/justinas/nosurf"
    	"github.com/keep94/weblogs"
    	"github.com/mattn/go-stackbuilder"
    	"net/http"
    )
    
    func main() {
        // Do something to handle requests
    
    	http.ListenAndServe(":8888", stackbuilder.Build(
    		nosurf.New,
    		weblogs.Handler,
    	))
    }

## License

MIT

## Author

Yasuhiro Matsumoto
