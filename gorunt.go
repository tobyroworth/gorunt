package gorunt

import (
	"errors"
	"strconv"
	"os/exec"
	log "github.com/Sirupsen/logrus"
)

func Copy(targets map[string][]string) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Copy",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	for dest, src := range targets {
		for _, src := range src {
			cmd := exec.Command("cp", "-r", "-L", "-T", src, dest)
			// log.Printf("S: %s -> D: %s", src, dest)
			if err:= cmd.Run(); err != nil {
				logger.Error(err)
			}
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " erros")
	} else {
		return nil
	}
}

func Clean(targets []string) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Clean",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	for _, target := range targets {
		cmd := exec.Command("rm", "-r", target)
		if err:= cmd.Run(); err != nil {
			logger.Error(err)
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " erros")
	} else {
		return nil
	}
}
