import { useRef, useState } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import '../Visualizer.css'

const NUMS = [2, 7, 11, 15, 3, 6]
const TARGET = 9
const STEP_DELAY_MS = 700

function* twoSumSteps(nums, target) {
  const seen = {}
  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i]
    yield { index: i, seen: { ...seen }, found: null }
    if (complement in seen) {
      yield { index: i, seen: { ...seen }, found: [seen[complement], i] }
      return
    }
    seen[nums[i]] = i
  }
}

export default function TwoSum() {
  const [index, setIndex] = useState(-1)
  const [seen, setSeen] = useState({})
  const [found, setFound] = useState(null)
  const [running, setRunning] = useState(false)
  const stopRef = useRef(false)

  async function play() {
    if (running) return
    setRunning(true)
    stopRef.current = false
    reset(false)
    for (const step of twoSumSteps(NUMS, TARGET)) {
      if (stopRef.current) break
      setIndex(step.index)
      setSeen(step.seen)
      setFound(step.found)
      await new Promise((resolve) => setTimeout(resolve, STEP_DELAY_MS))
    }
    setRunning(false)
  }

  function reset(stop = true) {
    if (stop) stopRef.current = true
    setIndex(-1)
    setSeen({})
    setFound(null)
    setRunning(false)
  }

  return (
    <div className="visualizer">
      <p className="visualizer-goal">
        Target sum: <strong>{TARGET}</strong>
      </p>
      <div className="visualizer-controls">
        <button onClick={play} disabled={running}>
          {running ? 'Scanning…' : 'Play'}
        </button>
        <button onClick={() => reset(true)} disabled={running}>
          Reset
        </button>
      </div>

      <div className="visualizer-row" aria-label="input array">
        {NUMS.map((value, i) => (
          <motion.div
            key={i}
            layout
            className="visualizer-cell"
            data-active={i === index}
            data-matched={found?.includes(i) ?? false}
            animate={{ scale: i === index ? 1.1 : 1 }}
            transition={{ type: 'spring', stiffness: 300, damping: 20 }}
          >
            <span className="visualizer-cell-index">{i}</span>
            {value}
          </motion.div>
        ))}
      </div>

      <div className="visualizer-panel">
        <h3>Seen (value → index)</h3>
        <div className="visualizer-hashmap">
          <AnimatePresence>
            {Object.entries(seen).map(([value, i]) => (
              <motion.div
                key={value}
                className="visualizer-chip"
                initial={{ opacity: 0, y: -10, scale: 0.8 }}
                animate={{ opacity: 1, y: 0, scale: 1 }}
                exit={{ opacity: 0, scale: 0.8 }}
              >
                {value} → {i}
              </motion.div>
            ))}
          </AnimatePresence>
        </div>
      </div>

      {found && (
        <motion.p
          className="visualizer-result"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
        >
          Found: indices [{found[0]}, {found[1]}] — {NUMS[found[0]]} + {NUMS[found[1]]} = {TARGET}
        </motion.p>
      )}
    </div>
  )
}
