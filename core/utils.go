package core

import (
	"os/exec"
	"io/ioutil"
)

func UserAdd(username,password string) error {
	pkgs := []string{"bumblebee","wireshark","tomcat7","tomcat8","sambashare"}

	groups := "wheel"
	for _,pkg := range pkgs {
		if SearchPackage(pkg) {
			groups += ","+pkg
		}
	}

	cmd := exec.Command("useradd","-m","-g","users","-G",groups,username)
	if err:= cmd.Run();err!=nil{
		return err
	}

	cmd = exec.Command("bash","-c","echo \""+username+":"+password+"\" | chpasswd")
	if err:= cmd.Run();err!=nil{
		return err
	}

	cmd = exec.Command("bash","-c","echo \" root:"+password+"\" | chpasswd")
	if err:= cmd.Run();err!=nil{
		return err
	}

	return nil
}

func SetHomeName(hostname string) error {
	return ioutil.WriteFile("/etc/hostname",[]byte(hostname),0644)
}

func SetLocale(username string) error {
	profile := `#!/bin/bash
export LANG=zh_CN.UTF-8
export LANGUAGE=zh_CN:en_US
export LC_CTYPE=zh_CN.UTF-8
`

	return ioutil.WriteFile("/home/"+username+"/.xprofile",[]byte(profile),0644)
}