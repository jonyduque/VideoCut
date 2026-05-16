import React from 'react';
import { useStore } from './store';
import { SelectDirectory, SplitVideo } from '../wailsjs/go/main/App';

const Settings: React.FC = () => {
  const { 
    filePath, cutPoint, overlap, outputDir, 
    part1Name, part2Name, useSourceFolder, isProcessing,
    setOverlap, setOutputDir, setPart1Name, setPart2Name, 
    setUseSourceFolder, setIsProcessing 
  } = useStore();

  const handleBrowse = async () => {
    try {
      const path = await SelectDirectory();
      if (path) {
        setOutputDir(path);
      }
    } catch (err) {
      console.error("SelectDirectory error:", err);
    }
  };

  const handleSplit = async () => {
    if (!filePath) return;
    
    try {
      setIsProcessing(true);
      await SplitVideo(
        filePath, 
        cutPoint, 
        overlap, 
        useSourceFolder ? "" : outputDir, 
        part1Name, 
        part2Name
      );
      alert("Vídeo dividido com sucesso!");
    } catch (err) {
      console.error("SplitVideo error:", err);
      alert("Erro ao dividir vídeo: " + err);
    } finally {
      setIsProcessing(false);
    }
  };

  return (
    <div className="settings-panel">
      <div className="form-group">
        <label>Sobra (segundos):</label>
        <input 
          type="number" 
          value={overlap} 
          onChange={(e) => setOverlap(parseFloat(e.target.value))}
          min="0"
          step="0.5"
        />
      </div>

      <div className="form-row">
        <div className="form-group">
          <label>Nome Parte 1:</label>
          <input 
            type="text" 
            value={part1Name} 
            onChange={(e) => setPart1Name(e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Nome Parte 2:</label>
          <input 
            type="text" 
            value={part2Name} 
            onChange={(e) => setPart2Name(e.target.value)}
          />
        </div>
      </div>

      <div className="destination-group">
        <label className="checkbox-label">
          <input 
            type="checkbox" 
            checked={useSourceFolder} 
            onChange={(e) => setUseSourceFolder(e.target.checked)}
          />
          Salvar na mesma pasta do original
        </label>
        
        {!useSourceFolder && (
          <div className="custom-dest">
            <input type="text" readOnly value={outputDir} placeholder="Selecione a pasta..." />
            <button onClick={handleBrowse} className="secondary-btn">Procurar</button>
          </div>
        )}
      </div>

      <button 
        onClick={handleSplit} 
        className="primary-btn" 
        disabled={isProcessing || !filePath}
      >
        {isProcessing ? "Processando..." : "CORTAR VÍDEO"}
      </button>
    </div>
  );
};

export default Settings;
