/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-25
 * @fileoverview Backward count by 5s from 100 to 0.
 */

package main

import "fmt"

func main() {
	fmt.Println("Backward Count by 5s:")

	for counter := 100; counter >= 0 ; counter -=5 {
		fmt.Println(counter)
	}
	fmt.Println("\nDone.")
}
