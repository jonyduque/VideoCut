import React from 'react';
import { useStore } from './store';

const Timeline: React.FC = () => {
  const { duration, cutPoint, setCutPoint } = useStore();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    let val = parseFloat(e.target.value);
    const middle = duration / 2;
    
    // Snap to middle (1s threshold)
    if (Math.abs(val - middle) < 1) {
      val = middle;
    }
    
    setCutPoint(val);
  };

  const formatTime = (seconds: number) => {
    const mins = Math.floor(seconds / 60);
    const secs = (seconds % 60).toFixed(2);
    return `${mins}:${secs.padStart(5, '0')}`;
  };

  return (
    <div className="timeline-container">
      <div className="timeline-labels">
        <span>0:00.00</span>
        <span className="current-time">{formatTime(cutPoint)}</span>
        <span>{formatTime(duration)}</span>
      </div>
      <input 
        type="range" 
        min="0" 
        max={duration} 
        step="0.01"
        value={cutPoint} 
        onChange={handleChange}
        className="timeline-slider"
      />
      <div className="timeline-info">
        Ponto de Corte: <strong>{cutPoint.toFixed(2)}s</strong>
      </div>
    </div>
  );
};

export default Timeline;
