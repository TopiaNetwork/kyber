package bn256

import "github.com/TopiaNetwork/kyber/v3"

func UnmarshalBinaryPG2(pBytes []byte) (kyber.Point, error) {
	var pg2 pointG2
	err := pg2.UnmarshalBinary(pBytes)
	if err != nil {
		return nil, err
	}
	return &pg2, nil
}
