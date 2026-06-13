// Restaurant Reservation System in Go
// This program allows users to manage restaurant reservations by adding tables and customers, changing their information, reserving tables, searching for tables, sorting tables by capacity, displaying reservation statistics, and deleting tables and customers.

// Made by Rehaan Zulfikar Parkar 103012560001 & Steve Nathaniel

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
	// Main user interface loop
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
	fmt.Println("10. Delete Table")
	fmt.Println("11. Delete Customer")
	fmt.Println("12. Exit")
	fmt.Println(" ")
	fmt.Scan(&opt)
	// Loop with switch case until the user chooses to exit
	for opt != "12" {
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
			reserveTable(&t, nT, &c, nC, &count)
		case "6":
			sortCapacity(&t, nT)
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
			deleteTable(&t, &nT)
		case "11":
			deleteCustomer(&c, &nC)
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
		fmt.Println("10. Delete Table")
		fmt.Println("11. Delete Customer")
		fmt.Println("12. Exit")
		fmt.Println(" ")
		fmt.Scan(&opt)
	}
}

// Function to add a new table to the system ( Steve Nathaniel )
func addTable(t *Table, n *int) {
	var i int = *n
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number: ")
	fmt.Scan(&t[i].tableNumber)
	fmt.Print("Enter table capacity: ")
	fmt.Scan(&t[i].capacity)
	fmt.Print("=======================================\n")
	t[i].times_reserved = 0
	// Initialize reservation times for the new table
	for j := 0; j < MAXTIME; j++ {
		t[i].reservetime[j].reserve_start = -1
		t[i].reservetime[j].reserve_end = -1
		t[i].reservetime[j].reservestatus.status = "available"
	}
	*n++
}

// Function to add a new customer to the system ( Steve Nathaniel )
func addCustomer(c *Customer, n *int) {
	var i int = *n
	fmt.Print("=======================================\n")
	fmt.Print("Enter customer ID: ")
	fmt.Scan(&c[i].custID)
	fmt.Print("Enter customer name: ")
	fmt.Scan(&c[i].name)
	fmt.Print("Enter customer phone number: ")
	fmt.Scan(&c[i].phone)
	// Initialize reservation information for the new customer
	c[i].reservedtable = -1
	c[i].reserveinfo.reserve_start = -1
	c[i].reserveinfo.reserve_end = -1
	c[i].reserveinfo.reservestatus.status = "no reservation"
	fmt.Print("=======================================\n")
	*n++
}

// Function to change information about an existing table ( Steve Nathaniel )
func changeTable(t *Table, n int) {
	var tableNum, i int
	var foundtable bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to change information about: ")
	fmt.Scan(&tableNum)
	// Search for the table by number
	for i = 0; i < n && !foundtable; i++ {
		if t[i].tableNumber == tableNum {
			foundtable = true
			i--
		}
	}
	// If the table is found, asks the user to enter new information
	if foundtable {
		fmt.Print("Enter the new table number: ")
		fmt.Scan(&t[i].tableNumber)
		fmt.Print("Enter the new table capacity: ")
		fmt.Scan(&t[i].capacity)
	} else {
		fmt.Println("Table not found.")
	}
	fmt.Print("=======================================\n")
}

