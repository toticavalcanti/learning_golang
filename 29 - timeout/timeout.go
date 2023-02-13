// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	timeout := make(chan bool, 1)
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		timeout <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Timeout reached")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"time"
// )

// func main() {
// 	client := &http.Client{
// 		Timeout: 1 * time.Second,
// 	}

// 	resp, err := client.Get("http://localhost:8080/doesntexist")
// 	if err != nil {
// 		if err, ok := err.(*url.Error); ok && err.Timeout() {
// 			fmt.Println("Timeout reached")
// 		}
// 	} else {
// 		defer resp.Body.Close()
// 		fmt.Println("Response received")
// 		_, err := io.ReadAll(resp.Body)
// 		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		//fmt.Println(string(b))
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/fluent_admin")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()

// 	db.SetConnMaxLifetime(5 * time.Second)

// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id int64
// 		var first_name string
// 		var last_name string
// 		var email string
// 		var password string

// 		if err := rows.Scan(&id, &first_name, &last_name, &email, &password); err != nil {
// 			log.Fatalf("Could not scan the result: %v", err)
// 		}

// 		fmt.Println(first_name)

// 		if err := rows.Err(); err != nil {
// 			log.Fatalf("Error while processing the rows: %v", err)
// 		}
// 	}
// }

// _Timeouts_ are important for programs that connect to
// external resources or that otherwise need to bound
// execution time. Implementing timeouts in Go is easy and
// elegant thanks to channels and `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// For our example, suppose we're executing an external
	// call that returns its result on a channel `c1`
	// after 2s. Note that the channel is buffered, so the
	// send in the goroutine is nonblocking. This is a
	// common pattern to prevent goroutine leaks in case the
	// channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Here's the `select` implementing a timeout.
	// `res := <-c1` awaits the result and `<-time.After`
	// awaits a value to be sent after the timeout of
	// 1s. Since `select` proceeds with the first
	// receive that's ready, we'll take the timeout case
	// if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive
	// from `c2` will succeed and we'll print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
