package rpc

// rpc client

// func RpcClient() {

// 	hostname := "localhost"
// 	port := ":9525"

// 	var reply string

// 	args := entities.Message{}

// 	client, err := rpc.DialHTTP("tcp", hostname+port)
// 	if err != nil {
// 		log.Fatal("dialing: ", err)
// 	}

// 	// Call normally takes service name.function name, args and
// 	// the address of the variable that hold the reply. Here we
// 	// have no args in the demo therefore we can pass the empty
// 	// args struct.
// 	err = client.Call("RpcService.SendMessage", args, &reply)
// 	if err != nil {
// 		log.Fatal("error", err)
// 	}

// 	// log the result
// 	log.Printf("%s\n", reply)
// }

// func serveJSONRPC(w http.ResponseWriter, req *http.Request) {
//     if req.Method != "CONNECT" {
//         http.Error(w, "method must be connect", 405)
//         return
//     }
//     conn, _, err := w.(http.Hijacker).Hijack()
//     if err != nil {
//         http.Error(w, "internal server error", 500)
//         return
//     }
//     defer conn.Close()
//     io.WriteString(conn, "HTTP/1.0 Connected\r\n\r\n")
//     jsonrpc.ServeConn(conn)
// }
// http.HandleFunc("/rpcendpoint", serveJSONRPC)
// NODESIGNER 2c2387845a0e17281653050892d3095e7fc99ad32d79b7fbdf11c9a87671daca; Signature: 66cf0db878b0af9d75c1f16942d3275e8d9c6b6cb99680d114fd597715a070a1f7b576b1e21761cbc1cb67cb8bb8dc888be23579e14f91d62dd2f4cd471a4a0d; message: 6c4d9e17290783589b1699bfb48ab7997b7f845d31384ba15ce4dabef870759c00000000000003e96469643a636f736d6f73317a3770757836706574663666766e67646b6170306370796e657a746a3577776d6c76377a39662c2387845a0e17281653050892d3095e7fc99ad32d79b7fbdf11c9a87671daca00000000000000000000018f0baa0beb00000000000003e9b8e053894de3e555b8521915603dc66812c9d5579dcaadd95a9ce771da064e61000000000012ddfc0000018f0baa0c21 