import { Link, useParams } from 'react-router-dom'
import { findProblem } from '../problems/registry'

export default function ProblemPage() {
  const { category: categorySlug, problem: problemSlug } = useParams()
  const found = findProblem(categorySlug, problemSlug)

  if (!found) {
    return (
      <div className="page">
        <Link to="/">&larr; Home</Link>
        <p>Problem not found.</p>
      </div>
    )
  }

  const { category, problem } = found
  const Visualizer = problem.Visualizer

  return (
    <div className="page">
      <Link to={`/${category.slug}`} className="breadcrumb">
        &larr; {category.title}
      </Link>
      <h1>{problem.title}</h1>
      <p>
        <span className="difficulty" data-level={problem.difficulty}>
          {problem.difficulty}
        </span>{' '}
        &middot;{' '}
        <a href={problem.leetcodeUrl} target="_blank" rel="noreferrer">
          View on LeetCode
        </a>
      </p>

      {Visualizer ? (
        <Visualizer />
      ) : (
        <div className="visualizer-placeholder">Visualizer coming soon.</div>
      )}
    </div>
  )
}
