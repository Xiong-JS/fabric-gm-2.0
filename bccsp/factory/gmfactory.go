/**
* @Author: XJS
* @Description:
* @File: gmfactory.go
* @Date: 2022/6/12 10:50
*/

package factory


const (
	// SoftwareGMFactoryName is the name of the factory of the software-based BCCSP implementation
	SoftwareGMFactoryName = "GW"
)

type GMFactory struct {}


// Name returns the name of this factory
func (f *GMFactory) Name() string {
	return SoftwareBasedFactoryName
}

