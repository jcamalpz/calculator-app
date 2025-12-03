
import React, { useState } from 'react';
import { Calculator, Trash2, Divide, Plus, Minus, X } from 'lucide-react';

const API_BASE_URL = 'http://localhost:8080/api/v1';

type Operation = 'add' | 'subtract' | 'multiply' | 'divide' | 'power' | 'sqrt' | 'percentage';

interface CalculationResult {
  result: number;
  operation: string;
}

interface ApiError {
  error: string;
}

const CalculatorApp: React.FC = () => {
  const [display, setDisplay] = useState('0');
  const [firstOperand, setFirstOperand] = useState<number | null>(null);
  const [operation, setOperation] = useState<Operation | null>(null);
  const [waitingForSecond, setWaitingForSecond] = useState(false);
  const [history, setHistory] = useState<string[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleNumber = (num: string) => {
    setError(null);
    if (waitingForSecond) {
      setDisplay(num);
      setWaitingForSecond(false);
    } else {
      setDisplay(display === '0' ? num : display + num);
    }
  };

  const handleDecimal = () => {
    if (waitingForSecond) {
      setDisplay('0.');
      setWaitingForSecond(false);
    } else if (!display.includes('.')) {
      setDisplay(display + '.');
    }
  };

  const callApi = async (endpoint: string, payload: any): Promise<number> => {
    setLoading(true);
    try {
      const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });

      const data = await response.json();

      if (!response.ok) {
        const errorData = data as ApiError;
        throw new Error(errorData.error || 'API request failed');
      }

      const result = data as CalculationResult;
      return result.result;
    } catch (err) {
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const performCalculation = async (op: Operation, a: number, b?: number) => {
    try {
      let result: number;
      let historyEntry: string;

      switch (op) {
        case 'add':
          result = await callApi('/calculate/add', { a, b });
          historyEntry = `${a} + ${b} = ${result}`;
          break;
        case 'subtract':
          result = await callApi('/calculate/subtract', { a, b });
          historyEntry = `${a} - ${b} = ${result}`;
          break;
        case 'multiply':
          result = await callApi('/calculate/multiply', { a, b });
          historyEntry = `${a} × ${b} = ${result}`;
          break;
        case 'divide':
          result = await callApi('/calculate/divide', { a, b });
          historyEntry = `${a} ÷ ${b} = ${result}`;
          break;
        case 'power':
          result = await callApi('/calculate/power', { a, b });
          historyEntry = `${a} ^ ${b} = ${result}`;
          break;
        case 'sqrt':
          result = await callApi('/calculate/sqrt', { a });
          historyEntry = `√${a} = ${result}`;
          break;
        case 'percentage':
          result = await callApi('/calculate/percentage', { a, b });
          historyEntry = `${a}% of ${b} = ${result}`;
          break;
        default:
          throw new Error('Invalid operation');
      }

      setDisplay(result.toString());
      setHistory([historyEntry, ...history].slice(0, 5));
      setFirstOperand(null);
      setOperation(null);
      setError(null);
      return result;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Calculation failed';
      setError(errorMessage);
      setDisplay('Error');
      setFirstOperand(null);
      setOperation(null);
    }
  };

  const handleOperation = (op: Operation) => {
    const current = parseFloat(display);
    
    if (op === 'sqrt') {
      performCalculation('sqrt', current);
      return;
    }

    if (firstOperand === null) {
      setFirstOperand(current);
      setOperation(op);
      setWaitingForSecond(true);
    } else if (operation) {
      performCalculation(operation, firstOperand, current);
    }
  };

  const handleEquals = () => {
    if (firstOperand !== null && operation) {
      const current = parseFloat(display);
      performCalculation(operation, firstOperand, current);
    }
  };

  const handleClear = () => {
    setDisplay('0');
    setFirstOperand(null);
    setOperation(null);
    setWaitingForSecond(false);
    setError(null);
  };

  const handleClearHistory = () => {
    setHistory([]);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex items-center justify-center p-4">
      <div className="w-full max-w-4xl">
        <div className="text-center mb-8">
          <div className="flex items-center justify-center gap-3 mb-2">
            <Calculator className="w-10 h-10 text-purple-400" />
            <h1 className="text-4xl font-bold text-white">CalcAPI Pro</h1>
          </div>
          <p className="text-purple-300">Full-Stack Calculator with REST API Backend</p>
        </div>

        <div className="grid md:grid-cols-3 gap-6">
          {/* History Panel */}
          <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-6 border border-white/20">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-lg font-semibold text-white">History</h2>
              <button
                onClick={handleClearHistory}
                className="text-purple-300 hover:text-purple-200 transition-colors"
                title="Clear history"
              >
                <Trash2 className="w-4 h-4" />
              </button>
            </div>
            <div className="space-y-2">
              {history.length === 0 ? (
                <p className="text-purple-300/50 text-sm">No calculations yet</p>
              ) : (
                history.map((item, idx) => (
                  <div
                    key={idx}
                    className="text-purple-200 text-sm bg-white/5 rounded-lg p-2 font-mono"
                  >
                    {item}
                  </div>
                ))
              )}
            </div>
          </div>

          {/* Calculator */}
          <div className="md:col-span-2 bg-white/10 backdrop-blur-lg rounded-2xl p-6 border border-white/20">
            {/* Display */}
            <div className="mb-6">
              <div className="bg-slate-800/50 rounded-xl p-6 mb-2">
                {operation && firstOperand !== null && (
                  <div className="text-purple-300 text-sm mb-1 font-mono">
                    {firstOperand} {operation === 'add' ? '+' : operation === 'subtract' ? '-' : operation === 'multiply' ? '×' : operation === 'divide' ? '÷' : operation === 'power' ? '^' : operation === 'percentage' ? '%' : ''}
                  </div>
                )}
                <div className="text-white text-4xl font-bold text-right font-mono break-all">
                  {loading ? 'Loading...' : display}
                </div>
              </div>
              {error && (
                <div className="bg-red-500/20 border border-red-500/50 rounded-lg p-3 text-red-200 text-sm">
                  {error}
                </div>
              )}
            </div>

            {/* Buttons */}
            <div className="grid grid-cols-4 gap-3">
              {/* Row 1 */}
              <button onClick={handleClear} className="col-span-2 bg-red-500/20 hover:bg-red-500/30 text-red-200 font-semibold py-4 rounded-xl transition-all">
                AC
              </button>
              <button onClick={() => handleOperation('percentage')} className="bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                %
              </button>
              <button onClick={() => handleOperation('divide')} className="bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                ÷
              </button>

              {/* Row 2 */}
              <button onClick={() => handleNumber('7')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                7
              </button>
              <button onClick={() => handleNumber('8')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                8
              </button>
              <button onClick={() => handleNumber('9')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                9
              </button>
              <button onClick={() => handleOperation('multiply')} className="bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                ×
              </button>

              {/* Row 3 */}
              <button onClick={() => handleNumber('4')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                4
              </button>
              <button onClick={() => handleNumber('5')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                5
              </button>
              <button onClick={() => handleNumber('6')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                6
              </button>
              <button onClick={() => handleOperation('subtract')} className="bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                -
              </button>

              {/* Row 4 */}
              <button onClick={() => handleNumber('1')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                1
              </button>
              <button onClick={() => handleNumber('2')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                2
              </button>
              <button onClick={() => handleNumber('3')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                3
              </button>
              <button onClick={() => handleOperation('add')} className="bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                +
              </button>

              {/* Row 5 */}
              <button onClick={() => handleOperation('sqrt')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                √
              </button>
              <button onClick={() => handleNumber('0')} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                0
              </button>
              <button onClick={handleDecimal} className="bg-slate-700/50 hover:bg-slate-600/50 text-white font-semibold py-4 rounded-xl transition-all">
                .
              </button>
              <button onClick={handleEquals} className="bg-green-500/30 hover:bg-green-500/40 text-green-200 font-semibold py-4 rounded-xl transition-all">
                =
              </button>

              {/* Row 6 */}
              <button onClick={() => handleOperation('power')} className="col-span-4 bg-purple-500/20 hover:bg-purple-500/30 text-purple-200 font-semibold py-4 rounded-xl transition-all">
                x^y (Power)
              </button>
            </div>

            <div className="mt-6 text-center text-purple-300/60 text-xs">
              Connected to API: {API_BASE_URL}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CalculatorApp;