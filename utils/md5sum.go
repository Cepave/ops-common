package utils

import (
	"log"
	"os/exec"
)

func Md5sumCheck(workdir, md5file string) bool {
	cmd := exec.Command("md5sum", "-c", md5file)
	cmd.Dir = workdir
	err := cmd.Run()
	if err != nil {
		log.Printf("cd %s; md5sum -c %s fail", workdir, md5file)
		return false
	}
	return true
}
