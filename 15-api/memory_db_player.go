package main

func NewMemoryPlayerStorage() *MemoryPlayerStorage {
	return &MemoryPlayerStorage{map[string]int{}}
}

type MemoryPlayerStorage struct {
	storage map[string]int
}

func (m *MemoryPlayerStorage) VictoryRegister(name string) {
	m.storage[name]++
}

func (m *MemoryPlayerStorage) GetPlayerScore(name string) int {
	return m.storage[name]
}

func (m *MemoryPlayerStorage) GetLeagueTable() []Player {
	var league []Player
	for name, wins := range m.storage {
		league = append(league, Player{name, wins})
	}
	return league
}
