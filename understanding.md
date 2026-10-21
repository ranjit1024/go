## Go interface 
```go
type react interface{
    area()
}
```

now anything that has a area function that is belong to react interface 

now if we have two differect struct where both have a area function then we have bike then in put in the interface then it will give bike wheel data and for car it will differnet it will call that car data


-> any type that implements all the metod reqired by an intreface automatically satisfy intreface 

## Range over Itarator
-> the prbolem is that think like that we can range where we consume an array but what if we do sometng we have an slices and then teh use rage 
what if we have a function that can produce a number on the fly and and pass to the range loop body

for that we have someting called as `iter.Seq[int]` this is return type of funciton 
`yeild()` this is the actutl callback function that deliver the nums to the range to indecate that is it do we still need to continue or not

## Go Timeout 
if we dont receive a some data from channel for certain time then we do print timeout 

```go
ch1 := make(chan string, 1)
go func(){
    time.Sleep(2 * time.Second)
    ch1 <- "Data"
}()
select{
case msg := <- ch1:
    fmt.Prinln(msg)
case <- time.After(1 * time.Second):
    fmt.Prinln("Timeout")
}
```
## Go Timer
go timer is used to perform after certain time taht timer will be expire and the notification send to the channel
```go
import (
    "fmt",
    "time"
)
func main(){
    timer1 := time.NewTimer(2 * time.Second)
    <- timer1.C;
}
```