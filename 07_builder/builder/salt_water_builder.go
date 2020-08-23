package builder

import("fmt")

type SaltWaterBuilder SaltWater

func CreateSaltWaterBuilder() *SaltWaterBuilder{
	saltWater := SaltWaterBuilder{0, 0}
	return &saltWater
}

func (swb *SaltWaterBuilder) AddSolute(saltAmount float64) {
	swb.salt = saltAmount
}
func (swb *SaltWaterBuilder) AddSolvert(waterAmount float64) {
	swb.water = waterAmount
}
func (swb *SaltWaterBuilder) AbandonSolution(saltWaterAmount float64) {
	saltDelta := saltWaterAmount * (swb.salt / (swb.salt + swb.water))
	waterDelta := saltWaterAmount * (swb.water / (swb.salt + swb.water))

	swb.salt = swb.salt - saltDelta
	swb.water = swb.water - waterDelta
}
func (swb *SaltWaterBuilder) GetResult() {
	fmt.Println(swb)
}