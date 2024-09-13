package p2p

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/quic-go/quic-go"
)

var ValidCerts  = map[string]string{}

func  SendSecureQuicRequest(config *configs.MainConfiguration, addr string, validSigner entities.PublicKeyString,  message []byte) ([]byte, error) {
	if ValidCerts[addr] == "" {
		// get the cert
		certPayload := NewP2pPayload(config, P2pActionGetCert, []byte{'0'})
		err := certPayload.Sign(config.PrivateKeyEDD)
		if err != nil {
			return nil, err
		}
		certResponseByte, err := SendInsecureQuicRequest(addr, certPayload.MsgPack())
		if err != nil {
			return nil, err
		}
		certResponse, err := UnpackP2pPayload(certResponseByte)
		if err != nil {
			return nil, err
		}
		if certResponse.IsValid(config.ChainId) && strings.EqualFold(hex.EncodeToString(certResponse.Signer), string(validSigner)) {
			ValidCerts[addr] = hex.EncodeToString(certResponse.Data)
		} else {
			return nil, fmt.Errorf("quic: invalid signer")
		}
	}
	b, err := sendQuicRequest(addr, message, false)
	if err == ErrInvalidCert {
		ValidCerts[addr] = ""
	}
	return b, err
}
func  SendInsecureQuicRequest(addr string, message []byte) ([]byte, error) {
	return sendQuicRequest(addr, message, true)
}
var ErrInvalidCert = fmt.Errorf("invalid certficate")
func  sendQuicRequest(addr string, message []byte, insecure bool) ([]byte, error) {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"mlayer-p2p"},
		VerifyPeerCertificate: func (rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error  {
			if insecure {
				return nil
			}
			for _, cert := range rawCerts {
				if((ValidCerts)[addr] != hex.EncodeToString(crypto.Keccak256Hash(cert))) {
					return ErrInvalidCert
				}
			}
			return nil
		},
	}
	conn, err := quic.DialAddr(context.Background(), addr, tlsConf, nil)
	if err != nil {
		logger.Debugf("sendQuicRequest/DialAddr: %v", err)
		return nil, err
	}
	defer conn.CloseWithError(0, "")

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		logger.Debugf("sendQuicRequest.OpenStream: %v", err)
		return nil, err
	}
	defer stream.Close()
	_, err = stream.Write([]byte(message))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	data := bytes.Buffer{}
	for {
		n, err := stream.Read(buf)  // Read into the buffer
		data.Write(buf[:n])
		if n == 0 || err == io.EOF {
			break  // End of file, stop reading
		}
		if err != nil  {
			return nil, err
		}
	}
	return data.Bytes(), nil
}