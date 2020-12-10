package metals

// Metal exported
type Metal struct {
	Mass   float64
	Volume float64
}

// Density Exported
func (m *Metal) Density() float64 {
	return m.Mass / m.Volume
}

// IsDenser exported
func IsDenser(a, b *Metal) bool {
	return a.Density() > b.Density()
}
