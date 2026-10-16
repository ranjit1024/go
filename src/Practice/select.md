# Select
select is used to wait and choose among the channel oparations that are ready to be processed 

``go
select {
    case msg : <-  c1{
        fmt.Prinln(msg)
    }
    case msg :<- c3{
        fmt.Prinln(msg)
    }
}
``