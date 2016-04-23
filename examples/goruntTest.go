package main

import (
	"github.com/tobyroworth/gorunt"
	"github.com/tobyroworth/gorunt/npm"
	"github.com/tobyroworth/gorunt/bower"
	log "github.com/Sirupsen/logrus"
)

func init() {
    
	log.SetLevel(log.DebugLevel)
	
    npm.Install()
}

func main() {
    
    clean := gorunt.FileList {
        "bower_components",
        "*/bower_components",
        // "proj2/bower_components",
    }
    
    
    links := gorunt.FileMap {
        // "proj1": {"local_components/ele1"},
        "proj*": {"local_components/ele*"},
    }
    
    install := gorunt.FileList {
        "proj1",
        "proj2",
    }
    
    allBowers := gorunt.FileMap {
        "bower_components": {"*/bower_components"},
    }
    
    _ = gorunt.Clean(clean)
    _ = bower.Link(links)
    _ = bower.Install(install)
    _ = gorunt.Copy(allBowers)
}