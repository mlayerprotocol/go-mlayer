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
// http.HandleFunc("/rpcendpoint", serveJSONRPC)\