package main

import (
	"fmt"
	"log"
	"time"

	client "github.com/lazybark/go-tls-client/v1"
)

func main() {
	ipsum := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Viverra nibh cras pulvinar mattis nunc sed. Congue nisi vitae suscipit tellus. Enim sit amet venenatis urna cursus. Egestas integer eget aliquet nibh. Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt. Feugiat vivamus at augue eget arcu dictum varius. Tincidunt praesent semper feugiat nibh sed pulvinar proin gravida. Neque gravida in fermentum et sollicitudin. Purus in massa tempor nec feugiat nisl. Vitae purus faucibus ornare suspendisse. Viverra tellus in hac habitasse. Aliquam sem et tortor consequat id porta nibh. Ipsum suspendisse ultrices gravida dictum fusce. Fermentum iaculis eu non diam phasellus. Ultrices eros in cursus turpis massa. Ut ornare lectus sit amet est placerat in. Id ornare arcu odio ut sem nulla pharetra.`

	c := client.New(nil)

	err := c.DialTo("localhost", 5555, `cert.pem`)
	if err != nil {
		log.Fatal(err)
		time.Sleep(time.Minute)
	}

	err = c.SendString(ipsum)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-c.ErrChan:
			fmt.Println(err)
		case m := <-c.MessageChan:
			fmt.Println("Got message:", string(m.Bytes()))
			time.Sleep(2 * time.Second)
			c.ClientDoneChan <- true
		}
	}

}
