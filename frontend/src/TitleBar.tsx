import React from 'react';
import { WindowMinimise, Quit } from '../wailsjs/runtime/runtime';

const TitleBar: React.FC = () => {
  return (
    <div className="title-bar" style={{ '--wails-draggable': 'drag' } as any}>
      <div className="title-spacer"></div>
      <div className="title-content">
        <span className="app-icon">🎬</span>
        <span className="app-name">VideoCut</span>
      </div>
      <div className="window-controls" style={{ '--wails-draggable': 'no-drag' } as any}>
        <button onClick={() => WindowMinimise()} title="Minimizar">─</button>
        <button onClick={() => Quit()} className="close-btn" title="Fechar">✕</button>
      </div>
    </div>
  );
};

export default TitleBar;
