package cmd

import (
	"gorm.io/gorm"
)

type MultiString []string

// "github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
type ExtendedInstanceIdentityDocument struct {
	gorm.Model
	AccountId               string
	Architecture            string
	AvailabilityZone        string
	Epochtime               int64
	ImageId                 string
	InstanceId              string
	InstanceType            string
	KernelId                string
	PendingTime             string
	PrivateIp               string
	RamdiskId               string
	Region                  string
	Version                 string
	BillingProducts         MultiString `gorm:"type:text"` // []string
	DevpayProductCodes      MultiString `gorm:"type:text"` // []string
	MarketplaceProductCodes MultiString `gorm:"type:text"` // []string
}
