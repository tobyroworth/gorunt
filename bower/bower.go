package bower

import (
	"errors"
	"strconv"
	"os/exec"
	"path/filepath"
	"github.com/tobyroworth/gorunt"
	"github.com/tobyroworth/gorunt/npm"
	log "github.com/Sirupsen/logrus"
)

func init() {
    npm.Global["bower"] = struct{}{}
}

func Install(targets gorunt.FileList) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Bower Install",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	targets.Glob()
	
	for _, target := range targets {
		cmd := exec.Command("bower", "install")
		cmd.Dir = target
		
		if err:= cmd.Run(); err != nil {
			logger.Error(err)
			errs++
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " errors")
	} else {
		return nil
	}
}

func Link(targets gorunt.FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Bower Link",
	})
	
	logger.Info("Started")
	
	froms := make(map[string]*exec.Cmd)
	tos := make(map[link]*exec.Cmd)
	errs := 0
	
	targets.Glob()
	
	for to, from := range targets {
		
		
		for _, from := range from {
    		cmdF := exec.Command("bower", "link")
    		cmdF.Dir = from
    		froms[from] = cmdF
		    
		    _, file := filepath.Split(from)
		    
		    l := link{
		        From: file,
		        To: to,
		    }
    		cmdT := exec.Command("bower", "link", l.From)
    		cmdT.Dir = l.To
    		tos[l] = cmdT
		}
	}
	
	for _, cmdF := range froms {
	    logger.Debugf("Setting up link: %s", cmdF.Dir)
		err:= cmdF.Run()
		if (err != nil) {
			logger.Error(err)
			errs++
		}
	}
	
	for _, cmdT := range tos {
	    logger.Debugf("Linking: %s in %s", cmdT.Args[2], cmdT.Dir)
		err:= cmdT.Run()
		if (err != nil) {
			logger.Error(err)
			errs++
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " errors")
	} else {
		return nil
	}
	
}

type link struct {
    From string
    To string
}