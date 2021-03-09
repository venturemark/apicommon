package key

type Key struct {
	id  *ID
	ele string
	lis string
}

func (k *Key) ID() *ID {
	return k.id
}

func (k *Key) Elem() string {
	return k.ele
}

func (k *Key) List() string {
	return k.lis
}
