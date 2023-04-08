package main

import "testing"

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func TestRun(t *testing.T) {
	t.Run("Reflection tests", func(t *testing.T) {

		cases := []struct {
			name      string
			enter     interface{}
			wantCalls []string
		}{
			{"String Test",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct com dois campos tipo string",
				struct {
					Nome   string
					Cidade string
				}{"Chris", "Londres"},
				[]string{"Chris", "Londres"},
			},
			{
				"Struct sem campo tipo string",
				struct {
					Nome  string
					Idade int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				"Campos aninhados",
				Pessoa{
					"Chris",
					Perfil{33, "Londres"},
				},
				[]string{"Chris", "Londres"},
			},
			{
				"Ponteiros para coisas",
				&Pessoa{
					"Chris",
					Perfil{33, "Londres"},
				},
				[]string{"Chris", "Londres"},
			},
			{
				"Slices",
				[]Perfil{
					{33, "Londres"},
					{34, "Reykjavík"},
				},
				[]string{"Londres", "Reykjavík"},
			},
			{
				"Arrays",
				[2]Perfil{
					{33, "Londres"},
					{34, "Reykjavík"},
				},
				[]string{"Londres", "Reykjavík"},
			},
		}

		for _, test := range cases {
			t.Run(test.name, func(t *testing.T) {
				var got []string
				run(test.enter, func(field string) {
					got = append(got, field)
				})
				if len(got) != len(test.wantCalls) {
					t.Fatalf("got %d calls, want %d", len(got), len(test.wantCalls))
				}
				for i, want := range test.wantCalls {
					if got[i] != want {
						t.Errorf("got %q, want %q", got[i], want)
					}
				}
			})
		}
	})

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		run(aMap, func(field string) {
			got = append(got, field)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
