package main

import (
	"fmt"
	"github.com/robfig/cron"
	"transferfiles.com/clearFolders"
	"transferfiles.com/wrapper"
)

func main() {
	c := cron.New()
	c.AddFunc("0 05 18 * *", func() {
		wrapper.Wrapper("../chatpy/dsp/superchats", "../superchatExc/dsp/superchats", "_superchats", "json")
	})
	c.AddFunc("0 07 18 * *", func() {
		wrapper.Wrapper("../chatpy/doody/superchats", "../superchatExc/doody/superchats", "_superchats", "json")
	})

	c.AddFunc("0 00 11 * *", func() {
		wrapper.Wrapper("../chatpy/dsp/chatlogs", "../chatlogs/channels/DSP", "", "txt")
	})
	c.AddFunc("0 01 11 * *", func() {
		wrapper.Wrapper("../chatpy/ddm/chatlogs", "../chatlogs/channels/DDM", "", "txt")
	})
	c.AddFunc("0 02 11 * *", func() {
		wrapper.Wrapper("../chatpy/raw/chatlogs", "../chatlogs/channels/RAW", "", "txt")
	})
	c.AddFunc("0 03 11 * *", func() {
		wrapper.Wrapper("../chatpy/aqua/chatlogs", "../chatlogs/channels/AQUA", "", "txt")
	})
	c.AddFunc("0 04 11 * *", func() {
		wrapper.Wrapper("../chatpy/beam/chatlogs", "../chatlogs/channels/BEAM", "", "txt")
	})
	c.AddFunc("0 05 11 * *", func() {
		wrapper.Wrapper("../chatpy/doody/chatlogs", "../chatlogs/channels/DOODY", "", "txt")
	})
	c.AddFunc("0 06 11 * *", func() {
		wrapper.Wrapper("../chatpy/decepticron/chatlogs", "../chatlogs/channels/DECEPTICRON", "", "txt")
	})
	c.AddFunc("0 07 11 * *", func() {
		wrapper.Wrapper("../chatpy/meerkat/chatlogs", "../chatlogs/channels/MEERKAT", "", "txt")
	})
	c.AddFunc("0 08 11 * *", func() {
		wrapper.Wrapper("../chatpy/proper/chatlogs", "../chatlogs/channels/PROPER", "", "txt")
	})
	c.AddFunc("0 09 11 * *", func() {
		wrapper.Wrapper("../chatpy/reacts/chatlogs", "../chatlogs/channels/REACTS", "", "txt")
	})
	c.AddFunc("0 11 11 * *", func() {
		wrapper.Wrapper("../chatpy/shinko/chatlogs", "../chatlogs/channels/SHINKO", "", "txt")
	})
	c.AddFunc("0 12 11 * *", func() {
		wrapper.Wrapper("../chatpy/tbs/chatlogs", "../chatlogs/channels/TBS", "", "txt")
	})
	c.AddFunc("0 13 11 * *", func() {
		wrapper.Wrapper("../chatpy/wpig/chatlogs", "../chatlogs/channels/WPIG", "", "txt")
	})

	// Clear folders
	c.AddFunc("0 01 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/DSP")
	})
	c.AddFunc("0 02 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/DDM")
	})
	c.AddFunc("0 03 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/RAW")
	})
	c.AddFunc("0 04 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/AQUA")
	})
	c.AddFunc("0 05 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/BEAM")
	})
	c.AddFunc("0 06 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/DECEPTICRON")
	})
	c.AddFunc("0 07 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/DOODY")
	})
	c.AddFunc("0 08 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/MEERKAT")
	})
	c.AddFunc("0 09 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/PROPER")
	})
	c.AddFunc("0 10 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/REACTS")
	})
	c.AddFunc("0 11 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/SHINKO")
	})
	c.AddFunc("0 12 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/TBS")
	})
	c.AddFunc("0 13 10 * *", func() {
		clearfolders.ClearFolders("../chatlogs/channels/WPIG")
	})
	c.AddFunc("0 14 10 * *", func() {
		clearfolders.ClearFolders("../superchatExc/dsp/excel")
	})
	c.AddFunc("0 15 10 * *", func() {
		clearfolders.ClearFolders("../superchatExc/doody/excel")
	})
	c.AddFunc("0 16 10 * *", func() {
		clearfolders.ClearFolders("../superchatExc/dsp/superchats")
	})
	c.AddFunc("0 17 10 * *", func() {
		clearfolders.ClearFolders("../superchatExc/doody/superchats")
	})
	c.Start()
	fmt.Println("Running transferFiles!!!!!")
	<-make(chan struct{})
}
