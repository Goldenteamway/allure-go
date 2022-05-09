package teardown

import (
	"errors"
	"github.com/Goldenteamway/allure-go"
	"testing"
)

func TestDeferredAfterFail(t *testing.T) {
	allure.Test(t,
		allure.Description("Break now"),
		allure.Action(func() {
			//panic("panic at the before statement! (disco)")
			allure.Fail(errors.New("break now statement"))

			allure.Step(allure.Description("Step after break"),
				allure.Action(func() {}))
		}))
	defer func() {
		allure.AfterTest(t, allure.Action(func() {
			allure.Step(allure.Description("Running after a panic"))
		}))
	}()
}
