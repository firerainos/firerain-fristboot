package core

import (
	"os/exec"
	"io/ioutil"
	"os"
	"fmt"
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

func SetIM(username string) error {
	profile := `
export GTK_IM_MODULE=%s
export QT_IM_MODULE=%s
export XMODIFIERS="@im=%s"
`
	if SearchPackage("fcitx5-git") {
		profile = fmt.Sprintf(profile,"fcitx5","fcitx5","fcitx")
	} else if SearchPackage("fcitx") {
		profile = fmt.Sprintf(profile,"fcitx","fcitx","fcitx")
	} else if SearchPackage("ibus") {
		profile = fmt.Sprintf(profile,"ibus","ibus","ibus")

		profile += "\nibus-daemon -x -d"
	}

	f, err := os.OpenFile("/home/"+username+"/.xprofile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_,err=f.WriteString(profile)
	return err
}