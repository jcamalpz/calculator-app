import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import App from './App';

// Mock fetch
global.fetch = jest.fn();

describe('Calculator App', () => {
  beforeEach(() => {
    (fetch as jest.Mock).mockClear();
  });

  test('renders calculator title', () => {
    render(<App />);
    const titleElement = screen.getByText(/CalcAPI Pro/i);
    expect(titleElement).toBeInTheDocument();
  });

  test('displays initial value of 0', () => {
    render(<App />);
    // The display should show 0 initially - there might be multiple "0" elements (button and display)
    const zeros = screen.getAllByText('0');
    expect(zeros.length).toBeGreaterThan(0);
  });

  test('handles number input', () => {
    render(<App />);
    const button5 = screen.getByRole('button', { name: '5' });
    fireEvent.click(button5);
    // The display should show 5 (there might be multiple "5" texts, but display should contain it)
    const displayElements = screen.getAllByText('5');
    expect(displayElements.length).toBeGreaterThan(0);
  });

  test('performs addition via API', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ result: 15, operation: 'add' }),
    });

    render(<App />);
    
    // Click number buttons to enter 10
    fireEvent.click(screen.getByRole('button', { name: '1' }));
    fireEvent.click(screen.getByRole('button', { name: '0' }));
    // Click operation button
    fireEvent.click(screen.getByRole('button', { name: '+' }));
    // Click second number
    fireEvent.click(screen.getByRole('button', { name: '5' }));
    // Click equals
    fireEvent.click(screen.getByRole('button', { name: '=' }));

    await waitFor(() => {
      expect(screen.getByText('15')).toBeInTheDocument();
    });

    // Verify API was called
    expect(fetch).toHaveBeenCalled();
    const callArgs = (fetch as jest.Mock).mock.calls[0];
    expect(callArgs[0]).toBe('http://localhost:8080/api/v1/calculate/add');
    expect(callArgs[1].method).toBe('POST');
    expect(callArgs[1].headers['Content-Type']).toBe('application/json');
    const body = JSON.parse(callArgs[1].body);
    expect(body.a).toBe(10);
    expect(body.b).toBe(5);
  });

  test('handles division by zero error', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
      ok: false,
      json: async () => ({ error: 'division by zero' }),
    });

    render(<App />);
    
    fireEvent.click(screen.getByRole('button', { name: '1' }));
    fireEvent.click(screen.getByRole('button', { name: '0' }));
    fireEvent.click(screen.getByRole('button', { name: '÷' }));
    fireEvent.click(screen.getByRole('button', { name: '0' }));
    fireEvent.click(screen.getByRole('button', { name: '=' }));

    await waitFor(() => {
      const errorElement = screen.getByText(/division by zero/i);
      expect(errorElement).toBeInTheDocument();
    });
  });

  test('clears display on AC button', () => {
    render(<App />);
    fireEvent.click(screen.getByRole('button', { name: '5' }));
    // Verify display shows 5 (might be multiple elements with "5")
    const fives = screen.getAllByText('5');
    expect(fives.length).toBeGreaterThan(0);
    // Click AC button
    fireEvent.click(screen.getByRole('button', { name: 'AC' }));
    // Display should be back to 0 - there might be multiple "0" elements (button and display)
    const zeros = screen.getAllByText('0');
    expect(zeros.length).toBeGreaterThan(0);
  });

  test('handles square root operation', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ result: 5, operation: 'sqrt' }),
    });

    render(<App />);
    
    fireEvent.click(screen.getByRole('button', { name: '2' }));
    fireEvent.click(screen.getByRole('button', { name: '5' }));
    fireEvent.click(screen.getByRole('button', { name: '√' }));

    await waitFor(() => {
      expect(screen.getByText('5')).toBeInTheDocument();
    });

    // Verify API was called correctly
    expect(fetch).toHaveBeenCalled();
    const callArgs = (fetch as jest.Mock).mock.calls[0];
    expect(callArgs[0]).toBe('http://localhost:8080/api/v1/calculate/sqrt');
    expect(callArgs[1].method).toBe('POST');
    const body = JSON.parse(callArgs[1].body);
    expect(body.a).toBe(25);
  });

  test('concatenates multiple number inputs', () => {
    render(<App />);
    
    fireEvent.click(screen.getByRole('button', { name: '1' }));
    fireEvent.click(screen.getByRole('button', { name: '2' }));
    fireEvent.click(screen.getByRole('button', { name: '3' }));
    
    // Display should show 123
    expect(screen.getByText('123')).toBeInTheDocument();
  });

  test('handles decimal point input', () => {
    render(<App />);
    
    fireEvent.click(screen.getByRole('button', { name: '5' }));
    fireEvent.click(screen.getByRole('button', { name: '.' }));
    fireEvent.click(screen.getByRole('button', { name: '2' }));
    
    // Display should show 5.2
    expect(screen.getByText('5.2')).toBeInTheDocument();
  });
});
