package gomakemkv

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/DrWalrus1/gomakemkv/events"
)

type MakeMkvCommandHandler struct {
	ExecutablePath string
}

func (m MakeMkvCommandHandler) TriggerDiskInfo(source string, ctx context.Context) (<-chan MkvDiscInfo, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(ctx)
	cmd := exec.CommandContext(ctx, m.ExecutablePath, "-r", "--progress=-stdout", "info", source)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error creating stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error starting command: %w", err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.Canceled {
				return
			}
			cancel()
			log.Printf("error waiting for command: %s", err.Error())
		}
	}()
	_, discInfo := ParseMakeMkvInfoCommandLogs(outputPipe)
	return discInfo, cancel, nil
}

func (m MakeMkvCommandHandler) TriggerInitialInfoLoad(ctx context.Context) (<-chan MkvDiscInfo, context.CancelFunc, error) {
	return m.TriggerDiskInfo("disc:9999", ctx)
}

func (m MakeMkvCommandHandler) TriggerSaveMkv(source string, title string, destination string) (<-chan events.MakeMkvOutput, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, m.ExecutablePath, "-r", "--progress=-stdout", "--debug=-stdout", "mkv", source, title, destination)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error executing makemkvcon: %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.Canceled {
				return
			}
			log.Printf("error waiting for command: %s", err.Error())
		}
	}()
	outputs := ParseMakeMkvLogs(outputPipe)
	return outputs, cancel, nil
}

func (m MakeMkvCommandHandler) TriggerDiskBackup(decrypt bool, source string, destination string) (<-chan events.MakeMkvOutput, context.CancelFunc, error) {
	var cmd *exec.Cmd
	ctx, cancel := context.WithCancel(context.Background())
	if decrypt {
		cmd = exec.CommandContext(ctx, m.ExecutablePath, "-r", "backup", "--decrypt", source, destination)
	} else {
		cmd = exec.CommandContext(ctx, m.ExecutablePath, "-r", "backup", source, destination)
	}
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("error creating stdout pipe: %w", err)
		return nil, cancel, err
	}
	if err := cmd.Start(); err != nil {
		err = fmt.Errorf("error starting command: %w", err)
		return nil, cancel, err
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.Canceled {
				return
			}
			log.Printf("error waiting for command: %s", err.Error())
		}
	}()
	outputs := ParseMakeMkvLogs(outputPipe)
	return outputs, cancel, nil
}

func (m MakeMkvCommandHandler) RegisterMakeMkv(key string) error {
	arguments := "-r"
	command := "reg"

	cmd := exec.Command(m.ExecutablePath, arguments, command, key)
	outputPipe, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("error creating pipe to command. %s", err.Error())
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("error executing command. %s", err.Error())
	}
	return HandleRegisterMakeMkvEvents(outputPipe)

}
