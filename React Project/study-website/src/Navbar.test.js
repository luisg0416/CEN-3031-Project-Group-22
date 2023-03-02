import { render, fireEvent, screen } from "@testing-library/react";
import '@testing-library/jest-dom'

test('should navigate to home when link is clicked', () => {
    const { getByText } = render(<a href="/">Click Me</a>);

    const link = getByText('Click Me');

    fireEvent.click(link);
    
    expect(screen.getByRole('link')).toHaveAttribute('href', '/');
  });

  test('should navigate to about when link is clicked', () => {
    const { getByText } = render(<a href="/about">Click Me</a>);

    const link = getByText('Click Me');

    fireEvent.click(link);
    
    expect(screen.getByRole('link')).toHaveAttribute('href', '/about');
  });

  test('should navigate to sign-in when link is clicked', () => {
    const { getByText } = render(<a href="/sign-in">Click Me</a>);

    const link = getByText('Click Me');

    fireEvent.click(link);
    
    expect(screen.getByRole('link')).toHaveAttribute('href', '/sign-in');
  });