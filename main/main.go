package main 
import ("fmt"
  "net/http"
  "io/ioutil")

func main(){
	fmt.Println("Creating a server")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hello)
	err:= http.ListenAndServe(":8080",mux)
	if err !=nil{
		fmt.Println("failed to start server")
	}
}

func hello(w http.ResponseWriter,req *http.Request){
	fmt.Println(" Request received")

	body,_:=ioutil.ReadAll(req.Body)
	responseString:= "Hello, "+string(body)
	w.Write([]byte(responseString))


}