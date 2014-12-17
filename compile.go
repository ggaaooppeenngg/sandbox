package sandbox

import (
	"os/exec"
)

const (
	C uint64 = iota
	CPP
	GO
)

//default comiple options
func compile(src string, des string, lan uint64) error {
	var cmd = new(exec.Cmd)
	switch lan {
	case C:
		//-lm for gcc math link option
		cmd = exec.Command("gcc", "-o", des, src, "-lm")
	case CPP:
		cmd = exec.Command("g++", "-o", des, src)
	case GO:
		cmd = exec.Command("go", "build", "-o", des, src)
	}
	if err := cmd.Run(); err != nil {
		return err
	} else {
		return nil
	}
}
