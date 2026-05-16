import { useStore } from './store';
import DropZone from './DropZone';
import Timeline from './Timeline';
import Settings from './Settings';
import './App.css';

function App() {
  const { isProcessing, filePath, duration } = useStore();

  return (
    <div id="App">
      <header>
        <h1>VideoCut</h1>
      </header>

      <main>
        <DropZone />
        
        {filePath && (
          <>
            <div className="info-panel">
              <p><strong>Arquivo:</strong> {filePath.split(/[\\/]/).pop()}</p>
              <p><strong>Duração:</strong> {duration.toFixed(2)}s</p>
            </div>

            <Timeline />
            
            <Settings />
          </>
        )}
      </main>

      {isProcessing && (
        <div className="loading-overlay">
          <div className="spinner"></div>
          <p>Processando vídeo...</p>
        </div>
      )}
    </div>
  )
}

export default App
