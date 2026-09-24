import TwoSum from './arrays-hashing/TwoSum'
import ContainerWithMostWater from './two-pointers/ContainerWithMostWater'

export const categories = [
  {
    slug: 'arrays-hashing',
    title: 'Arrays & Hashing',
    problems: [
      { slug: 'contains-duplicate', title: 'Contains Duplicate', difficulty: 'Easy', leetcodeUrl: 'https://leetcode.com/problems/contains-duplicate/', Visualizer: null },
      { slug: 'valid-anagram', title: 'Valid Anagram', difficulty: 'Easy', leetcodeUrl: 'https://leetcode.com/problems/valid-anagram/', Visualizer: null },
      { slug: 'two-sum', title: 'Two Sum', difficulty: 'Easy', leetcodeUrl: 'https://leetcode.com/problems/two-sum/', Visualizer: TwoSum },
      { slug: 'group-anagrams', title: 'Group Anagrams', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/group-anagrams/', Visualizer: null },
      { slug: 'top-k-frequent-elements', title: 'Top K Frequent Elements', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/top-k-frequent-elements/', Visualizer: null },
      { slug: 'encode-and-decode-strings', title: 'Encode and Decode Strings', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/encode-and-decode-strings/', Visualizer: null },
      { slug: 'products-of-array-except-self', title: 'Product of Array Except Self', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/product-of-array-except-self/', Visualizer: null },
      { slug: 'valid-sudoku', title: 'Valid Sudoku', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/valid-sudoku/', Visualizer: null },
      { slug: 'longest-consecutive-sequence', title: 'Longest Consecutive Sequence', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/longest-consecutive-sequence/', Visualizer: null },
    ],
  },
  {
    slug: 'two-pointers',
    title: 'Two Pointers',
    problems: [
      { slug: 'valid-palindrome', title: 'Valid Palindrome', difficulty: 'Easy', leetcodeUrl: 'https://leetcode.com/problems/valid-palindrome/', Visualizer: null },
      { slug: 'two-integers-sum-2', title: 'Two Sum II - Input Array Is Sorted', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/', Visualizer: null },
      { slug: '3sum', title: '3Sum', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/3sum/', Visualizer: null },
      { slug: 'container-with-most-water', title: 'Container With Most Water', difficulty: 'Medium', leetcodeUrl: 'https://leetcode.com/problems/container-with-most-water/', Visualizer: ContainerWithMostWater },
    ],
  },
]

export function findCategory(categorySlug) {
  return categories.find((c) => c.slug === categorySlug)
}

export function findProblem(categorySlug, problemSlug) {
  const category = findCategory(categorySlug)
  const problem = category?.problems.find((p) => p.slug === problemSlug)
  return category && problem ? { category, problem } : null
}
