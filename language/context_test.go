package gotutorial

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"net/http"
	"testing"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work done!")
			return
		default:
			fmt.Println("Working...")
		}
	}
}

func TestContextDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished
	go doWork(ctx)
	// Wait for a while before canceling the context
	select {
	case <-ctx.Done():
		fmt.Println("ctx.Done")
	case <-time.After(time.Second * 3):
		cancel()
	}
}

func TestContextWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	go contextWithTimeout1(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Work completed")
	case <-time.After(10 * time.Second):
		fmt.Println("Work took longer than 10 seconds")
	}
}

func contextWithTimeout1(ctx context.Context) {
	time.Sleep(time.Second * 3)
	fmt.Println("after 3 seconds")
	testContext(ctx)
	time.Sleep(time.Second * 3)
	fmt.Println("after 3 seconds")
	testContext(ctx)
}

func testContext(ctx context.Context) {
	fmt.Println("ctx:", ctx)
}

func TestContextWithDeadLine(t *testing.T) {
	ctx := context.Background()
	deadline := time.Now().Add(time.Second * 5)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestContextSQLQueryTimeOut(t *testing.T) {
	// Open a connection to the database
	db, _ := sql.Open("driverName", "dataSourceName")
	// Create a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Execute the query with the context
	rows, err := db.QueryContext(ctx, "SELECT * FROM table")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	// Handle the query results
	// ...
}

func TestHttpContextTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		fmt.Println("Error:", err) //Error: Get "https://example.com": context deadline exceeded
		return
	}
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
}

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := ctxhttp.Get(ctx, nil, "https://example.com")
	if err != nil {
		fmt.Println("Error:", err) //Error: context deadline exceeded
		return
	}
	defer resp.Body.Close()
}

func TestContextAsKVStore(t *testing.T) {
	oldContext := context.TODO()
	o2 := context.WithValue(oldContext, "user", "12345_old")
	ctx := context.WithValue(o2, "user_id", "12345")
	// use the context in a function
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("user_id").(string)
	fmt.Println("User ID:", userID)

	userID = ctx.Value("user").(string)
	fmt.Println("User ID:", userID)
}
