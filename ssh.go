package st

import (
	"fmt"
	"os/exec"
)

func SSHTunnel(serverAddr string, serverPort int, user string, remoteAddr string, remotePort int, localPort int) (err error) {
	cmd := exec.Command("ssh", "-N", "-p",
		fmt.Sprintf("%d", serverPort),
		fmt.Sprintf("%s@%s", user, serverAddr),
		"-L",
		fmt.Sprintf("%d:%s:%d", localPort, remoteAddr, remotePort),
	)

	if err = cmd.Start(); err != nil {
		return
	}

	go cmd.Wait()
	return
}