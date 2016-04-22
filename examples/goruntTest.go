package main

import (
	"github.com/tobyroworth/gorunt"
	"github.com/tobyroworth/gorunt/npm"
	"github.com/tobyroworth/gorunt/bower"
	log "github.com/Sirupsen/logrus"
)

func init() {
    
	log.SetLevel(log.InfoLevel)
	
    npm.Install()
}

func main() {
    
    clean := []string {
        "bower_components",
        "proj1/bower_components",
        "proj2/bower_components",
    }
    
    
    links := map[string][]string {
        "proj1": {"local_components/ele1"},
        "proj2": {"local_components/ele1", "local_components/ele2"},
    }
    
    install := []string {
        "proj1",
        "proj2",
    }
    
    allBowers := map[string][]string {
        "bower_components": {"proj2/bower_components", "proj1/bower_components"},
    }
    
    _ = gorunt.Clean(clean)
    _ = bower.Link(links)
    _ = bower.Install(install)
    _ = gorunt.Copy(allBowers)
}