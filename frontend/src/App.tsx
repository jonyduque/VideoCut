import { useStore } from './store';
import TitleBar from './TitleBar';
import DropZone from './DropZone';
import Timeline from './Timeline';
import Settings from './Settings';
import './App.css';

function App() {
  const { isProcessing, filePath, duration, setFilePath, setDuration, setSuccessMessage } = useStore();

  const clearFile = () => {
    setFilePath('');
    setDuration(0);
    setSuccessMessage('');
  };

  return (
    <div id="App">
      <TitleBar />

      <main>
        <DropZone />

        {filePath && (
          <>

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
