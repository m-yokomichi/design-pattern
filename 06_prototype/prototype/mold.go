package prototype

type Mold struct {
	mold     *MoldInterface
	moldName string
}

// 型名を指定して作成する（星・三角・丸など）
func CreateMold(moldName string) MoldInterface {
	var mold = Mold{
		mold: &Mold{},
	}
	mold.moldName = moldName

	return mold
}

func (m *Mold) CreateClone() Mold {
	return *m
}
