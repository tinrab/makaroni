package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

func main() {
	r := newEngine()
	if err := r.Run(":3000"); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}

func newEngine() *gin.Engine {
	// Init sonyflake
	st := sonyflake.Settings{}
	st.MachineID = machineID
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		log.Fatal("failed to initialize sonyflake")
	}
	// Init router
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// Generate new ID
		id, err := sf.NextID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			// Return ID as string
			c.JSON(http.StatusOK, gin.H{
				"id": fmt.Sprint(id),
			})
		}
	})
	return r
}

func machineID() (uint16, error) {
	ipStr := os.Getenv("MY_IP")
	if len(ipStr) == 0 {
		return 0, errors.New("'MY_IP' environment variable not set")
	}
	ip := net.ParseIP(ipStr)
	if ip == nil || len(ip) < 16 {
		return 0, errors.New("invalid IP")
	}
	return uint16(ip[8])<<7 +uint16(ip[9])<<6 +
	  uint16(ip[10])<<5 + uint16(ip[11])<<4+
    uint16(ip[12])<<3 +uint16(ip[13])<<2 +
	  uint16(ip[14])<<1 + uint16(ip[15]),
	nil
}
