package main

import (
//	"bytes"
	"strings"
	"context"
	"encoding/json"
	"fmt"
	"io"
"io/ioutil"
	"log"
	"net/http"
//	"net/url"
	"strconv"

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
        ID int32 `json:"id"`
        Parent string `json:"parent"`
}
// https://mholt.github.io/json-to-go/ for generating json

type hashcontent struct {
	Hashcode   [][]interface{} `json:"hashcode"`
	Parenthash string          `json:"parenthash"`
	Account    string          `json:"account"`
}

func (*server) Store(ctx context.Context, req *storepb.StoreRequest) (*storepb.StoreResponse, error) {
	fmt.Printf("store was invoked by %v ", req)

	content := req.GetMsg().GetContent()
	account := req.GetMsg().GetAccount()
    //    id := req.GetMsg().GetId()
     //   parent := req.GetMsg().GetParent()

	result := "content " + content + " " + account 
	res := &storepb.StoreResponse{
		Result: result,
	}

	var a StoreData
	a.Account = account
	a.Content = content
	a.CID = "none"
	storeBTFSFull(&a)

	return res, nil
}

func (*server) SendManyTimes(req *storepb.SendManyTimesRequest, stream storepb.StoreService_SendManyTimesServer) error {

		 fmt.Printf("send many times invoked ")
// read testfile then btfs content and return
	    content := req.GetMsg().GetContent()
            account := req.GetMsg().GetAccount()
            id := req.GetMsg().GetId()
            parent := req.GetMsg().GetParent()
	    for i:=0; i<10; i++ {
		    id1:=int(id)
		result := content + " " + account + " " + strconv.Itoa(id1) + " " + parent
		res := &storepb.SendManyTimesResponse {
			Result: result,
		}
		stream.Send(res)
	    }
	    return nil

}


func (*server) LongStore(stream storepb.StoreService_LongStoreServer) error {
        fmt.Printf("long store was invoked  ")
        
    storeJson := make(map[string]interface{})

        result := ""
      	var b hashcontent
		m := make(map[string]string)
        for {
                req, err := stream.Recv()
                if err == io.EOF {
					storeJson["hashes"] = m


//					Hashcode2, er2 := json.Marshal(m)
  // 				     if er2 != nil {
    //                    log.Fatalf("Error in maps %v ", er2)
//                	 }//if
//b.Hashcode = Hashcode2
    //				 fmt.Printf("json %s", Hashcode2)

	                var jsonData []byte
    	            jsonData, err11 := json.Marshal(storeJson)
	

    	            if err11 != nil {
        	            log.Fatalf("Error while stream: %v ", err11)
            	    }//if
                	fmt.Printf("my function ended")
                	fmt.Printf("json %s", jsonData)
					//filename:=storeJson["parent"]
  					d1 := []byte(jsonData)
    				errx := ioutil.WriteFile("testfile", d1, 0644)
					if errx != nil {
                        log.Fatalf("Error writing file: %v ", errx)
                    }//if

                        // end of stream
                        return stream.SendAndClose(&storepb.LongStoreResponse{
                                Result: result,
                        })

                }
                if err !=nil {
                        log.Fatalf("Error while stream: %v ", err)
                }
                content := req.GetMsg().GetContent()
                account := req.GetMsg().GetAccount()
				id := req.GetMsg().GetId()
	       		 parent := req.GetMsg().GetParent()

                result +=" "+  content + " " + account
       			var a StoreData
        		a.Account = account
        		a.Content = content
        //		a.CID = "none"
				a.Parent = parent
				a.ID = id
				b.Parenthash=parent
				b.Account=account
			    storeJson["parent"] = parent
 				storeJson["account"] = account
 

	
        		storeBTFS(&a)
				m[strconv.Itoa(int(id))]=a.CID


        }
	 	return nil
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
    //fmt.Printf("\nstore value : %s", content)

    sh := shell.NewShell("localhost:5001")
//  r:=[]byte(content)
    r:=strings.NewReader(content)
    cid, err := sh.Add(r)
    if err != nil {
        fmt.Printf("\nerror: %s", err)
    } else {
        fmt.Printf("\n worked: %s ", cid)
	 fmt.Printf("\n worked: %s ", bd.Parent)
 fmt.Printf("\n worked: %d ", bd.ID)

        bd.CID = cid
    }
    //bd.ID
    //cid
  //  hc.Hashcode[][] = [cid, bd.ID]

}

func storeBTFSFull(bd *StoreData) {

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
