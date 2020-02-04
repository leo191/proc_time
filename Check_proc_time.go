package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
			if x == n {
					return true
			}
	}
	return false
}

func checkPid(pid string){
	psCmd := exec.Command("ps","-eo", "pid")
	var stdout bytes.Buffer
	psCmd.Stdout = &stdout
	err := psCmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr := string(stdout.Bytes())
	allPid := strings.Split(outStr,"\n")
	allPid = allPid[:len(allPid)-1]
	if Contains(allPid, pid){
		fmt.Printf("%s exists\n", pid)
	}
}

func checkProc(processName string) {
	pgrepCmd := exec.Command("pgrep", processName)
	var stdout, stderr bytes.Buffer
	pgrepCmd.Stdout = &stdout
	pgrepCmd.Stderr = &stderr
	err := pgrepCmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:%s", outStr, errStr)
	pidSl := strings.Split(outStr, "\n")
	pidSl = pidSl[:len(pidSl)-1]
	for i:=0; i<len(pidSl); i++{
		time.Sleep(time.Second * 5)

		checkPid(pidSl[i])
	}

}

func main() {
	processName := os.Args[1]
	checkProc(processName)

}
