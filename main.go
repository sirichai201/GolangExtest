package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for {
		var functionToRun int
		fmt.Println("Enter the function to run")
		fmt.Println("1. function1")
		fmt.Println("1_2 function1_2")
		fmt.Println("2. function2")
		fmt.Println("3. function3")
		fmt.Println("3_1. function3_1")
		fmt.Println("4. function4")
		fmt.Println("4_1. function4_1")
		fmt.Println("5. function5")
		fmt.Println("6. function6")
		fmt.Println("7. function7")
		fmt.Println("0. Exit program")
		fmt.Scanln(&functionToRun)

		switch functionToRun {
		case 1:
			runFunctionWithConfirmation("function1")
		case 1_2:
			runFunctionWithConfirmation("function1_2")

		case 2:
			runFunctionWithConfirmation("function2")
		case 3:
			runFunctionWithConfirmation("function3")
		case 3_1:
			runFunctionWithConfirmation("function3_1")
		case 4:
			runFunctionWithConfirmation("function4")
		case 4_1:
			runFunctionWithConfirmation("function4_1")
		case 5:
			runFunctionWithConfirmation("function5")
		case 6:
			runFunctionWithConfirmation("function6")
		case 7:
			runFunctionWithConfirmation("function7")
		case 7_1:
		case 0:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func runFunctionWithConfirmation(functionName string) {
	for {
		var confirmRun string
		fmt.Printf("do you want to run %s (Y/N):", functionName)
		fmt.Scanln(&confirmRun)
		if confirmRun == "Y" || confirmRun == "y" {
			switch functionName {
			case "function1":
				function1()
			case "function1_2":
				function1_2()
			case "function2":
				function2()
			case "function3":
				function3()
			case "function3_1":
				function3_1()
			case "function4":
				function4()
			case "function4_1":
				function4_1()
			case "function5":
				function5()
			case "function6":
				function6()
			case "function7":
				function7()
			}
		} else if confirmRun == "N" || confirmRun == "n" {
			return
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func function1() {
	fmt.Println("Running function1")
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			count++
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\n\nจำนวนทั้งหมด: ", count)
}

func function1_2() {
	fmt.Println("Running function1_2")

	var base, exponent float64
	fmt.Print("Enter the base: ")
	fmt.Scanln(&base)
	fmt.Print("Enter the exponent: ")
	fmt.Scanln(&exponent)

	result := math.Pow(base, exponent)
	fmt.Printf("%.f ยกกำลัง %.f = %.f\n", base, exponent, result)
}

func function2() {
	fmt.Println("Running function2")
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	min := x[0]
	max := x[0]

	for _, v := range x {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Printf("Minimum: %d\n", min)
	fmt.Printf("Maximum: %d\n", max)
}

func function3() {
	fmt.Println("Running function3")
	count := 0
	for i := 0; i <= 1000; i++ {
		numberStr := strconv.Itoa(i)
		for _, digit := range numberStr {
			if digit == '9' {
				count += 1
			}
		}
	}
	fmt.Println("Total number of occurrences of digit 9 from 1 to 1000:", count)
}

func function3_1() {
	fmt.Println("Running function3_1")
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	count := 0
	for i := 1; i <= n; i++ {
		numberStr := strconv.Itoa(i)
		for _, digit := range numberStr {
			if digit == '9' {
				count++
			}
		}
	}
	fmt.Printf("Total occurrences of digit 9 from 1 to %d: %d\n", n, count)
}
func function4() {
	fmt.Println("Running function4")

	myWords := "AW SOME GO!"
	result := ""

	for _, char := range myWords {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Println(result)
}

func cutText(input string) string {
	return strings.ReplaceAll(input, " ", "")
}
func function4_1() {
	fmt.Println("Running function4_1")
	result := cutText("ine t")
	fmt.Println(result)
}
func function5() {
	fmt.Println("Running function5")
	type Person struct {
		Name    string // ชื่อของผู้ใช้
		Age     int    // อายุของผู้ใช้
		Address string // ที่อยู่ของผู้ใช้
		iphone  int    // หมายเลขโทรศัพท์ของผู้ใช้
	}
	people := []Person{
		{"นาย ขนมหวาน", 25, "สกนคร 12356", 1233558889},
		{"นาย โค้ก", 22, "อุดร, 12304", 158888489},
		{"นาย เค้ก", 22, "ขอนแก่น 12344", 33354555},
		{"นางสาว  กิฟ", 22, "ขอนแก่น 12344", 33354555},
	}

	var input string
	fmt.Println("Enter your choice:")        // แสดงข้อความให้ผู้ใช้เลือก
	fmt.Println("0. Print all data")         // แสดงตัวเลือกที่ 0
	fmt.Println("1. Print data of person 1") // แสดงตัวเลือกที่ 1
	fmt.Println("2. Print data of person 2") // แสดงตัวเลือกที่ 2
	fmt.Println("3. Print data of person 3") // แสดงตัวเลือกที่ 3
	fmt.Println("4. Exit")                   // แสดงตัวเลือกที่ 4
	fmt.Print("Choice: ")                    // แสดงข้อความให้ผู้ใช้ป้อนตัวเลือก
	fmt.Scanf("%s", &input)                  // รับข้อมูลจากผู้ใช้

	// ใช้ switch เพื่อตรวจสอบตัวเลือกของผู้ใช้และดำเนินการตามเงื่อนไข
	switch input {
	case "0": // แสดงข้อมูลทั้งหมด
		for i, person := range people {
			fmt.Printf("Person %d\n", i+1)
			fmt.Printf("Name: %s\n", person.Name)
			fmt.Printf("Age: %d\n", person.Age)
			fmt.Printf("Address: %s\n", person.Address)
			fmt.Printf("iphone: %d\n", person.iphone)
			fmt.Println()
		}

	case "1": // แสดงข้อมูลของคนที่ 1
		if len(people) >= 1 {
			person := people[0]
			fmt.Printf("Name: %s\n", person.Name)
			fmt.Printf("Age: %d\n", person.Age)
			fmt.Printf("Address: %s\n", person.Address)
			fmt.Printf("iphone: %d\n", person.iphone)
		} else {
			fmt.Println("Person 1 data not found")
		}

	case "2": // แสดงข้อมูลของคนที่ 2
		if len(people) >= 2 {
			person := people[1]
			fmt.Printf("Name: %s\n", person.Name)
			fmt.Printf("Age: %d\n", person.Age)
			fmt.Printf("Address: %s\n", person.Address)
			fmt.Printf("iphone: %d\n", person.iphone)
		} else {
			fmt.Println("Person 2 data not found")
		}

	case "3": // แสดงข้อมูลของคนที่ 3
		if len(people) >= 3 {
			person := people[2]
			fmt.Printf("Name: %s\n", person.Name)
			fmt.Printf("Age: %d\n", person.Age)
			fmt.Printf("Address: %s\n", person.Address)
			fmt.Printf("iphone: %d\n", person.iphone)
		} else {
			fmt.Println("Person 3 data not found")
		}

	case "4": // ออกจากรายการ
		fmt.Println("Exit")
		return

	default: // กรณีที่ไม่ระบุตัวเลือก
		fmt.Println("Invalid choice")
	}
}
func function6() {
	fmt.Println("Running function6")
	type company struct {
		Name     string
		Location string
		Revenue  float64
	}
	fmt.Println("Running function6")
	var c company
	fmt.Println("Enter company name:")
	fmt.Scanln(&c.Name)
	fmt.Println("Enter company location:")
	fmt.Scanln(&c.Location)
	fmt.Println("Enter company revenue:")
	fmt.Scanln(&c.Revenue)
	fmt.Println()
	fmt.Println("Company Name:", c.Name)
	fmt.Println("Location:", c.Location)
	fmt.Println("Revenue:", c.Revenue)
	fmt.Println()

}
func function7() {
	fmt.Println("Running function7")

	for i := 1; i <= 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
