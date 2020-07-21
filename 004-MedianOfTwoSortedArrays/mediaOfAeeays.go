package main

import (
    "fmt"
    "math"
)



func find(a, b []int) float64 {
    if len(a) > len(b) {
        return find(b, a)
    }
    
    low, high, midab, mida, midb := 0, len(a), (len(a) + len(b) + 1) >> 1, 0, 0
    for low <= high {
        mida = low + (high - low)>>1
        midb = midab - mida
        if mida > 0 && a[mida - 1] > b[midb] {
            high = mida - 1
        } else if mida != len(a) && a[mida] < b[midb - 1] {
            low = mida + 1
        } else {
            break
        }
    }

    var midLeft, midRight float64
    if mida == 0 {
        midLeft = float64(b[midb - 1])
    } else if midb == 0 {
        midLeft = float64(a[mida - 1])
    } else {
        midLeft = math.Max(float64(a[mida-1]), float64(b[midb-1]))
    }
    if (len(a) + len(b))&1 == 1 { 
        return midLeft
    }

    if mida == len(a) {
        midRight = float64(b[midb])
    } else if midb == len(b) {
        midRight = float64(a[mida])
    } else {
        midRight = math.Min(float64(a[mida]), float64(b[midb]))
    }
    return midLeft + midRight / 2.0
}

func main() {
    mid := find([]int{7, 11, 18, 19, 21, 25}, []int{1, 3, 8, 9, 15} )
    fmt.Println(mid)
}