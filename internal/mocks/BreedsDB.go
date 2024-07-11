// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/japhy-tech/backend-test/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// BreedsDB is an autogenerated mock type for the BreedsDB type
type BreedsDB struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *BreedsDB) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *BreedsDB) Insert(_a0 context.Context, _a1 entity.Breeds) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Breeds) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectById provides a mock function with given fields: ctx, id
func (_m *BreedsDB) SelectById(ctx context.Context, id int) ([]entity.Breeds, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SelectById")
	}

	var r0 []entity.Breeds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Breeds, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Breeds); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Breeds)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByName provides a mock function with given fields: ctx, name
func (_m *BreedsDB) SelectByName(ctx context.Context, name string) ([]entity.Breeds, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for SelectByName")
	}

	var r0 []entity.Breeds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Breeds, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Breeds); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Breeds)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByPetSize provides a mock function with given fields: ctx, petSize
func (_m *BreedsDB) SelectByPetSize(ctx context.Context, petSize string) ([]entity.Breeds, error) {
	ret := _m.Called(ctx, petSize)

	if len(ret) == 0 {
		panic("no return value specified for SelectByPetSize")
	}

	var r0 []entity.Breeds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Breeds, error)); ok {
		return rf(ctx, petSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Breeds); ok {
		r0 = rf(ctx, petSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Breeds)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, petSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectBySpecies provides a mock function with given fields: ctx, species
func (_m *BreedsDB) SelectBySpecies(ctx context.Context, species string) ([]entity.Breeds, error) {
	ret := _m.Called(ctx, species)

	if len(ret) == 0 {
		panic("no return value specified for SelectBySpecies")
	}

	var r0 []entity.Breeds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Breeds, error)); ok {
		return rf(ctx, species)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Breeds); ok {
		r0 = rf(ctx, species)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Breeds)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, species)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByWeight provides a mock function with given fields: ctx, weight
func (_m *BreedsDB) SelectByWeight(ctx context.Context, weight int) ([]entity.Breeds, error) {
	ret := _m.Called(ctx, weight)

	if len(ret) == 0 {
		panic("no return value specified for SelectByWeight")
	}

	var r0 []entity.Breeds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Breeds, error)); ok {
		return rf(ctx, weight)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Breeds); ok {
		r0 = rf(ctx, weight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Breeds)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, weight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, breeds
func (_m *BreedsDB) Update(ctx context.Context, breeds entity.Breeds) error {
	ret := _m.Called(ctx, breeds)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Breeds) error); ok {
		r0 = rf(ctx, breeds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBreedsDB creates a new instance of BreedsDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBreedsDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *BreedsDB {
	mock := &BreedsDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}