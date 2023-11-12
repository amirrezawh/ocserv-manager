package generator

import (
	"fmt"
	"os"
	"os/exec"
)

func jsonGenerator() {
	cmd := exec.Command("bash", "-c", "occtl -j show users > $JSON_PATH")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to create general.json ", err)
	}
}

func lockUser(passwdPath string, username string) {
	lockCommand := fmt.Sprintf("ocpasswd -c %s -u %s", passwdPath, username)
	cmd := exec.Command("bash", "-c", lockCommand)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to lock user ", err)
	}
}

func unlockUser(passwdPath string, username string) {
	unlockCommand := fmt.Sprintf("ocpasswd -c %s -u %s", passwdPath, username)
	cmd := exec.Command("bash", "-c", unlockCommand)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to unlock user ", err)
	}
}
