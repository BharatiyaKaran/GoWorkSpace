package main

import (
	"fmt"
	"os/exec"
	"log"
)

func  DateCmd() {
	out, err := exec.Command("date", "+%Z").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The TimeZone is %s", out)
}

func getBIOSInfo() {
    fmt.Printf("BIOSInfo\n")
    fmt.Printf("#########\n")
    //BIOSName
    cmdName := "sudo dmidecode -t bios | grep SMBIOS"
    out, _ := exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BIOSName: %s", out)

    //BIOSManufacturer
    cmdName = "sudo dmidecode -t bios | grep Vendor | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BIOSManufacturer: %s", out)

    //BIOSVersion
    cmdName = "sudo dmidecode -t bios | grep Version | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BIOSVersion: %s", out)

    //SMBIOSVersion
    cmdName = "sudo dmidecode -t bios | grep SMBIOS | cut -d \" \" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("SMBIOSVersion: %s", out)
}

func getBaseBoardInfo() {
    fmt.Printf("BaseBoardInfo\n")
    fmt.Printf("#############\n")
    //Product
    cmdName := "sudo dmidecode -t baseboard | grep \"Product Name\" | cut -d \":\" -f2"
    out, _ := exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("Product: %s", out)

    //BaseBoardManufacturer
    cmdName = "sudo dmidecode -t baseboard | grep Manufacturer | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BaseBoardManufacturer: %s", out)

    //BaseBoardSrNo
    cmdName = "sudo dmidecode -t baseboard | grep \"Serial Number\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BaseBoardSrNo: %s", out)

    //BaseBoardVersion
    cmdName = "sudo dmidecode -t baseboard |grep \"Version\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("BaseBoardVersion: %s", out)
}

func getProcessorInfo() {
    fmt.Printf("ProcessorInfo\n")
    fmt.Printf("#############\n")
    //ClockSpeed
    cmdName := "cat /proc/cpuinfo | grep \"model name\" | cut -d \":\" -f2 | cut -d \" \" -f7"
    out, _ := exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("ClockSpeed: %s", out)

    //Family
    cmdName = "cat /proc/cpuinfo | grep \"cpu family\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("Family: %s", out)

    //Manufacturer
    cmdName = "cat /proc/cpuinfo | grep \"vendor_id\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("Manufacturer: %s", out)

    //ProcessorType
    cmdName = "cat /proc/cpuinfo | grep \"model name\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("ProcessorType: %s", out)

    //ProcessorLevel
    cmdName = "cat /proc/cpuinfo | grep \"cpuid level\" | cut -d \":\" -f2"
    out, _ = exec.Command("bash", "-c", cmdName).Output()
    fmt.Printf("ProcessorLevel: %s", out)
}

func main() {
	//DateCmd()
	getBIOSInfo()
	getBaseBoardInfo()
	getProcessorInfo()
}
