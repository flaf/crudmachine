package exec

import(
	"os/exec"
	"log"
	"syscall"
)

func Run(i int, name string, arg... string) {

	cmd := exec.Command(name, arg...)

	log.Printf("[#%d] Run of command %v", i, cmd.Args)

	if err := cmd.Start(); err != nil {
		log.Fatalf("[#%d] cmd.Start: %v", i, err)
	}

	if err := cmd.Wait(); err != nil {
		exiterr := err.(*exec.ExitError)
		status := exiterr.Sys().(syscall.WaitStatus)
		log.Printf("[#%d] Exit code: %d", i, status.ExitStatus())
	} else {
		log.Printf("[#%d] Exit code: %d", i, 0)
	}

}


