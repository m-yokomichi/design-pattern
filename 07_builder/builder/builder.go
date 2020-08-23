package builder

type SaltWaterInterface interface {
	AddSolute(float64)
	AddSolvert(float64)
	AbandonSolution(float64)
	GetResult()
}