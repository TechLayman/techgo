package gases

// Gas exported
type Gas struct {
	Pressure      float64
	Temerature    float64
	MolecularMass float64
}

// Density Exported
func (g *Gas) Density() float64 {
	var density float64
	density = (g.MolecularMass * g.Pressure) / (0.0821 * (g.Temerature + 273))
	return density
}

// IsDenser exported
func IsDenser(a, b *Gas) bool {
	return a.Density() > b.Density()
}
