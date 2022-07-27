package main

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	s := mkSMF()
	err := s.WriteFile("test.mid")
	if err != nil {
		fmt.Printf("error")
		return
	}

}

func mkSMF() *smf.SMF {
	var (
		clock = smf.MetricTicks(96)
		tr    smf.Track
	)

	tr.Add(0, smf.MetaMeter(3, 4))
	tr.Add(0, smf.MetaTempo(140))
	tr.Add(0, smf.MetaInstrument("Trumpet"))
	tr.Add(0, midi.ProgramChange(0, gm.Instr_Trumpet.Value()))
	tr.Add(0, midi.NoteOn(0, midi.Ab(3), 120))
	tr.Add(clock.Ticks8th(), midi.NoteOn(0, midi.C(4), 120))
	tr.Add(clock.Ticks4th()*2, midi.NoteOff(0, midi.Ab(3)))
	tr.Add(0, midi.NoteOff(0, midi.C(4)))
	tr.Close(0)

	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	return s
}
