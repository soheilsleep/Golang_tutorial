package main

import "fmt"

type Room struct {
	Id       int
	Type     string
	Status   bool
	BedCount int
	Price    int
}

var rooms = GenerateRooms()

func main() {
	input := ""
	for input != "exit" {
		fmt.Println("Please enter a number:")
		fmt.Println("1: Room List")
		fmt.Println("2: Add Room")
		fmt.Println("3: Reserve Room")

		fmt.Scanln(&input)

		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRooms()
		case "exit":
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid input")
		}
	}
}

func GetRoomList() {
	for _, room := range rooms {
		fmt.Printf("%+v\n", room)
	}
}
func GetRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room information line by line (Id,Type,BedCount,Price)")
	fmt.Scanln(&room.Id)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)
	return room
}

func AddRoom() {
	room := GetRoomFromInput()
	rooms = append(rooms, room)
}

func ReserveRooms() {
	id := 0
	nights := 0
	personCount := 0
	fmt.Println("Enter room id for reservation")
	fmt.Scanln(&id)

	room := GetRoom(id)
	if room == nil {
		fmt.Println("Room Not Found")
		return
	}
	if room.Status {
		println("Room Already Reserved")
		return
	}
	fmt.Println("Enter reserve information line by line ( nights, personCount)")
	fmt.Scanln(&nights)
	fmt.Scanln(&personCount)
	roomPrice, tax, discountAmount, finalPrice := CalculateRoomPrice(*room, nights, personCount)
	room.Status = true
	fmt.Printf("roomPrice is %f, tax is :%f , discountAmount is : %f, finalPrice is : %f\n", roomPrice, tax, discountAmount, finalPrice)
}
func GetRoom(id int) *Room {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Id == id {
			return &rooms[i]
		}
	}
	return nil
}

func CalculateRoomPrice(room Room, nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, finalPrice float64) {
	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights >= 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}
	switch room.Type {
	case "Single":
		roomPrice = float64(nights*room.Price*personCount) * 1.0
	case "Double":
		roomPrice = float64(nights*room.Price*personCount) * 1.2
	case "Suite":
		roomPrice = float64(nights*room.Price*personCount) * 2.0

	}
	tax = roomPrice * 0.09
	discountAmount = roomPrice * discountPercentage
	finalPrice = roomPrice + tax - discountAmount
	return
}

func GenerateRooms() []Room {
	return []Room{
		{Id: 1, Type: "Single", Status: false, BedCount: 1, Price: 100},
		{Id: 2, Type: "Double", Status: false, BedCount: 2, Price: 120},
		{Id: 3, Type: "Suite", Status: false, BedCount: 1, Price: 200},
		{Id: 4, Type: "Single", Status: false, BedCount: 3, Price: 110},
		{Id: 5, Type: "Single", Status: false, BedCount: 4, Price: 300},
		{Id: 6, Type: "Suite", Status: false, BedCount: 2, Price: 210},
		{Id: 7, Type: "Double", Status: false, BedCount: 1, Price: 150},
		{Id: 8, Type: "Single", Status: false, BedCount: 1, Price: 160},
		{Id: 9, Type: "Double", Status: false, BedCount: 3, Price: 230},
		{Id: 10, Type: "Suite", Status: false, BedCount: 3, Price: 320},
	}
}
