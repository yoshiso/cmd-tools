package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	lastExecCmdGetter := "cat ~/.zsh_history | tail -n 2 | head -n 1"
	o, err := exec.Command("sh", "-c", lastExecCmdGetter).Output()
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(o), ";")

	cmd := strings.TrimSpace(strs[len(strs)-1])

	fmt.Println(fmt.Sprintf("Alias for `%v` :", cmd))

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	h := os.Getenv("HOME") + "/.aliases.rc"

	f, err := os.OpenFile(h, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0750)
	if err != nil {
		panic(err)
	}

	if strings.Contains(cmd, "'") {
		f.WriteString(fmt.Sprintf("alias %s=\"%v\"\n", strings.TrimSpace(name), cmd))
		return
	}

	f.WriteString(fmt.Sprintf("alias %s='%v'\n", strings.TrimSpace(name), cmd))

	exec.Command("source", h)
}
