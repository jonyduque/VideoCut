import React from 'react';
import { useStore } from './store';

const Timeline: React.FC = () => {
  const { duration, cutPoint, setCutPoint, overlap } = useStore();

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

  // Calculate overlap bar position and width
  const overlapWidth = (overlap * 2 / duration) * 100;
  const overlapLeft = ((cutPoint - overlap) / duration) * 100;

  return (
    <div className="timeline-container">
      <div className="timeline-labels">
        <span>0:00.00</span>
        <span className="current-time">{formatTime(cutPoint)}</span>
        <span>{formatTime(duration)}</span>
      </div>
      <div className="slider-wrapper">
        <div
          className="overlap-visual"
          style={{
            left: `${Math.max(0, overlapLeft)}%`,
            width: `${Math.min(100 - overlapLeft, overlapWidth)}%`
          }}
        ></div>
        <input
          type="range"
          min="0"
          max={duration}
          step="0.01"
          value={cutPoint}
          onChange={handleChange}
          className="timeline-slider"
        />
      </div>
    </div>
  );
};

export default Timeline;
