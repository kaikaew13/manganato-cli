package main

import "testing"

func TestValidateCommand(t *testing.T) {
	type test struct {
		description string
		command     string
		valid       bool
		cmd         string
		args        string
	}

	tests := []test{
		{
			description: "command for searching for mangas by name",
			command:     "search chainsaw man",
			valid:       true,
			cmd:         searchCommand,
			args:        "chainsaw man",
		},
		{
			description: "command for searching for mangas by author name",
			command:     "search-author tatsuki fujimoto",
			valid:       true,
			cmd:         searchByAuthorCommand,
			args:        "tatsuki fujimoto",
		},
		{
			description: "command for searching for mangas by genre",
			command:     "search-genre action",
			valid:       true,
			cmd:         searchByGenreCommand,
			args:        "action",
		},
		{
			description: "invalid command",
			command:     "seach chainsaw man",
			valid:       false,
			cmd:         "",
			args:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			valid, cmd, args := validateCommand(test.command)

			if valid != test.valid {
				t.Errorf("wanted valid to be %t, got %t", test.valid, valid)
			}

			if cmd != test.cmd {
				t.Errorf("wanted cmd to be %s, got %s", test.cmd, cmd)
			}

			if args != test.args {
				t.Errorf("wanted args to be %s, got %s", test.args, args)
			}
		})
	}
}
