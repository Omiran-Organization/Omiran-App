import React from 'react'
import { render, screen } from '@testing-library/react'
import Index from '../pages/index'

test('renders deploy link', () => {
  const { getByText } = render(<Index />)
  const linkElement = getByText(
    /Omiran/
  )
  expect(linkElement).toBeInTheDocument()
})