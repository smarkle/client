package libkb

import (
	"io"
	"strings"

	"golang.org/x/crypto/openpgp"
)

func PGPEncrypt(source io.Reader, sink io.WriteCloser, signer *PGPKeyBundle, recipients []*PGPKeyBundle) error {
	to := make([]*openpgp.Entity, len(recipients))
	for i, r := range recipients {
		to[i] = r.Entity
	}
	w, err := openpgp.Encrypt(sink, to, signer.Entity, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	n, err := io.Copy(w, source)
	if err != nil {
		return err
	}
	G.Log.Debug("PGPEncrypt: wrote %d bytes", n)
	if err := w.Close(); err != nil {
		return err
	}
	if err := sink.Close(); err != nil {
		return err
	}
	return nil
}

func PGPEncryptString(input string, signer *PGPKeyBundle, recipients []*PGPKeyBundle) ([]byte, error) {
	source := strings.NewReader(input)
	sink := NewBufferCloser()
	if err := PGPEncrypt(source, sink, signer, recipients); err != nil {
		return nil, err
	}
	return sink.Bytes(), nil
}
