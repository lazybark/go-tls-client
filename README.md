# go-tls-client
Simple TLS client that connects to specified host & port to send and recieve messages in stream. Message delimeter (terminator, separator) can be specified in config. <br>
Client is reusable and can be set up to connect to servers with unknown authority certs (for example, issued by [this app](https://github.com/lazybark/cert-generator)).
## Usage
```
import client "github.com/lazybark/go-tls-client/v1"
```
Create new client with default config:
```
c := client.New(nil)
```
Or customize some parameters:
```
conf := client.Config{SuppressErrors: true,MessageTerminator:'\n'}
c := client.New(&conf)
```
Dial to desired server.
Use path to cert in case server has cert that can not be verified. In other cases - just leave empty
```
err := c.DialTo("localhost", 5555, `cert.pem`)
if err != nil {
  log.Fatal(err)
}
```
Send message:
```
err = c.SendString("Hi!")
if err != nil {
  log.Fatal(err)
}
```
Recieve incoming messages or errors from client via `c.MessageChan` and `c.ErrChan`:
```
for {
  select {
  case err := <-c.ErrChan:
    fmt.Println(err)
  case m := <-c.MessageChan:
    fmt.Println("Got message:", string(m.Bytes()))
  }
}
```
Close the client:
```
c.ClientDoneChan <- true
```
