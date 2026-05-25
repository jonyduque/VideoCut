import React from 'react';
import { WindowMinimise, Quit } from '../wailsjs/runtime/runtime';

const TitleBar: React.FC = () => {
  return (
    <div className="title-bar" style={{ '--wails-draggable': 'drag' } as any}>
      <div className="title-spacer"></div>
      <div className="title-content">
        <div className="app-icon">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" width="48px" height="48px">
            <g transform="translate(256, 245) rotate(-10) scale(1.2)">
              <path d="M 0 -90 H -140 A 20 20 0 0 0 -160 -70 V 70 A 20 20 0 0 0 -140 90 H 0 Z" fill="#4d6780" stroke="#2c3e50" strokeWidth="10" strokeLinejoin="round" strokeLinecap="round"/>
              <path d="M 0 -65 H -120 A 12 12 0 0 0 -132 -53 V 53 A 12 12 0 0 0 -120 65 H 0 Z" fill="#1e608d" stroke="#2c3e50" strokeWidth="8" strokeLinejoin="round" strokeLinecap="round"/>
              <rect x="-129" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-95" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-61" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-27" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-129" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-95" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-61" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="-27" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <path d="M -45 -45 L 0 -18 L 0 18 L -45 45 Z" fill="#97bdd7" stroke="#2c3e50" strokeWidth="8" strokeLinejoin="round" strokeLinecap="round"/>
            </g>
            <g transform="translate(256, 295) rotate(10) scale(1.2)">
              <path d="M 0 -90 H 140 A 20 20 0 0 1 160 -70 V 70 A 20 20 0 0 1 140 90 H 0 Z" fill="#4d6780" stroke="#2c3e50" strokeWidth="10" strokeLinejoin="round" strokeLinecap="round"/>
              <path d="M 0 -65 H 120 A 12 12 0 0 1 132 -53 V 53 A 12 12 0 0 1 120 65 H 0 Z" fill="#1e608d" stroke="#2c3e50" strokeWidth="8" strokeLinejoin="round" strokeLinecap="round"/>
              <rect x="19" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="51" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="83" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="114" y="-80" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="19" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="51" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="83" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <rect x="114" y="66" width="16" height="16" rx="4" fill="#2c3e50"/>
              <path d="M 0 -18 L 30 0 L 0 18 Z" fill="#97bdd7" stroke="#2c3e50" strokeWidth="8" strokeLinejoin="round" strokeLinecap="round"/>
            </g>
          </svg>
        </div>
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
