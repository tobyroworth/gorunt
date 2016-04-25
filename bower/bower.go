package bower

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
	npm.Global["bower"] = struct{}{}
}

func Install(targets gorunt.FileList) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Bower Install",
	})
	
	logger.Info("Started")
	
	targets.Glob()
	
	cmds := make([]exec.Cmd, 0, len(targets))
	
	for _, target := range targets {
	
		cmd := *exec.Command("bower", "install")
		cmd.Dir = target
		cmds = append(cmds, cmd)
	}
	
	err := gorunt.RunParallel(cmds, logger)
	
	return err
}

func Link(targets gorunt.FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Bower Link",
	})
	
	logger.Info("Started")
	
	froms := make(map[string]exec.Cmd)
	tos := make(map[link]exec.Cmd)
	
	targets.Glob()
	
	for to, from := range targets {
		
		for _, from := range from {
			cmdF := exec.Command("bower", "link")
			cmdF.Dir = from
			froms[from] = *cmdF
			
			_, file := filepath.Split(from)
			
			l := link{
				From: file,
				To: to,
			}
			cmdT := exec.Command("bower", "link", l.From)
			cmdT.Dir = l.To
			tos[l] = *cmdT
		}
	}
	
	fromS := make([]exec.Cmd, 0, len(froms))
	toS := make([]exec.Cmd, 0, len(tos))
	
	for _, cmdF := range froms {
		fromS = append(fromS, cmdF)
	}
	
	for _, cmdT := range tos {
		toS = append(toS, cmdT)
	}
	
	logger.Info("Creating Links")
	err := gorunt.RunParallel(fromS, logger)
	
	logger.Info("Linking into projects")
	err = gorunt.RunParallel(toS, logger)
	
	return err
	
}

type link struct {
	From string
	To string
}