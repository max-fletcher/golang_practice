package main

// The following structs are unoptimized. They are consuming memory in between fields to maintain the defined memory structures
// type stats struct {
// 	NumPosts uint8
// 	Reach    uint16
// 	NumLikes uint8
// }

// type contact struct {
// 	sendingLimit int32
// 	userID       string
// 	age          int32
// }

// type perms struct {
// 	canSend         bool
// 	canReceive      bool
// 	permissionLevel int
// 	canManage       bool
// }

// The following structs are optimized. They are consuming extra memory in between fields
// Usual rule of thumb is to order the fields in order of their memory consumption in descending(high to low) order
type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {

}
