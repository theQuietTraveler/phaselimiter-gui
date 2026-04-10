package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ai-mastering/phaselimiter-gui/internal/parsing"
	"os/exec"
	"strconv"
	"sync"
)

type MasteringStatus string

const (
	MasteringStatusWaiting    = MasteringStatus("waiting")
	MasteringStatusProcessing = MasteringStatus("processing")
	MasteringStatusFailed     = MasteringStatus("failed")
	MasteringStatusSucceeded  = MasteringStatus("succeeded")
)

type Mastering struct {
	Id                 int
	Input              string
	Output             string
	Ffmpeg             string
	PhaselimiterPath   string
	SoundQuality2Cache string
	Loudness           float64
	Level              float64
	BassPreservation   bool
	Progression        float64
	Status             MasteringStatus
	Message            string
}

type MasteringRunner struct {
	MasteringUpdate chan Mastering
	mastering       chan Mastering
	ctx             context.Context
	cancel          context.CancelFunc
	terminateOnce   sync.Once
}

func (m Mastering) execute(ctx context.Context, update chan Mastering) {
	formatFloat := func(x float64) string {
		return strconv.FormatFloat(x, 'f', 7, 64)
	}
	formatBool := func(x bool) string {
		if x {
			return "true"
		}
		return "false"
	}

	args := []string{
		"--input", m.Input,
		"--output", m.Output,
		"--ffmpeg", m.Ffmpeg,
		"--mastering", "true",
		"--mastering_mode", "mastering5",
		"--sound_quality2_cache", m.SoundQuality2Cache,
		"--mastering_matching_level", formatFloat(m.Level),
		"--mastering_ms_matching_level", formatFloat(m.Level),
		"--mastering5_mastering_level", formatFloat(m.Level),
		"--erb_eval_func_weighting", formatBool(m.BassPreservation),
		"--reference", formatFloat(m.Loudness),
	}
	cmd := exec.CommandContext(ctx, m.PhaselimiterPath, args...)
	CmdHideWindow(cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		m.Status = MasteringStatusFailed
		m.Message = "failed to create stdout pipe: " + err.Error()
		update <- m
		return
	}
	cmd.Stderr = cmd.Stdout

	m.Status = MasteringStatusProcessing
	update <- m

	err = cmd.Start()
	if err != nil {
		m.Status = MasteringStatusFailed
		m.Message = "failed to start command: " + err.Error()
		update <- m
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if progression, ok := parsing.ExtractProgression(line); ok {
			m.Progression = progression
			update <- m
		}
	}

	if err := scanner.Err(); err != nil {
		m.Status = MasteringStatusFailed
		m.Message = "failed to read process output: " + err.Error()
		update <- m
		return
	}

	err = cmd.Wait()
	if err != nil {
		if ctx.Err() != nil {
			m.Status = MasteringStatusFailed
			m.Message = "processing cancelled"
		} else {
			m.Status = MasteringStatusFailed
			m.Message = "command failed: " + err.Error()
		}
		update <- m
		return
	}

	m.Progression = 1
	m.Status = MasteringStatusSucceeded
	update <- m
}

func CreateMasteringRunner() *MasteringRunner {
	ctx, cancel := context.WithCancel(context.Background())
	m := &MasteringRunner{}
	m.mastering = make(chan Mastering, 1000)
	m.MasteringUpdate = make(chan Mastering, 1000)
	m.ctx = ctx
	m.cancel = cancel
	return m
}

func (m *MasteringRunner) Run() {
	defer close(m.MasteringUpdate)
	for {
		select {
		case <-m.ctx.Done():
			return
		case x := <-m.mastering:
			x.execute(m.ctx, m.MasteringUpdate)
		}
	}
}

func (m *MasteringRunner) Add(mastering Mastering) {
	select {
	case <-m.ctx.Done():
		return
	case m.mastering <- mastering:
	}
}

func (m *MasteringRunner) Terminate() {
	m.terminateOnce.Do(func() {
		m.cancel()
	})
}
