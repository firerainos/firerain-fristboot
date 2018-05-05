package core

import "os/exec"

func SearchPackage(pkgname string) bool {
	cmd := exec.Command("pacman","-Qi",pkgname)

	return cmd.Run() == nil
}

func RemoveFristBoot() bool {
	cmd := exec.Command("pacman","-Rscn","firerain-fristboot")

	return cmd.Run() == nil
}