package main

import (
    "fmt"
    "os"
    redis "gopkg.in/redis.v3"
    cmd "github.com/jessevdk/go-flags"
)

func main()  {

    var opts struct {
        Host        string `short:"h" long:"host" description:"Host"`
        Port        string `short:"p" long:"port" description:"Port"`
        Password    string `short:"P" long:"password" description:"Password"`
        Command     string `short:"c" long:"command" description:"Command" required:"true"`
        Arguments   string `short:"a" long:"arguments" description:"Command Arguments"`
    }

    argsWithoutProg := os.Args[1:]

    _, err := cmd.ParseArgs(&opts, argsWithoutProg)

    if nil != err {
        panic(err)
        os.Exit(1)
    }

    if "" == opts.Host {
        opts.Host = "localhost"
    }

    if "" == opts.Port {
        opts.Port = "6379"
    }

    fmt.Println(opts.Host)
    fmt.Println(opts.Port)
    fmt.Println(opts.Password)
    fmt.Println(opts.Command)
    fmt.Println(opts.Arguments)

    client := redis.NewClient(&redis.Options{
        Addr        : opts.Host + ":" + opts.Port,
        Password    : opts.Password,
    })

    if opts.Command == "del" {
        results, err := client.Keys(opts.Arguments).Result()
        if nil == err {
            result := client.Del(results...)
            fmt.Println(result)
        }
    } else if opts.Command == "keys" {
        results, err := client.Keys(opts.Arguments).Result()

        if nil == err {
            for i := 0; i < len(results); i++  {
                fmt.Println(results[i])
            }
        }

    }

}
