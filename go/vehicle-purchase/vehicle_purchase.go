package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " " + "is clearly the better choice."
	} else {
		return option2 + " " + "is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var basePrice1 = 0.80 * originalPrice
	var basePrice2 = 0.50 * originalPrice
	var basePrice3 = 0.70 * originalPrice
	if age > 3 && age < 10 {
		return basePrice3
	} else if age >= 10 {
		return basePrice2
	} else if age < 3 {
		return basePrice1
	} else  {
		return originalPrice
	}
}
