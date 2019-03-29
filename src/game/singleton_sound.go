package game

import (
	"sync"

	"../gologger"
	"github.com/veandco/go-sdl2/mix"
)

// Enum
const (
	SoundMusic = iota
	SoundSFX
)

// Singleton
var (
	sm     *SoundManager
	smOnce sync.Once
)

// SSoundManager .. singleton
var SSoundManager = newSoundManager()

func newSoundManager() *SoundManager {
	gologger.SLogger.Println("Init New Sound Manager")

	smOnce.Do(func() {
		sm = &SoundManager{}

		sm.sfxs = make(map[string]*mix.Chunk)

		sm.music = make(map[string]*mix.Music)
	})
	return sm
}

// SoundManager ...
type SoundManager struct {
	sfxs  map[string]*mix.Chunk
	music map[string]*mix.Music
}

// Load ...
// filename - name of the file to load
// id - id for map to find
// soundType - int to enum
func (sm *SoundManager) Load(filename, id string, soundType int) bool {
	if soundType == SoundMusic {

		music, err := mix.LoadMUS(filename)
		checkError(err)
		sm.music[id] = music
		return true
	} else if soundType == SoundSFX {

		chunk, err := mix.LoadWAV(filename)
		checkError(err)
		sm.sfxs[id] = chunk
		return true
	}

	return false
}

// PlaySound ...
func (sm *SoundManager) PlaySound(id string, loop int) {
	_, err := sm.sfxs[id].Play(-1, loop)
	checkError(err)
}

// PlayMusic ...
func (sm *SoundManager) PlayMusic(id string, loop int) {
	err := sm.music[id].Play(loop)
	checkError(err)
}

// Close ...
func (sm *SoundManager) Close() {
	mix.CloseAudio()
}
