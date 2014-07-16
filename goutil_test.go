package goutil
import (
    "testing"
)

func TestMergeChannels(t *testing.T){
    ch1 := make(chan string)
    ch2 := make(chan string)
    go func(){
        ch1 <- "1"
        ch2 <- "2"
        ch1 <- "3"
        ch2 <- "4"
    }()
    ch_list := [] chan string{ch1, ch2}
    out := MergeChannels(ch_list)
    out1 := <- out
    if out1 != "1"{
        t.Error("First message from channel should be 1, got", out1)
    }
    out2 := <- out
    if out2 != "2"{
        t.Error("Second message from channel should be 2, got", out2)
    }
    out3 := <- out
    if out3 != "4"{
        t.Error("Third message from channel should be 3, got", out3)
    }
    out4 := <- out
    if out4 != "3"{
        t.Error("Last message from channel should be 4, got", out4)
    }



}
