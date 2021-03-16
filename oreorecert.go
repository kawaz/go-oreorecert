package oreorecert

import (
	"crypto/tls"
	"io/ioutil"

	"github.com/adrg/xdg"
	"github.com/levigross/grequests"
)

type KeyPair struct {
	CertFile string
	KeyFile  string
	Domain   string
}

func (p KeyPair) Certificate() (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(p.CertFile, p.KeyFile)
	if err == nil {
		return &cert, nil
	}
	return p.update()
}

func (p KeyPair) update() (*tls.Certificate, error) {
	err := download("https://oreore.net/crt.pem", p.CertFile)
	if err != nil {
		return nil, err
	}
	err = download("https://oreore.net/key.pem", p.KeyFile)
	if err != nil {
		return nil, err
	}

	cert, err := tls.LoadX509KeyPair(p.CertFile, p.KeyFile)
	return &cert, err
}

func GetCertificateOreoreNet() (*tls.Certificate, error) {
	return GetKeyPairOreoreNet().Certificate()
}

func GetKeyPairOreoreNet() KeyPair {
	d := "oreore.net"
	certFile, _ := xdg.CacheFile("oreorecert/certificates/_." + d + ".crt.pem")
	keyFile, _ := xdg.CacheFile("oreorecert/certificates/_." + d + ".key.pem")
	p := KeyPair{
		Domain:   d,
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	return p
}

func download(url string, file string) error {
	res, err := grequests.Get(url, nil)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, res.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
