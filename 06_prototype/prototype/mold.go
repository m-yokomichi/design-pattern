package prototype

type Mold struct {
	*MoldPrototype
	moldName string
}

// 型名を指定して作成する（星・三角・丸など）
func CreateMold(moldName string) Mold {
	mold := &Mold{
		MoldPrototype: &MoldPrototype{},
		moldName:      moldName,
	}
	mold.mold = mold

	return *mold
}

func (m *Mold) CreateClone() Mold {
	newMold := Mold{
		moldName: m.moldName,
	}
	return newMold
}
