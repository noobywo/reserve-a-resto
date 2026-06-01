package main

import "fmt"

const NMAX int = 9999

const MAXTIME int = 25

type reservationtime struct {
	reserve_start int
	reserve_end   int
	reservestatus reservationstatus
}

type reservationstatus struct {
	status string
}

type arrCustomer struct {
	custID        int
	name          string
	phone         int
	reservedtable int
	reserveinfo   reservationtime
}

type Customer [NMAX]arrCustomer

type arrTable struct {
	tableNumber    int
	capacity       int
	times_reserved int
	reservetime    [MAXTIME]reservationtime
}

type Table [NMAX]arrTable

func main() {
	var c Customer
	var t Table
	var opt string
	var nC, nT, count, i int
	fmt.Println("Welcome to the restaurant reservation system!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Add table")
	fmt.Println("2. Add customer")
	fmt.Println("3. Change table information")
	fmt.Println("4. Change customer information")
	fmt.Println("5. Reserve table")
	fmt.Println("6. Search table by capacity")
	fmt.Println("7. Search table by number")
	fmt.Println("8. Sort tables by capacity")
	fmt.Println("9. Display statistics")
	fmt.Println("10. Exit")
	fmt.Println(" ")
	fmt.Scan(&opt)
	for opt != "10" {
		switch opt {
		case "1":
			addTable(&t, &nT)
		case "2":
			addCustomer(&c, &nC)
		case "3":
			changeTable(&t, nT)
		case "4":
			changeCustomer(&c, nC)
		case "5":
			reserveTable(&t, nT, &count)
		case "6":
			searchTablecapacity(&t, nT)
		case "7":
			searchTableNumber(&t, nT)
		case "8":
			sortCapacity(&t, nT)
			fmt.Print("=======================================\n")
			fmt.Printf("The tables sorted by capacity are: \n")
			for i = 0; i < nT; i++ {
				fmt.Printf("Table number: %d, Capacity: %d\n", t[i].tableNumber, t[i].capacity)
			}
			fmt.Print("=======================================\n")
		case "9":
			sortTimesReservedDesc(&t, nT)
			fmt.Print("=======================================\n")
			fmt.Printf("The top most reserved tables are:\n %d with %d reservations\n %d with %d reservations\n %d with %d reservations\n", t[0].tableNumber, t[0].times_reserved, t[1].tableNumber, t[1].times_reserved, t[2].tableNumber, t[2].times_reserved)
			fmt.Printf("The total number of reservations made is: %d\n", count)
			fmt.Print("=======================================\n")
		case "10":
			fmt.Println("Thank you for using the restaurant reservation system. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
		fmt.Println(" ")
		fmt.Println("Please select an option:")
		fmt.Println("1. Add table")
		fmt.Println("2. Add customer")
		fmt.Println("3. Change table information")
		fmt.Println("4. Change customer information")
		fmt.Println("5. Reserve table")
		fmt.Println("6. Search table by capacity")
		fmt.Println("7. Search table by number")
		fmt.Println("8. Sort tables by capacity")
		fmt.Println("9. Display statistics")
		fmt.Println("10. Exit")
		fmt.Println(" ")
		fmt.Scan(&opt)
	}
}

func addTable(t *Table, n *int) {
	var i int = *n
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number: ")
	fmt.Scan(&t[i].tableNumber)
	fmt.Print("Enter table capacity: ")
	fmt.Scan(&t[i].capacity)
	fmt.Print("=======================================\n")
	t[i].times_reserved = 0
	for j := 0; j < MAXTIME; j++ {
		t[i].reservetime[j].reserve_start = -1
		t[i].reservetime[j].reserve_end = -1
		t[i].reservetime[j].reservestatus.status = "available"
	}
	*n++
}

func addCustomer(c *Customer, n *int) {
	var i int = *n
	fmt.Print("=======================================\n")
	fmt.Print("Enter customer ID: ")
	fmt.Scan(&c[i].custID)
	fmt.Print("Enter customer name: ")
	fmt.Scan(&c[i].name)
	fmt.Print("Enter customer phone number: ")
	fmt.Scan(&c[i].phone)
	c[i].reservedtable = -1
	c[i].reserveinfo.reserve_start = -1
	c[i].reserveinfo.reserve_end = -1
	c[i].reserveinfo.reservestatus.status = "no reservation"
	fmt.Print("=======================================\n")
	*n++
}

func changeTable(t *Table, n int) {
	var tableNum, i int
	var foundtable bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to change information about: ")
	fmt.Scan(&tableNum)
	fmt.Print("Enter the new table number: ")
	for i = 0; i < n && !foundtable; i++ {
		if t[i].tableNumber == tableNum {
			foundtable = true
			i--
		}
	}
	if foundtable {
		fmt.Scan(&t[i].tableNumber)
		fmt.Print("Enter the new table capacity: ")
		fmt.Scan(&t[i].capacity)
	} else {
		fmt.Println("Table not found.")
	}
	fmt.Print("=======================================\n")
}

func changeCustomer(c *Customer, n int) {
	var custID, i int
	var custfound bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter customer ID to change information about: ")
	fmt.Scan(&custID)
	for i = 0; i < n; i++ {
		custfound = c[i].custID == custID && !custfound
	}
	if custfound {
		fmt.Println("Enter the new customer name: ")
		fmt.Scan(&c[i].name)
		fmt.Print("Enter the new customer phone number: ")
		fmt.Scan(&c[i].phone)
	} else {
		fmt.Println("Customer not found.")
	}
	fmt.Print("=======================================\n")
}

func reserveTable(t *Table, nT int, count *int) {
	var tableNum, custID, i, start, end, st int
	var available bool = true
	var found bool = false
	var invalidTime bool = false
	var tableIdx int = -1
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to reserve: ")
	fmt.Scan(&tableNum)
	fmt.Print("Enter reservation start hour (24 hour format): ")
	fmt.Scan(&start)
	fmt.Print("Enter reservation end hour (24 hour format): ")
	fmt.Scan(&end)
	fmt.Print("Enter customer ID for reservation: ")
	fmt.Scan(&custID)

	invalidTime = start < 0 || start >= MAXTIME || end <= 0 || end > MAXTIME || end <= start

	for i = 0; i < nT; i++ {
		if t[i].tableNumber == tableNum && tableIdx == -1 {
			found = true
			tableIdx = i
		}
	}

	if found && !invalidTime {
		for st = start; st < end; st++ {
			available = t[tableIdx].reservetime[st].reservestatus.status == "available"
		}
		if available {
			for st = start; st < end; st++ {
				t[tableIdx].reservetime[st].reserve_start = start
				t[tableIdx].reservetime[st].reserve_end = end
				t[tableIdx].reservetime[st].reservestatus.status = "reserved"
			}
			t[tableIdx].times_reserved++
			(*count)++
		}
	}

	if invalidTime {
		fmt.Println("Invalid reservation time.")
	} else if !found {
		fmt.Println("Table not found.")
	} else if available {
		fmt.Println("Table is reserved successfully.")
	} else {
		fmt.Println("Table is not available at this time.")
	}
	fmt.Print("=======================================\n")
}

func searchTablecapacity(t *Table, n int) {
	var capacity, i int
	var foundcapacity bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table capacity to search for: ")
	fmt.Scan(&capacity)
	for i = 0; i < n && !foundcapacity; i++ {
		if t[i].capacity == capacity {
			foundcapacity = true
			i--
		}
	}
	if foundcapacity {
		fmt.Printf("Table number: %d\n", t[i].tableNumber)
		fmt.Printf("Table capacity: %d\n", t[i].capacity)
	} else {
		fmt.Println("No tables found with the specified capacity.")
	}
	fmt.Print("=======================================\n")
}

func searchTableNumber(t *Table, n int) {
	var tableNum, i int
	var foundtable bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to search for: ")
	fmt.Scan(&tableNum)
	for i = 0; i < n && !foundtable; i++ {
		if t[i].tableNumber == tableNum {
			foundtable = true
			i--
		}
	}
	if foundtable {
		fmt.Printf("Table capacity: %d\n", t[i].capacity)
		fmt.Printf("Table Number: %d\n", t[i].tableNumber)
	} else {
		fmt.Println("Table not found.")
	}
	fmt.Print("=======================================\n")
}

func sortCapacity(t *Table, n int) {
	var temp arrTable
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if t[j].capacity > t[j+1].capacity {
				temp = t[j]
				t[j] = t[j+1]
				t[j+1] = temp
			}
		}
	}
}

func sortTimesReservedDesc(t *Table, n int) {
	var temp arrTable
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if t[j].times_reserved < t[j+1].times_reserved {
				temp = t[j]
				t[j] = t[j+1]
				t[j+1] = temp
			}
		}
	}
}
