package main
    
import(
    "fmt"
    "os/exec"
)
func main() {
//   cmdStr1 := "docker run -v /home/dev:/home/dev -w /home/dev/ iggymacd/arch-nodejs ng new mytestcmd"
//   cmdStr2 := "docker run -v /home/dev:/home/dev -w /home/dev/mytestcmd iggymacd/arch-nodejs npm run postinstall"
projectName := "mytestcmd"
image := "iggymacd/arch-nodejs"
//   cmdStr3 := fmt.Sprintf("docker kill $(docker ps | awk '/arch-nodejs/{print $1}')",image)
  cmdStr4 := fmt.Sprintf("docker run -v /home/dev:/home/dev -w /home/dev/%s %s ng serve",projectName,image)

//   out, _ := exec.Command("/bin/sh", "-c", cmdStr1).Output()  
//   fmt.Printf("%s", out)
//   out, _ = exec.Command("/bin/sh", "-c", cmdStr2).Output()  
//   fmt.Printf("%s", out)
//   out, _ := exec.Command("/bin/sh", "-c", cmdStr4).Output()  
//   fmt.Printf("%s", out)
    start(cmdStr4)
}

func start(cmd string){
    out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
    fmt.Printf("%s", out)
}
 