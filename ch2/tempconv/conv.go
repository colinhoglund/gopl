package tempconv

// CToF converts celsius values to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts celsius values to kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts fahrenheit values to celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts fahrenheit values to kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * (5 / 9)) }

// KToC converts kelvin values to celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts kelvin values to fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*(9/5) - 459.67) }
