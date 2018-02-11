package cards

const (
	MaxStrength int64 = 100
	MaxHealth   int64 = 500
)

type Card struct {
	Name     string
	Strength int64
	Health   int64
}
