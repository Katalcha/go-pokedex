package commands

type pokedexCommand struct {
	name        string
	description string
	Callback    func(*Config, ...string) error
}
