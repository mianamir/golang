package main

import (
	"fmt"
	"math"
	"time"
)

// constants

var incrementAmount int = 20

func main() {

	var firstName string = "Muhammad"
	var lastName string = "Amir"
	fmt.Println("Welcome " + firstName + " " + lastName)
	//fmt.Printf("User first name type is: %T\n", firstName)

	var totalExperince int = 10

	fmt.Println("Total experience is ", totalExperince, " years!")

	var defaultIntValue int

	fmt.Println(defaultIntValue)

	// type inference

	var techStack = "Golang"

	fmt.Println(techStack)

	// initialize the data points with := then don't use var keyword

	// := use inside the functions and use var outside of functions
	//
	//for declaring the data points

	departmentName := "Software Development"

	fmt.Println(departmentName)

	fmt.Println("This user has following increment plan quarterly: ", incrementAmount, "%")

	// User IO data points

	//reader := bufio.NewReader(os.Stdin)
	//
	//fmt.Println("Enter your Docker experience")

	// for ignoring the variable use _

	//dockerExperienceStr, _ := reader.ReadString('\n')

	// type conversion for data point from string to float

	//dockerExperienceFloat, err := strconv.ParseFloat(strings.TrimSpace(dockerExperienceStr), 64)

	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Because you have experience", dockerExperienceFloat,
	//		"years in Docker. You can use AWS ECS.")
	//}

	// declare the multiple data points with same type

	pythonExperience, goLangExperience, rustExperience := 6, 3, 2

	sumExperiences := pythonExperience + goLangExperience + rustExperience

	fmt.Println("Sum of programming experiences", sumExperiences)

	pythonExperienceFloat, goLangExperienceFloat := 66.2, 304.5

	sumExperiencesFloat := pythonExperienceFloat + goLangExperienceFloat

	fmt.Println("Sum of programming experiences float", math.Round(sumExperiencesFloat*100)/100)

	// math calculations

	circleRadius := 12.5

	circumference := circleRadius * 2 * math.Pi

	fmt.Printf("Circumference %.2f\n", circumference)

	// Date time
	timeNow := time.Now()

	fmt.Println("Time now is: ", timeNow)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	fmt.Println("Golang launched at: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTIme, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")

	fmt.Printf("Type of parsed time is %T\n", parsedTIme)

}