// Function to change information about an existing customer ( Steve Nathaniel )
func changeCustomer(c *Customer, n int) {
	var custID, i int
	var custfound bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter customer ID to change information about: ")
	fmt.Scan(&custID)
	// Search for the customer by ID
	for i = 0; i < n; i++ {
		custfound = c[i].custID == custID && !custfound
	}
	// If the customer is found, asks the user to enter new information
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

// Function to reserve a table for a customer ( Rehaan Zulfikar Parkar )
func reserveTable(t *Table, nT int, c *Customer, nC int, count *int) {
	var tableNum, custID, i, start, end, st int
	var available bool = true
	var found bool = false
	var invalidTime bool = false
	var foundCustomer bool = false
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

	// Check for invalid time inputs
	invalidTime = start < 0 || start >= MAXTIME || end <= 0 || end > MAXTIME || end <= start

	// Check if the customer exists
	for i = 0; i < nC; i++ {
		if c[i].custID == custID {
			foundCustomer = true
		}
	}

	// Search for the table by number and get its index
	for i = 0; i < nT; i++ {
		if t[i].tableNumber == tableNum && tableIdx == -1 {
			found = true
			tableIdx = i
		}
	}

	// If the table is found, the time inputs are valid, and the customer exists, check if the table is available for the requested time
	if found && !invalidTime && foundCustomer {
		for st = start; st < end; st++ {
			available = t[tableIdx].reservetime[st].reservestatus.status == "available"
		}
		// If the table is available, reserve it for the customer by updating the reservation times and status
		if available {
			for st = start; st < end; st++ {
				t[tableIdx].reservetime[st].reserve_start = start
				t[tableIdx].reservetime[st].reserve_end = end
				t[tableIdx].reservetime[st].reservestatus.status = "reserved"
			}
			// Counts the number of times the table has been reserved and increments the total reservation count
			t[tableIdx].times_reserved++
			(*count)++
		}
	}

	// Print appropriate messages based on the results of the reservation attempt
	if invalidTime {
		fmt.Println("Invalid reservation time.")
	} else if !found {
		fmt.Println("Table not found.")
	} else if !foundCustomer {
		fmt.Println("Customer not found.")
	} else if available {
		fmt.Println("Table is reserved successfully.")
	} else {
		fmt.Println("Table is not available at this time.")
	}
	fmt.Print("=======================================\n")
}

// Function to search for tables by capacity using binary search ( Rehaan Zulfikar Parkar )
func searchTablecapacity(t *Table, n int) {
	var left, right, mid int
	var foundtable bool = false
	var capacity int
	fmt.Print("=======================================\n")
	fmt.Print("Enter table capacity to search for: ")
	fmt.Scan(&capacity)
	// Perform binary search on the sorted tables by capacity
	left = 0
	right = n - 1
	for left <= right && !foundtable {
		mid = left + (right-left)/2
		if t[mid].capacity == capacity {
			foundtable = true
		} else if t[mid].capacity < capacity {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if foundtable {
		fmt.Printf("Table number: %d\n", t[mid].tableNumber)
		fmt.Printf("Table capacity: %d\n", t[mid].capacity)
	} else {
		fmt.Println("No tables found with the specified capacity.")
	}
	fmt.Print("=======================================\n")
}

// Function to search for a table by its number using linear search ( Rehaan Zulfikar Parkar )
func searchTableNumber(t *Table, n int) {
	var tableNum, i int
	var foundtable bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to search for: ")
	fmt.Scan(&tableNum)
	// Perform sequential search to find the table by its number
	for i = 0; i < n && !foundtable; i++ {
		if t[i].tableNumber == tableNum {
			foundtable = true
			i--
		}
	}
	// If the table is found, print its information; otherwise, print a not found message
	if foundtable {
		fmt.Printf("Table capacity: %d\n", t[i].capacity)
		fmt.Printf("Table Number: %d\n", t[i].tableNumber)
	} else {
		fmt.Println("Table not found.")
	}
	fmt.Print("=======================================\n")
}

// Function to sort tables by their capacity using insertion sort ( Rehaan Zulfikar Parkar )
func sortCapacity(t *Table, n int) {
	var i, j int
	var key arrTable
	for i = 1; i < n; i++ {
		key = t[i]
		j = i - 1
		for j >= 0 && t[j].capacity > key.capacity {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

// Function to sort tables by the number of times they have been reserved in descending order using selection sort ( Rehaan Zulfikar Parkar )
func sortTimesReservedDesc(t *Table, n int) {
	var temp arrTable
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if t[j].times_reserved > t[maxIdx].times_reserved {
				maxIdx = j
			}
		}
		temp = t[i]
		t[i] = t[maxIdx]
		t[maxIdx] = temp
	}
}

// Function to delete a table from the system ( Rehaan Zulfikar Parkar )
func deleteTable(t *Table, n *int) {
	var tableNum, i int
	var temp arrTable
	var foundtable bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter table number to delete: ")
	fmt.Scan(&tableNum)
	// Search for the table by number
	for i = 0; i < *n && !foundtable; i++ {
		if t[i].tableNumber == tableNum {
			foundtable = true
			// If the table is found, delete it by replacing it with the last table in the array and decrementing the count of tables
			if foundtable {
				temp = t[*n-1]
				t[*n-1] = t[i]
				t[i] = temp
				(*n)--
				i--
				fmt.Println("Table deleted successfully.")
			} else {
				fmt.Println("Table not found.")
			}
		}
	}
	fmt.Print("=======================================\n")
}

// Function to delete a customer from the system ( Rehaan Zulfikar Parkar )
func deleteCustomer(c *Customer, n *int) {
	var custNum, i int
	var temp arrCustomer
	var foundcustomer bool = false
	fmt.Print("=======================================\n")
	fmt.Print("Enter customer ID to delete: ")
	fmt.Scan(&custNum)
	// Search for the customer by ID
	for i = 0; i < *n && !foundcustomer; i++ {
		if c[i].custID == custNum {
			foundcustomer = true
			// If the customer is found, delete it by replacing it with the last customer in the array and decrementing the count of customers
			if foundcustomer {
				temp = c[*n-1]
				c[*n-1] = c[i]
				c[i] = temp
				(*n)--
				i--
				fmt.Println("Customer deleted successfully.")
			} else {
				fmt.Println("Customer not found.")
			}
		}
	}
	fmt.Print("=======================================\n")
}
