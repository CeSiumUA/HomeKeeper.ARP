package main

var storedAddresses []string = make([]string, 0)

func GetDifference(scannedAddresses *[]string) (*[]string, *[]string) {
	newlyCreatedAddresses := make([]string, 0)
	newlyDeletedAddresses := make([]string, 0)
	for _, addr := range *scannedAddresses {
		var found bool = false
		for _, storedAddress := range storedAddresses {
			if storedAddress == addr {
				found = true
			}
		}
		if !found {
			newlyCreatedAddresses = append(newlyCreatedAddresses, addr)
		}
	}

	for _, storedAddress := range storedAddresses {
		var found bool = false
		for _, addr := range *scannedAddresses {
			if storedAddress == addr {
				found = true
			}
		}
		if !found {
			newlyDeletedAddresses = append(newlyDeletedAddresses, storedAddress)
		}
	}

	storedAddresses = *scannedAddresses

	return &newlyCreatedAddresses, &newlyDeletedAddresses
}
