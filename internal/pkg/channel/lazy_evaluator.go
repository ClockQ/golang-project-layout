package channel

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)

	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	go loopFunc()

	retFunc := func() Any {
		return <-retValChan
	}
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState int) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
