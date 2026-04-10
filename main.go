package main

import (
	"fmt"
	"github.com/ai-mastering/phaselimiter-gui/internal/parsing"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	COLUMN_ID = iota
	COLUMN_INPUT
	COLUMN_OUTPUT
	COLUMN_STATUS
	COLUMN_MESSAGE
)

func getExecDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func getDefaultOutputDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "/tmp"
	}
	downloads := filepath.Join(home, "Downloads")
	_, err = os.Stat(downloads)
	if err == nil {
		return downloads
	}
	desktop := filepath.Join(home, "Desktop")
	_, err = os.Stat(desktop)
	if err == nil {
		return desktop
	}
	return home
}

func createTreeViewColumn(title string, order int) *gtk.TreeViewColumn {
	renderer, _ := gtk.CellRendererTextNew()
	tvc, _ := gtk.TreeViewColumnNewWithAttribute(
		title, renderer, "text", order)
	return tvc
}

func chooseOutputDirectory(parent *gtk.Window, current string) (string, bool) {
	dialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"Choose output directory",
		parent,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		"Cancel", gtk.RESPONSE_CANCEL,
		"Select", gtk.RESPONSE_ACCEPT,
	)
	if err != nil {
		return "", false
	}
	defer dialog.Destroy()

	if current != "" {
		_ = dialog.SetCurrentFolder(current)
	}

	if dialog.Run() != gtk.RESPONSE_ACCEPT {
		return "", false
	}

	path, err := dialog.GetFilename()
	if err != nil || path == "" {
		return "", false
	}
	return path, true
}

func updateListItem(model *gtk.ListStore, iter *gtk.TreeIter, m Mastering) {
	status := string(m.Status)
	if m.Status == MasteringStatusProcessing {
		status = strconv.FormatFloat(m.Progression*100, 'f', 0, 64) + "%"
	}
	model.Set(iter, []int{COLUMN_ID, COLUMN_INPUT, COLUMN_OUTPUT, COLUMN_STATUS, COLUMN_MESSAGE},
		[]interface{}{m.Id, m.Input, m.Output, status, m.Message})
}

