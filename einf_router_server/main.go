package main

import (
	"context"
	"log"
	"net/http"

	"../pb"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

// grpc server port..
const port = ":9000"

func createExternalId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// no certificates so far for prototype
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("localhost"+port, opts...)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := messages.NewEinf_ServiceClient(conn)
	ctx := context.Background()

	externalId, err := client.CreateExternalIdentity(ctx, &messages.CreateExternalId{Imsi: ps.ByName("imsi")})

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(externalId.GetLocalIdentifier() + "@" + externalId.GetDomain()))
}

func getImsifromExternalId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("localhost"+port, opts...)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := messages.NewEinf_ServiceClient(conn)
	ctx := context.Background()

	imsi, err := client.GetImsiFromExternalId(ctx, &messages.GetImsiFromExtId{ExtId: ps.ByName("extid")})

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(imsi.GetImsi()))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome!\n"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/:authority/neinf_serv/v1/:imsi/externalid", createExternalId)
	router.GET("/:authority/neinf_serv/v1/:extid/imsi", getImsifromExternalId)

	log.Fatal(http.ListenAndServe(":8088", router))
}
