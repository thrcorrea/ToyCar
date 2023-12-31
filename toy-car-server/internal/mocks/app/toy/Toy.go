// Code generated by mockery v2.16.0. DO NOT EDIT.

package toy

import (
	mock "github.com/stretchr/testify/mock"
	apptoy "github.com/toy-simulator/internal/app/toy"
)

// Toy is an autogenerated mock type for the Toy type
type Toy struct {
	mock.Mock
}

type Toy_Expecter struct {
	mock *mock.Mock
}

func (_m *Toy) EXPECT() *Toy_Expecter {
	return &Toy_Expecter{mock: &_m.Mock}
}

// GetPosition provides a mock function with given fields:
func (_m *Toy) GetPosition() (*apptoy.ToyPosition, error) {
	ret := _m.Called()

	var r0 *apptoy.ToyPosition
	if rf, ok := ret.Get(0).(func() *apptoy.ToyPosition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apptoy.ToyPosition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Toy_GetPosition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPosition'
type Toy_GetPosition_Call struct {
	*mock.Call
}

// GetPosition is a helper method to define mock.On call
func (_e *Toy_Expecter) GetPosition() *Toy_GetPosition_Call {
	return &Toy_GetPosition_Call{Call: _e.mock.On("GetPosition")}
}

func (_c *Toy_GetPosition_Call) Run(run func()) *Toy_GetPosition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Toy_GetPosition_Call) Return(_a0 *apptoy.ToyPosition, _a1 error) *Toy_GetPosition_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Left provides a mock function with given fields:
func (_m *Toy) Left() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Toy_Left_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Left'
type Toy_Left_Call struct {
	*mock.Call
}

// Left is a helper method to define mock.On call
func (_e *Toy_Expecter) Left() *Toy_Left_Call {
	return &Toy_Left_Call{Call: _e.mock.On("Left")}
}

func (_c *Toy_Left_Call) Run(run func()) *Toy_Left_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Toy_Left_Call) Return(_a0 error) *Toy_Left_Call {
	_c.Call.Return(_a0)
	return _c
}

// Move provides a mock function with given fields:
func (_m *Toy) Move() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Toy_Move_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Move'
type Toy_Move_Call struct {
	*mock.Call
}

// Move is a helper method to define mock.On call
func (_e *Toy_Expecter) Move() *Toy_Move_Call {
	return &Toy_Move_Call{Call: _e.mock.On("Move")}
}

func (_c *Toy_Move_Call) Run(run func()) *Toy_Move_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Toy_Move_Call) Return(_a0 error) *Toy_Move_Call {
	_c.Call.Return(_a0)
	return _c
}

// Place provides a mock function with given fields: x, y, direction
func (_m *Toy) Place(x int, y int, direction apptoy.Direction) error {
	ret := _m.Called(x, y, direction)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, apptoy.Direction) error); ok {
		r0 = rf(x, y, direction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Toy_Place_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Place'
type Toy_Place_Call struct {
	*mock.Call
}

// Place is a helper method to define mock.On call
//   - x int
//   - y int
//   - direction apptoy.Direction
func (_e *Toy_Expecter) Place(x interface{}, y interface{}, direction interface{}) *Toy_Place_Call {
	return &Toy_Place_Call{Call: _e.mock.On("Place", x, y, direction)}
}

func (_c *Toy_Place_Call) Run(run func(x int, y int, direction apptoy.Direction)) *Toy_Place_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(apptoy.Direction))
	})
	return _c
}

func (_c *Toy_Place_Call) Return(_a0 error) *Toy_Place_Call {
	_c.Call.Return(_a0)
	return _c
}

// Report provides a mock function with given fields:
func (_m *Toy) Report() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Toy_Report_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Report'
type Toy_Report_Call struct {
	*mock.Call
}

// Report is a helper method to define mock.On call
func (_e *Toy_Expecter) Report() *Toy_Report_Call {
	return &Toy_Report_Call{Call: _e.mock.On("Report")}
}

func (_c *Toy_Report_Call) Run(run func()) *Toy_Report_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Toy_Report_Call) Return(_a0 string, _a1 error) *Toy_Report_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Right provides a mock function with given fields:
func (_m *Toy) Right() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Toy_Right_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Right'
type Toy_Right_Call struct {
	*mock.Call
}

// Right is a helper method to define mock.On call
func (_e *Toy_Expecter) Right() *Toy_Right_Call {
	return &Toy_Right_Call{Call: _e.mock.On("Right")}
}

func (_c *Toy_Right_Call) Run(run func()) *Toy_Right_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Toy_Right_Call) Return(_a0 error) *Toy_Right_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewToy interface {
	mock.TestingT
	Cleanup(func())
}

// NewToy creates a new instance of Toy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewToy(t mockConstructorTestingTNewToy) *Toy {
	mock := &Toy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
