
package main

import (
    "fmt"
    //"time"
    //"math/rand"
)



func collatz_next(i int) int{

    // calculates the next element in the sequence
    if i%2 == 0 {
        return i/2
    } else {
        return i*3+1
    }

}

func collatz_seq(i int){

    mynum := i
    counter := 0

    for mynum > 1{
        counter += 1
        fmt.Printf("counter: %d, number: %d\n",counter,mynum)

        switch {
            case 0 >= mynum:
                break
            case 1 == mynum:
                break
            case 0 == mynum%2:
                mynum = mynum/2
            case 1 == mynum%2:
                mynum = mynum*3+1
        }
    }

    fmt.Println()

}

func collatz_seq_length(number int, map_def map[int] int) int {

    if number == 1  {
        return 1
    }

    if map_def[number] != 0 {
        return map_def[number]
    }

    seq_size := 1 + collatz_seq_length(collatz_next(number),map_def)
    if map_def[number] == 0 {
        map_def[number] = seq_size
    }
    return seq_size

}

func print_first_100_num(map_def map[int] int) {
    
    for i:= 1; i< 50; i++ {
        fmt.Printf("%d --> %d\n",i,map_def[i])
        //collatz_seq(i)
    }

}

func longest(map_def map[int] int) (int,int){
    max := 0
    ind := 0
    for count := 1000000; count > 0; count-- {
        if max < map_def[count]{
            ind = count
            max = map_def[count]
        }
    }
    return ind, max
}

func main() {

    //rand.Seed( time.Now().UTC().UnixNano())
    //randnumber := rand.Intn(1000)

    map_defined := make(map[int] int)

    for count := 1000000; count > 0; count-- {
        collatz_seq_length(count,map_defined)
    }

    max_ind, max_val := longest(map_defined)
    fmt.Printf("Longest seq: %d at %d items long\n",max_ind,max_val)
    //print_first_100_num(map_defined)

}
