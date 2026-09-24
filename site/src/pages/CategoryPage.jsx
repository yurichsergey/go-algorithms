import { Link, useParams } from 'react-router-dom'
import { findCategory } from '../problems/registry'

export default function CategoryPage() {
  const { category: categorySlug } = useParams()
  const category = findCategory(categorySlug)

  if (!category) {
    return (
      <div className="page">
        <Link to="/">&larr; Home</Link>
        <p>Category not found.</p>
      </div>
    )
  }

  return (
    <div className="page">
      <Link to="/" className="breadcrumb">
        &larr; All categories
      </Link>
      <h1>{category.title}</h1>
      <ul className="problem-list">
        {category.problems.map((problem) => (
          <li key={problem.slug}>
            <Link to={`/${category.slug}/${problem.slug}`}>{problem.title}</Link>
            <span className="difficulty" data-level={problem.difficulty}>
              {problem.difficulty}
            </span>
            {!problem.Visualizer && <span className="badge-soon">coming soon</span>}
          </li>
        ))}
      </ul>
    </div>
  )
}
