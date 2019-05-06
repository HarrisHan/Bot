package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"os/exec"
	"time"
)




func main() {
	r := gin.New()

	 r.GET("/qrcode/:num", func(c *gin.Context) {

		 c.Writer.Header().Set("Content-Type", "image/jpeg")

		 number := c.Param("num")

		 imgPath := "/Users/harris/Desktop/wechat_write_path/processNum" + number + ".jpeg"

		 if _, err := os.Stat(imgPath); err == nil {
			 // QRCode image exists
			 QRCodeImgFile, _ := os.Open(imgPath)
			 defer QRCodeImgFile.Close()
			 io.Copy(c.Writer,QRCodeImgFile)
			 return
		 }

		openCmd := "open"
		cmd := exec.Command(openCmd, "-n", "/Applications/WeChat.app", "--args", "processNum"+number)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}

		time.Sleep(0.1 * time.Second)

		QRCodeImgFile, err := os.Open(imgPath)
		if err != nil {
			fmt.Println("load image error", err)
		}
		defer QRCodeImgFile.Close()


		io.Copy(c.Writer,QRCodeImgFile)

	})

	r.Run(":3001")
}
