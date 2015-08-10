package helpers

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)


var adjustments = [][]int64 {{0,0,0},{0,12,13},{36,45,50}}

var predefinedColors = []string {"008bb9", "FF00FF","FF0000","8B4513","FF8C00", "A9A9A9",
	"FFFF00", "4B0082", "006400","808000","000000","800000",
	"0000FF", "00008B","B8860B", "ADFF2F", "B0C4DE", "FF1493", "8FBC8F", "660000"}

var colorHexStrings []string

func init(){

	indexColor := os.Getenv("INS_CLR")

	var colorString string

	if indexColor != "" && indexColor =="true" {
		clrIndex, err := FetchIndex()
		if err != nil || clrIndex < 0 {
			clrIndex = 0
		}
		
		for clrIndex > len(predefinedColors)-1{
			clrIndex = clrIndex%len(predefinedColors)
		}

		colorString = predefinedColors[clrIndex]
		//do some logic to deal with more indexes than in predefined
	}else{
		colorString = os.Getenv("HEX_COLOR")
	}

        if colorString == "" {
		colorString = os.Getenv("COLOR_NUM")
		if colorString != "" {
			colorIndex, indexErr := strconv.ParseInt(colorString,10,0) 
			if indexErr != nil || colorIndex < 0 || colorIndex > int64(len(predefinedColors))  {
				fmt.Println("Invalid color index using default")
				colorString = predefinedColors[0]
			}else{
				colorString = predefinedColors[colorIndex-1]
			}
		}else{
			fmt.Println("No color set using default")
			colorString = predefinedColors[0]
		}
        }else if strings.ToLower(colorString) == "ffffff"{
		fmt.Println("No color set to white not valid using default")
		colorString = predefinedColors[0]

	}

	colors, err := parseColorString(colorString)
	
	if err != nil {
		fmt.Printf("Cannot parse string %s using default color scheme\n", colorString)

		colors, err = parseColorString(predefinedColors[0])
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

	colorHexStrings = createColorScheme(colors)
}

func FetchColors()[]string{
	
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




