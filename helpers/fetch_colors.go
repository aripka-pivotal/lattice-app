package helpers

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

const defaultColor = "008bb9"

var adjustments = [][]int64 {{0,0,0},{0,12,13},{36,45,50}}

func FetchColors()[]string{

	colorString := os.Getenv("HEX_COLOR")

        if colorString == "" {
		fmt.Println("No color set using default")
                colorString = defaultColor
        }else if strings.ToLower(colorString) == "ffffff"{
		fmt.Println("No color set to white not valid using default")
                colorString = defaultColor

	}

	colors, err := parseColorString(colorString)
	
	if err != nil {
		fmt.Printf("Cannot parse string %s using default color scheme\n", colorString)

		colors, err = parseColorString("008bb9")
	}else if colors[0] == 0 && colors[1] == 0 && colors[2] == 0 {
		fmt.Printf("Color String set to Black")
	}


	if colors[0] > 219 {
	  colors[0] = 219
	}

	if colors[1] > 210 {
	  colors[1] = 210
	}

	if colors[2] > 205 {
	  colors[2] = 205
	} 

	colorHexStrings := createColorScheme(colors)
	
	return colorHexStrings
}

func createColorScheme(colors []int64)[]string{
	colorHex := make([]string, 3)

	colorHex[0] = createHexColorString(colors, adjustments[0])
	colorHex[1] = createHexColorString(colors, adjustments[1])
	colorHex[2] = createHexColorString(colors, adjustments[2])

	return colorHex
}

func createHexColorString(vals []int64, adjust []int64) string{
	
	colors := make([]string, 3)
	
	for index, _ := range colors {
		colors[index] = "0" + fmt.Sprintf("%x", vals[index]+adjust[index])
		colors[index] = colors[index][len(colors[index])-2:len(colors[index])]	
		
	}

	return "#"+colors[0]+colors[1]+colors[2]
}

func parseColorString(hexColor string)([]int64, error){
	
	if len(hexColor) < 6 {
		hexColor = "000000" + hexColor
		hexColor = hexColor[len(hexColor)-6:len(hexColor)]
	}
	
	colorRGB := make([]int64, 3)
	
	var err error
	
	for i := 0; i < 3; i++{
		colorRGB[i], err = strconv.ParseInt(hexColor[2*i:(2*i)+2], 16, 0)		
		if err != nil{
			break
		}
	}

	return colorRGB, err
}




