package main

import "encoding/json"

type Product struct {
	Name   string
	Color  string
	Length int
	Price  int
	Width  int
	Weight int
	Brand  string
	MadeIn string
}
type ElectronicProduct struct {
	Product
	Ram                    int
	Cpu                    string
	ScreenSize             int
	OperatingSystem        string
	OperatingSystemVersion string
}
type Mobile struct {
	ElectronicProduct
	SimcardCapacity int
	SimcardType     string
	NetworkType     string
	CameraCount     int
}
type Laptop struct {
	ElectronicProduct
	UsbPortCount int
	HasCdRom     bool
	KeyboardType string
}

func main() {
	mobile := Mobile{}
	//inherit fields
	mobile.Name = "Samsung A56"
	mobile.Color = "Red"
	mobile.Brand = "Samsung"
	mobile.ScreenSize = 7
	mobile.Ram = 4
	mobile.Cpu = "Snapdragon 8 elite"
	mobile.OperatingSystem = "Android"
	mobile.OperatingSystemVersion = "16"
	//Exclusive fields
	mobile.CameraCount = 3
	mobile.NetworkType = "5G"
	mobile.SimcardCapacity = 2
	mobile.SimcardType = "Nano"

	laptop := Laptop{}
	//inherit fields
	laptop.Name = "Hp Victus 15 fa2013dx"
	laptop.Color = "black"
	laptop.Brand = "Hp"
	laptop.ScreenSize = 15
	laptop.Ram = 16
	laptop.Cpu = "intel core i5 13420h"
	laptop.OperatingSystem = "Windows"
	laptop.OperatingSystemVersion = "11"
	//Exclusive fields
	laptop.HasCdRom = true
	laptop.KeyboardType = "Light"
	laptop.UsbPortCount = 3

	mobileJson, _ := json.Marshal(mobile)
	laptopJson, _ := json.Marshal(laptop)
	println(string(mobileJson))
	println(string(laptopJson))
}
