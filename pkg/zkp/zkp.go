package keystore

type Prover interface {
	GenerateProof() error
}

type ZKP struct {
}

func (z *ZKP) GenerateProof() error {
	return nil
}
