package segwit

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	// rsv, err := secp256k1.Sign(crypto.Sha256(msg), privKey.Key)
	// if err != nil {
	//	 return nil, err
	// }
	// // we do not need v  in r||s||v:
	// rs := rsv[:len(rsv)-1]
	// return rs, nil
	return nil, nil
}

// VerifySignature validates the signature.
// The msg will be hashed prior to signature verification.
func (pubKey *PubKey) VerifySignature(msg, sigStr []byte) bool {
	// return secp256k1.VerifySignature(pubKey.Bytes(), crypto.Sha256(msg), sigStr)
	return false
}