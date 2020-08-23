package builder

type SaltWater struct {
	salt float64
	water float64
}

func CreateSaltWater(water , salt float64) SaltWater {
	return SaltWater{salt: salt, water: water}
}