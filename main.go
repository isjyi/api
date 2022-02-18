package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	b := make([]byte, 20)
	fmt.Println(b) //

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}

	id := hex.EncodeToString(b)
	fmt.Println(id)
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})
	//
	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//
	//go func() {
	//	// 服务连接
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//// 优雅关闭
	//shutdown.NewHook().Close(
	//	// 关闭 http server
	//	func() {
	//		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//		defer cancel()
	//
	//		if err := srv.Shutdown(ctx); err != nil {
	//			log.Fatal("Server Shutdown:", err)
	//		}
	//	},
	//)
}
