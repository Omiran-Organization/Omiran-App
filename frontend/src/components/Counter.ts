import { useSelector, useDispatch } from 'react-redux';

const useCounter: object = () => {
  const count: any = useSelector((state) => state.count)
  const dispatch: any = useDispatch()
  const increment: object = () =>
    dispatch({
      type: 'INCREMENT',
    })
  const decrement: object = () =>
    dispatch({
      type: 'DECREMENT',
    })
  const reset: object = () =>
    dispatch({
      type: 'RESET',
    })
  return { count, increment, decrement, reset }
}

const Counter: any = () => {
  const { count, increment, decrement, reset } = useCounter()
  return (
    <div>
      <h1>
        Count: <span>{count}</span>
      </h1>
      <button onClick={increment}>+1</button>
      <button onClick={decrement}>-1</button>
      <button onClick={reset}>Reset</button>
    </div>
  )
}

export default Counter;
