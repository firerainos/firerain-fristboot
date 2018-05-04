package core

import "os/exec"

func UserAdd(username,password string) error {
	cmd := exec.Command("useradd","-m","-g","user","-G","wheel",username)
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