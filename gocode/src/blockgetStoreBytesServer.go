package main

import (
	"bytes"
//"encoding/base64"
"sort"
"reflect"
//	"strings"
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
	"github.com/blockget-grpc/storebytes/storebytespb"
	"github.com/blockget-grpc/audit/auditpb"
	"google.golang.org/grpc"
)

type server struct{}

type StoreData struct {
        Account string `json:"account"`
Content []byte `json:"content"`
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


func (*server) Store(ctx context.Context, req *storebytespb.StoreRequest) (*storebytespb.StoreResponse, error) {
	fmt.Printf("store was invoked by %v ", req)

	content := req.GetMsg().GetContent()
	account := req.GetMsg().GetAccount()
fmt.Println("account %s", account)
    //    id := req.GetMsg().GetId()
     //   parent := req.GetMsg().GetParent()
//	s := string([]byte{content})
//	result := "content " + s+ " " + account 
	res := &storebytespb.StoreResponse{
		Result: content,
	}

	var a StoreData
	a.Account = account
	a.Content = content
	a.CID = "none"
	storeBTFSFull(&a)

	return res, nil
}

  type hashIndex struct {
       Hash string
       Index int
  }
   
  type ByIndex []hashIndex


func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }

func (*server) SendManyTimes(req *storebytespb.SendManyTimesRequest, stream storebytespb.StoreService_SendManyTimesServer) error {

		 fmt.Printf("send many times invoked ")
// read testfile then btfs content and return
	   // content := req.GetMsg().GetContent()
         //   account := req.GetMsg().GetAccount()
           // id := req.GetMsg().GetId()
//            parent := req.GetMsg().GetParent()

		dat, err := ioutil.ReadFile("testfile")
        if err !=nil {
                   log.Fatalf("failed to open  %w", err)
           } 
    	storeJson := make(map[string]interface{})
    	storeJson["parent"] = ""
		storeJson["account"] = ""
    	storeJson["hashes"] = ""
 

    	fmt.Printf(string(dat))
		json.Unmarshal([]byte(dat), &storeJson)

		fmt.Printf("send many times invoked %s ", storeJson["account"])

  //      var m  map[string]interface{}

//		json.Unmarshal([]byte(storeJson["hashes"]), &m)
	//	m = storeJson["hashes"].(map[string]string)
         
		hashes := storeJson["hashes"].(map[string]interface{})

//    	fmt.Printf("send many times invoked %s ", m)

var indArray ByIndex

for key, value := range hashes {
   	// Each value is an interface{} type, that is type asserted as a string
   	fmt.Println(key, value.(string))
	x, err:=strconv.Atoi(key)
    if err !=nil {
 	   log.Fatalf("failed to open  %w", err)
    } 

  	s := hashIndex{Hash: value.(string), Index:x}
	indArray = append(indArray,s)
}

sort.Sort(ByIndex(indArray))

fmt.Println(indArray)

 sh := shell.NewShell("localhost:5001")

//n:=0
//QmXEQPcQsPfGfFoannQG9jiH84P1BxwVAJ5sLapZDZfvWp
 for _, elem := range indArray {
//         elem.Hash
	fmt.Printf("hash code  %s ", elem.Hash)
	rc, err := sh.Cat(fmt.Sprintf("/ipfs/%s",elem.Hash))

      // cid,err := sh.Cat(elem.Hash)
	if err!=nil {
		 fmt.Printf("\nerror: %s ", err)
	}
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(rc)
//	byteArray, err1 = w.Write(buf.Bytes())

 //   if err1!=nil {
 //         fmt.Printf("\nerror: %s ", err1)
 //   }
   var foo = []byte { 0, 0, 0, 0, 0  }

 fmt.Println(reflect.TypeOf(rc))

// buf := []byte(rc)


// used when strings were used	fileBytes := buf.String()

//	fmt.Printf("%+v", fileBytes)



//birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
//var result map[string]interface{}
//json.Unmarshal([]byte(birdJson), &result)

// The object stored in the "birds" key is also stored as 
// a map[string]interface{} type, and its type is asserted from
// the interface{} type
//birds := result["birds"].(map[string]interface{})

//for key, value := range birds {
  // Each value is an interface{} type, that is type asserted as a string
  //fmt.Println(key, value.(string))
//}

	//	    id1:=int(id)
	//	result := content + " " + account + " " + strconv.Itoa(id1) + " " + parent
		res := &storebytespb.SendManyTimesResponse {
			Result: foo,
		}
		stream.Send(res)
	    }
	    return nil

}


func (*server) LongStore(stream storebytespb.StoreService_LongStoreServer) error {
        fmt.Printf("long store was invoked  ")
        
    storeJson := make(map[string]interface{})


     sh := shell.NewShell("localhost:5001")
 //  r:=[]byte(content)

    //    result := ""
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
                        return stream.SendAndClose(&storebytespb.LongStoreResponse{
                                Result: d1,
                        })

                }
                if err !=nil {
                        log.Fatalf("Error while stream: %v ", err)
                }
                content := req.GetMsg().GetContent()
                account := req.GetMsg().GetAccount()
				id := req.GetMsg().GetId()
	       		parent := req.GetMsg().GetParent()

//                result +=" "+  content + " " + account
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
 

	
//        		storeBTFS(&a)

  			    r:=bytes.NewReader(content)
   				cid, err := sh.Add(r)
   				if err != nil {
         	 		fmt.Printf("\nerror: %s", err)
     			} else {
					a.CID = cid
					m[strconv.Itoa(int(id))]=a.CID
      			}


//				m[strconv.Itoa(int(id))]=a.CID
//				m[int(id)] = a.CID

        }
	 	return nil
}


//func Add(data []byte) (string, error) {
	// my have ipfs daemon running first
//	shell := ipfs.NewLocalShell()

//	reader := bytes.NewReader(data)
//	hash, err := shell.Add(reader)
//	if err != nil {
//		log.Fatal(err)
//	}

//	return hash, nil
//}



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
	       storebytespb.RegisterStoreServiceServer(s, &server{})
	       if err:=s.Serve(lis); err !=nil {
	               log.Fatalf("failed to serve %w", err)

		} else {
			 fmt.Printf("registering worked")
		}

}

func storeBTFS(bd *StoreData) {
    content := bd.Content
    //fmt.Printf("\nstore value : %s", content)
     r:=bytes.NewReader(content)

    sh := shell.NewShell("localhost:5001")
//	r:=[]byte(content)
//    r:=strings.NewReader(content)
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
	r:=bytes.NewReader(content)
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
