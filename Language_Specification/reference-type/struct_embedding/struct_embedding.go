package main

import "fmt"

type Notifier interface {
	notify()
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// struct not embedding
type Channel struct {
	Owner       User
	ChannelName string
}

// struct embedding, inner type promotion
type Merchant struct {
	User
	ShopName string
}

func (u *User) notify() {
	fmt.Printf("Type: %T, Sending user email to %s<%s>\n", u, u.Name, u.Email)
}

// polymorphic function
func sendNotification(n Notifier) {
	n.notify()
}

func main() {
	// struct method, call method direct
	user := User{
		Name:  "Naray",
		Email: "naray@gmail.com",
	}
	user.notify()
	sendNotification(&user)
	// struct not embedding, call method through the field owner
	channel := &Channel{
		Owner:       user,
		ChannelName: "channel",
	}
	channel.Owner.notify()
	sendNotification(&channel.Owner)
	// struct embedding, inner type promotion, call method directly
	merchant := &Merchant{
		User: User{
			Name:  "Nurul",
			Email: "nurul@gmail.com",
		},
		ShopName: "shopname",
	}
	merchant.notify()
	sendNotification(merchant)
}
