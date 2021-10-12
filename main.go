package main

import(
	"fmt"
  "github.com/bwmarrin/discordgo"
  "strconv"
  "strings"
  "errors"
  "os"
)

type account struct {
  balance float64
}

var  mySecret float64
var BotID string 

func main() {
mySecret := os.Getenv("token")
var a account
  dg, err := discordgo.New("Bot " + mySecret)

  if err != nil {
    fmt.Println(err.Error())
  }

  u, err := dg.User("@me")

  if err != nil {
    fmt.Println(err.Error())
  }

  BotID = u.ID


  err = dg.Open()

   if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("Bot is running!")
  
  dg.AddHandler(messageHandler)
  fmt.Println(a.balance)
  <-make(chan struct{})
  return
}

func (a *account) withdraw(out float64) error {
    if out > a.balance {
        return errors.New("Insufficient Funds")
    }
    if out < 0 {
        return errors.New("Must have positive value")
    }
    a.balance -= out
    return nil
}

func (a *account) deposit(in float64) error {
    if in < 0 {
        return errors.New("Deposits must have money")
    }

    a.balance += in
    return nil
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
  var a account
  input := m.Content
  input = strings.TrimSpace(input)
  inputs := strings.Fields(input)
   
	if m.Author.ID == BotID {
		return
	}

	if inputs[0] == "!bank" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Welcome to our Bank user number " + m.Author.ID + ", enter your request")
      }
  if inputs[0] == "!withdraw" {
    str := inputs[1]
    fmt.Println(str)
    result, err := strconv.ParseFloat(str, 64)
		  if err == nil {
        a.withdraw(result)
        _, _ = s.ChannelMessageSend(m.ChannelID, strconv.FormatFloat(a.balance, 'f', 6, 64))
     }
    } 
  if inputs[0] == "!deposit" {
    str := inputs[1]
    fmt.Println(str)
    result, err := strconv.ParseFloat(str, 64)
		  if err == nil {
        a.deposit(result)
        _, _ = s.ChannelMessageSend(m.ChannelID, strconv.FormatFloat(a.balance, 'f', 6, 64))
    }
  } 
	
    if inputs[0] == "!balance" {
    _, _ = s.ChannelMessageSend(m.ChannelID, strconv.FormatFloat(a.balance, 'f', 6, 64))
  }
}
