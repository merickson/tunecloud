/*
Package musicdirectory handles the physical resting location of the data on disk.

Right now, this means being given a directory and being able to return all the
metadata associated with that directory.
*/
package tunecloud

import (
	"errors"
	"os"
	"path/filepath"
	
	"github.com/vbatts/go-taglib/taglib"
)

type MusicDirectory struct {
	pathName string
}

// scanner is a closure to build a filepath.WalkFunc with our Track slice included.
func scanner(tslice *[]Track) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If processing a directory, just carry on.
		if info.IsDir() == true {
			return nil
		}

		t, lerr := scanFile(path)
		if lerr != nil {
			return nil
		}

		*tslice = append(*tslice, t)
		return nil
	}
}

// Scan returns a slice of tracks discovered.
func (md *MusicDirectory) Scan() ([]Track, error) {
	var tslice []Track
	
	err := filepath.Walk(md.pathName, scanner(&tslice))

	return tslice, err
}

// PathName returns the name of the directory to be scanned.
func (md *MusicDirectory) PathName() (string) {
	return md.pathName
}

// NewMusicDirectory returns an initialized instance of a MusicDirectory.
func NewMusicDirectory(pathName string) *MusicDirectory {
	absPath, _ := filepath.Abs(pathName)
	cleanPath := filepath.Clean(absPath)
	
	return &MusicDirectory{pathName: cleanPath}
}

// scanFile provides a Track for an individual song.
func scanFile(filePath string) (Track, error) {
	f := taglib.Open(filePath)

	if f == nil {
		return Track{}, errors.New("Could not open file in taglib")
	}
	defer f.Close()
	
	fTags := f.GetTags()
	fProps := f.GetProperties()

	if fTags == nil || fProps == nil {
		return Track{}, errors.New("Unable to open taglib tags and/or properties")
	}
	
	return buildTrack(filePath, fTags, fProps), nil
}

// buildTrack inflates a Track object given tags and properties from taglib.
func buildTrack(filePath string, tags *taglib.Tags, props *taglib.Properties) Track {
	t := Track{}
	t.Title = tags.Title
	t.Album = tags.Album
	t.Artist = tags.Artist
	t.PathName = filePath
	t.Length = props.Length
	t.Bitrate = props.Bitrate

	return t
}
