/*
 * Copyright 2016 DGraph Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/dgraph-io/dgraphgoclient/client"
	"github.com/dgraph-io/dgraphgoclient/graph"
)

var ip = flag.String("ip", "127.0.0.1:8081", "Port to communicate with server")

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	conn, err := grpc.Dial(*ip, grpc.WithInsecure())
	if err != nil {
		log.Fatal("DialTCPConnection")
	}
	defer conn.Close()

	c := graph.NewDgraphClient(conn)

	req := client.NewRequest()
	if err := req.SetMutation("alice", "name", "", "Alice", ""); err != nil {
		log.Fatal(err)
	}
	if err := req.SetMutation("alice", "falls.in", "", "rabbithole", ""); err != nil {
		log.Fatal(err)
	}

	resp, err := c.Query(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}

	req = client.NewRequest()
	req.SetQuery("{ me(_xid_: alice) { name falls.in } }")
	resp, err = c.Query(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}

	fmt.Println("alice", resp.N.Properties)

	req = client.NewRequest()
	if err := req.DelMutation("alice", "name", "", "Alice", ""); err != nil {
		log.Fatal(err)
	}
	resp, err = c.Query(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}

	req = client.NewRequest()
	req.SetQuery("{ me(_xid_: alice) { name falls.in } }")
	resp, err = c.Query(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}
	fmt.Println("alice", resp.N.Properties)
}
