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
conf := client.Config{SuppressErrors: true, MessageTerminator: '\n'}
c := client.New(&conf)
```
Dial to desired server.
Use path to cert in case server has cert that can not be verified. In other cases - just leave empty.
```
err := c.DialTo("localhost", 5555, `cert.pem`)
if err != nil {
  log.Fatal(err)
}
```
Send message using string or bytes. Note that strings will be converted into bytes before sending anyway.
```
err = c.SendString("Hi!")
if err != nil {
  log.Fatal(err)
}
err = c.SendByte([]byte{'H', 'i', '!'})
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
Get stats:
```
fmt.Println(c.Stats())
fmt.Println(c.ErrorsCount())
fmt.Println(c.RecievedBytes())
fmt.Println(c.SentBytes())
```
## Config parameters

* SuppressErrors - set true to avoid errors being sent into ErrChan
* MaxMessageSize - max length of an incoming message in bytes. In case it's reached, client will break the connecton and send error. Also, you may want to consider using maximal message size for your own messages that you send via client. [TLS could be tricky](https://hpbn.co/transport-layer-security-tls/#optimize-tls-record-size).
* MessageTerminator - a byte that signals end of message in TLS stream.
* BufferSize - number of bytes that message reciever can process at once. Depends on purposes of your project.
* DropOldStats - set true if you want to set all client stats to 0 after new connection is made. By default client will compile stats of all connections it has made.
