package key

type Key struct {
	ele string
	lis string
	rol string
}

func (k *Key) Elem() string {
	return k.ele
}

func (k *Key) List() string {
	return k.lis
}

func (k *Key) Role() string {
	return k.rol
}
