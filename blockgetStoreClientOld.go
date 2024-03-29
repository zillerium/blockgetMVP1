package main

import (
//	"bytes"
	"context"
//	"encoding/json"
	"fmt"
//	"io/ioutil"
	"log"
//	"net/http"
//	"net/url"
//	"strconv"

//	"github.com/gorilla/mux"
//	shell "github.com/ipfs/go-ipfs-api"
//	webFile "github.com/ipfs/go-ipfs-files"

	//	"net"
	"github.com/blockget-grpc/store/storepb"
	"google.golang.org/grpc"
)

type server struct{}

type StoreData struct {
        Account string `json:"account"`
        Content     string `json:"content"`
        CID     string `json:"cid"`
}


func (*server) Store(ctx context.Context, req *storepb.StoreRequest) (*storepb.StoreResponse, error) {
        fmt.Printf("audit was invoked by %v ", req)

        content := req.GetMsg().GetContent()
        account := req.GetMsg().GetAccount()
        result := "content " + content + " " + account
        res := &storepb.StoreResponse{
                Result: result,
        }
        return res, nil
}





func handleRequests() {
	storeBTFS()

}

func storeBTFS() {
	content := "This is a test message from, trevor"

        cc, err := grpc.Dial("54.92.219.108:50052", grpc.WithInsecure())

        if err != nil {
                log.Fatalf("could not connect: %w", err)
        }

        defer cc.Close()

        c := storepb.NewStoreServiceClient(cc)

        log.Printf("created client: %f", c)

        fmt.Println("Starting unary RPC")
        req := &storepb.StoreRequest{
                Msg: &storepb.StoreMessage{
                        Content: content,
                        Account:  "trevor3",
                },
        }

        res, err := c.Store(context.Background(), req)
        if err != nil {
                log.Fatalf("error while calling audit rpc %v", err)
        }

        log.Printf("response from Store %v", res.Result)

//	fmt.Printf("\ncid: %s", cid)
}

func main() {
	handleRequests()
}
