package main

import (
	"bytes"
	"context"
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
//	"net/http"
//	"net/url"
//	"strconv"
"crypto/sha256"
"encoding/hex"
//	"github.com/gorilla/mux"
//	shell "github.com/ipfs/go-ipfs-api"
//	webFile "github.com/ipfs/go-ipfs-files"

	//	"net"
	"github.com/blockget-grpc/store/storepb"
	"google.golang.org/grpc"
)

type server struct{}

type StoreData struct {
        ID int32 `json:"id"`
	Parent string `json:"parent"`
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
        return res, nil
}



func handleRequests() {
	storeBTFS()

}

func buildRequest(s []string, hash string) []*storepb.LongStoreRequest {
//	ss:= []storepb.LongStoreRequest
// int32 i =0
//stringv := s[0]
//	for i := 0; i < len(s); i++ {
        //              println(i, apps[i])
//	ss := []*storepb.LongStoreRequest{
/*		&storepb.LongStoreRequest {
                        Msg: &storepb.StoreMessage{
                                Content: stringv,
                                Account:  "trevor3",
                                Parent: "parentrec",
                                Id: 0,
                        },
                        }, 
		}
//	}
	return ss
*/
ss := []*storepb.LongStoreRequest{}
for i,str:=range s {
  ss=append(ss,&storepb.LongStoreRequest {
                  Msg: &storepb.StoreMessage{
                            Content: str,
                            Account:  "trevor3",
                            Parent: hash,
                            Id: int32(i),
                    }})
}
return ss


}

func clientStreaming(c storepb.StoreServiceClient, s []string, hash string) {

	fmt.Println("starting client")

	requests := buildRequest(s, hash)

	stream, err := c.LongStore(context.Background())	
	if err!=nil {
		log.Fatalf("error client stream %v", err)
	}

	for _, req := range requests {
      		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
    	if err!=nil {
                log.Fatalf("error client stream %v", err)
                log.Fatalf("error client stream %v", res)
        }
//	fmt.Printf("long store response %v", res)


}

func splitString(s string, n int) []string {
    sub := ""
    subs := []string{}

    runes := bytes.Runes([]byte(s))
    l := len(runes)
    for i, r := range runes {
        sub = sub + string(r)
        if (i + 1) % n == 0 {
            subs = append(subs, sub)
            sub = ""
        } else if (i + 1) == l {
            subs = append(subs, sub)
        }
    }

    return subs
}

func storeBTFS() {
//	content := "This is a test message from, trevor"

	dat, err := ioutil.ReadFile("sample.pdf")
    
	if err!=nil {
                log.Fatalf("error opening file %v", err)
        }
aStringToHash := []byte(string(dat))
	sha256Bytes := sha256.Sum256(aStringToHash)
 hash := hex.EncodeToString(sha256Bytes[:])
	l:=len(string(dat))
	blocksize := 8192
	segments := l/blocksize
	segments++

	subs:=splitString(string(dat),blocksize)
	
	fmt.Print(len(string(dat)))
	fmt.Print(" segments %d\n ", segments)
fmt.Print(" len array subs %d\n ", len(subs))


 	
	cc, err := grpc.Dial("54.92.219.108:50052", grpc.WithInsecure())

        if err != nil {
                log.Fatalf("could not connect: %w", err)
        }

        defer cc.Close()

        c := storepb.NewStoreServiceClient(cc)

	clientStreaming(c, subs, hash)

        log.Printf("created client: %f", c)

    /*    fmt.Println("Starting unary RPC")
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
*/
//	fmt.Printf("\ncid: %s", cid)
}

func main() {
	handleRequests()
}
