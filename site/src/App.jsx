import { Routes, Route } from 'react-router-dom'
import HomePage from './pages/HomePage'
import CategoryPage from './pages/CategoryPage'
import ProblemPage from './pages/ProblemPage'
import './App.css'

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/:category" element={<CategoryPage />} />
      <Route path="/:category/:problem" element={<ProblemPage />} />
    </Routes>
  )
}

export default App
