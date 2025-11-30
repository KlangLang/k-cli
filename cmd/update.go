package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func updateCommand() {
	// URL oficial do updater
	updaterURL := "https://raw.githubusercontent.com/KlangLang/loom/main/helpers/update.sh"
	tmpFile := "/tmp/loom_update.sh"
	l := NewLog()

	// Baixa o updater
	curl := exec.Command("curl", "-sL", updaterURL, "-o", tmpFile)
	if err := curl.Run(); err != nil {
		fmt.Printf("%s✖%s Failed to download update script: %v", l.PRIMARY_COLOR, l.RESET_COLOR, err)
	}

	// Torna executável
	if err := os.Chmod(tmpFile, 0755); err != nil {
		fmt.Printf("%s✖%s Failed to chmod update script: %v", l.PRIMARY_COLOR, l.RESET_COLOR, err)
	}

	// Executa
	cmd := exec.Command("bash", tmpFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s✖%s Update failed: %v", l.PRIMARY_COLOR, l.RESET_COLOR, err)
	}

	fmt.Printf("%s✔%s Loom updated successfully!\n", l.PRIMARY_COLOR, l.RESET_COLOR)
}
