package gorunt

import (
	"errors"
	"strconv"
	"os/exec"
	"path/filepath"
	log "github.com/Sirupsen/logrus"
)

func Copy(targets FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Copy",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	// do not glob keys, as destinations may not exist
	targets.GlobValues()
	
	for dest, src := range targets {
		for _, src := range src {
			cmd := exec.Command("cp", "-r", "-L", "-T", src, dest)
			logger.Debugf("S: %s -> D: %s", src, dest)
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

func Clean(targets FileList) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Clean",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	targets.Glob()
	
	for _, target := range targets {
		logger.Debugf("Removing: %s", target)
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

type FileList []string

func (g *FileList)Glob() {
	
	o := make(FileList, 0, len(*g))
	
	for _, val := range *g {
		vals, err := filepath.Glob(val)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		for _, val := range vals {
			o = append(o, val)
		}
	}
	
	*g = o
}

type FileMap map[string]FileList

func (g *FileMap)Glob() {
	g.GlobKeys()
	g.GlobValues()
}

func (g *FileMap)GlobKeys() {
	o := make(FileMap)
	for key, val := range *g {
		keys, err := filepath.Glob(key)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		for _, key := range keys {
			o[key] = val
		}
	}
	
	*g = o
}

func (g *FileMap)GlobValues() {
	for key, val := range *g {
		val.Glob()
		(*g)[key] = val
	}
}

func RunParallel(cmds []exec.Cmd, logger log.FieldLogger) error{
	
	goLeft := len(cmds)
	
	errs := 0
	err := make(chan error) 
	
	for _, cmd := range cmds {
		go Run(cmd, err)
	}
	
	for e := range err {
		if e == RunDone {
			goLeft--
			if goLeft == 0 {
				break
			}
		} else {
			logger.Error(e)
			errs++
		}
	}
	
	close(err)
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " errors")
	} else {
		return nil
	}
}

func Run(cmd exec.Cmd, e chan error) {
	
	if err:= cmd.Run(); err != nil {
		e <- err
	}
	
	e <- RunDone
}

var RunDone error