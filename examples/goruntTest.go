package main

import (
	"github.com/tobyroworth/gorunt"
	// "github.com/tobyroworth/gorunt/npm"
	"github.com/tobyroworth/gorunt/bower"
	"github.com/tobyroworth/gorunt/polymer"
	log "github.com/Sirupsen/logrus"
)

func init() {
	
	log.SetLevel(log.DebugLevel)
	
	// npm.Install()
}

func main() {
	
	clean := gorunt.FileList {
		"build",
		"bower_components",
		"*/bower_components",
	}
	
	
	links := gorunt.FileMap {
		"proj*": {"local_components/ele*"},
	}
	
	install := gorunt.FileList {
		"proj1",
		"proj2",
	}
	
	allBowers := gorunt.FileMap {
		"build/bower_components": {"*/bower_components"},
	}
	
	indexCp := gorunt.FileMap {
		"build/index.html": {"*/index.html"},
	}
	
	indexV := gorunt.FileMap {
		"build/index.html": {"build/vulc.html"},
	}
	
	indexC := gorunt.FileMap {
		"build/vulc.html": {"build/crisp.html", "build/crisp.js"},
	}
	
	_ = gorunt.Clean(clean)
	_ = bower.Link(links)
	_ = bower.Install(install)
	_ = gorunt.Copy(allBowers)
	_ = gorunt.Copy(indexCp)
	_ = polymer.Vulcanize(indexV)
	_ = polymer.Crisper(indexC)
	
	_ = gorunt.Clean(clean)
}