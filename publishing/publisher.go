package publishing

import "homekeeperarp/models"

type Publisher interface {
	Publish(*models.ScanResult) error
}
