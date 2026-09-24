import { Link } from 'react-router-dom'
import { categories } from '../problems/registry'

export default function HomePage() {
  return (
    <div className="page">
      <h1>LeetCode Tutoring</h1>
      <p>Animated, step-by-step walkthroughs of algorithm problems.</p>
      <div className="category-grid">
        {categories.map((category) => (
          <Link key={category.slug} to={`/${category.slug}`} className="category-card">
            <h2>{category.title}</h2>
            <p>{category.problems.length} problems</p>
          </Link>
        ))}
      </div>
    </div>
  )
}
