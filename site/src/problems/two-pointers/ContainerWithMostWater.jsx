import { useRef, useState } from 'react'
import { motion } from 'framer-motion'
import '../Visualizer.css'

const HEIGHTS = [1, 8, 6, 2, 5, 4, 8, 3, 7]
const STEP_DELAY_MS = 500

function* containerSteps(heights) {
  let left = 0
  let right = heights.length - 1
  let best = { area: 0, left, right }
  while (left < right) {
    const area = Math.min(heights[left], heights[right]) * (right - left)
    if (area > best.area) best = { area, left, right }
    yield { left, right, area, best }
    if (heights[left] < heights[right]) left++
    else right--
  }
}

export default function ContainerWithMostWater() {
  const [pointers, setPointers] = useState({ left: 0, right: HEIGHTS.length - 1 })
  const [area, setArea] = useState(0)
  const [best, setBest] = useState({ area: 0, left: 0, right: HEIGHTS.length - 1 })
  const [running, setRunning] = useState(false)
  const stopRef = useRef(false)
  const max = Math.max(...HEIGHTS)

  async function play() {
    if (running) return
    setRunning(true)
    stopRef.current = false
    reset(false)
    for (const step of containerSteps(HEIGHTS)) {
      if (stopRef.current) break
      setPointers({ left: step.left, right: step.right })
      setArea(step.area)
      setBest(step.best)
      await new Promise((resolve) => setTimeout(resolve, STEP_DELAY_MS))
    }
    setRunning(false)
  }

  function reset(stop = true) {
    if (stop) stopRef.current = true
    setPointers({ left: 0, right: HEIGHTS.length - 1 })
    setArea(0)
    setBest({ area: 0, left: 0, right: HEIGHTS.length - 1 })
    setRunning(false)
  }

  return (
    <div className="visualizer">
      <div className="visualizer-controls">
        <button onClick={play} disabled={running}>
          {running ? 'Searching…' : 'Play'}
        </button>
        <button onClick={() => reset(true)} disabled={running}>
          Reset
        </button>
      </div>

      <div className="visualizer-bars" aria-label="height bars">
        {HEIGHTS.map((value, i) => (
          <motion.div
            key={i}
            className="visualizer-bar"
            data-pointer={i === pointers.left || i === pointers.right}
            data-in-container={i >= pointers.left && i <= pointers.right}
            style={{ height: `${(value / max) * 100}%` }}
            animate={{ opacity: i >= pointers.left && i <= pointers.right ? 1 : 0.35 }}
          >
            {value}
          </motion.div>
        ))}
      </div>

      <div className="visualizer-panel visualizer-panel-row">
        <div>
          Current area: <strong>{area}</strong>
        </div>
        <div>
          Best area: <strong>{best.area}</strong> (left {best.left}, right {best.right})
        </div>
      </div>
    </div>
  )
}
