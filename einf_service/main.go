package main

import (
	"context"
	"einf_demo/pb"
	"einf_demo/util"
	"encoding/hex"
	"log"
	"net"
	"strings"

	"github.com/capitalone/fpe/ff1"

	"google.golang.org/grpc"
)

// gprc server port
const port = ":9000"

//number of added chars to imsi to get different external ids for an imsi
const numberrandchars = 4

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	messages.RegisterEinf_ServiceServer(s, new(Einf_Service))
	log.Println("Starting server on port: " + port)
	s.Serve(lis)
}

type Einf_Service struct{}

func (s *Einf_Service) CreateExternalIdentity(ctx context.Context,
	req *messages.CreateExternalId) (*messages.ExternalId, error) {

	key, _ := hex.DecodeString("34ab865acb5678fe7cb98013645feb99")
	plaintext := req.Imsi + util.GetRandomChars(numberrandchars)

	tweak, err := hex.DecodeString("D8E7920AFA330A73")
	if err != nil {
		log.Fatal(err)
	}

	FF1, err := ff1.NewCipher(36, 8, key, tweak)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := FF1.Encrypt(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	var externalId = messages.ExternalId{}
	externalId.LocalIdentifier = ciphertext
	externalId.Domain = "mydomain"
	return &externalId, nil
}

func (s *Einf_Service) GetImsiFromExternalId(ctx context.Context,
	req *messages.GetImsiFromExtId) (*messages.Imsi, error) {

	var imsi = messages.Imsi{}

	ciphertext := strings.Split(req.ExtId, "@")[0]

	key, _ := hex.DecodeString("34ab865acb5678fe7cb98013645feb99")

	tweak, err := hex.DecodeString("D8E7920AFA330A73")
	if err != nil {
		log.Fatal(err)
	}

	FF1, err := ff1.NewCipher(36, 8, key, tweak)
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := FF1.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	tempstring := string(plaintext)[:len(plaintext)-numberrandchars]

	imsi.Imsi = tempstring
	return &imsi, nil
}
