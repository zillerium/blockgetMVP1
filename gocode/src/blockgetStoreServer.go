package main

import (
//	"bytes"
	"strings"
	"context"
//	"encoding/json"
	"fmt"
//	"io/ioutil"
	"log"
	"net/http"
//	"net/url"
//	"strconv"

//	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"
//	webFile "github.com/ipfs/go-ipfs-files"

		"net"
	"github.com/blockget-grpc/store/storepb"
	"github.com/blockget-grpc/audit/auditpb"
	"google.golang.org/grpc"
)

type server struct{}

type StoreData struct {
	Account string `json:"account"`
	Content     string `json:"content"`
	CID     string `json:"cid"`
}


func (*server) Store(ctx context.Context, req *storepb.StoreRequest) (*storepb.StoreResponse, error) {
	fmt.Printf("store was invoked by %v ", req)

	content := req.GetMsg().GetContent()
	account := req.GetMsg().GetAccount()
	result := "content " + content + " " + account
	res := &storepb.StoreResponse{
		Result: result,
	}

	var a StoreData
	a.Account = account
	a.Content = content
	a.CID = "none"
	storeBTFS(&a)

	return res, nil
}



func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")

}

func handleRequests() {


	// withinsecure for testing only - add ssl later

	     lis, errs := net.Listen("tcp", "0.0.0.0:50052")
	     if errs !=nil {
	              log.Fatalf("failed to listen %w", errs)
	      } else {
			fmt.Printf("listening on 50052")
	      }
	       s :=grpc.NewServer()
	       storepb.RegisterStoreServiceServer(s, &server{})
	       if err:=s.Serve(lis); err !=nil {
	               log.Fatalf("failed to serve %w", err)

		} else {
			 fmt.Printf("registering worked")
		}

}

func storeBTFS(bd *StoreData) {

	content := bd.Content
	fmt.Printf("\nstore value : %s", content)

	sh := shell.NewShell("localhost:5001")
//	r:=[]byte(content)
	r:=strings.NewReader(content)
	cid, err := sh.Add(r)
	if err != nil {
		fmt.Printf("\nerror: %s", err)
	} else {
		fmt.Printf("\n worked: %s ", cid)
		bd.CID = cid
	}
        cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

        if err != nil {
                log.Fatalf("could not connect: %w", err)
        }

        defer cc.Close()

        c := auditpb.NewAuditServiceClient(cc)

        log.Printf("created client: %f", c)

        fmt.Println("Starting unary RPC")
        req := &auditpb.AuditRequest{
                Msg: &auditpb.AuditMessage{
                        Hashcode: cid,
                        Account:  "trevor3",
                },
        }

        res, err := c.Audit(context.Background(), req)
        if err != nil {
                log.Fatalf("error while calling audit rpc %v", err)
        }

        log.Printf("response from Audit %v", res.Result)

	fmt.Printf("\ncid: %s", cid)
}

func main() {
	handleRequests()
}
