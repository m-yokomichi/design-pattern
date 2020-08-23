package builder

type Director struct {
	builder SaltWaterInterface
}

func CreateDirector(builder SaltWaterInterface) *Director {
	director := &Director{
		builder : builder,
	}

	return director
}

func (d *Director) Construct() {
	d.builder.AddSolvert(100)
	d.builder.AddSolute(40)
	d.builder.AbandonSolution(70)
	d.builder.AddSolvert(100)
	d.builder.AddSolute(15)
}