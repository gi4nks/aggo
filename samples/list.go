package samples

import "fmt"

type Channel struct {
	name string
}

func main() {
	var channels []Channel  // an empty list
	channels = append(channels, Channel{name:"some channel name"})
	channels = append(channels, Channel{name:"some other name"})

	// "%+v" prints fields names of structs along with values.
	fmt.Printf("%+v\n", channels)

	// or, create a list of a pre-determined length:
	ten_channels := make([]Channel, 10)
	for i := 0; i < 10; i++ {
		ten_channels[i].name = fmt.Sprintf("chan %d", i)
	}
	fmt.Println(ten_channels)
}
