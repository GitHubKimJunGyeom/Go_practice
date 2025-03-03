package post

import "fmt"

type PostSender struct {
}

func (f *PostSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