func main() {
	masteringRunner := CreateMasteringRunner()
	go masteringRunner.Run()
	masteringId := 0

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("phaselimiter-gui")
	win.SetDefaultSize(920, 620)
	win.Connect("destroy", func() {
		masteringRunner.Terminate()
		gtk.MainQuit()
	})

	targets, err := gtk.TargetEntryNew("text/uri-list", gtk.TARGET_OTHER_APP, 1)
	if err != nil {
		log.Fatal("Unable to create target entry:", err)
	}
	win.DragDestSet(gtk.DEST_DEFAULT_ALL, []gtk.TargetEntry{*targets}, gdk.ACTION_LINK)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(12)
	box.SetMarginBottom(12)
	box.SetMarginStart(12)
	box.SetMarginEnd(12)
	win.Add(box)

	header, err := gtk.LabelNew("")
	header.SetMarkup("<b>PhaseLimiter GUI</b>\nDrag audio files into the window to start mastering.")
	header.SetXAlign(0)
	box.Add(header)

	entryLabel, err := gtk.LabelNew("Output directory")
	entryLabel.SetXAlign(0)
	box.Add(entryLabel)

	outputRow, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 8)
	box.Add(outputRow)

	entry, err := gtk.EntryNew()
	entry.SetText(getDefaultOutputDir())
	entry.SetHexpand(true)
	outputRow.Add(entry)

	browseButton, err := gtk.ButtonNewWithLabel("Browse…")
	browseButton.Connect("clicked", func() {
		current, _ := entry.GetText()
		if selected, ok := chooseOutputDirectory(win, strings.TrimSpace(current)); ok {
			entry.SetText(selected)
		}
	})
	outputRow.Add(browseButton)

	loudnessLabel, err := gtk.LabelNew("Target loudness")
	loudnessLabel.SetXAlign(0)
	box.Add(loudnessLabel)
	loudness, err := gtk.SpinButtonNewWithRange(-20, 0.0, 0.01)
	loudness.SetValue(-9)
	box.Add(loudness)

	masteringLevelLabel, err := gtk.LabelNew("Mastering intensity")
	masteringLevelLabel.SetXAlign(0)
	box.Add(masteringLevelLabel)
	masteringLevel, err := gtk.SpinButtonNewWithRange(0.0, 1.0, 0.01)
	masteringLevel.SetValue(1)
	box.Add(masteringLevel)

	bassPreservation, err := gtk.CheckButtonNewWithLabel("Preserve bass")
	box.Add(bassPreservation)

	notes, err := gtk.LabelNew(`Drop audio files.

Process
1. The input audio files are mastered
2. The output files are saved to output directory

Notes
- Uses the same algorithm as bakuage.com / aimastering.com
- No internet access`)
	notes.SetXAlign(0)
	box.Add(notes)

	ls, err := gtk.ListStoreNew(glib.TYPE_INT, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)

	tv, err := gtk.TreeViewNewWithModel(ls)
	tv.AppendColumn(createTreeViewColumn("input file", COLUMN_INPUT))
	tv.AppendColumn(createTreeViewColumn("output file", COLUMN_OUTPUT))
	tv.AppendColumn(createTreeViewColumn("status", COLUMN_STATUS))
	tv.AppendColumn(createTreeViewColumn("message", COLUMN_MESSAGE))
	tv.SetVExpand(true)

	scroll, err := gtk.ScrolledWindowNew(nil, nil)
	scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scroll.Add(tv)
	scroll.SetVExpand(true)
	box.Add(scroll)

	var destInData = func(lbi *gtk.Window,
		context *gdk.DragContext,
		x, y int,
		data_ptr *gtk.SelectionData,
		info, time uint) {

		s := string(data_ptr.GetData())
		fmt.Println(s)
		lines := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")

		outputDir, _ := entry.GetText()
		outputDir = strings.TrimSpace(outputDir)
		if outputDir == "" {
			m := Mastering{}
			m.Status = MasteringStatusFailed
			m.Id = masteringId
			masteringId += 1
			m.Message = "output directory is empty"
			iter := ls.Insert(0)
			updateListItem(ls, iter, m)
			return
		}
		if err := os.MkdirAll(outputDir, 0o755); err != nil {
			m := Mastering{}
			m.Status = MasteringStatusFailed
			m.Id = masteringId
			masteringId += 1
			m.Message = "failed to create output directory: " + err.Error()
			iter := ls.Insert(0)
			updateListItem(ls, iter, m)
			return
		}

		for _, line := range lines {
			inputPath, parseErr := parsing.ParseDroppedFilePath(line, runtime.GOOS)
			if parseErr != nil {
				if strings.TrimSpace(line) == "" {
					continue
				}
				m := Mastering{}
				m.Status = MasteringStatusFailed
				m.Id = masteringId
				masteringId += 1
				m.Input = strings.TrimSpace(line)
				m.Output = ""
				m.Message = parseErr.Error()
				iter := ls.Insert(0)
				updateListItem(ls, iter, m)
				continue
			}

			m := Mastering{}
			m.Status = MasteringStatusWaiting
			m.Id = masteringId
			masteringId += 1
			m.Ffmpeg = "ffmpeg"
			m.PhaselimiterPath = filepath.Join(getExecDir(), "phaselimiter/bin/phase_limiter")
			m.SoundQuality2Cache = filepath.Join(getExecDir(), "phaselimiter/resource/sound_quality2_cache")

			m.Input = inputPath
			m.Output = filepath.Base(m.Input)
			m.Output = strings.TrimSuffix(m.Output, filepath.Ext(m.Output))
			m.Output += "_output.wav"
			m.Output = filepath.Join(outputDir, m.Output)

			m.Loudness = loudness.GetValue()
			m.Level = masteringLevel.GetValue()
			m.BassPreservation = bassPreservation.GetActive()

			masteringRunner.Add(m)

			iter := ls.Insert(0)
			updateListItem(ls, iter, m)
		}
	}
	win.Connect("drag-data-received", destInData)

	go func() {
		for m := range masteringRunner.MasteringUpdate {
			fmt.Printf("%#v\n", m)

			glib.IdleAdd(func() {
				iter, _ := ls.GetIterFirst()
				if iter == nil {
					return
				}
				for {
					v, _ := ls.GetValue(iter, COLUMN_ID)
					id, _ := v.GoValue()
					if m.Id == id {
						updateListItem(ls, iter, m)
					}
					if ls.IterNext(iter) == false {
						break
					}
				}
			})
		}
	}()

	win.ShowAll()
	gtk.Main()
}
