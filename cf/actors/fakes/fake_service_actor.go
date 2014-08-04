// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/actors"
	"github.com/cloudfoundry/cli/cf/models"
	"sync"
)

type FakeServiceActor struct {
	FilterBrokersStub        func(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error)
	filterBrokersMutex       sync.RWMutex
	filterBrokersArgsForCall []struct {
		brokerFlag  string
		serviceFlag string
		orgFlag     string
	}
	filterBrokersReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	AttachPlansToServiceStub        func(models.ServiceOffering) (models.ServiceOffering, error)
	attachPlansToServiceMutex       sync.RWMutex
	attachPlansToServiceArgsForCall []struct {
		arg1 models.ServiceOffering
	}
	attachPlansToServiceReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
}

func (fake *FakeServiceActor) FilterBrokers(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error) {
	fake.filterBrokersMutex.Lock()
	defer fake.filterBrokersMutex.Unlock()
	fake.filterBrokersArgsForCall = append(fake.filterBrokersArgsForCall, struct {
		brokerFlag  string
		serviceFlag string
		orgFlag     string
	}{brokerFlag, serviceFlag, orgFlag})
	if fake.FilterBrokersStub != nil {
		return fake.FilterBrokersStub(brokerFlag, serviceFlag, orgFlag)
	} else {
		return fake.filterBrokersReturns.result1, fake.filterBrokersReturns.result2
	}
}

func (fake *FakeServiceActor) FilterBrokersCallCount() int {
	fake.filterBrokersMutex.RLock()
	defer fake.filterBrokersMutex.RUnlock()
	return len(fake.filterBrokersArgsForCall)
}

func (fake *FakeServiceActor) FilterBrokersArgsForCall(i int) (string, string, string) {
	fake.filterBrokersMutex.RLock()
	defer fake.filterBrokersMutex.RUnlock()
	return fake.filterBrokersArgsForCall[i].brokerFlag, fake.filterBrokersArgsForCall[i].serviceFlag, fake.filterBrokersArgsForCall[i].orgFlag
}

func (fake *FakeServiceActor) FilterBrokersReturns(result1 []models.ServiceBroker, result2 error) {
	fake.filterBrokersReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceActor) AttachPlansToService(arg1 models.ServiceOffering) (models.ServiceOffering, error) {
	fake.attachPlansToServiceMutex.Lock()
	defer fake.attachPlansToServiceMutex.Unlock()
	fake.attachPlansToServiceArgsForCall = append(fake.attachPlansToServiceArgsForCall, struct {
		arg1 models.ServiceOffering
	}{arg1})
	if fake.AttachPlansToServiceStub != nil {
		return fake.AttachPlansToServiceStub(arg1)
	} else {
		return fake.attachPlansToServiceReturns.result1, fake.attachPlansToServiceReturns.result2
	}
}

func (fake *FakeServiceActor) AttachPlansToServiceCallCount() int {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return len(fake.attachPlansToServiceArgsForCall)
}

func (fake *FakeServiceActor) AttachPlansToServiceArgsForCall(i int) models.ServiceOffering {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return fake.attachPlansToServiceArgsForCall[i].arg1
}

func (fake *FakeServiceActor) AttachPlansToServiceReturns(result1 models.ServiceOffering, result2 error) {
	fake.attachPlansToServiceReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

var _ ServiceActor = new(FakeServiceActor)