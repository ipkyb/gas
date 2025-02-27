package api

type I_Dependencies struct {
	TempsInterface
}

func I_Register(deps I_Dependencies) {
	Temps.TempsInterface = deps.TempsInterface
}
