package main

import (
  "fmt"
  "time"
)

const (
  a = 1
  b
  c = 2
  d
)

type Weekday int

const (
  Sunday Weekday = iota
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)

type Flags uint

const (
  FlagUp Flags = 1 << iota
  FlagBoradcast
  FlagLoopback
  FlagPointToPoint
  FlagMulticast
)

type Bytes int64
const (
  B Bytes = 1 << (10* iota)
  KiB
  MiB
  GiB
)

func main() {
  const noDelay time.Duration = 0
  const timeout = 5 * time.Minute

  fmt.Printf("%T %[1]v\n", noDelay)
  fmt.Printf("%T %[1]v\n", timeout)
  fmt.Printf("%T %[1]v\n", time.Minute)

  fmt.Println(a, b, c, d)
  fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

  fmt.Println(FlagUp, FlagBoradcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

  fmt.Println(B, KiB, MiB, GiB)

  var v Flags = FlagMulticast | FlagUp
  fmt.Printf("%b %t\n", v, IsUp(v))
  TurnDown(&v)
  fmt.Printf("%b %t\n", v, IsUp(v))
  SetBroadcast(&v)
  fmt.Printf("%b %t\n", v, IsUp(v))
  fmt.Printf("%b %t\n", v, IsCast(v))
}

func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBoradcast }
func IsCast(v Flags) bool { return v&(FlagBoradcast|FlagMulticast) != 0 }
