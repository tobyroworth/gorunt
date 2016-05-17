package polymer

import (
	// "errors"
	// "strconv"
	// "sync"
	"os/exec"
	"path/filepath"
	"github.com/tobyroworth/gorunt"
	"github.com/tobyroworth/gorunt/npm"
	log "github.com/Sirupsen/logrus"
)

func init() {
	npm.Global["vulcanize"] = struct{}{}
	npm.Global["crisper"] = struct{}{}
}

func Vulcanize(targets gorunt.FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Vulcanize",
	})
	
	logger.Info("Started")
	
	cmds := make([]exec.Cmd, 0, len(targets))
	
	targets.GlobKeys()
	
	for src, dest := range targets {
		
		for _, dest := range dest {
		    cmd := exec.Command("vulcanize", src, "--out-html", dest)
		    logger.Debugf("S: %s -> D: %s", src, dest)
			cmds = append(cmds, *cmd)
		}
	}
	
	logger.Info("Vulcanizing")
	err := gorunt.RunParallel(cmds, logger)
	
	return err
	
}

func Crisper(targets gorunt.FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Crisper",
	})
	
	logger.Info("Started")
	
	cmds := make([]exec.Cmd, 0, len(targets))
	
	targets.GlobKeys()
	
	for src, dest := range targets {
		
		opts := crisperOpts{
			Source: src,
		}
		
		for _, dest := range dest {
		    logger.Debugf("Ext: %s", filepath.Ext(dest))
			switch filepath.Ext(dest) {
				case ".js":
					opts.Js = dest
				case ".html":
					opts.Html = dest
			}
		}
		
		cmd := exec.Command("crisper", "--source", opts.Source, "--html", opts.Html, "--js", opts.Js)
		logger.Debugf("S: %s -> H: %s, J: %s", opts.Source, opts.Html, opts.Js)
		cmds = append(cmds, *cmd)
	}
	
	logger.Info("Crisping")
	err := gorunt.RunParallel(cmds, logger)
	
	return err
	
}

type crisperOpts struct {
	Source string
	Html string
	Js string
}