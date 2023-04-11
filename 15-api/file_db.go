package api

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileDB struct {
	database *json.Encoder
	league   League
}

func NewFileDB(file *os.File) (*FileDB, error) {
	err := startDatabaseFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem starting player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player db file, %v", err)
	}

	return &FileDB{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func (f *FileDB) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileDB) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileDB) RecordWin(name string) {
	player := f.GetLeague().Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	err := f.database.Encode(f.league)
	if err != nil {
		return
	}
}

func FileSystemPlayerStoreFromFile(path string) (*FileDB, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		err := db.Close()
		if err != nil {
			return
		}
	}

	storage, err := NewFileDB(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return storage, closeFunc, nil
}

func startDatabaseFile(file *os.File) error {
	_, err := file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("could not seek to start of file %s, %v", file.Name(), err)
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not get file info for %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, err := file.Write([]byte("[]"))
		if err != nil {
			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
