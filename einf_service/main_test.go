package main

import (
	"einf_demo/pb"
	"net/http"
	"testing"
)

func TestCreateExternalIdandGetIMSI(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := req.Context()
	externalidmssg := messages.CreateExternalId{Imsi: "123456789"}

	s := Einf_Service{}

	externalId, err := s.CreateExternalIdentity(ctx, &externalidmssg)
	if err != nil {
		t.Fatal(err)
	}

	imsimssg := messages.GetImsiFromExtId{ExtId: externalId.GetLocalIdentifier() + "@" + externalId.GetDomain()}

	imsi, err := s.GetImsiFromExternalId(ctx, &imsimssg)
	if err != nil {
		t.Fatal(err)
	}

	if imsi.GetImsi() != "123456789" {
		t.Errorf("imsi obtained not correct, got: %v, want 123456789", imsi.GetImsi())
	}
}
